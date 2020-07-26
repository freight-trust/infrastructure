$ helm install intranet haproxytech/kubernetes-ingress \
  --set controller.ingressClass=intranet
