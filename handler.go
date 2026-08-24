package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/idl/kitex_gen/user_position"
	"github.com/kouleen/user-center/handle"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Captcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) Captcha(ctx context.Context, req *user.LoginRequest) (resp *user.CaptchaResponse, err error) {
	return handle.Captcha(ctx, req)
}

// SmsCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) SmsCode(ctx context.Context, phone string) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	return handle.Login(ctx, req)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// ResetPwd implements the UserServiceImpl interface.
func (s *UserServiceImpl) ResetPwd(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LoginRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderPage implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderPage(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderList implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderList(ctx context.Context, req *user.UserHeaderRequest) (resp []*user.UserHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderInfo(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserPositionPage implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserPositionPage(ctx context.Context, req *user_position.UserPositionRequest) (resp *user_position.UserPositionPageResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) SaveUserPosition(ctx context.Context, req *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserPosition(ctx context.Context, req *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// DeleteUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUserPosition(ctx context.Context, req *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
