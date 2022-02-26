FROM mongo:latest
COPY bap_db/init /docker-entrypoint-initdb.d/
ENV MONGO_INITDB_ROOT_USERNAME_FILE=/run/secrets/mongo-root
ENV MONGO_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongo-root-password
ENV MONGO_INITDB_DATABASE=database
LABEL "project"="bap"
