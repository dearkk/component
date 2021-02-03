module component

go 1.15

require (
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.45.0
	github.com/sirupsen/logrus v1.6.0
	github.com/tektoncd/pipeline v0.20.1
	istio.io/client-go v1.8.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
//k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
)

// Pin k8s deps to 0.18.9
replace (
	k8s.io/api => k8s.io/api v0.18.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.9
	//k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.9
	k8s.io/client-go => k8s.io/client-go v0.18.9
)

replace github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.45.0 => ./prometheus-operator-0.42.1/pkg/apis/monitoring
