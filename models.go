package assetdb

import (
	"encoding/json"
	"fmt"
	"time"

	oam "github.com/owasp-amass/open-asset-model"
	"github.com/owasp-amass/open-asset-model/domain"
	"github.com/owasp-amass/open-asset-model/network"

	"gorm.io/datatypes"
)

type Asset struct {
	ID        int64     `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"type:datetime"`
	Type      string
	Content   datatypes.JSON
}

// TODO: not thrilled with this .. doesn't scale well when new types
// are to open-asset-oam
func (a Asset) Parse() (oam.Asset, error) {
	var asset oam.Asset
	// TODO: should transform a.Type to AssetType ?
	switch a.Type {
	case string(oam.FQDN):
		var fqdn domain.FQDN
		err := json.Unmarshal(a.Content, &fqdn)
		if err != nil {
			return domain.FQDN{}, err
		}
		asset = fqdn
	case string(oam.IPAddress):
		var ip network.IPAddress
		err := json.Unmarshal(a.Content, &ip)
		if err != nil {
			return network.IPAddress{}, err
		}
		asset = ip
	case string(oam.ASN):
		var asn network.AutonomousSystem
		err := json.Unmarshal(a.Content, &asn)
		if err != nil {
			return network.AutonomousSystem{}, err
		}
		asset = asn
	case string(oam.RIROrg):
		var rir network.RIROrganization
		err := json.Unmarshal(a.Content, &rir)
		if err != nil {
			return network.RIROrganization{}, err
		}
		asset = rir
	case string(oam.Netblock):
		var netblock network.Netblock
		err := json.Unmarshal(a.Content, &netblock)
		if err != nil {
			return network.Netblock{}, err
		}
		asset = netblock
	default:
		return nil, fmt.Errorf("unknown asset type: %s", a.Type)
	}

	return asset, nil
}

func (a Asset) JSONQuery() (*datatypes.JSONQueryExpression, error) {
	switch a.Type {
	case string(oam.FQDN):
		asset, err := a.Parse()
		if err != nil {
			return nil, err
		}
		assetData := asset.(domain.FQDN)
		return datatypes.JSONQuery("content").Equals(assetData.Name, "name").Equals(assetData.Name, "name2"), nil
	default:
		return nil, fmt.Errorf("unknown asset type: %s", a.Type)
	}
}

type Relation struct {
	ID          int64     `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt   time.Time `gorm:"type:datetime"`
	Type        string
	FromAssetID int64
	ToAssetID   int64
	FromAsset   Asset
	ToAsset     Asset
}
