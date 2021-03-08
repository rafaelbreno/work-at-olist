# https://hub.docker.com/_/postgres
FROM postgres:alpine

# https://www.postgresql.org/docs/9.1/auth-pg-hba-conf.html
# Config file
COPY /.docker/postgres/conf/pg_hba.conf /var/lib/postgres/data/pg_hba.conf

# Running the database
CMD ["postgres"]

# Exposing PG port
EXPOSE 5432
