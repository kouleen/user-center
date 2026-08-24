package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
	"github.com/kouleen/user-center/repository"
	"github.com/kouleen/user-center/utils"
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
	if err = utils.Del(ctx, "login:phone:"+loginRequest.GetPhone()); err != nil {
		return
	}
	if common.BaseStatus(userHeader.Status) == common.BaseStatus_DISABLE {
		return nil, errors.New("this account has been deactivated")
	}
	userHeaderByte, err := json.Marshal(userHeader)
	if err != nil {
		return nil, err
	}
	generateUUID := utils.GenerateUUID()
	duration := time.Duration(24) * time.Hour
	idStr := strconv.FormatInt(userHeader.ID, 10)
	token, err := utils.Get(ctx, idStr)
	if err != nil {
		return nil, err
	}
	if err = utils.Del(ctx, token); err != nil {
		return nil, err
	}
	if err = utils.Set(ctx, generateUUID, string(userHeaderByte), duration); err != nil {
		return nil, err
	}
	if err = utils.Set(ctx, idStr, generateUUID, duration); err != nil {
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
	if err = utils.Del(ctx, "login:password:"+loginRequest.GetUuid()); err != nil {
		return
	}
	// 对比密码
	if err = bcrypt.CompareHashAndPassword([]byte(userHeader.Password), []byte(loginRequest.Password)); err != nil {
		return
	}
	if common.BaseStatus(userHeader.Status) == common.BaseStatus_DISABLE {
		return nil, errors.New("this account has been deactivated")
	}
	userHeaderByte, err := json.Marshal(userHeader)
	if err != nil {
		return nil, err
	}
	generateUUID := utils.GenerateUUID()
	duration := time.Duration(24) * time.Hour
	idStr := strconv.FormatInt(userHeader.ID, 10)
	token, err := utils.Get(ctx, idStr)
	if err != nil {
		return nil, err
	}
	if err = utils.Del(ctx, token); err != nil {
		return nil, err
	}
	if err = utils.Set(ctx, generateUUID, string(userHeaderByte), duration); err != nil {
		return nil, err
	}
	if err = utils.Set(ctx, idStr, generateUUID, duration); err != nil {
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
	if err = utils.Set(ctx, uuid, code, duration); err != nil {
		return
	}
	return &user.CaptchaResponse{CaptchaEnabled: true, Img: uuid, Uuid: imgBase64}, nil
}
