package kube

import (
	"component/klog"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	tektoncd "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	client "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Context struct {
	Kube       *client.Clientset
	Ist        *versionedclient.Clientset
	Prometheus *monitoringv1.MonitoringV1Client
	Tekton     *tektoncd.Clientset
}

func NewContext(kubeconfigPath string) *Context {
	var err error
	var config *rest.Config
	var context Context
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.Panicf(err.Error())
		}
		klog.Infof("init from cluster config\n")
	} else {
		klog.Infof("init from file config: %s\n", kubeconfigPath)
	}
	context.Kube, err = client.NewForConfig(config)
	if err != nil {
		klog.Errorf(err.Error())
	} else {
		klog.Infof("init k8s ok.")
	}

	context.Ist, err = versionedclient.NewForConfig(config)
	if err != nil {
		klog.Errorf("Failed to create istio client: %s", err)
	} else {
		klog.Infof("init istio ok.")
	}
	context.Prometheus, err = monitoringv1.NewForConfig(config)
	if err != nil {
		klog.Errorf(err.Error())
	} else {
		klog.Infof("init prometheus ok.")
	}
	context.Tekton, err = tektoncd.NewForConfig(config)
	if err != nil {
		klog.Errorf(err.Error())
	} else {
		klog.Infof("init tektoncd ok.")
	}
	return &context
}
