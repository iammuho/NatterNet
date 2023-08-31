#!/bin/bash

# Check if the script is running as root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

# Generate the RSA private key
openssl genrsa -out /etc/ssl/certs/natternet.private.pem 2048

# Generate the RSA public key from the private key
openssl rsa -in /etc/ssl/certs/natternet.private.pem -pubout -out /etc/ssl/certs/natternet.public.pem

# Set read permissions for all users
chmod 644 /etc/ssl/certs/natternet.private.pem
chmod 644 /etc/ssl/certs/natternet.public.pem

# Provide feedback that the keys were created
echo "RSA keys generated and placed in /etc/ssl/certs/"
