package main

import (
	"fmt"

	"github.com/owasp-amass/asset-db/open_asset_model"
	"github.com/owasp-amass/asset-db/repository/gorm"
	"github.com/owasp-amass/asset-db/service"
)

func main() {
	// Example of asset-db usage

	// initialize database
	database, err := gorm.NewDatabase("sqlite", "amassdb.sqlite3")
	if err != nil {
		fmt.Println(err)
		return
	}

	// initialize repositories
	assetRepository := gorm.NewAssetRepository(database)
	relationRepository := gorm.NewRelationRepository(database)

	// initialize services
	relationService := service.NewRelationService(relationRepository)
	assetService := service.NewAssetService(assetRepository, relationService)

	// discover new asset
	newAsset := open_asset_model.FQDN{Name: "domain.com"}

	// create asset into database
	firstAsset, err := assetService.Create(newAsset, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("asset Createed:", firstAsset)

	// discover new asset and its relation to the previous asset
	newAsset2 := open_asset_model.FQDN{Name: "domain.subdomain.com"}

	// Create asset into database and create relation to the previous asset
	relationType := "subdomain"
	secondAsset, err := assetService.Create(newAsset2, &firstAsset, &relationType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("asset Createed:", secondAsset)

	// get asset by id from database
	storedAsset, err := assetService.FindById(firstAsset.ID)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("asset retrieved:", storedAsset)
}
