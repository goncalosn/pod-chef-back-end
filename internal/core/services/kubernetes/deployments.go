package kubernetes

import (
	uuid "github.com/satori/go.uuid"
)

func (srv *Service) GetDeployments() (interface{}, error) {
	response, err := srv.kubernetesRepository.GetDeployments()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (srv *Service) CreateDeployment(token string, replicas *int32, image string) (interface{}, error) {

	appUuid := uuid.NewV4().String() //generate uuid for this deployment

	namespaceUuid := "namespace-" + appUuid //generate name for namespace
	//create new namespace for user app
	_, err := srv.kubernetesRepository.CreateNamespace(namespaceUuid)
	if err != nil {
		return nil, err
	}

	deploymentUuid := "deployment-" + appUuid //generate name for the deployment

	//create deployment in namespace
	_, err = srv.kubernetesRepository.CreateDeployment(namespaceUuid, deploymentUuid, replicas, image)
	if err != nil {
		//delete namespace
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUuid)
		return nil, err
	}

	serviceUuid := "service-" + appUuid //generate name for the service
	//create service to expose the deployment
	_, err = srv.kubernetesRepository.CreateClusterIPService(namespaceUuid, serviceUuid)
	if err != nil {
		//delete namespace
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUuid)
		return nil, err
	}

	ingressUuid := "ingress-" + appUuid //generate name for the service
	//create ingress to expose the service
	_, err = srv.kubernetesRepository.CreateIngress(namespaceUuid, ingressUuid, "app-"+appUuid)
	if err != nil {
		//delete namespace
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUuid)
		return nil, err
	}

	//TODO: save on the database the name of the namespace and app

	//TODO: return link to app
	return deploymentUuid, err
}

func (srv *Service) DeleteDeployment(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.DeleteDeployment(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
