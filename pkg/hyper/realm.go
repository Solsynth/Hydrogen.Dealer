package hyper

import (
	"errors"
	"fmt"
	"reflect"

	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BaseRealm struct {
	BaseModel

	Alias        string            `json:"alias"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Avatar       string            `json:"avatar"`
	Banner       string            `json:"banner"`
	IsPublic     bool              `json:"is_public"`
	IsCommunity  bool              `json:"is_community"`
	AccessPolicy datatypes.JSONMap `json:"access_policy"`
}

// LinkRealm will help you build a BaseRealm model from proto.RealmInfo
func LinkRealm(tx *gorm.DB, table string, info *proto.RealmInfo) (BaseRealm, error) {
	var realm BaseRealm
	if info == nil {
		return realm, fmt.Errorf("remote realm info was not found")
	}
	if err := tx.
		Where("id = ?", info.GetId()).
		Table(table).
		First(&realm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			realm = BaseRealm{
				Alias:        info.Alias,
				Name:         info.Name,
				Description:  info.Description,
				Avatar:       info.Avatar,
				Banner:       info.Banner,
				IsPublic:     info.IsPublic,
				IsCommunity:  info.IsCommunity,
				AccessPolicy: DecodeAccessPolicy(info.AccessPolicy),
			}

			return realm, tx.Table(table).Save(&realm).Error
		}
		return realm, err
	}

	prev := realm
	realm.Alias = info.Alias
	realm.Name = info.Name
	realm.Description = info.Description
	realm.Avatar = info.Avatar
	realm.Banner = info.Banner
	realm.IsPublic = info.IsPublic
	realm.IsCommunity = info.IsCommunity
	realm.AccessPolicy = DecodeAccessPolicy(info.AccessPolicy)

	var err error
	if !reflect.DeepEqual(prev, realm) {
		err = tx.Table(table).Save(&realm).Error
	}

	return realm, err
}

func EncodeAccessPolicy(policy map[string]any) []byte {
	raw, _ := jsoniter.Marshal(policy)
	return raw
}

func DecodeAccessPolicy(raw []byte) map[string]any {
	var policy map[string]any
	_ = jsoniter.Unmarshal(raw, &policy)
	return policy
}
