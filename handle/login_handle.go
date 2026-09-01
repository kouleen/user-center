package handle

import (
	"context"
	"errors"

	"github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/service"
)

type LoginProcess interface {
	HandleLogin(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error)
}

var loginProcessMap = map[user.LoginType]LoginProcess{
	user.LoginType_LOGIN_PHONE: &service.LoginPhone{},
	user.LoginType_LOGIN_PWD:   &service.LoginPwd{},
}

func Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	if err = checkLoginParams(ctx, req); err != nil {
		return nil, err
	}
	loginProcess, ok := loginProcessMap[user.LoginType(req.GetLoginType())]
	if !ok {
		return nil, errors.New("not found")
	}
	if resp, err = loginProcess.HandleLogin(ctx, req); err != nil {
		return nil, err
	}
	return resp, nil
}

func Captcha(ctx context.Context, req *user.LoginRequest) (resp *user.CaptchaResponse, err error) {
	return service.Captcha(ctx, req)
}

func SendSmsCode(ctx context.Context, loginRequest *user.LoginRequest) (resp bool, err error) {
	return
}

func checkLoginParams(ctx context.Context, req *user.LoginRequest) error {
	if req.LoginType == nil {
		return errors.New("login_type is nil")
	}
	if req.GetCode() == "" {
		return errors.New("code is empty")
	}
	switch user.LoginType(req.GetLoginType()) {
	case user.LoginType_LOGIN_PHONE:
		if req.GetPhone() == "" {
			return errors.New("phone is empty")
		}
		captchaCode, err := redis.Get(ctx, "login:phone:"+req.GetPhone())
		if err != nil {
			return err
		}
		if captchaCode != req.GetCode() {
			return errors.New("captcha code error")
		}
	case user.LoginType_LOGIN_PWD:
		if req.GetUsername() == "" {
			return errors.New("username is empty")
		}
		if req.GetPassword() == "" {
			return errors.New("password is empty")
		}
		if req.GetUuid() == "" {
			return errors.New("uuid is empty")
		}
		captchaCode, err := redis.Get(ctx, "login:password:"+req.GetUuid())
		if err != nil {
			return err
		}
		if captchaCode != req.GetCode() {
			return errors.New("captcha code error")
		}
	default:
		return errors.New("login_type error")
	}
	return nil
}
