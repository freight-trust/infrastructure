# Create a certificate signing request with CN=bob and O=dev-group
# This creates bob.csr and bob.key
$ openssl req -newkey rsa:2048 -nodes -keyout bob.key -out bob.csr -subj "/CN=bob/O=dev-group"