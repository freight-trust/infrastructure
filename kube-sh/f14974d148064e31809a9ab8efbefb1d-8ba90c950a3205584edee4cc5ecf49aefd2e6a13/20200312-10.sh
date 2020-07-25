$ kubectl config set-credentials bob --client-certificate=bob.crt --client-key=bob.key

$ kubectl config set-context minikube-bob --cluster=minikube --user=bob

$ kubectl config  use-context minikube-bob
