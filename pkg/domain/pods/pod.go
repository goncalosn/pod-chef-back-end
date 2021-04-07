package pods

import v1 "k8s.io/api/core/v1"

type Pod struct {
	State        v1.PodPhase `json:"state"`
	RestartCount int32       `json:"restartCount"`
	Name         string      `json:"name"`
}
