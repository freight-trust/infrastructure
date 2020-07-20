# nginx setup

#### Generate Diffie-Hellman keys by running this command on your server:
openssl dhparam -out /etc/nginx/dhparam.pem 2048

#### Create a common ACME-challenge directory (for Let's Encrypt):
mkdir -p /var/www/_letsencrypt
chown www-data /var/www/_letsencrypt

#### Comment out SSL related directives in the configuration:

sed -i -r 's/(listen .*443)/\1;#/g; s/(ssl_(certificate|certificate_key|trusted_certificate) )/#;#\1/g' /etc/nginx/sites-available/freight-trust.com.conf

#### Reload your NGINX server:

sudo nginx -t && sudo systemctl reload nginx
##### Obtain SSL certificates from Let's Encrypt using Certbot (SSLMATE or ZeroSSL)

certbot certonly --webroot -d freight-trust.com --email admin@freighttrust.com -w /var/www/_letsencrypt -n --agree-tos --force-renewal

#### Uncomment SSL related directives in the configuration:

sed -i -r 's/#?;#//g' /etc/nginx/sites-available/freight-trust.com.conf
#### Reload your NGINX server:

sudo nginx -t && sudo systemctl reload nginx

#### Configure Certbot to reload NGINX when it successfully renews certificates:

echo -e '#!/bin/bash\nnginx -t && systemctl reload nginx' | sudo tee /etc/letsencrypt/renewal-hooks/post/nginx-reload.sh
sudo chmod a+x /etc/letsencrypt/renewal-hooks/post/nginx-reload.sh

Reload NGINX to load in your new configuration:
sudo nginx -t && sudo systemctl reload nginx

