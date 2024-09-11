package hyper

import (
	"errors"
	"fmt"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/database"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"reflect"
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

func LinkAccount(model any, userinfo *proto.UserInfo) (BaseUser, error) {
	var account BaseUser
	if userinfo == nil {
		return account, fmt.Errorf("remote userinfo was not found")
	}
	if err := database.C.Where("id = ?", userinfo.GetId()).Model(model).First(&account).Error; err != nil {
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
			return account, database.C.Model(model).Save(&account).Error
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
		err = database.C.Model(model).Save(&account).Error
	}

	return account, err
}
