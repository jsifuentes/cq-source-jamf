package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryApplicationsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_applications",
		Description: "Jamf Computers Applications",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetApplications{},
			transformers.WithPrimaryKeys("Path"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchApplications,
	}
}

func fetchApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, application := range *computer.Applications {
		res <- application
	}

	return nil
}
