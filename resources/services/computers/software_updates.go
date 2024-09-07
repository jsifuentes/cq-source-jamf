package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventorySoftwareUpdatesTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_software_updates",
		Description: "Jamf Computers Software Updates",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetSoftwareUpdates{},
			transformers.WithPrimaryKeys("Name"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchSoftwareUpdates,
	}
}

func fetchSoftwareUpdates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, update := range *computer.SoftwareUpdates {
		res <- update
	}

	return nil
}
