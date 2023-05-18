package assetdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	model "github.com/owasp-amass/open-asset-model"
)

const (
	Postgres DBType = "postgres"
	SQLite   DBType = "sqlite"
)

type DBConfig struct {
	DBType DBType
	DSN    string
}

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(config *DBConfig) (*AssetRepository, error) {
	gdb, err := newDb(config)
	if err != nil {
		return nil, err
	}

	return &AssetRepository{
		db: gdb,
	}, nil

}

func newDb(config *DBConfig) (*gorm.DB, error) {
	switch config.DBType {
	case Postgres:
		return postgresDatabase(config.DSN)
	case SQLite:
		return sqliteDatabase(config.DSN)
	default:
		panic("Unknown db type")
	}
}

func postgresDatabase(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}

func sqliteDatabase(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}

func (ar *AssetRepository) CreateAsset(asset model.Asset) (*Asset, error) {
	jc, err := asset.JSON()
	if err != nil {
		return nil, err
	}

	as := &Asset{
		Type:    string(asset.AssetType()),
		Content: jc,
	}

	result := ar.db.Create(as)
	if result.Error != nil {
		return nil, result.Error
	}

	return as, nil
}

func (ar *AssetRepository) FindById(id int64) (*Asset, error) {
	asset := &Asset{ID: id}
	result := ar.db.First(asset)
	if result.Error != nil {
		return nil, result.Error
	}

	return asset, nil
}

func (ar *AssetRepository) FindByContent(oam model.Asset) (*Asset, error) {
	jsonContent, err := oam.JSON()
	if err != nil {
		return nil, err
	}

	asset := &Asset{
		Type:    string(oam.AssetType()),
		Content: jsonContent,
	}

	jsonQuery, err := asset.JSONQuery()
	if err != nil {
		return nil, err
	}

	result := ar.db.First(&asset, jsonQuery)
	if result.Error != nil {
		return nil, result.Error
	}

	return asset, nil
}

func (ar *AssetRepository) CreateRelation(from Asset, relationType string, to Asset) (*Relation, error) {
	relation := Relation{
		Type:        relationType,
		FromAssetID: from.ID,
		ToAssetID:   to.ID,
	}

	result := ar.db.Create(&relation)
	if result.Error != nil {
		return nil, result.Error
	}

	return &relation, nil
}
