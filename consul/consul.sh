#!/bin/bash -e

curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -

sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"

sudo apt-get update && sudo apt-get install consul

 export CONSUL_VERSION="1.8.0"
 export CONSUL_URL="https://releases.hashicorp.com/consul"

 curl --silent --remote-name \
  ${CONSUL_URL}/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_linux_amd64.zip

curl --silent --remote-name \
  ${CONSUL_URL}/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_SHA256SUMS

curl --silent --remote-name \
  ${CONSUL_URL}/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_SHA256SUMS.sig

unzip consul_${CONSUL_VERSION}_linux_amd64.zip

sudo chown root:root consul

sudo mv consul /usr/local/bin/

consul -autocomplete-install

complete -C /usr/local/bin/consul consul

sudo useradd --system --home /etc/consul.d --shell /bin/false consul

sudo mkdir --parents /opt/consul

sudo chown --recursive consul:consul /opt/consul

consul keygen

consul tls ca create

# where <dc_name> is data center name/region, etc
consul tls cert create -server -dc <dc_name>

# where <dc_name> is data center name/region, etc
consul tls cert create -client -dc <dc_name>

consul tls cert create -client -dc dc1

# Below is an SCP example which will send the CA certificate, agent certificate and private key to the IP address you specify, and put it into the /etc/consul.d/ directory.
scp consul-agent-ca.pem <dc-name>-<server/client>-consul-<cert-number>.pem <dc-name>-<server/client>-consul-<cert-number>-key.pem <USER>@<PUBLIC_IP>:/etc/consul.d/