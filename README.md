# drone-azure-storage

Example .drone.yml
```
pipeline:
  upload:
    image: kaey/drone-azure-storage
    storage_account_key: 123889asd89u8hsfdjh98128hh
    storage_account: my-storage-account
    container: my-storage-container
    source: file/to/upload
    target: remote/file/path
```

Building:
```
$ CGO_ENABLED=0 go build -v
$ docker build -t kaey/drone-azure-storage .
```
