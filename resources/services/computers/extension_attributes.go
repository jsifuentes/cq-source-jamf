package computers

import (
	"context"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yohan460/go-jamf-api"
)

type ExtensionAttributeRow struct {
	jamf.ComputerInventoryDataSubsetExtensionAttributes
	Section string `json:"section"`
}

func ComputerInventoryExtensionAttributesTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory_extension_attributes",
		Description: "Jamf Computers Extension Attributes",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventoryDataSubsetExtensionAttributes{},
			transformers.WithPrimaryKeys("DefinitionID"),
		),
		Columns: []schema.Column{
			computerIDColumn,
			{
				Name:       "section",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
		Resolver: fetchExtensionAttributes,
	}
}

func fetchExtensionAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	computer := parent.Item.(jamf.ComputerInventory)

	// Send all extension attributes
	// We have to do multiple for loops because the extension attributes properties
	// are defined as inline structs, so none of them are the same type.

	for _, extensionAttribute := range *computer.ExtensionAttributes {
		row := ExtensionAttributeRow{
			ComputerInventoryDataSubsetExtensionAttributes: extensionAttribute,
			Section: "root",
		}
		res <- row
	}

	for _, extensionAttribute := range computer.General.ExtensionAttributes {
		row := ExtensionAttributeRow{
			ComputerInventoryDataSubsetExtensionAttributes: extensionAttribute,
			Section: "general",
		}
		res <- row
	}

	for _, extensionAttribute := range computer.OperatingSystem.ExtensionAttributes {
		row := ExtensionAttributeRow{
			ComputerInventoryDataSubsetExtensionAttributes: extensionAttribute,
			Section: "operating_system",
		}
		res <- row
	}

	for _, extensionAttribute := range computer.Purchasing.ExtensionAttributes {
		row := ExtensionAttributeRow{
			ComputerInventoryDataSubsetExtensionAttributes: extensionAttribute,
			Section: "purchasing",
		}
		res <- row
	}

	for _, extensionAttribute := range computer.UserAndLocation.ExtensionAttributes {
		row := ExtensionAttributeRow{
			ComputerInventoryDataSubsetExtensionAttributes: extensionAttribute,
			Section: "user_and_location",
		}
		res <- row
	}

	for _, extensionAttribute := range computer.Hardware.ExtensionAttributes {
		row := ExtensionAttributeRow{
			ComputerInventoryDataSubsetExtensionAttributes: extensionAttribute,
			Section: "hardware",
		}
		res <- row
	}

	return nil
}
