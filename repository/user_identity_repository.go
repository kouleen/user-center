package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
	"github.com/kouleen/user-center/utils"
	"gorm.io/gorm"
)

func QueryIdentityPage(ctx context.Context, req *user.UserIdentityRequest) (resp []*modle.UserIdentity, total int64, err error) {
	query := getUserIdentityQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func getUserIdentityQuery(ctx context.Context, req *user.UserIdentityRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Where("is_deleted = 0")
	if req.Gender != nil {
		query = query.Where("gender = ?", req.Gender)
	}
	if req.GetUserCode() != "" {
		query = query.Where("user_code like ?", req.GetUserCode()+"%")
	}
	if req.GetUsername() != "" {
		query = query.Where("username like ?", "%"+req.GetUsername()+"%")
	}
	return query
}

func QueryIdentityList(ctx context.Context, req *user.UserIdentityRequest) (resp []*modle.UserIdentity, err error) {
	query := getUserIdentityQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	query = query.Order("create_time desc")
	if err = query.Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryIdentityById(ctx context.Context, id int64) (resp *modle.UserIdentity, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&resp, id).Error; err != nil {
		return
	}
	return
}

func CreateIdentity(ctx context.Context, entity *modle.UserIdentity) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error; err != nil {
		return
	}
	return
}

func UpdateIdentity(ctx context.Context, entity *modle.UserIdentity) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Save(entity).Error; err != nil {
		return
	}
	return
}

func BatchDeleteIdentity(ctx context.Context, idList []int64) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(idList) > 0 {
			if err = tx.Where("id in (?)", idList).Delete(&modle.UserIdentity{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
