package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryIbeaconsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_ibeacons",
		Description: "Jamf Computers iBeacons",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetIbeacons{},
			transformers.WithPrimaryKeys("Name"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchIbeacons,
	}
}

func fetchIbeacons(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, ibeacon := range *computer.Ibeacons {
		res <- ibeacon
	}

	return nil
}
