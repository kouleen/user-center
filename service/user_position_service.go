package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/repository"
)

func QueryPositionPage(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionPageResponse, err error) {
	list, total, err := repository.QueryPositionPage(ctx, req)
	if err != nil {
		return
	}
	userPositionList := make([]*user.UserPositionResponse, len(list))
	for i, item := range list {
		userPositionList[i] = item.ConvertResp()
	}
	return &user.UserPositionPageResponse{
		Total:   total,
		Records: userPositionList,
	}, nil
}

func QueryPositionList(ctx context.Context, req *user.UserPositionRequest) (resp []*user.UserPositionResponse, err error) {
	list, err := repository.QueryPositionList(ctx, req)
	if err != nil {
		return
	}
	userPositionList := make([]*user.UserPositionResponse, len(list))
	for i, item := range list {
		userPositionList[i] = item.ConvertResp()
	}
	return userPositionList, nil
}

func QueryPosition(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionResponse, err error) {
	return
}

func SaveUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return
}

func UpdateUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return
}

func DeleteUserPosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	return
}
