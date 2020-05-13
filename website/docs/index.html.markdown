---
layout: "hdns"
page_title: "Provider: Hetzner DNS"
sidebar_current: "docs-hdns-index"
description: |-
  The Hetzner DNS (hdns) provider is used to interact with the resources supported by Hetzner DNS.
---

# Hetzner DNS Provider

The Hetzner DNS (hdns) provider is used to interact with the resources supported by [Hetzner Cloud](https://www.hetzner.com/cloud). The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Set the variable value in *.tfvars file
# or using the -var="hdns_token=..." CLI option
variable "hdns_token" {}

# Configure the Hetzner DNS Provider
provider "hdns" {
  token = "${var.hdns_token}"
}

# Create a zone
resource "hdns_zone" "example.com" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

- `token` - (Required, string) This is the Hetzner DNS API Token, can also be specified with the `HDNS_TOKEN` environment variable.
- `endpoint` - (Optional, string) Hetzner Cloud API endpoint, can be used to override the default API Endpoint `https://dns.hetzner.com/api/v1`.
- `poll_interval` -  (Optional, string) Configures the interval in which actions are polled by the client. Default `500ms`. Increase this interval if you run into rate limiting errors.
