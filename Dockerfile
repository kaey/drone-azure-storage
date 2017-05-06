FROM alpine:3.5
RUN apk -Uuv add ca-certificates python py-pip build-base python-dev libffi-dev openssl-dev \
  && pip install blobxfer \
  && apk del build-base python-dev libffi-dev openssl-dev \
  && rm -rf /var/cache/apk/*
ADD drone-azure-storage /drone-azure-storage
ENTRYPOINT ["/drone-azure-storage"]
