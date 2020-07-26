$ helm install onlydev haproxytech/kubernetes-ingress \
  --set-string "controller.extraArgs={--namespace-whitelist=dev}"
