package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryPrintersTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_printers",
		Description: "Jamf Computers Printers",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetPrinters{},
			transformers.WithPrimaryKeys("Name"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchPrinters,
	}
}

func fetchPrinters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, printer := range *computer.Printers {
		res <- printer
	}

	return nil
}
