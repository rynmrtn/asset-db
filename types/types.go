package types

import model "github.com/owasp-amass/open-asset-model"

// type Asset interface {
// 	AssetType() model.AssetType
// 	JSON() ([]byte, error)
// }

type StoredAsset struct {
	ID    string
	Asset model.Asset
}

type StoredRelation struct {
	ID        string
	Type      string
	FromAsset StoredAsset
	ToAsset   StoredAsset
}
