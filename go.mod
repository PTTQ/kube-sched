module github.com/pttq/kube-sched

go 1.16

require (
	github.com/kubesys/client-go v1.1.2
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/sys v0.0.0-20210324051608-47abb6519492 // indirect
)

replace k8s.io/kubernetes => k8s.io/kubernetes v0.20.2
