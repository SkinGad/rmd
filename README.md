# Routes Mikrotik Daemon
This is a service that manages the routing table on Mikrotik devices.
It retrieves IP addresses from DNS servers and adds them to the routing table with a specified lifetime

# Run

## Configure config
You need to copy `config.yml.example` and remove `example`
```yaml
rmd:
  # Frequency of requests to DNS servers
  lookup: 5s
  # Duration of storage of records in the cache
  ttl: 240h

routers:
  - address: localhost
    port: 8728
    login: admin
    password: admin
    # This gateway will be set to routes
    gateway: 192.168.0.1
    # Optional field. Which routing table should I work with?
    table: rmd

# You can specify domains, IP addresses and subnets
domain:
  - nvcr.io
  - 8.8.8.8/24
```

## Start
```
go run ./cmd/main.go
```

# Logic of work
During operation, data is retrieved from the saved cache, from DNS server requests, and from Mikrotik

The lifetime of entries written to the config is constantly updated, and if we stop receiving a certain IP address from DNS, it will be disabled when its expiration date expires

We can manually add entries to Mikrotik, but keep in mind that these entries will not be updated and a countdown begins for them until they are disabled

# Bugs
- The service crashes when the connection to Mikrotik is lost
- Duplicate routes may be created if the TCP connection is unstable