# Table: jamf_computer_inventory_configuration_profiles

This table shows data for Jamf Computer Inventory Configuration Profiles.

Jamf Computers Configuration Profiles

The composite primary key for this table is (**computer_id**, **id**).

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|id (PK)|`utf8`|
|username|`utf8`|
|last_installed|`timestamp[us, tz=UTC]`|
|removable|`bool`|
|display_name|`utf8`|
|profile_identifier|`utf8`|