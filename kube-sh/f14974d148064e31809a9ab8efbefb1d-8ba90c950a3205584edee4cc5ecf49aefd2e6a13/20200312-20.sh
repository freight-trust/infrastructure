$ kubectl describe resourcequota dev-resources -n dev
Name:            dev-resources
Namespace:       dev
Resource         Used  Hard
--------         ----  ----
limits.cpu       0     2
limits.memory    0     2Gi
requests.cpu     500m  1
requests.memory  50Mi  1Gi
