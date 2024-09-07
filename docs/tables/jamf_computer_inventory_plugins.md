# Table: jamf_computer_inventory_plugins

This table shows data for Jamf Computer Inventory Plugins.

Jamf Computers Plugins

The composite primary key for this table is (**computer_id**, **path**).

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|name|`utf8`|
|version|`utf8`|
|path (PK)|`utf8`|