package nodes

import v1 "k8s.io/api/core/v1"

type Node struct {
	MemoryPressure v1.NodeConditionType `json:"MemoryPressure"`
	DiskPressure   v1.NodeConditionType `json:"DiskPressure"`
	PIDPressure    v1.NodeConditionType `json:"PIDPressure"`
	Ready          v1.NodeConditionType `json:"Ready"`
}
