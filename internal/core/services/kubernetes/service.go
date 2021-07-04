package kubernetes

import (
	ports "pod-chef-back-end/internal/core/ports"
)

//Service kubernetes repository
type Service struct {
	kubernetesRepository ports.KubernetesRepository
	mongoRepository      ports.MongoRepository
	cloudflareRepository ports.CloudflareRepository
}

//NewKubernetesService where the kubernetes repository is injected
func NewKubernetesService(kubernetesRepository ports.KubernetesRepository, mongoRepository ports.MongoRepository, cloudflareRepository ports.CloudflareRepository) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
		mongoRepository:      mongoRepository,
		cloudflareRepository: cloudflareRepository,
	}
}
