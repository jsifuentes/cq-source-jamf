package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryGroupMembershipsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_group_memberships",
		Description: "Jamf Computers Group Memberships",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetGroupMemberships{},
			transformers.WithPrimaryKeys("GroupID"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchGroupMemberships,
	}
}

func fetchGroupMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, group := range *computer.GroupMemberships {
		res <- group
	}

	return nil
}
