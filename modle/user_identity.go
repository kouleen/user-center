package modle

import (
	"regexp"
	"time"

	"github.com/kouleen/idl/kitex_gen/user"
)

type UserIdentity struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	UserCode   string     `json:"userCode" gorm:"column:user_code;not null;default:'';uniqueIndex"`
	Username   string     `json:"username" gorm:"column:user_name;not null;default:'';"`
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
	idCard := desensitizeIDCard(p.UserCode)
	name := desensitizeName(p.Username)
	return &user.UserIdentityResponse{
		Id:         p.ID,
		UserCode:   idCard,
		Username:   name,
		Gender:     p.Gender,
		Phone:      p.Phone,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}

// desensitizeIDCard 身份证脱敏
// 18位: 前6 + 8* + 后4
// 15位: 前6 + 6* + 后3
func desensitizeIDCard(id string) string {
	// 去除空格
	id = regexp.MustCompile(`\s+`).ReplaceAllString(id, "")
	l := len(id)
	switch l {
	case 18:
		return id[:6] + "********" + id[14:]
	case 15:
		return id[:6] + "******" + id[12:]
	default:
		// 长度非法，直接返回原值，也可以返回空
		return id
	}
}

// desensitizeName 姓名脱敏
// 规则：
// 2字：李*
// 3字：张*三
// 4字：王**华
func desensitizeName(name string) string {
	name = trimSpace(name)
	runes := []rune(name)
	l := len(runes)
	if l == 0 {
		return ""
	}
	switch l {
	case 1:
		// 单字姓名，直接返回原字 如：赵
		return name
	case 2:
		return string(runes[0]) + "*"
	default:
		// 首字 + (l-2)个* + 尾字
		mid := make([]rune, l-2)
		for i := range mid {
			mid[i] = '*'
		}
		return string(runes[0]) + string(mid) + string(runes[l-1])
	}
}

func trimSpace(s string) string {
	runes := []rune(s)
	start := 0
	for ; start < len(runes); start++ {
		if runes[start] != ' ' {
			break
		}
	}
	end := len(runes)
	for ; end > start; end-- {
		if runes[end-1] != ' ' {
			break
		}
	}
	return string(runes[start:end])
}
