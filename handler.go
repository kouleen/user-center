package main

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/idl/kitex_gen/user_position"
	"github.com/kouleen/user-center/service"
	"github.com/kouleen/user-center/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error) {
	return service.Login(ctx, loginRequest)
}

// SendSmsMessage implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendSmsMessage(ctx context.Context, loginRequest *user.LoginRequest) (resp bool, err error) {
	return
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, loginRequest *user.LoginRequest) (resp bool, err error) {
	traceId, _ := metainfo.GetPersistentValue(ctx, "x-trace-id")
	userId, _ := metainfo.GetPersistentValue(ctx, "x-user-id")
	logger.CtxInfof(ctx, "Logout request traceId: %s,userId: %s", traceId, userId)
	token, _ := metainfo.GetPersistentValue(ctx, "x-token")
	if err = utils.Del(ctx, "auth:token:"+token); err != nil {
		return false, err
	}
	return true, nil
}

// QueryUserHeaderPage implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderPage(ctx context.Context, userHeaderRequest *user.UserHeaderRequest) (resp *user.UserHeaderPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderList implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderList(ctx context.Context, userHeaderRequest *user.UserHeaderRequest) (resp []*user.UserHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserHeaderInfo(ctx context.Context, userHeaderRequest *user.UserHeaderRequest) (resp *user.UserHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserPositionPage implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserPositionPage(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp *user_position.UserPositionPageResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) SaveUserPosition(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserPosition(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// DeleteUserPosition implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUserPosition(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
