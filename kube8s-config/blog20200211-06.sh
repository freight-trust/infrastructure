$ kubectl get svc haproxy-ingress -n haproxy-controller
NAME              TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)                                     AGE
haproxy-ingress   NodePort   10.101.75.28   <none>        80:31179/TCP,443:31923/TCP,1024:30430/TCP   98s