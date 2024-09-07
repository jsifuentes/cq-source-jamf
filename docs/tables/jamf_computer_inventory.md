# Table: jamf_computer_inventory

This table shows data for Jamf Computer Inventory.

Jamf Computers

The primary key for this table is **id**.

## Relations

The following tables depend on jamf_computer_inventory:
  - [jamf_computer_inventory_applications](jamf_computer_inventory_applications.md)
  - [jamf_computer_inventory_attachments](jamf_computer_inventory_attachments.md)
  - [jamf_computer_inventory_certificates](jamf_computer_inventory_certificates.md)
  - [jamf_computer_inventory_configuration_profiles](jamf_computer_inventory_configuration_profiles.md)
  - [jamf_computer_inventory_extension_attributes](jamf_computer_inventory_extension_attributes.md)
  - [jamf_computer_inventory_fonts](jamf_computer_inventory_fonts.md)
  - [jamf_computer_inventory_group_memberships](jamf_computer_inventory_group_memberships.md)
  - [jamf_computer_inventory_ibeacons](jamf_computer_inventory_ibeacons.md)
  - [jamf_computer_inventory_licensed_software](jamf_computer_inventory_licensed_software.md)
  - [jamf_computer_inventory_local_user_accounts](jamf_computer_inventory_local_user_accounts.md)
  - [jamf_computer_inventory_plugins](jamf_computer_inventory_plugins.md)
  - [jamf_computer_inventory_printers](jamf_computer_inventory_printers.md)
  - [jamf_computer_inventory_services](jamf_computer_inventory_services.md)
  - [jamf_computer_inventory_software_updates](jamf_computer_inventory_software_updates.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|instance_domain|`utf8`|
|id (PK)|`utf8`|
|udid|`utf8`|
|general_name|`utf8`|
|general_last_ip_address|`utf8`|
|general_last_reported_ip|`utf8`|
|general_jamf_binary_version|`utf8`|
|general_platform|`utf8`|
|general_barcode1|`utf8`|
|general_barcode2|`utf8`|
|general_asset_tag|`utf8`|
|general_remote_management|`json`|
|general_supervised|`bool`|
|general_mdm_capable|`json`|
|general_report_date|`timestamp[us, tz=UTC]`|
|general_last_contact_time|`timestamp[us, tz=UTC]`|
|general_last_cloud_backup_date|`timestamp[us, tz=UTC]`|
|general_last_enrolled_date|`timestamp[us, tz=UTC]`|
|general_mdm_profile_expiration|`timestamp[us, tz=UTC]`|
|general_initial_entry_date|`utf8`|
|general_distribution_point|`utf8`|
|general_enrollment_method|`json`|
|general_site|`json`|
|general_itunes_store_account_active|`bool`|
|general_enrolled_via_automated_device_enrollment|`bool`|
|general_user_approved_mdm|`bool`|
|general_extension_attributes|`json`|
|disk_encryption_boot_partition_encryption_details|`json`|
|disk_encryption_individual_recovery_key_validity_status|`utf8`|
|disk_encryption_institutional_recovery_key_present|`bool`|
|disk_encryption_disk_encryption_configuration_name|`utf8`|
|disk_encryption_file_vault2_enabled_user_names|`list<item: utf8, nullable>`|
|disk_encryption_file_vault2_eligibility_message|`utf8`|
|purchasing_leased|`bool`|
|purchasing_purchased|`bool`|
|purchasing_po_number|`utf8`|
|purchasing_po_date|`utf8`|
|purchasing_vendor|`utf8`|
|purchasing_warranty_date|`utf8`|
|purchasing_apple_care_id|`utf8`|
|purchasing_lease_date|`utf8`|
|purchasing_purchase_price|`utf8`|
|purchasing_life_expectancy|`int64`|
|purchasing_purchasing_account|`utf8`|
|purchasing_purchasing_contact|`utf8`|
|purchasing_extension_attributes|`json`|
|applications|`json`|
|storage_boot_drive_available_space_megabytes|`int64`|
|storage_disks|`json`|
|user_and_location_username|`utf8`|
|user_and_location_realname|`utf8`|
|user_and_location_email|`utf8`|
|user_and_location_position|`utf8`|
|user_and_location_phone|`utf8`|
|user_and_location_department_id|`utf8`|
|user_and_location_building_id|`utf8`|
|user_and_location_room|`utf8`|
|user_and_location_extension_attributes|`json`|
|configuration_profiles|`json`|
|printers|`json`|
|services|`json`|
|hardware_make|`utf8`|
|hardware_model|`utf8`|
|hardware_model_identifier|`utf8`|
|hardware_serial_number|`utf8`|
|hardware_processor_speed_mhz|`int64`|
|hardware_processor_count|`int64`|
|hardware_core_count|`int64`|
|hardware_processor_type|`utf8`|
|hardware_processor_architecture|`utf8`|
|hardware_bus_speed_mhz|`int64`|
|hardware_cache_size_kilobytes|`int64`|
|hardware_network_adapter_type|`utf8`|
|hardware_mac_address|`utf8`|
|hardware_alt_network_adapter_type|`utf8`|
|hardware_alt_mac_address|`utf8`|
|hardware_total_ram_megabytes|`int64`|
|hardware_open_ram_slots|`int64`|
|hardware_battery_capacity_percent|`int64`|
|hardware_smc_version|`utf8`|
|hardware_nic_speed|`utf8`|
|hardware_optical_drive|`utf8`|
|hardware_boot_rom|`utf8`|
|hardware_ble_capable|`bool`|
|hardware_supports_ios_app_installs|`bool`|
|hardware_apple_silicon|`bool`|
|hardware_extension_attributes|`json`|
|local_user_accounts|`json`|
|certificates|`json`|
|attachments|`json`|
|plugins|`json`|
|package_receipts_installed_by_jamf_pro|`list<item: utf8, nullable>`|
|package_receipts_installed_by_installer_swu|`list<item: utf8, nullable>`|
|package_receipts_cached|`list<item: utf8, nullable>`|
|fonts|`json`|
|security_sip_status|`utf8`|
|security_gatekeeper_status|`utf8`|
|security_xprotect_version|`utf8`|
|security_auto_login_disabled|`bool`|
|security_remote_desktop_enabled|`bool`|
|security_activation_lock_enabled|`bool`|
|security_recovery_lock_enabled|`bool`|
|security_firewall_enabled|`bool`|
|security_secure_boot_level|`utf8`|
|security_external_boot_level|`utf8`|
|security_bootstrap_token_allowed|`bool`|
|operating_system_name|`utf8`|
|operating_system_version|`utf8`|
|operating_system_build|`utf8`|
|operating_system_active_directory_status|`utf8`|
|operating_system_file_vault2_status|`utf8`|
|operating_system_software_update_device_id|`utf8`|
|operating_system_extension_attributes|`json`|
|licensed_software|`json`|
|ibeacons|`json`|
|software_updates|`json`|
|extension_attributes|`json`|
|content_caching_computer_content_caching_information_id|`utf8`|
|content_caching_parents|`json`|
|content_caching_alerts|`json`|
|content_caching_activated|`bool`|
|content_caching_active|`bool`|
|content_caching_actual_cache_bytes_used|`int64`|
|content_caching_cache_details|`json`|
|content_caching_cache_bytes_free|`int64`|
|content_caching_cache_bytes_limit|`int64`|
|content_caching_cache_status|`utf8`|
|content_caching_cache_bytes_used|`int64`|
|content_caching_data_migration_completed|`bool`|
|content_caching_data_migration_progress_percentage|`int64`|
|content_caching_data_migration_error|`json`|
|content_caching_max_cache_pressure_last1_hour_percentage|`int64`|
|content_caching_personal_cache_bytes_free|`int64`|
|content_caching_personal_cache_bytes_limit|`int64`|
|content_caching_personal_cache_bytes_used|`int64`|
|content_caching_port|`int64`|
|content_caching_public_address|`utf8`|
|content_caching_registration_error|`utf8`|
|content_caching_registration_response_code|`int64`|
|content_caching_registration_started|`timestamp[us, tz=UTC]`|
|content_caching_registration_status|`utf8`|
|content_caching_restricted_media|`bool`|
|content_caching_server_guid|`utf8`|
|content_caching_startup_status|`utf8`|
|content_caching_tetherator_status|`utf8`|
|content_caching_total_bytes_are_since|`timestamp[us, tz=UTC]`|
|content_caching_total_bytes_dropped|`int64`|
|content_caching_total_bytes_imported|`int64`|
|content_caching_total_bytes_returned_to_children|`int64`|
|content_caching_total_bytes_returned_to_clients|`int64`|
|content_caching_total_bytes_returned_to_peers|`int64`|
|content_caching_total_bytes_stored_from_origin|`int64`|
|content_caching_total_bytes_stored_from_parents|`int64`|
|content_caching_total_bytes_stored_from_peers|`int64`|
|group_memberships|`json`|