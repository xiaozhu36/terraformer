// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	util "github.com/alibabacloud-go/tea-utils/service"
)

// VpcGenerator Struct for generating AliCloud Elastic Compute Service
type VpcGenerator struct {
	AliCloudService
}

func resourceFromVpcResponse(vpc map[string]interface{}) terraformutils.Resource {
	return terraformutils.NewResource(
		vpc["VpcId"].(string),
		vpc["VpcId"].(string),
		"alicloud_vpc",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all vpc Vpc ids and generates resources
func (g *VpcGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	conn, err := client.NewVpcClient()
	if err != nil {
		return err
	}
	action := "DescribeVpcs"
	request := map[string]interface{}{
		"RegionId":   client.RegionId,
		"PageNumber": 1,
		"PageSize":   50,
	}
	allVpcs := make([]interface{}, 0, 50)
	for {
		response, err := conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			return err
		}
		vpcs := response["Vpcs"].(map[string]interface{})["Vpc"].([]interface{})
		allVpcs = append(allVpcs, vpcs...)
		if len(vpcs) < 50 {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	for _, vpc := range allVpcs {
		resource := resourceFromVpcResponse(vpc.(map[string]interface{}))
		g.Resources = append(g.Resources, resource)
	}
	return nil
}
