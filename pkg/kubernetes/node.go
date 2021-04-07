package kubernetes

import v1 "k8s.io/api/core/v1"

type Node struct {
	State        v1.PodPhase
	RestartCount int32
	Name         string
}
