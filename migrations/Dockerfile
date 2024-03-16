FROM ubuntu:jammy

#update system package
#and install golang-migrate tool
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl && \
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash && \
    apt-get update && \
    apt-get install -y migrate

WORKDIR /usr/var/migrations

#copy the migrations file inside container
COPY ./ ./

#allow to execute the migration script
RUN chmod +x migrate.sh

#run the migration
CMD ["./migrate.sh"]