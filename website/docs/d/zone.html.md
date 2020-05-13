---
layout: "hdns"
page_title: "Hetzner DNS: hdns_zone"
sidebar_current: "docs-hdns-datasource-zone"
description: |-
  Provides details about a specific Hetzner DNS Zone.
---
# Data Source: hdns_zone
Provides details about a Hetzner DNS Zone.
This resource is useful if you want to use a non-terraform managed DNS zone.

## Example Usage
```hcl
data "hdns_zone" "z_1" {
  name = "example.com"
}
data "hdns_zone" "z_2" {
  id = "123"
}
```

## Argument Reference
- `id` - ID of the zone.
- `name` - Name of the zone.

## Attributes Reference
- `id` - (int) Unique ID of the zone.
- `name` - (string) Name of the zone.
