package modle

import "time"

type UserPosition struct {
	ID          int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	UserID      int64      `json:"user_id,string" gorm:"column:user_id;not null"`      // 用户id
	Longitude   float64    `json:"longitude" gorm:"column:longitude;not null"`         // 经度
	Latitude    float64    `json:"latitude" gorm:"column:latitude;default:''"`         // 纬度
	Province    string     `json:"province" gorm:"column:province;default:''"`         // 省/州
	City        string     `json:"city" gorm:"column:city;default:''"`                 // 城市
	District    string     `json:"district" gorm:"column:district;default:''"`         // 区县
	Street      string     `json:"street" gorm:"column:street;default:''"`             // 街道
	FullAddress string     `json:"full_address" gorm:"column:full_address;default:''"` // 完整地址
	Country     string     `json:"country" gorm:"column:country;default:''"`           // 国家
	CountryCode string     `json:"country_code" gorm:"column:country_code;default:''"` // 国家代码
	IsDelete    uint8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy   int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy   int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime  *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime  *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (UserPosition) TableName() string {
	return "user_position"
}
