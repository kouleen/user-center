package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
	"github.com/kouleen/user-center/utils"
	"gorm.io/gorm"
)

func QueryPositionPage(ctx context.Context, req *user.UserPositionRequest) (list []modle.UserPosition, total int64, err error) {
	query := getPositionQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&list).Error; err != nil {
		return
	}
	return
}

func getPositionQuery(ctx context.Context, req *user.UserPositionRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.UserPosition{}).Where("is_delete = ?", 0)
	if req.UserId != nil {
		query = query.Where("user_id = ?", req.UserId)
	}
	if req.CountryCode != "" {
		query = query.Where("country_code = ?", req.CountryCode)
	}
	if req.Country != "" {
		query = query.Where("country like ?", req.Country+"%")
	}
	if req.Province != "" {
		query = query.Where("province like ?", req.Province+"%")
	}
	if req.City != "" {
		query = query.Where("city like ?", req.City+"%")
	}
	if req.District != "" {
		query = query.Where("district like ?", req.District+"%")
	}
	if req.Street != "" {
		query = query.Where("street like ?", req.Street+"%")
	}
	return query
}

func QueryPositionList(ctx context.Context, req *user.UserPositionRequest) (list []modle.UserPosition, err error) {
	query := getPositionQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	query.Order("create_time desc")
	if err = query.Find(&list).Error; err != nil {
		return
	}
	return
}
