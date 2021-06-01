package kubernetes

import (
	"flag"
	"path/filepath"

	"github.com/labstack/gommon/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

//KubernetesRepository kubernetes clientset responsible for acessing the cluster api
type KubernetesRepository struct {
	Clientset *kubernetes.Clientset
}

//NewKubernetesRepository new connection to the cluster
func NewKubernetesRepository() *KubernetesRepository {
	return &KubernetesRepository{
		Clientset: Client(),
	}
}

//Client responsible for creating the connection to the cluster
func Client() *kubernetes.Clientset {
	var kubeconfig *string

	log.Info("creating flag for kubeconfig")
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	log.Info("getting kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	// create the clientset
	log.Info("creating clientset")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return clientset
}
