# eventing-init

A [Kubernetes Init Container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) as a wait mechanism for Kafka availability and topic configuration.

## Usage

A ConfigMap should be provided with the required Kafka Topics.
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

Topic entries should be supplied through `values.yaml`
For Example:

```yaml
...
kafka:
  topics:
    - notification-events
    - request-events
```

An `initContainers` entry should be added to the service Deployment PodSpec with a `volumeMounts` entry specifying the associated ConfigMap.  
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
