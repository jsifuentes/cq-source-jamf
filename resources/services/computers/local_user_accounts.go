package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryLocalUserAccountsTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_local_user_accounts",
		Description: "Jamf Computers Local User Accounts",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetLocalUserAccounts{},
			transformers.WithPrimaryKeys("UID"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchLocalUserAccounts,
	}
}

func fetchLocalUserAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, account := range *computer.LocalUserAccounts {
		res <- account
	}

	return nil
}
