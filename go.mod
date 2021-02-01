module component

go 1.15

require (
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.45.0
	github.com/sirupsen/logrus v1.4.2
	istio.io/client-go v1.8.2
	k8s.io/client-go v0.19.2
)

replace (
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.45.0 => ./prometheus-operator-0.42.1/pkg/apis/monitoring
)
