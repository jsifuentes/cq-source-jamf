package computers

import (
	"context"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-jamf/client"
	"github.com/yohan460/go-jamf-api"
)

func ComputerInventoryTable() *schema.Table {
	return &schema.Table{
		Name:        "jamf_computer_inventory",
		Description: "Jamf Computers",
		Transform: transformers.TransformWithStruct(
			&jamf.ComputerInventory{},
			transformers.WithPrimaryKeys("ID"),
			transformers.WithUnwrapStructFields("General", "DiskEncryption", "Purchasing", "Storage", "UserAndLocation", "Hardware", "PackageReceipts", "Security", "OperatingSystem", "ContentCaching"),
		),
		Columns: []schema.Column{
			client.InstanceDomainColumn,
		},
		Relations: []*schema.Table{
			ComputerInventoryExtensionAttributesTable(),
			ComputerInventoryApplicationsTable(),
			ComputerInventoryAttachmentsTable(),
			ComputerInventoryCertificatesTable(),
			ComputerInventoryConfigurationProfilesTable(),
			ComputerInventoryFontsTable(),
			ComputerInventoryGroupMembershipsTable(),
			ComputerInventoryIbeaconsTable(),
			ComputerInventoryLicensedSoftwareTable(),
			ComputerInventoryLocalUserAccountsTable(),
			ComputerInventoryPluginsTable(),
			ComputerInventoryPrintersTable(),
			ComputerInventoryServicesTable(),
			ComputerInventorySoftwareUpdatesTable(),
		},
		Resolver: fetchComputers,
	}
}

func fetchComputers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	allSections := []jamf.ComputerInventoryDataSubsetName{
		jamf.ComputerInventoryDataSubsetNameGeneral,
		jamf.ComputerInventoryDataSubsetNameLocation,
		jamf.ComputerInventoryDataSubsetNamePurchasing,
		jamf.ComputerInventoryDataSubsetNameApplications,
		jamf.ComputerInventoryDataSubsetNameStorage,
		jamf.ComputerInventoryDataSubsetNameUserAndLocation,
		jamf.ComputerInventoryDataSubsetNameConfigurationProfiles,
		jamf.ComputerInventoryDataSubsetNamePrinters,
		jamf.ComputerInventoryDataSubsetNameServices,
		jamf.ComputerInventoryDataSubsetNameHardware,
		jamf.ComputerInventoryDataSubsetNameLocalUserAccounts,
		jamf.ComputerInventoryDataSubsetNameCertificates,
		jamf.ComputerInventoryDataSubsetNameAttachments,
		jamf.ComputerInventoryDataSubsetNamePlugins,
		jamf.ComputerInventoryDataSubsetNamePackageReceipts,
		jamf.ComputerInventoryDataSubsetNameFonts,
		jamf.ComputerInventoryDataSubsetNameSecurity,
		jamf.ComputerInventoryDataSubsetNameOperatingSystem,
		jamf.ComputerInventoryDataSubsetNameLicensedSoftware,
		jamf.ComputerInventoryDataSubsetNameIBeacons,
		jamf.ComputerInventoryDataSubsetNameSoftwareUpdates,
		jamf.ComputerInventoryDataSubsetNameExtensionAttributes,
		jamf.ComputerInventoryDataSubsetNameContentCaching,
		jamf.ComputerInventoryDataSubsetNameGroupMemberships,
	}

	var totalCount int
	var results []jamf.ComputerInventory
	page := 1
	for done := false; !done; done = (len(results) >= totalCount) {
		computers, err := c.JamfClient.GetComputerInventories(&jamf.ComputerInventoriesQuery{
			Sections: &allSections,
			Page:     page,
			PageSize: 20,
			Sort:     &[]string{"general.name:asc"},
		})
		if err != nil {
			return err
		}

		totalCount = computers.TotalCount
		results = append(results, computers.Results...)
		page++
	}

	for _, computer := range results {
		res <- computer
	}

	return nil
}

var computerIDColumn = schema.Column{
	Name:       "computer_id",
	Type:       arrow.BinaryTypes.String,
	PrimaryKey: true,
	Resolver:   resolveComputerIdColumn,
}

func resolveComputerIdColumn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	parent := resource.Parent.Item.(jamf.ComputerInventory)
	return resource.Set(c.Name, parent.ID)
}
