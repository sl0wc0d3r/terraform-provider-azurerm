---
subcategory: "Messaging"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_eventgrid_topic"
description: |-
  Manages an EventGrid Topic

---

# azurerm_eventgrid_topic

Manages an EventGrid Topic

~> **Note:** at this time EventGrid Topic's are only available in a limited number of regions.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_eventgrid_topic" "example" {
  name                = "my-eventgrid-topic"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  tags = {
    environment = "Production"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Specifies the name of the EventGrid Topic resource. Changing this forces a new resource to be created.

* `resource_group_name` - (Required) The name of the resource group in which the EventGrid Topic exists. Changing this forces a new resource to be created.

* `location` - (Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.

* `identity` - (Optional) An `identity` block as defined below.

* `input_schema` - (Optional) Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `EventGridSchema`. Changing this forces a new resource to be created.

* `input_mapping_fields` - (Optional) A `input_mapping_fields` block as defined below.

* `input_mapping_default_values` - (Optional) A `input_mapping_default_values` block as defined below.

* `public_network_access_enabled` - (Optional) Whether or not public network access is allowed for this server. Defaults to `true`.

* `inbound_ip_rule` - (Optional) One or more `inbound_ip_rule` blocks as defined below.

* `tags` - (Optional) A mapping of tags to assign to the resource.

---

A `identity` block supports the following:

* `type` - Specifies the identity type of Event Grid Topic. Possible values are `SystemAssigned` (where Azure will generate a Principal for you) or `UserAssigned` where you can specify the User Assigned Managed Identity IDs in the `identity_ids` field.

~> **NOTE:** When `type` is set to `SystemAssigned`, The assigned `principal_id` and `tenant_id` can be retrieved after the Event Grid Topic has been created. More details are available below.

* `identity_ids` - (Optional) Specifies a list of user managed identity ids to be assigned. Required if `type` is `UserAssigned`.

---

A `input_mapping_fields` supports the following:

* `id` - (Optional) Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `topic` - (Optional) Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `event_type` - (Optional) Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `event_time` - (Optional) Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `data_version` - (Optional) Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `subject` - (Optional) Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

---

A `input_mapping_default_values` supports the following:

* `event_type` - (Optional) Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `data_version` - (Optional) Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

* `subject` - (Optional) Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.

---

A `inbound_ip_rule` block supports the following:

* `ip_mask` - (Required) The ip mask (CIDR) to match on.

* `action` - (Optional) The action to take when the rule is matched. Possible values are `Allow`.

## Attributes Reference

The following attributes are exported:

* `id` - The EventGrid Topic ID.

* `endpoint` - The Endpoint associated with the EventGrid Topic.

* `primary_access_key` - The Primary Shared Access Key associated with the EventGrid Topic.

* `secondary_access_key` - The Secondary Shared Access Key associated with the EventGrid Topic.

* `identity` - An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid Topic.

---

A `identity` block supports the following:

* `type` - Specifies the type of Managed Service Identity that is configured on this Event Grid Topic.

* `principal_id` - Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Event Grid Topic.

* `tenant_id` - Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Event Grid Topic.

* `identity_ids` - A list of IDs for User Assigned Managed Identity resources to be assigned.

---

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the EventGrid Topic.
* `update` - (Defaults to 30 minutes) Used when updating the EventGrid Topic.
* `read` - (Defaults to 5 minutes) Used when retrieving the EventGrid Topic.
* `delete` - (Defaults to 30 minutes) Used when deleting the EventGrid Topic.

## Import

EventGrid Topic's can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_eventgrid_topic.topic1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
```
