# Table: jamf_computer_inventory_services

This table shows data for Jamf Computer Inventory Services.

Jamf Computers Services

The composite primary key for this table is (**computer_id**, **name**).

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|name (PK)|`utf8`|