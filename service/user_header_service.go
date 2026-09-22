package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/repository"
)

func QueryUserHeaderPage(ctx context.Context, req *user.UserHeaderRequest) (resp *user.UserHeaderPageResponse, err error) {
	list, total, err := repository.QueryUserHeaderPage(ctx, req)
	if err != nil {
		return nil, err
	}
	respList := make([]*user.UserHeaderResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &user.UserHeaderPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryUserHeaderList(ctx context.Context, req *user.UserHeaderRequest) ([]*user.UserHeaderResponse, error) {
	list, err := repository.QueryUserHeaderList(ctx, req)
	if err != nil {
		return nil, err
	}
	respList := make([]*user.UserHeaderResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func QueryUserHeaderInfo(ctx context.Context, id int64) (resp *user.UserHeaderResponse, err error) {
	header, err := repository.GetUserHeaderById(ctx, id)
	if err != nil {
		return nil, err
	}
	return header.ConvertResp(), nil
}
