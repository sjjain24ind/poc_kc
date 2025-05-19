#!/bin/sh
# Directory for certificates
CERT_DIR="/opt/keycloak/certs"
mkdir -p $CERT_DIR
# IMPORT_REALM_DIR="/opt/keycloak/data/import"
# mkdir -p $IMPORT_REALM_DIR


# Generate self-signed certificates
openssl req -x509 -newkey rsa:4096 -keyout $CERT_DIR/tls.key -out $CERT_DIR/tls.crt -days 365 -nodes -subj "/CN=localhost"

# Set permissions for certificates
chmod 644 $CERT_DIR/tls.crt
chmod 600 $CERT_DIR/tls.key

# Import the realm configuration
if [ -f "/opt/keycloak/data/import/realm_account.json" ]; then
    echo "Importing realm configuration... for Partner accounts"
    /opt/keycloak/bin/kc.sh import --file /opt/keycloak/data/import/realm_account.json
else
    echo "Realm configuration file not found. Skipping import."
fi


if [ -f "/opt/keycloak/data/import/realm_endusers.json" ]; then
    echo "Importing realm configuration...for end users "
    /opt/keycloak/bin/kc.sh import --file /opt/keycloak/data/import/realm_endusers.json
else
    echo "Realm configuration realm_endusers.json file not found. Skipping import."
fi

# Start Keycloak
exec /opt/keycloak/bin/kc.sh start-dev --hostname-strict=false

