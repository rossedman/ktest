# Notes

### Prerequisites

Setup a `kind` cluster locally

```sh
export KUBECONFIG=$HOME/.kube/otk-local
kind create cluster --name testkube
kubectl cluster-info --context kind-testkube
brew install testkube
```

### Installing

Install `testkube` globally. The simplest way to do this is `helm install --create-namespace testkube kubeshop/testkube` but this will install `testkube` into the `default` namespace on a cluster. 

```sh
# install
helm repo add kubeshop https://kubeshop.github.io/helm-charts
helm install \
	--namespace twilio-system-testkube  \
	--create-namespace \
		testkube kubeshop/testkube
```

Once complete let's open a local port and allow communication to the testkube api server

```sh
# open direct port locally
kubectl port-forward \
	--namespace twilio-system-testkube \
		svc/testkube-api-server 8088:8088

# check status
kubectl testkube --client direct status

🚀  Getting status on the testkube CLI
🥇  Telemetry on CLI enabled
Handling connection for 8088
🥇  Telemetry on API enabled
❌  Oauth disabled
```

### Testing 

Run and execute a ginkgo test suite that has global RBAC perms to the cluster 

```sh
# create testkube role
kubectl apply -f ginkgo-rbac.yaml

# create ginkgo test suite
kubectl testkube create test \
	--git-uri https://github.com/rossedman/ktest.git \
	--type ginkgo/test \
	--name ktest \
	--git-branch main \
	--job-template ginkgo-job.yaml \
	--client direct

# run ginkgo test suite
kubectl testkube run test ktest -f --client direct

# show tests 
kubectl testkube get test --client --direct

# check results
kubectl testkube get execution 6632820c75f1a393d3900046 --client direct
```

### Gingko

Setting up our test suite

```
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
ginkgo bootstrap
```

Followed this document [here](https://docs.testkube.io/test-types/executor-ginkgo)