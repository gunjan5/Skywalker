# ![Skywalker](https://raw.githubusercontent.com/gunjan5/Skywalker/master/Skywalker_logo_70.png) Skywalker

### :cloud: Platform to help launch Docker, Kubernetes, Logging, DB, LB
------------------------------------------------------------------------
#### *One click to deploy pods with <span>prewired</span> components*
#### *Things <span>just work</span>* 

## Components
- [ ] Kubernetes
- [ ] UI
- [ ] Predefined environment 1 (Golang 1.5, Ubuntu 15.10)
- [ ] Predefined environment 2 (Python 2.7.5, Alpine)
- [ ] Predefined environment 3 (Java 7, Fedora 22)
- [ ] Predefined environment 4 (Java 8, Ubuntu 14.04)
- [ ] Elasticsearch
- [ ] Logstash
- [ ] Kibana
- [ ] Nginx
- [ ] HA Proxy
- [ ] Logging integration (with ELK)
- [ ] Redis
- [ ] PostgreSQL
- [ ] MySQL
- [ ] Cassandra
- [ ] CLI in web browser (with gotty)
- [ ] MEssagebus/etcd/Kafka?
- [ ] Loadbalancer (LBaaS)
- [ ] Service Discovery

## How To Start Skywalker
- Clone the project: `git clone https://github.com/gunjan5/Skywalker.git`
- `cd Skywalker`
- Build the image: `make build`
- Start the container /w privileged mode: ` make run`
- (optional) get into the container shell: `make cli` 
- Cleanup: `make clean`


### Up all night to get lucky!
-------------------------------
Like the legend of the Phoenix All ends with beginnings What keeps the planets spinning (uh) The force from the beginning We've come too far to give up who we are So let's raise the bar and our cups to the stars


## Scratch

KUBE_STATIC_OVERRIDES=kubectl hack/build-go.sh cmd/kubectl


/ # find . -name kubectl
./kubernetes/docs/user-guide/kubectl
./kubernetes/_output/local/go/pkg/linux_amd64_cgo/k8s.io/kubernetes/pkg/kubectl
./kubernetes/_output/local/go/bin/kubectl
./kubernetes/_output/local/bin/linux/amd64/kubectl
./kubernetes/pkg/kubectl
./kubernetes/contrib/completions/bash/kubectl
./kubernetes/cmd/kubectl
/ # cd /kubernetes/_output/local/bin/linux/amd64/
/kubernetes/_output/local/bin/linux/amd64 # ls -la
total 21044
drwxr-xr-x    2 root     root          4096 Sep 23 07:52 .
drwxr-xr-x    3 root     root          4096 Sep 23 07:52 ..
-rwxr-xr-x    1 root     root      21536880 Sep 23 07:52 kubectl
/kubernetes/_output/local/bin/linux/amd64 # ./kubectl 
kubectl controls the Kubernetes cluster manager.

Find more information at https://github.com/kubernetes/kubernetes.

Usage: 
  kubectl [flags]
  kubectl [command]