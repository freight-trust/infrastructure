# Sign it with the cluster's CA certificate
# This creates bob.crt
$ openssl x509 -req -in bob.csr -CA ~/.minikube/ca.crt -CAkey ~/.minikube/ca.key -CAcreateserial -out bob.crt -days 1000
