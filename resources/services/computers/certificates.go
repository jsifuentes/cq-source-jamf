package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryCertificatesTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_certificates",
		Description: "Jamf Computers Certificates",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetCertificates{},
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchCertificates,
	}
}

func fetchCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, certificate := range *computer.Certificates {
		res <- certificate
	}

	return nil
}
