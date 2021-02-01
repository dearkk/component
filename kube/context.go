package kube

import (
	"component/klog"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	client "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Context struct {
	kube       *client.Clientset
	ist        *versionedclient.Clientset
	prometheus *monitoringv1.MonitoringV1Client
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
	context.kube, err = client.NewForConfig(config)
	if err != nil {
		klog.Errorf(err.Error())
	} else {
		klog.Infof("init k8s ok.")
	}

	context.ist, err = versionedclient.NewForConfig(config)
	if err != nil {
		klog.Errorf("Failed to create istio client: %s", err)
	} else {
		klog.Infof("init istio ok.")
	}
	context.prometheus, err = monitoringv1.NewForConfig(config)
	if err != nil {
		klog.Errorf(err.Error())
	} else {
		klog.Infof("init prometheus ok.")
	}
	return &context
}
