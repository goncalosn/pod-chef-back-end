package kubernetes

import (
	uuid "github.com/satori/go.uuid"
)

//GetDeploymentsByUser service responsible for getting all deployments inside a namespace
func (srv *Service) GetDeploymentsByUser(token string) (interface{}, error) {
	//TODO: get user namespace
	//TODO: for each all namespaces from the data base

	//call driven adapter responsible for getting a deployment from the kubernetes cluster
	response, err := srv.kubernetesRepository.GetDeploymentByNameAndNamespace("name", "default")

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//GetDeploymentByUserAndName service responsible for getting a deployment inside a user's namespace
func (srv *Service) GetDeploymentByUserAndName(token string, name string) (interface{}, error) {
	//TODO: get user namespace

	//call driven adapter responsible for getting a deployment from the kubernetes cluster
	response, err := srv.kubernetesRepository.GetDeploymentByNameAndNamespace("name", "default")

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//CreateDeployment service responsible for creating a deployment inside a new namespace
func (srv *Service) CreateDeployment(token string, replicas *int32, image string) (interface{}, error) {

	appUUID := uuid.NewV4().String() //generate uuid for this deployment

	namespaceUUID := "namespace-" + appUUID //generate name for namespace

	//call driven adapter responsible for creating namespaces inside the kubernetes cluster
	_, err := srv.kubernetesRepository.CreateNamespace(namespaceUUID)
	if err != nil {
		//return error from the kubernetes repository method
		return nil, err
	}

	deploymentUUID := "deployment-" + appUUID //generate name for the deployment

	//call driven adapter responsible for creating deployments inside the kubernetes cluster
	_, err = srv.kubernetesRepository.CreateDeployment(namespaceUUID, deploymentUUID, replicas, image)
	if err != nil {
		//creation of the deployment went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	serviceUUID := "service-" + appUUID //generate name for the service
	//create service to expose the deployment
	_, err = srv.kubernetesRepository.CreateClusterIPService(namespaceUUID, serviceUUID)
	if err != nil {
		//creation of the service went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	ingressUUID := "ingress-" + appUUID //generate name for the service
	//create ingress to expose the service
	_, err = srv.kubernetesRepository.CreateIngress(namespaceUUID, ingressUUID, "app-"+appUUID)
	if err != nil {
		//creation of the ingress went wrong, delete everything inside it's namespace
		//call driven adapter responsible for deleting namespaces inside the kubernetes cluster
		_, _ = srv.kubernetesRepository.DeleteNamespace(namespaceUUID)

		//return error from the kubernetes repository method
		return nil, err
	}

	//TODO: save on the database the name of the namespace and app

	//TODO: return link to app
	return deploymentUUID, err
}
