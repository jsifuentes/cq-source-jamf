package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryServicesTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_services",
		Description: "Jamf Computers Services",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetServices{},
			transformers.WithPrimaryKeys("Name"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchServices,
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, service := range *computer.Services {
		res <- service
	}

	return nil
}
