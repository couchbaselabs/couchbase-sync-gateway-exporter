FROM gcr.io/distroless/base
COPY couchbase-sync-gateway-exporter /bin/couchbase-sync-gateway-exporter
ENTRYPOINT ["/bin/couchbase-sync-gateway-exporter"]
CMD [ "-h" ]
