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
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Banner struct {
	Device              types.String `tfsdk:"device"`
	Id                  types.String `tfsdk:"id"`
	DeleteMode          types.String `tfsdk:"delete_mode"`
	ExecBanner          types.String `tfsdk:"exec_banner"`
	LoginBanner         types.String `tfsdk:"login_banner"`
	PromptTimeoutBanner types.String `tfsdk:"prompt_timeout_banner"`
	MotdBanner          types.String `tfsdk:"motd_banner"`
}

type BannerData struct {
	Device              types.String `tfsdk:"device"`
	Id                  types.String `tfsdk:"id"`
	ExecBanner          types.String `tfsdk:"exec_banner"`
	LoginBanner         types.String `tfsdk:"login_banner"`
	PromptTimeoutBanner types.String `tfsdk:"prompt_timeout_banner"`
	MotdBanner          types.String `tfsdk:"motd_banner"`
}

func (data Banner) getPath() string {
	return "Cisco-IOS-XE-native:native/banner"
}

func (data BannerData) getPath() string {
	return "Cisco-IOS-XE-native:native/banner"
}

// if last path element has a key -> remove it
func (data Banner) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Banner) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.ExecBanner.IsNull() && !data.ExecBanner.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec.banner", data.ExecBanner.ValueString())
	}
	if !data.LoginBanner.IsNull() && !data.LoginBanner.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"login.banner", data.LoginBanner.ValueString())
	}
	if !data.PromptTimeoutBanner.IsNull() && !data.PromptTimeoutBanner.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prompt-timeout.banner", data.PromptTimeoutBanner.ValueString())
	}
	if !data.MotdBanner.IsNull() && !data.MotdBanner.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"motd.banner", data.MotdBanner.ValueString())
	}
	return body
}

func (data *Banner) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "exec.banner"); value.Exists() && !data.ExecBanner.IsNull() {
		data.ExecBanner = types.StringValue(value.String())
	} else {
		data.ExecBanner = types.StringNull()
	}
	if value := res.Get(prefix + "login.banner"); value.Exists() && !data.LoginBanner.IsNull() {
		data.LoginBanner = types.StringValue(value.String())
	} else {
		data.LoginBanner = types.StringNull()
	}
	if value := res.Get(prefix + "prompt-timeout.banner"); value.Exists() && !data.PromptTimeoutBanner.IsNull() {
		data.PromptTimeoutBanner = types.StringValue(value.String())
	} else {
		data.PromptTimeoutBanner = types.StringNull()
	}
	if value := res.Get(prefix + "motd.banner"); value.Exists() && !data.MotdBanner.IsNull() {
		data.MotdBanner = types.StringValue(value.String())
	} else {
		data.MotdBanner = types.StringNull()
	}
}

func (data *BannerData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "exec.banner"); value.Exists() {
		data.ExecBanner = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "login.banner"); value.Exists() {
		data.LoginBanner = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "prompt-timeout.banner"); value.Exists() {
		data.PromptTimeoutBanner = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "motd.banner"); value.Exists() {
		data.MotdBanner = types.StringValue(value.String())
	}
}

func (data *Banner) getDeletedListItems(ctx context.Context, state Banner) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *Banner) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *Banner) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.ExecBanner.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/exec/banner", data.getPath()))
	}
	if !data.LoginBanner.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/login/banner", data.getPath()))
	}
	if !data.PromptTimeoutBanner.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/prompt-timeout/banner", data.getPath()))
	}
	if !data.MotdBanner.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/motd/banner", data.getPath()))
	}
	return deletePaths
}