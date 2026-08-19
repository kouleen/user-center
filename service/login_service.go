package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kouleen/user-center/kitex_gen/user"
	"github.com/kouleen/user-center/utils"
)

func Login(ctx context.Context, loginRequest *user.LoginRequest) (resp *user.LoginResponse, err error) {
	newUUID, _ := uuid.NewUUID()
	if err = utils.Set(ctx, "auth:token:"+newUUID.String(), newUUID.String(), 12*time.Hour); err != nil {
		return nil, err
	}
	v := int64(7200)
	return &user.LoginResponse{
		AccessToken: newUUID.String(),
		ExpireTime:  &v,
	}, nil
}
