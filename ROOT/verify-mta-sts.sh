sudo certbot --apache -d mta-sts.your-domain
curl https://mta-sts.your-domain/.well-known/mta-sts.txt