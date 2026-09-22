package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/service"
)

func QueryPositionPage(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionPageResponse, err error) {
	return service.QueryPositionPage(ctx, req)
}

func QueryPositionList(ctx context.Context, req *user.UserPositionRequest) (resp []*user.UserPositionResponse, err error) {
	return service.QueryPositionList(ctx, req)
}

func QueryPosition(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionResponse, err error) {
	return service.QueryPosition(ctx, req)
}

func SaveUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return service.SaveUserPosition(ctx, req)
}

func UpdateUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return service.UpdateUserPosition(ctx, req)
}

func DeleteUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return service.DeleteUserPosition(ctx, req)
}
