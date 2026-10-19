package service

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/kouleen/common/client"
	"github.com/kouleen/common/pkg/ctxutil"
	rediscli "github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/repository"
	"golang.org/x/crypto/bcrypt"
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

func QueryUserProfile(ctx context.Context, id int64) (resp *user.UserHeaderResponse, err error) {
	header, err := repository.GetUserHeaderById(ctx, id)
	if err != nil {
		return nil, err
	}
	headerResponse := header.ConvertResp()
	req := &system.SystemRoleRequest{
		UserId: &id,
	}
	systemRoleList, err := client.GetSystemRpc().QuerySystemRoleList(ctx, req)
	if err != nil {
		return nil, err
	}
	roleNameArr := make([]string, len(systemRoleList))
	for index, response := range systemRoleList {
		roleNameArr[index] = response.RoleName
	}
	roleNames := strings.Join(roleNameArr, ",")
	headerResponse.RoleNames = roleNames
	return headerResponse, nil
}

func UpdateUserProfile(ctx context.Context, req *user.UserHeaderRequest) (r bool, err error) {
	header, err := repository.GetUserHeaderById(ctx, ctxutil.GetUserId(ctx))
	if err != nil {
		return false, err
	}
	header.Nickname = req.GetNickname()
	header.Phone = req.GetPhone()
	header.Gender = req.Gender
	if err = repository.UpdateUserHeader(ctx, header); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateUserAvatar(ctx context.Context, req *user.UserHeaderRequest) (r bool, err error) {
	header, err := repository.GetUserHeaderById(ctx, ctxutil.GetUserId(ctx))
	if err != nil {
		return false, err
	}
	header.Avatar = req.GetAvatar()
	if err = repository.UpdateUserHeader(ctx, header); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateUserPassword(ctx context.Context, req *user.UpdateUserPwd) (r bool, err error) {
	userHeader, err := repository.GetUserHeaderById(ctx, ctxutil.GetUserId(ctx))
	if err != nil {
		return false, err
	}
	// 对比密码
	if err = bcrypt.CompareHashAndPassword([]byte(userHeader.Password), []byte(req.OldPassword)); err != nil {
		return
	}
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword_), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	userHeader.Password = string(hashPwd)
	if err = repository.UpdateUserHeader(ctx, userHeader); err != nil {
		return false, err
	}
	// 退出登录
	userIdStr := strconv.FormatInt(userHeader.ID, 10)
	token, err := rediscli.Get(ctx, userIdStr)
	if _, err = rediscli.Unlink(ctx, token); err != nil {
		return false, err
	}
	if _, err = rediscli.Unlink(ctx, userIdStr); err != nil {
		return false, err
	}
	return true, nil
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
	for _, header := range respList {
		if common.BaseStatus_ENABLE == common.BaseStatus(header.GetStatus()) {
			return false, errors.New("user status is enabled")
		}
	}
	if err = repository.BatchDeleteUserHeader(ctx, req.GetUserIdList()); err != nil {
		return false, err
	}
	return true, nil
}
