package repository

import (
	"context"

	"github.com/kouleen/user-center/modle"
	"github.com/kouleen/user-center/utils"
)

func GetUserHeaderByUsername(ctx context.Context, username string) (resp *modle.UserHeader, err error) {
	if err = utils.GetReadMysqlDDB().WithContext(ctx).Model(&modle.UserHeader{}).Where("is_delete = 0 and username = ?", username).First(&resp).Error; err != nil {
		return
	}
	return resp, nil
}

func GetUserHeaderByPhone(ctx context.Context, phone string) (resp *modle.UserHeader, err error) {
	if err = utils.GetReadMysqlDDB().WithContext(ctx).Model(&modle.UserHeader{}).Where("is_delete = 0 and phone = ?", phone).First(&resp).Error; err != nil {
		return
	}
	return resp, nil
}

func CreateUserHeader(ctx context.Context, userHeader *modle.UserHeader) error {
	if err := utils.GetWriteMysqlDDB().WithContext(ctx).Create(userHeader).Error; err != nil {
		return err
	}
	return nil
}
