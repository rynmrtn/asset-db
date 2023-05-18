package repository

import "github.com/owasp-amass/asset-db/types"

type AssetRepository interface {
	Create(asset types.Asset) (types.StoredAsset, error)
	FindById(id int64) (types.StoredAsset, error)
	FindByContent(asset types.Asset) (types.StoredAsset, error)
}

type RelationRepository interface {
	Create(relationType string, newAssetId string, srcAssetId string) (types.StoredRelation, error)
}
