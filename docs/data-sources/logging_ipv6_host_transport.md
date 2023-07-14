---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_logging_ipv6_host_transport Data Source - terraform-provider-iosxe"
subcategory: "Management"
description: |-
  This data source can read the Logging IPv6 Host Transport configuration.
---

# iosxe_logging_ipv6_host_transport (Data Source)

This data source can read the Logging IPv6 Host Transport configuration.

## Example Usage

```terraform
data "iosxe_logging_ipv6_host_transport" "example" {
  ipv6_host = "2001::1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ipv6_host` (String)

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `transport_tcp_ports` (Attributes List) Port Number List (see [below for nested schema](#nestedatt--transport_tcp_ports))
- `transport_tls_ports` (Attributes List) Port Number List (see [below for nested schema](#nestedatt--transport_tls_ports))
- `transport_udp_ports` (Attributes List) Port Number List (see [below for nested schema](#nestedatt--transport_udp_ports))

<a id="nestedatt--transport_tcp_ports"></a>
### Nested Schema for `transport_tcp_ports`

Read-Only:

- `port_number` (Number) Specify the TCP port number (default=601)


<a id="nestedatt--transport_tls_ports"></a>
### Nested Schema for `transport_tls_ports`

Read-Only:

- `port_number` (Number) Specify the TLS port number (default=6514)
- `profile` (String) Specify the TLS profile


<a id="nestedatt--transport_udp_ports"></a>
### Nested Schema for `transport_udp_ports`

Read-Only:

- `port_number` (Number) Specify the UDP port number (default=514)