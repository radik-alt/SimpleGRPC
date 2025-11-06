FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

RUN apk add --no-cache bash curl
ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /usr/local/bin/goose
RUN chmod +x /usr/local/bin/goose

WORKDIR /app
COPY scripts/run_migrations.sh .
COPY migrations ./migrations

ENTRYPOINT ["bash", "run_migrations.sh"]