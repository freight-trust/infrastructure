$ kubectl get pods --namespace=dev
NAME                                       READY   STATUS    RESTARTS   AGE
app-66d9457bf5-vpbnw                       1/1     Running   1          22h

$ kubectl get pods --namespace=default
Error from server (Forbidden): pods is forbidden: User "bob" cannot list resource "pods" in API group "" in the namespace "default"
