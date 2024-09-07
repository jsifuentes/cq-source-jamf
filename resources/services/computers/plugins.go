package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryPluginsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_plugins",
		Description: "Jamf Computers Plugins",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetPlugins{},
			transformers.WithPrimaryKeys("Path"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchPlugins,
	}
}

func fetchPlugins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, plugin := range *computer.Plugins {
		res <- plugin
	}

	return nil
}
