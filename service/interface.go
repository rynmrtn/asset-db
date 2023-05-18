package service

import (
	"github.com/owasp-amass/asset-db/types"
)

type AssetService interface {
	Create(asset types.Asset, srcAsset *types.StoredAsset, relationType string) (types.StoredAsset, error)
	FindById(id int64) (types.StoredAsset, error)
	FindByContent(asset types.Asset) (types.StoredAsset, error)
}

type RelationService interface {
	Create(relationType string, newAssetId string, srcAssetId string) (types.StoredRelation, error)
}
