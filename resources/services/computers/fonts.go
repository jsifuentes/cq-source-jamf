package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryFontsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_fonts",
		Description: "Jamf Computers Fonts",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetFonts{},
			transformers.WithPrimaryKeys("Path"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchFonts,
	}
}

func fetchFonts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, font := range *computer.Fonts {
		res <- font
	}

	return nil
}
