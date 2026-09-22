package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/user"
)

type UserHeader struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Username   string     `json:"username" gorm:"column:username;not null;default:'';uniqueIndex"`
	Password   string     `json:"password" gorm:"column:password;not null;default:''"`
	Nickname   string     `json:"nickname" gorm:"column:nickname;not null;default:''"`
	Gender     *int8      `json:"gender" gorm:"column:gender;not null;default:1"`
	Avatar     string     `json:"avatar" gorm:"column:avatar;default:''"`
	Phone      string     `json:"phone" gorm:"column:phone;default:''"`
	Status     *int8      `json:"status" gorm:"column:status;not null;default:1"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *UserHeader) GetStatus() (status int8) {
	if p == nil {
		return
	}
	return *p.Status
}

func (p *UserHeader) TableName() string {
	return "user_header"
}

func (p *UserHeader) ConvertResp() *user.UserHeaderResponse {
	return &user.UserHeaderResponse{
		Id:         p.ID,
		Username:   p.Username,
		Password:   p.Password,
		Nickname:   p.Nickname,
		Gender:     p.Gender,
		Avatar:     p.Avatar,
		Phone:      p.Phone,
		Status:     p.Status,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
