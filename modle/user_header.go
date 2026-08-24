package modle

import "time"

type UserHeader struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Username   string     `json:"username" gorm:"column:username;not null;default:'';uniqueIndex"`
	Password   string     `json:"password" gorm:"column:password;not null;default:''"`
	Nickname   string     `json:"nickname" gorm:"column:nickname;not null;default:''"`
	Gender     uint8      `json:"gender" gorm:"column:gender;not null;default:1"`
	Avatar     string     `json:"avatar" gorm:"column:avatar;default:''"`
	Phone      string     `json:"phone" gorm:"column:phone;default:''"`
	Status     uint8      `json:"status" gorm:"column:status;not null;default:1"`
	IsDelete   uint8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (UserHeader) TableName() string {
	return "user_header"
}
