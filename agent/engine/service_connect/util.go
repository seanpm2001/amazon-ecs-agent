// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package serviceconnect

import (
	"fmt"

	apitask "github.com/aws/amazon-ecs-agent/agent/api/task"
)

// DNSConfigToDockerExtraHostsFormat converts a []DNSConfigEntry slice to a list of ExtraHost entries that Docker will
// recognize.
func DNSConfigToDockerExtraHostsFormat(dnsConfigs []apitask.DNSConfigEntry) []string {
	var hosts []string
	for _, dnsConf := range dnsConfigs {
		if len(dnsConf.Address) > 0 {
			hosts = append(hosts,
				fmt.Sprintf("%s:%s", dnsConf.HostName, dnsConf.Address))
		}
	}
	return hosts
}