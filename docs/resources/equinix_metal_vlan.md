---
subcategory: "Metal"
---

# equinix_metal_vlan (Resource)

Provides a resource to allow users to manage Virtual Networks in their projects.

To learn more about Layer 2 networking in Equinix Metal, refer to

* <https://metal.equinix.com/developers/docs/networking/layer2/>
* <https://metal.equinix.com/developers/docs/networking/layer2-configs/>

## Example Usage

```hcl
# Create a new VLAN in facility "sv15"
resource "equinix_metal_vlan" "vlan1" {
  description = "VLAN in New Jersey"
  facility    = "sv15"
  project_id  = local.project_id
}

# Create a new VLAN in metro "esv"
resource "equinix_metal_vlan" "vlan1" {
  description = "VLAN in New Jersey"
  metro       = "sv"
  project_id  = local.project_id
  vxlan       = 1040
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) ID of parent project.
* `facility` - (Required) Facility where to create the VLAN.
* `description` - (Optional) Description string.
* `vxlan` - (Optional) VLAN ID, must be unique in metro.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the virtual network.

## Import

This resource can be imported using an existing VLAN ID (UUID):

```sh
terraform import equinix_metal_vlan {existing_vlan_id}
```
