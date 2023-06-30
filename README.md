# k8s-tpm-device
Kubernetes [device plugin](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/) 
based on [intel-device-plugin-for-kubernetes](https://github.com/intel/intel-device-plugins-for-kubernetes)

## About
The device plugin runs as a `DaemonSet` to register a TPM with the kubelet.

## Usage
To install
```shell
helm repo add k8s-tpm-device https://boxboat.github.io/k8s-tpm-device/chart
helm repo update
helm upgrade install k8s-tpm-device --namespace tpm-device --create-namespace k8s-tpm-device/k8s-tpm-device 
```

Add this resource limit to grant the desired container access to `/dev/tpmrm0` 
```yaml
resources:
  limits:
    tpm.boxboat.io/tpmrm: '1'
```
