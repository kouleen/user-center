package handle

import (
	"context"
	"errors"

	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/service"
)

func QueryUserHeaderPage(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderPageResponse, err error) {
	return service.QueryUserHeaderPage(ctx, req)
}

func QueryUserHeaderList(ctx context.Context, req *user.UserHeaderRequest) (resp []*user.UserHeaderResponse, err error) {
	return service.QueryUserHeaderList(ctx, req)
}

func QueryUserHeaderInfo(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderResponse, err error) {
	userId := req.GetId()
	if req.Id == nil {
		userId = ctxutil.GetUserId(ctx)
	}
	return service.QueryUserHeaderInfo(ctx, userId)
}

func QueryUserProfile(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderResponse, err error) {
	userId := req.GetId()
	if req.Id == nil {
		userId = ctxutil.GetUserId(ctx)
	}
	return service.QueryUserProfile(ctx, userId)
}

func UpdateUserProfile(ctx context.Context, req *user.UserHeaderRequest) (r bool, err error) {
	if req.GetNickname() == "" {
		return false, errors.New("nickname is empty")
	}
	if req.GetPhone() == "" {
		return false, errors.New("phone is empty")
	}
	if req.Gender == nil {
		return false, errors.New("gender is empty")
	}
	return service.UpdateUserProfile(ctx, req)
}

func UpdateUserPassword(ctx context.Context, req *user.UpdateUserPwd) (r bool, err error) {
	if req.GetOldPassword() == "" || req.GetNewPassword_() == "" {
		return false, errors.New("password is empty")
	}
	return service.UpdateUserPassword(ctx, req)
}

func UpdateUserHeaderStatus(ctx context.Context, req *user.UserHeaderRequest) (resp bool, err error) {
	if err = checkUpdateUserHeaderStatus(req); err != nil {
		return
	}
	return service.UpdateUserHeaderStatus(ctx, req)
}

func DeleteUserHeader(ctx context.Context, req *user.UserHeaderRequest) (resp bool, err error) {
	if req.GetUserIdList() == nil || len(req.GetUserIdList()) == 0 {
		return false, errors.New("invalid user id list")
	}
	return service.DeleteUserHeader(ctx, req)
}

func checkUpdateUserHeaderStatus(req *user.UserHeaderRequest) (err error) {
	if req.Id == nil {
		return errors.New("id is required")
	}
	if req.Status == nil {
		return errors.New("status is required")
	}
	return
}
