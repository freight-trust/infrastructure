# SPF Records

* Toolchain: [easydmarc](https://easydmarc.com/tools/spf-record-generator/freighttrust.com)

## Old SPF

v=spf1 a include:"v=spf1 include:mx include:a include:ptr include:include:google.com,servers.mcsv.net include:~all" ~all

TXT	freighttrust.com	"v=spf1 mx a ptr include:google.com,servers.mcsv.net ~all"	


 SMTP AUTH or from the HELO or EHLO command.



## New SPF 
 v=spf1 include:_spf.google.com ~all

v=spf1 include:_spf.servers.mcsv.net ~all


# Fixed

v=spf1 mx a include:google.com include:servers.mcsv.net ~all
