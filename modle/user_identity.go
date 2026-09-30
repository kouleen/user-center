package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/user"
)

type UserIdentity struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	UserCode   string     `json:"userCode" gorm:"column:user_code;not null;default:'';uniqueIndex"`
	Username   string     `json:"username" gorm:"column:username;not null;default:'';"`
	Gender     *int8      `json:"gender" gorm:"column:gender;not null;default:1"`
	Phone      string     `json:"phone" gorm:"column:phone;default:''"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *UserIdentity) TableName() string {
	return "user_identity"
}

func (p *UserIdentity) ConvertResp() *user.UserIdentityResponse {
	return &user.UserIdentityResponse{
		Id:         p.ID,
		UserCode:   p.UserCode,
		Username:   p.Username,
		Gender:     p.Gender,
		Phone:      p.Phone,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
