---
layout: "hdns"
page_title: "Hetzner DNS: hdns_zone"
sidebar_current: "docs-hdns-resource-zone"
description: |-
  Provides a Hetzner DNS zone resource. This can be used to create, modify and delete DNS zones.
---

# hdns_zone

Provides a Hetzner DNS zone resource. This can be used to create, modify and delete DNS zones.

## Example Usage
```hcl
resource "hdns_zone" "zone1" {
  name = "example.com"
  ttl  = 84600
}
```

## Argument Reference

- `name` - (Required, string) Domain name for which DNS zone is created.
- `ttl` - (Required, int) TTL of the zone.

## Attributes Reference

- `id` - (int) Unique ID of the DNS Zone.
- `name` - (string) Domain name for which DNS zone is created.
- `ttl` - (int) TTL of the zone.

## Import

DNS Zones can be imported using the zone `id`:

```
terraform import hdns_zone.myzone 123
```
