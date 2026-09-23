package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
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
	item, err := repository.QueryPositionById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), nil
}

func SavePosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	entity := &modle.UserPosition{
		ID:          id,
		UserID:      req.GetUserId(),
		Latitude:    req.GetLatitude(),
		Longitude:   req.GetLongitude(),
		Province:    req.GetProvince(),
		City:        req.GetCity(),
		District:    req.GetDistrict(),
		Street:      req.GetStreet(),
		FullAddress: req.GetFullAddress(),
		Country:     req.GetCountry(),
		CountryCode: req.GetCountryCode(),
		CreatedBy:   ctxutil.GetUserId(ctx),
	}
	if err = repository.CreatePosition(ctx, entity); err != nil {
		return
	}
	return true, nil
}

func UpdatePosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	item, err := repository.QueryPositionById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.UserID = req.GetUserId()
	item.Latitude = req.GetLatitude()
	item.Longitude = req.GetLongitude()
	item.Province = req.GetProvince()
	item.City = req.GetCity()
	item.District = req.GetDistrict()
	item.Street = req.GetStreet()
	item.FullAddress = req.GetFullAddress()
	item.Country = req.GetCountry()
	item.CountryCode = req.GetCountryCode()
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdatePosition(ctx, item); err != nil {
		return
	}
	return true, nil
}

func DeletePosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	itemList, err := repository.QueryPositionByIdList(ctx, req.GetIdList())
	if err != nil {
		return
	}
	isDelete := int8(1)
	for _, position := range itemList {
		position.IsDelete = &isDelete
		position.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdatePosition(ctx, itemList); err != nil {
		return
	}
	return true, nil
}
