# Table: jamf_computer_inventory_licensed_software

This table shows data for Jamf Computer Inventory Licensed Software.

Jamf Computers Licensed Software

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
|name|`utf8`|