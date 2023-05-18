package service

import (
	"strconv"

	"github.com/owasp-amass/asset-db/repository"
	"github.com/owasp-amass/asset-db/types"
)

// assetService is a struct that implements the AssetService interface.
// It has a dependency on the AssetRepository and RelationService interfaces
// and it's the layer that handles the business logic of the amass application
type assetService struct {
	assetRepository repository.AssetRepository
	relationService RelationService
}

// NewAssetService is a function that returns a new assetService
func NewAssetService(assetRepository repository.AssetRepository, relationService RelationService) *assetService {
	return &assetService{
		assetRepository: assetRepository,
		relationService: relationService,
	}
}

// Create is a method that creates a new asset in the database.
// It receives an asset, a source asset and a relation type,
// it creates the asset in the database, creates a relation between the two assets
// and returns a new asset and an error if it exists
func (as *assetService) Create(asset types.Asset, srcAsset *types.StoredAsset, relationType *string) (types.StoredAsset, error) {
	if srcAsset == nil || relationType == nil {
		return as.assetRepository.Create(asset)
	}

	newAsset, err := as.assetRepository.Create(asset)
	if err != nil {
		return types.StoredAsset{}, err
	}

	_, err = as.relationService.Create(*relationType, newAsset.ID, srcAsset.ID)
	if err != nil {
		return types.StoredAsset{}, err
	}

	return newAsset, nil
}

func (as *assetService) Exist(asset types.Asset) (bool, error) {
	storedAsset, err := as.FindByContent(asset)
	if err != nil {
		return false, err
	}

	return storedAsset.ID != "", nil
}

func (as *assetService) FindByContent(asset types.Asset) (types.StoredAsset, error) {
	return as.assetRepository.FindByContent(asset)
}

func (as *assetService) FindById(id string) (types.StoredAsset, error) {
	assetId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return types.StoredAsset{}, err
	}

	return as.assetRepository.FindById(assetId)
}
