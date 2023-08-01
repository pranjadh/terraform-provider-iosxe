// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CryptoIKEv2KeyringDataSource{}
	_ datasource.DataSourceWithConfigure = &CryptoIKEv2KeyringDataSource{}
)

func NewCryptoIKEv2KeyringDataSource() datasource.DataSource {
	return &CryptoIKEv2KeyringDataSource{}
}

type CryptoIKEv2KeyringDataSource struct {
	clients map[string]*restconf.Client
}

func (d *CryptoIKEv2KeyringDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crypto_ikev2_keyring"
}

func (d *CryptoIKEv2KeyringDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Crypto IKEv2 Keyring configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"peers": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a Peer and associated keys",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Specify a description of this peer",
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname of peer",
							Computed:            true,
						},
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv4_mask": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv6_prefix": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"identity_key_id": schema.StringAttribute{
							MarkdownDescription: "proprietary types of identification (ID KEY ID)",
							Computed:            true,
						},
						"identity_address": schema.StringAttribute{
							MarkdownDescription: "IP address",
							Computed:            true,
						},
						"identity_email_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name string",
							Computed:            true,
						},
						"identity_email_domain": schema.StringAttribute{
							MarkdownDescription: "email Domain Name",
							Computed:            true,
						},
						"identity_fqdn_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name string",
							Computed:            true,
						},
						"identity_fqdn_domain": schema.StringAttribute{
							MarkdownDescription: "email Domain Name",
							Computed:            true,
						},
						"pre_shared_key_local_encryption": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pre_shared_key_local": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pre_shared_key_remote_encryption": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pre_shared_key_remote": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pre_shared_key_encryption": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pre_shared_key": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CryptoIKEv2KeyringDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *CryptoIKEv2KeyringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CryptoIKEv2KeyringData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = CryptoIKEv2KeyringData{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
