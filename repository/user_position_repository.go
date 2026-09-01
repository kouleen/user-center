package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/user_position"
	"github.com/kouleen/user-center/modle"
)

func QueryPositionPage(ctx context.Context, req *user_position.UserPositionRequest) (list []modle.UserPosition, total int64, err error) {
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
