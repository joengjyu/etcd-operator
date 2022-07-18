FROM alpine:3.6

RUN apk add --no-cache ca-certificates

COPY ./etcd-backup-operator /usr/local/bin/etcd-backup-operator
COPY ./etcd-restore-operator /usr/local/bin/etcd-restore-operator
COPY ./etcd-operator /usr/local/bin/etcd-operator

RUN adduser -D etcd-operator
USER etcd-operator
