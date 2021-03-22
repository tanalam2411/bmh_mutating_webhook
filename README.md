# bmh_mutating_webhook
Baremetal Mutating Webhook



```bash
$ make docker-build

$ make docker-push
```

- Create minikube k8s cluster
```bash
 minikube start --extra-config=apiserver.enable-admission-plugins=MutatingAdmissionWebhook,ValidatingAdmissionWebhook
```

- Apply webhook.yml
```bash
bmh_mutating_webhook$ k apply -f webhook.yaml 
mutatingwebhookconfiguration.admissionregistration.k8s.io/mymutatingwebhook.example.com created
namespace/mutatingwebhook created
namespace/testmutatingwebhook created
deployment.apps/mutatingwebhook created
service/mutatingwebhook created
```

- Get certificate
```bash
$ k get po -A | grep -i mutating
mutatingwebhook   mutatingwebhook-7f5497df59-szfzk   1/1     Running   0          20m

$ kubectl exec -it -n mutatingwebhook $(kubectl get pods --no-headers -o custom-columns=":metadata.name" -n mutatingwebhook) -- wget -q -O- localhost:8080/ca.pem?base64

LS0tLS1CRUdJT...0tCg==
```

- Update the webhook.yml's `caBundle` with new certificate

- Again apply webhook.yml
```bash
bmh_mutating_webhook$ k apply -f webhook.yaml 
mutatingwebhookconfiguration.admissionregistration.k8s.io/mymutatingwebhook.example.com configured
namespace/mutatingwebhook unchanged
namespace/testmutatingwebhook unchanged
deployment.apps/mutatingwebhook unchanged
service/mutatingwebhook unchanged
```

- Tail logs of mutatingwebhook POD
```bash
$ k logs -f mutatingwebhook-7f5497df59-szfzk -n mutatingwebhook
```

- Create or Update BareMetalHost resource and check logs
```bash
bmh_mutating_webhook$ k create -f bmh.yaml
```


- logs
```bash
debug: {UID:b1489b1a-69e8-48bf-8bec-6cf0c16dd254 Kind:{Group:metal3.io Version:v1alpha1 Kind:BareMetalHost} Resource:{Group:metal3.io Version:v1alpha1 Resource:baremetalhosts} SubResource: RequestKind:{Group:metal3.io Version:v1alpha1 Kind:BareMetalHost} RequestResource:{Group:metal3.io Version:v1alpha1 Resource:baremetalhosts} RequestSubResource: Name:example-baremetalhost Namespace:testmutatingwebhook Operation:CREATE UserInfo:{Username:minikube-user UID: Groups:[system:masters system:authenticated] Extra:{SomeKey:[]}} Object:{TypeMeta:{Kind:BareMetalHost APIVersion:metal3.io/v1alpha1} ObjectMeta:{Name:example-baremetalhost GenerateName: Namespace:testmutatingwebhook SelfLink: UID: ResourceVersion: Generation:0 CreationTimestamp:0001-01-01 00:00:00 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[{Manager:kubectl-create Operation:Update APIVersion:metal3.io/v1alpha1 Time:2021-03-22 11:06:26 +0000 UTC FieldsType:FieldsV1 FieldsV1:{"f:spec":{".":{},"f:bmc":{".":{},"f:address":{},"f:credentialsName":{}},"f:online":{}}}}]} Spec:{Taints:[] BMC:{Address:ipmi://192.168.122.1:6233 CredentialsName:example-baremetalhost-secret DisableCertificateVerification:false} RAID:<nil> HardwareProfile: RootDeviceHints:<nil> BootMode: BootMACAddress: Online:true ConsumerRef:nil Image:<nil> UserData:nil NetworkData:nil MetaData:nil Description: ExternallyProvisioned:false} Status:{OperationalStatus: ErrorType: LastUpdated:<nil> HardwareProfile: HardwareDetails:<nil> Provisioning:{State: ID: Image:{URL: Checksum: ChecksumType: DiskFormat:<nil>} RootDeviceHints:<nil> BootMode: RAID:<nil>} GoodCredentials:{Reference:nil Version:} TriedCredentials:{Reference:nil Version:} ErrorMessage: PoweredOn:false OperationHistory:{Register:{Start:0001-01-01 00:00:00 +0000 UTC End:0001-01-01 00:00:00 +0000 UTC} Inspect:{Start:0001-01-01 00:00:00 +0000 UTC End:0001-01-01 00:00:00 +0000 UTC} Provision:{Start:0001-01-01 00:00:00 +0000 UTC End:0001-01-01 00:00:00 +0000 UTC} Deprovision:{Start:0001-01-01 00:00:00 +0000 UTC End:0001-01-01 00:00:00 +0000 UTC}} ErrorCount:0}} OldObject:{APIVersion: Kind:} Options:{APIVersion:meta.k8s.io/v1 Kind:CreateOptions} DryRun:false}

```