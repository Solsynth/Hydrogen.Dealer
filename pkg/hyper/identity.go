package hyper

import (
	"errors"
	"fmt"
	"reflect"

	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type BaseUser struct {
	BaseModel

	Name         string `json:"name"`
	Nick         string `json:"nick"`
	Avatar       string `json:"avatar"`
	Banner       string `json:"banner"`
	Description  string `json:"description"`
	EmailAddress string `json:"email_address"`
	AffiliatedTo *uint  `json:"affiliated_to"`
	AutomatedBy  *uint  `json:"automated_by"`
}

// LinkAccount will help you build a BaseUser model from proto.UserInfo
// And also will help you to update the info in your database, so that this function requires a database context
func LinkAccount(tx *gorm.DB, table string, userinfo *proto.UserInfo) (BaseUser, error) {
	var account BaseUser
	if userinfo == nil {
		return account, fmt.Errorf("remote userinfo was not found")
	}
	if err := tx.
		Where("id = ?", userinfo.GetId()).
		Table(table).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			account = BaseUser{
				BaseModel: BaseModel{
					ID: uint(userinfo.GetId()),
				},
				Name:         userinfo.Name,
				Nick:         userinfo.Nick,
				Avatar:       userinfo.Avatar,
				Banner:       userinfo.Banner,
				Description:  userinfo.GetDescription(),
				EmailAddress: userinfo.Email,
			}
			if userinfo.AffiliatedTo != nil {
				account.AffiliatedTo = lo.ToPtr(uint(*userinfo.AffiliatedTo))
			}
			if userinfo.AutomatedBy != nil {
				account.AutomatedBy = lo.ToPtr(uint(*userinfo.AutomatedBy))
			}
			return account, tx.Table(table).Save(&account).Error
		}
		return account, err
	}

	prev := account
	account.Name = userinfo.Name
	account.Nick = userinfo.Nick
	account.Avatar = userinfo.Avatar
	account.Banner = userinfo.Banner
	account.Description = userinfo.GetDescription()
	account.EmailAddress = userinfo.Email
	if userinfo.AffiliatedTo != nil {
		account.AffiliatedTo = lo.ToPtr(uint(*userinfo.AffiliatedTo))
	}
	if userinfo.AutomatedBy != nil {
		account.AutomatedBy = lo.ToPtr(uint(*userinfo.AutomatedBy))
	}

	var err error
	if !reflect.DeepEqual(prev, account) {
		err = tx.Table(table).Save(&account).Error
	}

	return account, err
}
