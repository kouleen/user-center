package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/bytedance/gopkg/util/logger"
	rediscli "github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
	"github.com/kouleen/user-center/repository"
	"github.com/kouleen/user-center/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginPhone struct{}

func (p *LoginPhone) HandleLogin(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error) {
	userHeader, err := repository.GetUserHeaderByPhone(ctx, loginRequest.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if userHeader == nil {
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(loginRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		node, err := snowflake.NewNode(1)
		if err != nil {
			return nil, err
		}
		id := node.Generate().Int64()
		userHeader = &modle.UserHeader{
			ID:       id,
			Username: loginRequest.Phone,
			Password: string(hashPwd),
			Nickname: loginRequest.Phone,
			Phone:    loginRequest.Phone,
		}
		if err = repository.CreateUserHeader(ctx, userHeader); err != nil {
			return nil, err
		}
	}
	// 删除验证码
	if _, err = rediscli.Unlink(ctx, loginRequest.GetPhone()); err != nil {
		return
	}
	if common.BaseStatus_DISABLED == common.BaseStatus(userHeader.GetStatus()) {
		return nil, errors.New("this account has been deactivated")
	}
	userHeaderByte, err := json.Marshal(userHeader)
	if err != nil {
		return nil, err
	}
	generateUUID := utils.GenerateUUID()
	duration := time.Duration(24) * time.Hour
	idStr := strconv.FormatInt(userHeader.ID, 10)
	token, err := rediscli.Get(ctx, idStr)
	if err != nil {
		return nil, err
	}
	if _, err = rediscli.Unlink(ctx, token); err != nil {
		return nil, err
	}
	if err = rediscli.Set(ctx, generateUUID, string(userHeaderByte), duration); err != nil {
		return nil, err
	}
	if err = rediscli.Set(ctx, idStr, generateUUID, duration); err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		AccessToken: generateUUID,
		ExpireTime:  duration.Microseconds(),
	}, nil
}

type LoginPwd struct{}

func (p *LoginPwd) HandleLogin(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error) {
	allStart := time.Now()
	start := time.Now()
	userHeader, err := repository.GetUserHeaderByUsername(ctx, loginRequest.Username)
	logger.CtxInfof(ctx, "【1-Mysql查询】耗时=%v\n", time.Since(start))
	if err != nil {
		return
	}
	// 删除验证码
	start = time.Now()
	if _, err = rediscli.Unlink(ctx, loginRequest.GetUuid()); err != nil {
		return
	}
	logger.CtxInfof(ctx, "【2-Redis Unlink验证码】耗时=%v", time.Since(start))
	// 对比密码
	start = time.Now()
	if err = bcrypt.CompareHashAndPassword([]byte(userHeader.Password), []byte(loginRequest.Password)); err != nil {
		return
	}
	logger.CtxInfof(ctx, "【3-Bcrypt】耗时=%v", time.Since(start))
	start = time.Now()
	if common.BaseStatus_DISABLED == common.BaseStatus(userHeader.GetStatus()) {
		return nil, errors.New("this account has been deactivated")
	}
	logger.CtxInfof(ctx, "【4-账号状态判断】耗时=%v\n", time.Since(start))
	start = time.Now()
	userHeaderByte, err := json.Marshal(userHeader)
	logger.CtxInfof(ctx, "【5-JsonMarshal】耗时=%v\n", time.Since(start))
	if err != nil {
		return nil, err
	}
	generateUUID := utils.GenerateUUID()
	duration := time.Duration(24) * time.Hour
	idStr := strconv.FormatInt(userHeader.ID, 10)
	start = time.Now()
	token, err := rediscli.Get(ctx, idStr)
	logger.CtxInfof(ctx, "【6-Redis Get】耗时=%v, err=%v\n", time.Since(start), err)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	start = time.Now()
	if _, err = rediscli.Unlink(ctx, token); err != nil {
		return nil, err
	}
	logger.CtxInfof(ctx, "【7-Redis Unlink旧token】耗时=%v\n", time.Since(start))
	start = time.Now()
	if err = rediscli.Set(ctx, generateUUID, string(userHeaderByte), duration); err != nil {
		return nil, err
	}
	logger.CtxInfof(ctx, "【8-Redis Set user】耗时=%v", time.Since(start))
	start = time.Now()
	if err = rediscli.Set(ctx, idStr, generateUUID, duration); err != nil {
		return nil, err
	}
	logger.CtxInfof(ctx, "【9-Redis Set mapping】耗时=%v", time.Since(start))
	logger.CtxInfof(ctx, "【登录总耗时】%v\n", time.Since(allStart))
	return &user.LoginResponse{
		AccessToken: generateUUID,
		ExpireTime:  duration.Microseconds(),
	}, nil
}

func Captcha(ctx context.Context, req *user.LoginRequest) (resp *user.CaptchaResponse, err error) {
	code := utils.GenerateRandomCode(4)
	uuid := utils.GenerateUUID()
	imgBase64 := utils.CreateCaptchaSvg(code)
	// 验证码保留60秒
	duration := time.Duration(60) * time.Second
	if err = rediscli.Set(ctx, uuid, code, duration); err != nil {
		return
	}
	return &user.CaptchaResponse{
		CaptchaEnabled: true,
		Img:            imgBase64,
		Uuid:           uuid,
	}, nil
}

func SmsCode(ctx context.Context, phone string) (resp int64, err error) {
	limit, err := rediscli.Get(ctx, phone+"-limit")
	if err != nil {
		return
	}
	if limit != "" {
		duration, err := rediscli.Ttl(ctx, phone+"-limit")
		if err != nil {
			return 0, err
		}
		return duration.Milliseconds(), nil
	}
	smsCode := utils.GenerateRandomByNumCode(6)

	// 此处调用短信验证码发送接口 ，发送短信验证码

	// 设置过期时间5分钟
	expire := time.Duration(5) * time.Minute
	if err := rediscli.Set(ctx, phone, smsCode, expire); err != nil {
		return 0, err
	}
	// 限流一分钟
	limitExpire := time.Minute
	if err := rediscli.Set(ctx, phone+"-limit", "1", limitExpire); err != nil {
		return 0, err
	}
	return limitExpire.Milliseconds(), nil
}

func Register(ctx context.Context, req *user.RegisterRequest) (resp *user.LoginResponse, err error) {
	captchaCode, err := rediscli.Get(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	if captchaCode == "" || !strings.EqualFold(captchaCode, req.Code) {
		return nil, errors.New("验证码已过期")
	}
	userHeader, err := repository.GetUserHeaderByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if userHeader != nil {
		return nil, errors.New("该用户名已被注册！")
	}
	if req.Phone != "" {
		userHeader, err = repository.GetUserHeaderByPhone(ctx, req.Phone)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if userHeader != nil {
			return nil, errors.New("该手机号已被注册！")
		}
	}
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	password := string(hashPwd)
	if err = repository.CreateUserHeader(ctx, &modle.UserHeader{
		Username: req.GetUsername(),
		Password: password,
		Nickname: req.GetNickname(),
		Phone:    req.GetPhone(),
		Gender:   req.Gender,
	}); err != nil {
		return nil, err
	}
	userHeader, err = repository.GetUserHeaderByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if userHeader == nil {
		return nil, errors.New("注册失败！")
	}
	userHeader.ID = id
	bytes, err := json.Marshal(userHeader)
	if err != nil {
		return nil, err
	}
	uuid := utils.GenerateUUID()
	duration := time.Duration(24) * time.Hour
	if err := rediscli.Set(ctx, uuid, string(bytes), duration); err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		AccessToken: uuid,
		ExpireTime:  duration.Microseconds(),
	}, nil
}

func ResetPwd(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	smsCode, err := rediscli.Get(ctx, req.Phone)
	if err != nil {
		return false, err
	}
	if smsCode == "" || !strings.EqualFold(smsCode, req.Code) {
		return false, errors.New("验证码已过期")
	}
	userHeader, err := repository.GetUserHeaderByPhone(ctx, req.Phone)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	userHeader.Password = string(hashPwd)
	if err = repository.UpdateUserHeader(ctx, userHeader); err != nil {
		return false, err
	}
	return true, nil
}

func Logout(ctx context.Context, id int64) (resp bool, err error) {
	userIdStr := strconv.FormatInt(id, 10)
	token, err := rediscli.Get(ctx, userIdStr)
	if err != nil {
		return false, err
	}
	if _, err = rediscli.Unlink(ctx, token); err != nil {
		return false, err
	}
	if _, err = rediscli.Unlink(ctx, userIdStr); err != nil {
		return false, err
	}
	return true, nil
}
