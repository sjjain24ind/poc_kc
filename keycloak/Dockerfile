# # Stage 1: Install OpenSSL in a separate image
# FROM alpine:latest AS openssl-builder

# # Install OpenSSL
# RUN apk add --no-cache openssl

# # Generate self-signed certificates
# RUN mkdir -p /certs && \
#     openssl req -x509 -newkey rsa:4096 -keyout /certs/tls.key -out /certs/tls.crt -days 365 -nodes -subj "/CN=localhost"

# # Stage 2: Use the Keycloak image
# FROM quay.io/keycloak/keycloak:latest

# # Copy the certificates from the openssl-builder stage
# COPY --from=openssl-builder /certs /opt/keycloak/certs

# # Copy the entrypoint script
# COPY entrypoint.sh /opt/keycloak/bin/entrypoint.sh

# # Switch to root user to set executable permission
# USER root
# RUN chmod +x /opt/keycloak/bin/entrypoint.sh

# # Switch back to the keycloak user
# USER keycloak

# # Set the entrypoint
# ENTRYPOINT ["/opt/keycloak/bin/entrypoint.sh"]

# Stage 1: Install OpenSSL in a separate image
FROM alpine:latest AS openssl-builder

# Install OpenSSL
RUN apk add --no-cache openssl

# Generate self-signed certificates
RUN mkdir -p /certs && \
    openssl req -x509 -newkey rsa:4096 -keyout /certs/tls.key -out /certs/tls.crt -days 365 -nodes -subj "/CN=localhost"

# Stage 2: Use the Keycloak image
FROM quay.io/keycloak/keycloak:latest

# Copy the certificates from the openssl-builder stage
COPY --from=openssl-builder /certs /opt/keycloak/certs

# Copy the realm configuration file

COPY realm_account.json /opt/keycloak/data/import/realm_account.json
COPY realm_endusers.json /opt/keycloak/data/import/realm_endusers.json

# Copy the entrypoint script
COPY entrypoint.sh /opt/keycloak/bin/entrypoint.sh

# Switch to root user to set executable permission
USER root
RUN chmod +x /opt/keycloak/bin/entrypoint.sh

# Switch back to the keycloak user
USER keycloak

# Set the entrypoint
ENTRYPOINT ["/opt/keycloak/bin/entrypoint.sh"]
