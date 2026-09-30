package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
	"github.com/kouleen/user-center/repository"
)

func QueryIdentityPage(ctx context.Context, req *user.UserIdentityRequest) (resp *user.UserIdentityPageResponse, err error) {
	list, total, err := repository.QueryIdentityPage(ctx, req)
	if err != nil {
		return nil, err
	}
	respList := make([]*user.UserIdentityResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &user.UserIdentityPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryIdentityList(ctx context.Context, req *user.UserIdentityRequest) (resp []*user.UserIdentityResponse, err error) {
	list, err := repository.QueryIdentityList(ctx, req)
	if err != nil {
		return nil, err
	}
	respList := make([]*user.UserIdentityResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func QueryIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp *user.UserIdentityResponse, err error) {
	item, err := repository.QueryIdentityById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), nil
}

func SaveIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	entity := &modle.UserIdentity{
		ID:        id,
		UserCode:  req.GetUserCode(),
		Username:  req.GetUsername(),
		Gender:    req.Gender,
		Phone:     req.GetPhone(),
		Remark:    req.GetRemark(),
		IsDelete:  req.IsDelete,
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateIdentity(ctx, entity); err != nil {
		return
	}
	return true, nil
}

func UpdateIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp bool, err error) {
	item, err := repository.QueryIdentityById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.Username = req.GetUsername()
	item.Gender = req.Gender
	item.Phone = req.GetPhone()
	item.Remark = req.GetRemark()
	item.IsDelete = req.IsDelete
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateIdentity(ctx, item); err != nil {
		return
	}
	return true, nil
}

func DeleteIdentity(ctx context.Context, req *user.UserIdentityRequest) (resp bool, err error) {
	if err = repository.BatchDeleteIdentity(ctx, req.GetIdList()); err != nil {
		return
	}
	return true, nil
}
