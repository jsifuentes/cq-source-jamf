# Table: jamf_computer_inventory_group_memberships

This table shows data for Jamf Computer Inventory Group Memberships.

Jamf Computers Group Memberships

The composite primary key for this table is (**computer_id**, **group_id**).

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|group_id (PK)|`utf8`|
|group_name|`utf8`|
|smart_group|`bool`|