package hyper

import "git.solsynth.dev/hydrogen/dealer/pkg/proto"

type BaseRealm struct {
	BaseModel

	Alias       string `json:"alias"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Banner      string `json:"banner"`
	IsPublic    bool   `json:"is_public"`
	IsCommunity bool   `json:"is_community"`
}

// LinkRealm will help you build a BaseRealm model from proto.RealmInfo
// WARNING This function doesn't like the LinkAccount, it will not help you deal the database stuff
func LinkRealm(info *proto.RealmInfo) BaseRealm {
	return BaseRealm{
		Alias:       info.Alias,
		Name:        info.Name,
		Description: info.Description,
		Avatar:      info.Avatar,
		Banner:      info.Banner,
		IsPublic:    info.IsPublic,
		IsCommunity: info.IsCommunity,
	}
}
