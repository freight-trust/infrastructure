

## Toolchain

[MTA-STS](https://aykevl.nl/apps/mta-sts/)
[dmarc generator](https://tools.socketlabs.com/dmarc/generator)

## DMARC

[dmarc generator](https://tools.socketlabs.com/dmarc/generator)


`v=DMARC1; p=quarantine; ruf=mailto:reports@freighttrust.com; rua=mailto:reports@freighttrust.com;`

Forensic Reporting Address
Aggregate Reporting Address



## DKIM
Domain
freighttrust.com
Enter the domain that you will send email from.
Key Selector
dkim
Enter the key selector for your public key.

Escape records 

Split record 
Key Length: 2048



## MTA-STS Tool

MTA-STS (RFC8461) is a new standard that makes it possible to send downgrade-resistant email over SMTP. In that sense, it is like an alternative to DANE. It does this by piggybacking on the browser Certificate Authority model. This validator checks whether a domain adheres to the RFC. An alternative validator is Hardenize, which checks for much more than just MTA-STS

To enable Strict Transport Security on your mailserver configure the following things:

Check that your host has properly configured STARTTLS
Add a TLSRPT DNS TXT record at _smtp._tls on your domain, e.g. _smtp._tls.example.com, with something like v=TLSRPTv1; rua=mailto:mta-sts@example.com.
Add a MTA-STS DNS TXT record at _mta-sts on your domain, e.g. _mta-sts.example.com, with something like v=STSv1; id=20160831085700Z.
Add a subdomain mta-sts to your domain (note the lack of an underscore) and serve a policy file on https://mta-sts.example.com/.well-known/mta-sts.txt. Here is an example policy file:
version: STSv1
mode: enforce
max_age: 10368000
mx: mail.example.com
mx: *.example.net
mx: backupmx.example.com



## Remidiation 


### MTA-STS 

MTA-STS Verification
This tool will verify that MTA-STS is configured correctly for your domain. The tool will lookup the policy for your domain and connect to your email server to verify it will pass the published policy.

freighttrust.com
freighttrust.com
Policy Discovery Record
A valid DNS TXT resource record could not be found for the specified domain.

Policy Enforced
No policy is found so the policy mode will be treated as none. Senders that honor MTA-STS will ignore policy failures.

Policy Result
No MTA-STS policy could be found for the specified domain.



### SPF


v=spf1?all	

[dmarc survey](https://dmarcian.com/spf-survey/)



```text
4 +include:_spf.google.com
v=spf1 include:_netblocks.google.com include:_netblocks2.google.com include:_netblocks3.google.com ~all
6 +include:_spf.firebasemail.com
v=spf1 include:sendgrid.net include:_spf.google.com ~all
1 +include:sendgrid.net
v=spf1 ip4:167.89.0.0/17 ip4:208.117.48.0/20 ip4:50.31.32.0/19 ip4:198.37.144.0/20 ip4:198.21.0.0/21 ip4:192.254.112.0/20 ip4:168.245.0.0/17 ip4:149.72.0.0/16 ~all
4 +include:_spf.google.com
v=spf1 include:_netblocks.google.com include:_netblocks2.google.com include:_netblocks3.google.com ~all
Warning! The target name for "include:_spf.google.com" equals an already evaluated "include" mechanism / "redirect" modifier.
```

