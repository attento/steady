# Redundancy

Redundancy is a HTTP balancer and reverse proxy, it helps you to configure and
manage dynamic HTTP application.

This is a typical configuration file when `gianarb.it` is a frontend, you can
have more of one, each contains `port`, `bind` and a list of `nodes`.

```json
{
  "rconf": {
    "admin": {
      "port": 9992,
      "bind": "0.0.0.0"
    }
  },
  "frontends": {
    "gianarb": {
      "port": 8089,
      "bind": "0.0.0.0",
      "nodes": [
        {"host": "www.google.com"},
        {"host": "www.google.com"}
      ]
    }
  }
}
```

Each `node` is describe with an `host`.

## Try for dev
At the moment we don't serve any facility to compile and install this tool
because it's work in progress but you can try anyway with this easy flow:

```
go get github.com/gianarb/lb.git
cd $GOPATH/src/github/gianarb/lb
go get ./...
go run main.go -c ./lb.config.json
```

## rconf
`rconf` configuration node descibes the general configuration for `redundancy`
you can enable or disable the API system and also configure bind adress and
port for this service.

`admin` field enable or disable the JSON Api to manage `lb` you can change
`port` and `bind` address.

# Api 
The api system is the most important port with the external world, the goal is
support everything that you can do with this tool.

* `/ping` is a sanity call to understand if the service is up and running
