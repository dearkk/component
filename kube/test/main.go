package main

import (
	"component/klog"
	"component/kube"
	"context"
	"encoding/json"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	c := kube.NewContext("D:\\bean\\config")

	d, _ := c.Kube.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	klog.Infof("Deploymentslll: %+v", d.Items[0])
	//m := make(map[string]interface{})
	bys, _ := json.Marshal(d.Items[0])
	//klog.Infof("Marshal JSON: %+v", err)
	var item interface{}
	yaml.Unmarshal(bys, &item)

	klog.Infof("item: %+v", item)
}

// Convert JSON to YAML.
func JSONToYAML(j []byte) ([]byte, error) {
	// Convert the JSON to an object.
	var jsonObj interface{}
	// We are using yaml.Unmarshal here (instead of json.Unmarshal) because the
	// Go JSON library doesn't try to pick the right number type (int, float,
	// etc.) when unmarshalling to interface{}, it just picks float64
	// universally. go-yaml does go through the effort of picking the right
	// number type, so we can preserve number type throughout this process.
	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}

	// Marshal this object into YAML.
	return yaml.Marshal(jsonObj)
}
