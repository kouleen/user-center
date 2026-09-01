package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/user_position"
	"github.com/kouleen/user-center/repository"
)

func QueryPositionPage(ctx context.Context, req *user_position.UserPositionRequest) (resp *user_position.UserPositionPageResponse, err error) {
	list, total, err := repository.QueryPositionPage(ctx, req)
	if err != nil {
		return
	}
	userPositionList := make([]*user_position.UserPositionResponse, len(list))
	for i, item := range list {
		var updateTime int64
		if item.UpdateTime != nil {
			updateTime = item.UpdateTime.UnixMilli()
		}
		userPositionList[i] = &user_position.UserPositionResponse{
			Id:          item.ID,
			UserId:      item.UserID,
			Longitude:   item.Longitude,
			Latitude:    item.Latitude,
			Province:    item.Province,
			City:        item.City,
			District:    item.District,
			Street:      item.Street,
			FullAddress: item.FullAddress,
			Country:     item.Country,
			CountryCode: item.CountryCode,
			IsDelete:    int8(item.IsDelete),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   updateTime,
			CreateTime:  item.CreateTime.UnixMilli(),
			UpdateTime:  updateTime,
		}
	}
	return &user_position.UserPositionPageResponse{
		Total:   total,
		Records: userPositionList,
	}, nil
}
