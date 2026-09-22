package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/handle"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Captcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) Captcha(ctx context.Context, req *user.LoginRequest) (resp *user.CaptchaResponse, err error) {
	return handle.Captcha(ctx, req)
}

// SmsCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) SmsCode(ctx context.Context, phone string) (resp int64, err error) {
	return handle.SmsCode(ctx, phone)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	return handle.Login(ctx, req)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.LoginResponse, err error) {
	return handle.Register(ctx, req)
}

// ResetPwd implements the UserServiceImpl interface.
func (s *UserServiceImpl) ResetPwd(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	return handle.ResetPwd(ctx, req)
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	return handle.Logout(ctx, req)
}

// QueryUserHeaderPage implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderPage(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderPageResponse, err error) {
	return handle.QueryUserHeaderPage(ctx, req)
}

// QueryUserHeaderList implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderList(ctx context.Context, req *user.UserHeaderRequest) (resp []*user.UserHeaderResponse, err error) {
	return handle.QueryUserHeaderList(ctx, req)
}

// QueryUserHeaderInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderInfo(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderResponse, err error) {
	return handle.QueryUserHeaderInfo(ctx, req)
}

// QueryUserPositionPage implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserPositionPage(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionPageResponse, err error) {
	return handle.QueryPositionPage(ctx, req)
}

// QueryUserPositionList implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserPositionList(ctx context.Context, req *user.UserPositionRequest) (resp []*user.UserPositionResponse, err error) {
	return handle.QueryPositionList(ctx, req)
}

// QueryUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionResponse, err error) {
	return handle.QueryPosition(ctx, req)
}

// SaveUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) SaveUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return handle.SaveUserPosition(ctx, req)
}

// UpdateUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return handle.UpdateUserPosition(ctx, req)
}

// DeleteUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return handle.DeleteUserPosition(ctx, req)
}
