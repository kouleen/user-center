package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	rediscli "github.com/kouleen/common/pkg/redis"
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
	if err = rediscli.Del(ctx, loginRequest.GetPhone()); err != nil {
		return
	}
	if userHeader.Status == 0 {
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
	if err = rediscli.Del(ctx, token); err != nil {
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
	userHeader, err := repository.GetUserHeaderByUsername(ctx, loginRequest.Username)
	if err != nil {
		return
	}
	// 删除验证码
	if err = rediscli.Del(ctx, loginRequest.GetUuid()); err != nil {
		return
	}
	// 对比密码
	if err = bcrypt.CompareHashAndPassword([]byte(userHeader.Password), []byte(loginRequest.Password)); err != nil {
		return
	}
	if userHeader.Status == 0 {
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
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if err = rediscli.Del(ctx, token); err != nil {
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
		Username: req.Username,
		Password: password,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Gender:   uint8(req.Gender),
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
