package handle

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/google/uuid"
	"github.com/kouleen/common/message"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/service"
	"github.com/kouleen/user-center/utils"
)

type LoginProcess interface {
	HandleLogin(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error)
}

var loginProcessMap = map[user.LoginType]LoginProcess{
	user.LoginType_LOGIN_PHONE: &service.LoginPhone{},
	user.LoginType_LOGIN_PWD:   &service.LoginPwd{},
}

func Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	loginResponse, err := handleLogin(ctx, req)
	os, browser := utils.ParseUA(ctx)
	systemLoginLog := &system.SystemLoginLogRequest{
		Username: req.GetUsername(),
		Ip:       utils.GetClientIp(ctx),
		Os:       os,
		Browser:  browser,
	}
	status := int8(1)
	systemLoginLog.Remark = "登录成功"
	if err == nil {
		systemLoginLog.Token = loginResponse.AccessToken
	}
	if err != nil {
		status = int8(0)
		systemLoginLog.Remark = err.Error()
	}
	systemLoginLog.Status = &status
	messageByte, errMarshal := json.Marshal(&systemLoginLog)
	if errMarshal != nil {
		logger.CtxErrorf(ctx, "errMarshal: %v", errMarshal)
		return
	}
	if err = message.SendToQueue(ctx, "user.login.log.queue", &message.Message{
		TraceID:     ctxutil.GetTraceId(ctx),
		MessageID:   strconv.Itoa(int(uuid.New().ID())),
		ContentType: "application/json",
		Body:        messageByte,
		Headers:     make(map[string]any),
	}); err != nil {
		logger.CtxErrorf(ctx, "send login log error: %v", err)
		return
	}
	return loginResponse, err
}

func handleLogin(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
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

var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

func SmsCode(ctx context.Context, phone string) (resp int64, err error) {
	if phoneRegex.MatchString(phone) {
		return 0, errors.New("invalid phone number")
	}
	return service.SmsCode(ctx, phone)
}

func Register(ctx context.Context, req *user.RegisterRequest) (resp *user.LoginResponse, err error) {
	if err = checkRegister(req); err != nil {
		return nil, err
	}
	return service.Register(ctx, req)
}

func ResetPwd(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	if err = checkResetPwd(req); err != nil {
		return false, err
	}
	return service.ResetPwd(ctx, req)
}

func Logout(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	id := ctxutil.GetUserId(ctx)
	return service.Logout(ctx, id)
}

func checkRegister(req *user.RegisterRequest) error {
	if req.GetUsername() == "" {
		return errors.New("empty username")
	}
	if req.GetPassword() == "" {
		return errors.New("empty password")
	}
	if req.GetNickname() == "" {
		return errors.New("empty nickname")
	}
	if req.GetCode() == "" {
		return errors.New("empty code")
	}
	return nil
}

func checkResetPwd(req *user.LoginRequest) error {
	if req.GetPassword() == "" {
		return errors.New("empty password")
	}
	if req.GetPhone() == "" {
		return errors.New("empty phone")
	}
	if req.GetCode() == "" {
		return errors.New("empty code")
	}
	return nil
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
		captchaCode, err := redis.Get(ctx, req.GetPhone())
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
		captchaCode, err := redis.Get(ctx, req.GetUuid())
		if err != nil {
			return err
		}
		// 转成全小些对比
		if strings.ToLower(captchaCode) != strings.ToLower(req.GetCode()) {
			return errors.New("captcha code error")
		}
	default:
		return errors.New("login_type error")
	}
	return nil
}
