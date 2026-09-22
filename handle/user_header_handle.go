package handle

import (
	"context"

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
