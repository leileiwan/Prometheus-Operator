// Copyright 2018 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//定义了Prometheus 对象的集中类型，并且设置统一接口为几种类型赋值（但是具体spce由什么元素组成没有给出）
//但是真正的类型定义是由types.go完成的。
//不知道这个文件是作为什么用的
package v1

import (
	"fmt"
	"strings"
)

type CrdKind struct {
	Kind     string
	Plural   string
	SpecName string
}

type CrdKinds struct {
	KindsString    string
	Prometheus     CrdKind
	Alertmanager   CrdKind
	ServiceMonitor CrdKind
	PodMonitor     CrdKind
	PrometheusRule CrdKind
}

var DefaultCrdKinds = CrdKinds{
	KindsString:    "",
	Prometheus:     CrdKind{Plural: PrometheusName, Kind: PrometheusesKind, SpecName: "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1.Prometheus"},
	ServiceMonitor: CrdKind{Plural: ServiceMonitorName, Kind: ServiceMonitorsKind, SpecName: "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1.ServiceMonitor"},
	PodMonitor:     CrdKind{Plural: PodMonitorName, Kind: PodMonitorsKind, SpecName: "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1.PodMonitor"},
	Alertmanager:   CrdKind{Plural: AlertmanagerName, Kind: AlertmanagersKind, SpecName: "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1.Alertmanager"},
	PrometheusRule: CrdKind{Plural: PrometheusRuleName, Kind: PrometheusRuleKind, SpecName: "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1.PrometheusRule"},
}

// Implement the flag.Value interface
func (crdkinds *CrdKinds) String() string {
	return crdkinds.KindsString
}

// Set Implement the flag.Set interface
func (crdkinds *CrdKinds) Set(value string) error {
	*crdkinds = DefaultCrdKinds
	if value == "" {
		value = fmt.Sprintf("%s=%s:%s,%s=%s:%s,%s=%s:%s,%s=%s:%s,%s=%s:%s",
			PrometheusKindKey, PrometheusesKind, PrometheusName,
			AlertManagerKindKey, AlertmanagersKind, AlertmanagerName,
			ServiceMonitorKindKey, ServiceMonitorsKind, ServiceMonitorName,
			PodMonitorKindKey, PodMonitorsKind, PodMonitorName,
			PrometheusRuleKindKey, PrometheusRuleKind, PrometheusRuleName,
		)
	}
	splited := strings.Split(value, ",")
	for _, pair := range splited {
		sp := strings.Split(pair, "=")
		kind := strings.Split(sp[1], ":")
		crdKind := CrdKind{Plural: kind[1], Kind: kind[0]}
		switch kindKey := sp[0]; kindKey {
		case PrometheusKindKey:
			(*crdkinds).Prometheus = crdKind
		case ServiceMonitorKindKey:
			(*crdkinds).ServiceMonitor = crdKind
		case PodMonitorKindKey:
			(*crdkinds).PodMonitor = crdKind
		case AlertManagerKindKey:
			(*crdkinds).Alertmanager = crdKind
		case PrometheusRuleKindKey:
			(*crdkinds).PrometheusRule = crdKind
		default:
			fmt.Printf("Warning: unknown kind: %s... ignoring", kindKey)
		}

	}
	(*crdkinds).KindsString = value
	return nil
}
