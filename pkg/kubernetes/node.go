package kubernetes

import v1 "k8s.io/api/core/v1"

type Node struct {
	MemoryPressure v1.NodeConditionType
	DiskPressure   v1.NodeConditionType
	PIDPressure    v1.NodeConditionType
	Ready          v1.NodeConditionType
}
