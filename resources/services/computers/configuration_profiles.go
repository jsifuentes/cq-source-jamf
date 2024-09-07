package computers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryConfigurationProfilesTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_configuration_profiles",
		Description: "Jamf Computers Configuration Profiles",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetConfigurationProfiles{},
			transformers.WithPrimaryKeys("ID"),
		),
		Columns: []schema.Column{
			computerIDColumn,
		},
		Resolver: fetchConfigurationProfiles,
	}
}

func fetchConfigurationProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	for _, profile := range *computer.ConfigurationProfiles {
		res <- profile
	}

	return nil
}
