#!/bin/sh

while !(migrate -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable -path /usr/var/migrations up)
do
    echo "Waiting for postgres..."
done
