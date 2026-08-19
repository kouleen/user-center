package main

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/kouleen/user-center/kitex_gen/user"
	"github.com/kouleen/user-center/kitex_gen/user_position"
	"github.com/kouleen/user-center/service"
	"github.com/kouleen/user-center/utils"
)

// UserHeaderServiceImpl implements the last service interface defined in the IDL.
type UserHeaderServiceImpl struct{}

// Login implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) Login(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error) {
	traceId, _ := metainfo.GetPersistentValue(ctx, "x-trace-id")
	logger.CtxInfof(ctx, "[%s]-Login request loginRequest: %#v", traceId, loginRequest)
	return service.Login(ctx, loginRequest)

}

// SendSmsMessage implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) SendSmsMessage(ctx context.Context, loginRequest *user.LoginRequest) (resp bool, err error) {
	return
}

// Logout implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) Logout(ctx context.Context, loginRequest *user.LoginRequest) (resp bool, err error) {
	traceId, _ := metainfo.GetPersistentValue(ctx, "x-trace-id")
	logger.CtxInfof(ctx, "Logout request traceId: %s", traceId)
	token, _ := metainfo.GetPersistentValue(ctx, "token")
	if err = utils.Del(ctx, "auth:token:"+token); err != nil {
		return false, err
	}
	return true, nil
}

// QueryUserHeaderPage implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) QueryUserHeaderPage(ctx context.Context, userHeaderRequest *user.UserHeaderRequest) (resp *user.UserHeaderPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderList implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) QueryUserHeaderList(ctx context.Context, userHeaderRequest *user.UserHeaderRequest) (resp []*user.UserHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserHeaderInfo implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) QueryUserHeaderInfo(ctx context.Context, userHeaderRequest *user.UserHeaderRequest) (resp *user.UserHeaderResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserPositionPage implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) QueryUserPositionPage(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp *user_position.UserPositionPageResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveUserPosition implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) SaveUserPosition(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateUserPosition implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) UpdateUserPosition(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// DeleteUserPosition implements the UserHeaderServiceImpl interface.
func (s *UserHeaderServiceImpl) DeleteUserPosition(ctx context.Context, userPositionRequest *user_position.UserPositionRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
