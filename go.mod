module github.com/boxboat/k8s-tpm-device

go 1.20

require (
	github.com/intel/intel-device-plugins-for-kubernetes v0.27.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.7.0
	k8s.io/kubelet v1.27.3
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
)

// beause the intel-device-plugins-for-kubernetes insists on using their own versions instead of k8s.io versions
replace (
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.27.3
	k8s.io/kubectl => k8s.io/kubectl v0.27.3
	k8s.io/kubelet => k8s.io/kubelet v0.27.3
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.27.3
)
