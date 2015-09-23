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
- Build the image: `sudo docker build -t skywalker .`
- Start the container /w privileged mode: ` sudo docker run -it --privileged --name sky1 -d skywalker`
- (optional) get into the container shell: `sudo docker exec -it sky1 sh` 


### Up all night to get lucky!
-------------------------------
Like the legend of the Phoenix All ends with beginnings What keeps the planets spinning (uh) The force from the beginning We've come too far to give up who we are So let's raise the bar and our cups to the stars


## Scratch

KUBE_STATIC_OVERRIDES=kubectl hack/build-go.sh cmd/kubectl


