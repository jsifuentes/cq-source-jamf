# Table: jamf_computer_inventory_certificates

This table shows data for Jamf Computer Inventory Certificates.

Jamf Computers Certificates

The primary key for this table is **computer_id**.

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|common_name|`utf8`|
|identity|`bool`|
|expiration_date|`timestamp[us, tz=UTC]`|
|username|`utf8`|
|lifecycle_status|`utf8`|
|certificate_status|`utf8`|