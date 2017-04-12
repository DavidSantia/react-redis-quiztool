FROM redis:alpine

EXPOSE 4000

# On the container, hostname must be 0.0.0.0 not localhost
ENV WEBSOCKET_HOST 0.0.0.0

COPY redis-ws.tar /usr/local/bin

# Locate exec command, Insert command for redis-ws before that
RUN sed -i '/^exec/i\redis-ws &' /usr/local/bin/docker-entrypoint.sh 
