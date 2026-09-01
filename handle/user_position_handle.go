package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user_position"
	"github.com/kouleen/user-center/service"
)

func QueryPositionPage(ctx context.Context, req *user_position.UserPositionRequest) (resp *user_position.UserPositionPageResponse, err error) {
	return service.QueryPositionPage(ctx, req)
}
