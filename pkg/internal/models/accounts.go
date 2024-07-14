package models

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/viper"
	"gorm.io/datatypes"
)

type Account struct {
	BaseModel

	Name        string            `json:"name" gorm:"uniqueIndex"`
	Nick        string            `json:"nick"`
	Description string            `json:"description"`
	Avatar      *uint             `json:"avatar"`
	Banner      *uint             `json:"banner"`
	ConfirmedAt *time.Time        `json:"confirmed_at"`
	SuspendedAt *time.Time        `json:"suspended_at"`
	PermNodes   datatypes.JSONMap `json:"perm_nodes"`

	Contacts []AccountContact `json:"contacts"`
}

func (v Account) GetAvatar() *string {
	if v.Avatar != nil {
		return lo.ToPtr(fmt.Sprintf("%s/api/attachments/%d", viper.GetString("content_endpoint"), *v.Avatar))
	}
	return nil
}

func (v Account) GetBanner() *string {
	if v.Banner != nil {
		return lo.ToPtr(fmt.Sprintf("%s/api/attachments/%d", viper.GetString("content_endpoint"), *v.Banner))
	}
	return nil
}

func (v Account) GetPrimaryEmail() AccountContact {
	val, _ := lo.Find(v.Contacts, func(item AccountContact) bool {
		return item.Type == EmailAccountContact && item.IsPrimary
	})
	return val
}

type AccountContactType = int8

const (
	EmailAccountContact = AccountContactType(iota)
)

type AccountContact struct {
	BaseModel

	Type       int8       `json:"type"`
	Content    string     `json:"content" gorm:"uniqueIndex"`
	IsPublic   bool       `json:"is_public"`
	IsPrimary  bool       `json:"is_primary"`
	VerifiedAt *time.Time `json:"verified_at"`
	AccountID  uint       `json:"account_id"`
}
