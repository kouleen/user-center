package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/user-center/service"
)

func QueryPositionPage(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionPageResponse, err error) {
	return service.QueryPositionPage(ctx, req)
}

func QueryPositionList(ctx context.Context, req *user.UserPositionRequest) (resp []*user.UserPositionResponse, err error) {
	return service.QueryPositionList(ctx, req)
}

func QueryPosition(ctx context.Context, req *user.UserPositionRequest) (resp *user.UserPositionResponse, err error) {
	if req.Id == nil {
		return nil, errors.New("id is required")
	}
	return service.QueryPosition(ctx, req)
}

func SavePosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	if err = checkSavePosition(req); err != nil {
		return
	}
	return service.SavePosition(ctx, req)
}

func UpdatePosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	if err = checkUpdatePosition(req); err != nil {
		return
	}
	return service.UpdatePosition(ctx, req)
}

func DeletePosition(ctx context.Context, req *user.UserPositionRequest) (resp bool, err error) {
	if len(req.GetIdList()) == 0 {
		return false, errors.New("id is required")
	}
	return service.DeletePosition(ctx, req)
}

func checkSavePosition(req *user.UserPositionRequest) error {
	if req.UserId == nil {
		return errors.New("user_id is required")
	}
	if req.Latitude == nil {
		return errors.New("latitude is required")
	}
	if req.Longitude == nil {
		return errors.New("longitude is required")
	}
	if req.GetProvince() == "" {
		return errors.New("province is required")
	}
	if req.GetCity() == "" {
		return errors.New("city is required")
	}
	if req.GetDistrict() == "" {
		return errors.New("district is required")
	}
	if req.GetCountry() == "" {
		return errors.New("country is required")
	}
	if req.GetCountryCode() == "" {
		return errors.New("countryCode is required")
	}
	return nil
}

func checkUpdatePosition(req *user.UserPositionRequest) error {
	if req.Id == nil {
		return errors.New("user_id is required")
	}
	return checkSavePosition(req)
}
