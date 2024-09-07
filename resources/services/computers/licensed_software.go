package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryLicensedSoftwareTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_licensed_software",
		Description: "Jamf Computers Licensed Software",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetLicensedSoftware{},
			transformers.WithPrimaryKeys("ID"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchLicensedSoftware,
	}
}

func fetchLicensedSoftware(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, software := range *computer.LicensedSoftware {
		res <- software
	}

	return nil
}
