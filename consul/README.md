# env.consul
# https://releases.hashicorp.com/consul-k8s/0.17.0/
# Environment variables cannot be used to configure the Consul client. They can be used when running other consul CLI commands that connect with a running agent, e.g. CONSUL_HTTP_ADDR=192.168.0.1:8500 consul members
CONSUL_HTTP_ADDR=192.168.0.1:8500 consul members

CONSUL_HTTP_ADDR=127.0.0.1:8500
CONSUL_HTTP_SSL
CONSUL_HTTP_TOKEN=aba7cbe5-879b-999a-07cc-2efd9ac0ffe
CONSUL_HTTP_TOKEN_FILE=/path/to/consul.token
CONSUL_HTTP_AUTH=operations:JPIMCmhDHzTukgO6
CONSUL_HTTP_SSL=true
CONSUL_CACERT=ca.crt
CONSUL_CAPATH=ca_certs/
CONSUL_CLIENT_CERT=client.crt
CONSUL_CLIENT_KEY=client.key
CONSUL_TLS_SERVER_NAME=consulserver.domain
CONSUL_GRPC_ADDR=127.0.0.1:8502
CONSUL_GRPC_ADDR=unix://var/run/consul_grpc.sock
CONSUL_NAMESPACE=consul.freighttrust

# development
# CONSUL_HTTP_SSL_VERIFY=false

# CONSUL_HTTP_ADDR=unix:///var/run/consul_http.sock
