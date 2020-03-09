### Sriov Crd
base on code-generate, k8s api version is `b799cb063522` (if version different may run incorrectly)

#### Get Start
First register the custom resource definition
```bash
kubectl apply -f yamls/crd.yaml
```

Second import crd
```bash
go get -u github.com/fafucoder/sriov-crd
```