module component

go 1.15

require (
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.45.0
	github.com/sirupsen/logrus v1.6.0
	github.com/tektoncd/pipeline v0.20.1
	istio.io/client-go v1.8.2
	k8s.io/api v0.20.0 // indirect
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.20.0

replace github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.45.0 => ./kube/prometheus-operator-0.42.1/pkg/apis/monitoring
