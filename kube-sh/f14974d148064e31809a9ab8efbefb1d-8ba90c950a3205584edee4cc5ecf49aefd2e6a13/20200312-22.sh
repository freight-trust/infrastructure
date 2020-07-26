$ cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: bob
spec:
  request: $(cat bob.csr | base64 | tr -d '\n')
  usages:
  - digital signature
EOF
