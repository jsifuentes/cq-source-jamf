# Table: jamf_computer_inventory_extension_attributes

This table shows data for Jamf Computer Inventory Extension Attributes.

Jamf Computers Extension Attributes

The composite primary key for this table is (**computer_id**, **section**, **definition_id**).

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|section (PK)|`utf8`|
|definition_id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|enabled|`bool`|
|multi_value|`bool`|
|values|`list<item: utf8, nullable>`|
|data_type|`utf8`|
|options|`list<item: utf8, nullable>`|
|input_type|`utf8`|