package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryAttachmentsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_attachments",
		Description: "Jamf Computers Attachments",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetAttachments{},
			transformers.WithPrimaryKeys("ID"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchAttachments,
	}
}

func fetchAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, attachment := range *computer.Attachments {
		res <- attachment
	}

	return nil
}
