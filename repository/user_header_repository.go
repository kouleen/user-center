package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/modle"
	"gorm.io/gorm"
)

func QueryUserHeaderPage(ctx context.Context, req *user.UserHeaderRequest) (resp []*modle.UserHeader, total int64, err error) {
	query := getUserHeaderQuery(ctx, req)
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

func getUserHeaderQuery(ctx context.Context, req *user.UserHeaderRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.UserHeader{}).Where("is_deleted = 0")
	if req.Status != nil {
		query = query.Where("status = ?", req.GetStatus())
	}
	if req.Gender != nil {
		query = query.Where("gender = ?", req.GetGender())
	}
	if req.GetUsername() != "" {
		query = query.Where("username like ?", req.GetUsername()+"%")
	}
	if req.GetNickname() != "" {
		query = query.Where("nickname like ?", "%"+req.GetNickname()+"%")
	}
	return query
}

func QueryUserHeaderList(ctx context.Context, req *user.UserHeaderRequest) (resp []*modle.UserHeader, err error) {
	query := getUserHeaderQuery(ctx, req)
	query = query.Order("create_time desc")
	if err = query.Find(&resp).Error; err != nil {
		return
	}
	return
}

func GetUserHeaderById(ctx context.Context, id int64) (resp *modle.UserHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&resp, id).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func GetUserHeaderByUsername(ctx context.Context, username string) (resp *modle.UserHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.UserHeader{}).Where("is_delete = 0 and username = ?", username).First(&resp).Error; err != nil {
		return
	}
	return resp, nil
}

func GetUserHeaderByPhone(ctx context.Context, phone string) (resp *modle.UserHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.UserHeader{}).Where("is_delete = 0 and phone = ?", phone).First(&resp).Error; err != nil {
		return
	}
	return resp, nil
}

func CreateUserHeader(ctx context.Context, userHeader *modle.UserHeader) error {
	if err := mysql.GetWriteMysqlDDB().WithContext(ctx).Create(userHeader).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePassword(ctx context.Context, userHeader *modle.UserHeader) error {
	if err := mysql.GetWriteMysqlDDB().WithContext(ctx).Model(userHeader).Updates(userHeader).Error; err != nil {
		return err
	}
	return nil
}
