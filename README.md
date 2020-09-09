<p align="center">
  <img src="https://raw.githubusercontent.com/freight-trust/branding/master/static/logo-420.svg" alt="Freight Trust logo">
</p>

<h3 align="center">
  Freight Trust
</h3>

<p align="center">
  Frieght Trust & Clearing Corporation
</p>

<p align="center">
  <a href="https://www.npmjs.org/package/Freight Trust-lib">
    <img src="https://img.shields.io/github/license/freight-trust/omnibus" alt="MPL-2.0 License" />
  </a>
  <a href="https://circleci.com/gh/Freight Trust/Freight Trust-lib">
    <img src="https://img.shields.io/github/commits-since/freight-trust/omnibus/latest/master" alt="GitHub commits since latest release" />
  </a>
  <a href="https://circleci.com/gh/Freight Trust/Freight Trust-lib">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs welcome!" />
  </a>
  <a href="https://twitter.com/intent/follow?screen_name=Freight Trust">
    <img src="https://img.shields.io/twitter/url?label=%40FreightTrustNet&url=https%3A%2F%2Ftwitter.com%2Ffreighttrustnet" alt="Follow @Freight Trust" />
  </a>
</p>

<align=center><a href="https://check-your-website.server-daten.de/?q=freighttrust.com" target="_blank">Check this Site: freighttrust.com</a></align>

<br>

### Audits

See [dns audit](#) in releases


[ref/component](/com/freighttrust) <br />
[policy/component](/com/freighttrust/policy) <br />
[schemas/component](/com/freighttrust/schemas) <br />


Relevant Documentation

[dnsimple reference](https://developer.dnsimple.com/v2/certificates/)
[netlify](https://docs.netlify.com/domains-https/https-ssl/)
[SSL Mate](https://sslmate.com/caa/)


### Structure 
* ROOT
	- [CAA_POLICY](/ROOT/CAA_POLICY)
	- [DNSSEC](/ROOT/DNSSEC)
	- [MX](ROOT/MX)
	- CNMAE
	- DKIM
	- ALIAS
	- DNSZONE

* com
	- freighttrust
		- schemas
		- policy
			- besu
			- network
			- availability
		- service

* security
	- [defects](/security/defects.md)
	- [document retention policy](/security/document-retention-policy.md)
	- [pki](/security/pki.md)

## Known Issues

> SSL 

### Cerbot / Nginx

a `wontfix issue` [https://github.com/certbot/certbot/issues/5764](https://github.com/certbot/certbot/issues/5764)
see this github issue [7275#issue-473726625](https://github.com/certbot/certbot/issues/7275#issue-473726625)
`certbot renew -a webroot -w /var/www/html --dry-run`


### Support or Contact

support@freight.zendesk.com


## License 

MIT
