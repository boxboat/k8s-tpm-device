# k8s-tpm-device

Helm chart for k8s-tpm-device

## Values

| Key               | Type   | Default                  | Description                          |
|-------------------|--------|--------------------------|--------------------------------------|
| nameOverride      | string | `""`                     |                                      |
| fullnameOverride  | string | `""`                     |                                      |
| imagePullPolicy   | string | `IfNotPresent`           |                                      |
| debug             | bool   | `false`                  | enables debug logging                |
| image.registry    | string | `ghcr.io`                |                                      |
| image.repository  | string | `boxboat/k8s-tpm-device` |                                      |
| image.tag         | string | `master`                 |                                      |
| device.namespace  | string | `tpm.boxboat.io`         | device namespace                     |
| device.capacity   | int    | `1`                      | specifies the tpm capacity           |
| priorityClassName | string | `""`                     |                                      |
| securityContext   | object | `{}`                     | override the default securityContext |
| resources         | object | `{}`                     |                                      |
| tolerations       | object | `{}`                     |                                      |
| extraVolumes      | list   | `[]`                     |                                      |
| extraVolumeMounts | list   | `[]`                     |                                      |
| extraContainers   | list   | `[]`                     |                                      |
| initContainers    | list   | `[]`                     |                                      |

## TPM Device Usage

Add this resource limit to grant the desired container access to `/dev/tpmrm0`

```yaml
resources:
  limits:
    tpm.boxboat.io/tpmrm: '1'
```
