---
title: "Artifacts Connections"
sub_link: "connections/artifacts"
meta_title: "Connections for your dataset, artifacts, volumes and storage in Polyaxon - Configuration"
meta_description: "Polyaxon allows to use volumes as well as cloud stores for storing outputs and artifacts, and connecting datasets."
tags:
    - configuration
    - polyaxon
    - kubernetes
    - docker-compose
    - environment
    - orchestration
    - volumes
    - s3
    - gcp
    - azure-storage
sidebar: "setup"
---

You can connect as many datasets, volumes, and artifacts stores in Polyaxon.
 
It's better to set a connection for each datasets, or artifacts store, or volume path holding some data, 
to get more visibility and granular control over who is using that connection and how often.

Exposing each datasets or artifacts as a connection also give the possibility to 
effectively version you data, and expose information about the changes from one dataset version to another in the description.
By using connections you can also migrate and find jobs that use a dataset, and take necessary actions.

Using multiple connections is also very useful for large teams who need either to scale or
to have different teams to access different volumes and storage backends.

This section tries to explain how Polyaxon mounts these volumes for experiments and jobs.

## Default behaviour

When the user does not provide any connection, the default behaviour is to use a local path on the host node for storing outputs and logs. 
Oftentimes this default behaviour is sufficient for users who are just trying the platform, and don't want to deal with configuration steps.

## Host paths


### Schema Fields

You can use host paths to define storage connections:

  * hostPath: the host path.
  * mountPath: path where to mount the volume content in the container
  * readOnly: if th volume should be mounted in read only mode.

Users should be aware as well, that by losing the node where the host path is defined, all data will be lost as well.

### Example usage in the default artifactsStore

```yaml
artifactsStore:
  name: my-artifacts-store
  kind: host_path
  schema:
    mountPath: "/outputs/1"
    hostPath: "/path/to/outputs"
```

### Example usage in connections

```yaml
connections:
  ...
  - name: my-artifacts-store
    kind: host_path
    schema:
      mountPath: "/outputs/1"
      hostPath: "/path/to/outputs"
```

## Persistent Volumes

You can use a [Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) to store your outputs and artifacts, or to connect a dataset(s):

### Schema Fields

  * volumeClaim: volume claim name.
  * mountPath: path where to mount the volume content in the container
  * readOnly: if th volume should be mounted in read only mode.

### Example in the default artifactsStore

```yaml
artifactsStore:
  name: my-volume
  kind: volume_claim
  schema:
    mountPath: "/tmp/outputs"
    volumeClaim: "outputs-2-pvc"
```

### Example usage in connections

```yaml
connections:
  ...
  - name: my-volume
    kind: volume_claim
    schema:
      mountPath: "/outputs/1"
      volumeClaim: "outputs-2-pvc"
```

If you are using a persistent volume with one node access you need to be aware that you can only use it with experiment/jobs running on that same node at the same time.

There are some options that support multi-nodes access, e.g. a PVC backed with an [NFS](/integrations/outputs-on-nfs/)/Glusterfs server, 
where you can use multiple nodes and schedule experiments on all the nodes to access the artifacts/datasets. 
Please refer to [this section](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes) to learn more about access modes.

## Cloud stores

In order to mount a cloud storage, 
users need to provide authentication access to Polyaxon for the storage needed during the scheduling.

The way to do that is by creating a secret of your cloud storage access, 
and providing the secret name. 
(You can use the same k8s secret to manage multiple storage access auth).

### Schema Fields

    * bucket: the bucket you want to expose in this connection.


### Example in the default artifactsStore

```yaml
artifactsStore:
  name: azure-bucket
  kind: wasb
  schema: {"bucket": "wasbs://bucket@owner.blob.core.windows.net/"}
  secret:
    name: "az-secret"
```

### Example usage in connections

```yaml
connections:
  ...
  - name: azure-bucket
    kind: wasb
    schema: {"bucket": "wasbs://bucket@owner.blob.core.windows.net/"}
    secret:
      name: "az-secret"
  - name: s3-bucket
    kind: s3
    schema: {"bucket": "s3://bucket/"}
    secret:
      name: "s3-secret"
  - name: gcs-bucket2
    kind: gcs
    schema: {"bucket": "gs://bucket/"}
    secret:
      name: "gcs-secret"
```

Please refer to this integration sections for more details:

 * [Outputs on GCS](/integrations/outputs-on-gcs/)
 * [Outputs on AWS S3](/integrations/outputs-on-s3/)
 * [Outputs on Azure storage](/integrations/outputs-on-azure/)
