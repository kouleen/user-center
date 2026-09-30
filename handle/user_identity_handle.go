package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/service"
)

func QueryIdentityPage(ctx context.Context, req *user.UserIdentityRequest) (resp *user.UserIdentityPageResponse, err error) {
	return service.QueryIdentityPage(ctx, req)
}

func QueryIdentityList(ctx context.Context, req *user.UserIdentityRequest) (resp []*user.UserIdentityResponse, err error) {
	return service.QueryIdentityList(ctx, req)
}

func QueryIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp *user.UserIdentityResponse, err error) {
	return service.QueryIdentity(ctx, req)
}

func SaveIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp bool, err error) {
	return service.SaveIdentity(ctx, req)
}

func UpdateIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp bool, err error) {
	return service.UpdateIdentity(ctx, req)
}

func DeleteIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp bool, err error) {
	return service.DeleteIdentity(ctx, req)
}
