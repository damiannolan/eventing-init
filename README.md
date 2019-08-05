# eventing-init

A [Kubernetes Init Container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) as a wait mechanism for Kafka availability and topic configuration.

This project serves as a means to block service startup until Kafka eventing and topic dependencies have been fulfilled.

## Usage

The following example assumes the use of [Helm](https://helm.sh/) (A Package manager for Kubernetes), however this is not a requirement and appropriate K8s manifests can be constructed without the use of Helm's templating functionality.

1. A ConfigMap should be provided with the required Kafka Topics.
For Example:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: service-x-topics-config
data:
  topics.yml: |
    topics:
    {{- range .Values.kafka.topics }}
      - {{ . | quote }}
    {{- end }}
```

2. Topic entries should be supplied through `values.yaml`
For Example:

```yaml
...
kafka:
  topics:
    - notification-events
    - request-events
```

3. An `initContainers` entry should be added to the service Deployment PodSpec with a `volumeMounts` entry specifying the associated ConfigMap. Environment variables can be provided in the Container spec, if omitted the default configuration fallbacks are used for Kafka host, port and version.

For Example:

```yaml
spec:
    initContainers:
    - name: eventing-init
      image: damiannolan/eventing-init:0.1.0
      imagePullPolicy: Always
      volumeMounts:
      - name: topics-config
        mountPath: /etc/config
    ...
    volumes:
    - name: topics-config
      configMap:
      name: sevice-x-topics-config
```

### TODO's

- Cleanup config pkg - possibly use [viper](https://github.com/spf13/viper)
- Add more configurability (Create topics if not exists option)
- Use [logrus](https://github.com/sirupsen/logrus) for logging
- Use [testify](https://github.com/stretchr/testify) for assertions