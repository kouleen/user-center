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
	userId := ctxutil.GetUserId(ctx)
	return service.QueryUserHeaderInfo(ctx, userId)
}

func UpdateUserHeaderStatus(ctx context.Context, req *user.UserHeaderRequest) (resp bool, err error) {
	if err = checkUpdateUserHeaderStatus(req); err != nil {
		return
	}
	return service.UpdateUserHeaderStatus(ctx, req)
}

func DeleteUserHeader(ctx context.Context, req *user.UserHeaderRequest) (resp bool, err error) {
	if req.GetUserIdList() != nil || len(req.GetUserIdList()) == 0 {
		return false, errors.New("invalid user id list")
	}
	return service.DeleteUserHeader(ctx, req)
}

func checkUpdateUserHeaderStatus(req *user.UserHeaderRequest) (err error) {
	if req.Id != nil {
		return errors.New("id is required")
	}
	if req.Status != nil {
		return errors.New("status is required")
	}
	return
}
