package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/user-center/modle"
)

func GetUserHeader(ctx context.Context, id int64) (resp *modle.UserHeader, err error) {
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
