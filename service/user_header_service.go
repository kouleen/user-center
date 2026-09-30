package service

import (
	"context"
	"strconv"

	"github.com/kouleen/common/pkg/ctxutil"
	rediscli "github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/common"
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

func UpdateUserHeaderStatus(ctx context.Context, req *user.UserHeaderRequest) (resp bool, err error) {
	userHeader, err := repository.GetUserHeaderById(ctx, req.GetId())
	if err != nil {
		return false, err
	}
	userHeader.Status = req.Status
	userHeader.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateUserHeader(ctx, userHeader); err != nil {
		return false, err
	}
	// 如果是改成禁用 注销用户登录
	if common.BaseStatus_DISABLED == common.BaseStatus(req.GetStatus()) {
		userIdStr := strconv.FormatInt(userHeader.ID, 10)
		token, err := rediscli.Get(ctx, userIdStr)
		if _, err = rediscli.Unlink(ctx, token); err != nil {
			return false, err
		}
		if _, err = rediscli.Unlink(ctx, userIdStr); err != nil {
			return false, err
		}
	}
	return true, nil
}

func DeleteUserHeader(ctx context.Context, req *user.UserHeaderRequest) (resp bool, err error) {
	respList, err := repository.GetUserHeaderByIdList(ctx, req.GetUserIdList())
	if err != nil {
		return false, err
	}
	if len(respList) == 0 {
		return false, nil
	}
	isDelete := int8(1)
	for _, header := range respList {
		header.IsDelete = &isDelete
		header.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdateUserHeader(ctx, respList); err != nil {
		return false, err
	}
	return true, nil
}
