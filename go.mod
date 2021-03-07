module component

go 1.15

require (
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.45.0
	github.com/sirupsen/logrus v1.6.0
	github.com/tektoncd/pipeline v0.20.1
	istio.io/client-go v1.8.2
	github.com/emicklei/go-restful v2.9.6+incompatible
    github.com/emicklei/go-restful-openapi v1.4.1
    github.com/sirupsen/logrus v1.6.0
    gopkg.in/yaml.v2 v2.3.0
    gorm.io/driver/mysql v1.0.4
    gorm.io/gorm v1.21.2
)

replace k8s.io/client-go => k8s.io/client-go v0.20.0

replace github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.45.0 => ./kube/prometheus-operator-0.42.1/pkg/apis/monitoring
