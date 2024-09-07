# Table: jamf_computer_inventory_local_user_accounts

This table shows data for Jamf Computer Inventory Local User Accounts.

Jamf Computers Local User Accounts

The composite primary key for this table is (**computer_id**, **uid**).

## Relations

This table depends on [jamf_computer_inventory](jamf_computer_inventory.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|computer_id (PK)|`utf8`|
|uid (PK)|`utf8`|
|username|`utf8`|
|full_name|`utf8`|
|admin|`bool`|
|home_directory|`utf8`|
|home_directory_size_mb|`int64`|
|file_vault2_enabled|`bool`|
|user_account_type|`utf8`|
|password_min_length|`int64`|
|password_max_age|`int64`|
|password_min_complex_characters|`int64`|
|password_history_depth|`int64`|
|password_require_alphanumeric|`bool`|
|computer_azure_active_directory_id|`utf8`|
|user_azure_active_directory_id|`utf8`|
|azure_active_directory_id|`utf8`|