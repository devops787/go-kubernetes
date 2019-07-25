# go-kubernetes
Basic Go application with Gorialla/mux, Docker and Kubernetes.

## Installation
1. Install Docker (https://docs.docker.com/install/linux/docker-ce/ubuntu).
2. Install Kubectl and Minikube (https://kubernetes.io/docs/tasks/tools/install-minikube).
3. Clone Github repository: `git clone https://github.com/devops787/go-kubernetes.git`
4. Ensure that you have installed kubectl:
```
$ kubectl version
Client Version: version.Info{Major:"1", Minor:"15", GitVersion:"v1.15.1", GitCommit:"4485c6f18cee9a5d3c3b4e523bd27972b1b53892", GitTreeState:"clean", BuildDate:"2019-07-18T09:18:22Z", GoVersion:"go1.12.5", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"15", GitVersion:"v1.15.0", GitCommit:"e8462b5b5dc2584fdcd18e6bcfe9f1e4d970a529", GitTreeState:"clean", BuildDate:"2019-06-19T16:32:14Z", GoVersion:"go1.12.5", Compiler:"gc", Platform:"linux/amd64"}
```
5. Start Minikube: 
```
$ minikube start
😄  minikube v1.2.0 on linux (amd64)
💡  Tip: Use 'minikube start -p <name>' to create a new cluster, or 'minikube delete' to delete this one.
🔄  Restarting existing virtualbox VM for "minikube" ...
⌛  Waiting for SSH access ...
🐳  Configuring environment for Kubernetes v1.15.0 on Docker 18.09.6
🔄  Relaunching Kubernetes v1.15.0 using kubeadm ... 
⌛  Verifying: apiserver proxy etcd scheduler controller dns
🏄  Done! kubectl is now configured to use "minikube"
```
6. Apply backend deployment:
```
$ kubectl apply -f ./k8s/deployments/backend.yml
$ kubectl get deployments
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
backend   3/3     3            3           25m

$ kubectl get pods
NAME                       READY   STATUS    RESTARTS   AGE
backend-6c95dcfc5f-n6wwh   1/1     Running   1          26m
backend-6c95dcfc5f-thrsb   1/1     Running   1          26m
backend-6c95dcfc5f-vc8m8   1/1     Running   1          26m
```
7. Apply load-balancer service:
```
$ kubectl apply -f ./k8s/services/load-balancer.yml
$ kubectl get services
NAME            TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes      ClusterIP      10.96.0.1       <none>        443/TCP        110m
load-balancer   LoadBalancer   10.108.59.126   <pending>     80:32570/TCP   21m
```
7. Expose Minikube service url (optional):
```
$ minikube service load-balancer --url
http://192.168.99.100:32570
```
8. Run Kubernetes dashboard (optional):
```
$ minikube dashboard
```

### Example
```
$ curl -i http://192.168.99.100:32570
HTTP/1.1 200 OK
Date: Thu, 25 Jul 2019 10:05:47 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

Hello Kubernetes!
```