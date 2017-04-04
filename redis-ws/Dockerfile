FROM redis:alpine

COPY redis-ws /usr/local/bin

# Locate exec command, Insert command for redis-ws before that
RUN sed -i '/^exec/i\redis-ws &' /usr/local/bin/docker-entrypoint.sh 
