package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/repository"
)

func QueryUserHeaderInfo(ctx context.Context, id int64) (resp *user.UserHeaderResponse, err error) {
	header, err := repository.GetUserHeader(ctx, id)

	if err != nil {
		return nil, err
	}
	return &user.UserHeaderResponse{
		Username:   header.Username,
		Password:   header.Password,
		Nickname:   header.Nickname,
		Gender:     int8(header.Gender),
		Avatar:     header.Avatar,
		Phone:      header.Phone,
		Status:     int8(header.Status),
		Id:         header.ID,
		IsDelete:   int8(header.IsDelete),
		CreatedBy:  header.CreatedBy,
		UpdatedBy:  header.UpdatedBy,
		CreateTime: header.CreateTime.UnixMilli(),
		UpdateTime: header.CreateTime.UnixMilli(),
	}, nil
}
