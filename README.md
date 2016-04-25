# Steady
[![Build Status](https://travis-ci.org/gianarb/steady.svg?branch=master)](https://travis-ci.org/gianarb/steady)

Steady is a HTTP balancer and reverse proxy, it helps you to configure and
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
go get github.com/gianarb/steady.git
cd $GOPATH/src/github/gianarb/steady
go get ./...
go run main.go -c ./lb.config.json
```

## rconf
`rconf` configuration node descibes the general configuration for `Steady`
you can enable or disable the API system and also configure bind adress and
port for this service.

`admin` field enable or disable the JSON Api to manage `lb` you can change
`port` and `bind` address.

# Api 
The api system is the most important port with the external world, the goal is
support everything that you can do with this tool.

* `/ping` is a sanity call to understand if the service is up and running
* `/backup` returns the lb's configuration in the current status
* `/frontend` returns the list of frontends
* `/frontend/{name}` returns single frontends
* POST on `/frontend/{name}` add frontend `{"bind": "0.0.0.0", "port": 9129}`
* DELETE on `/frontend/{name}` delete single frontend
* POST on `/frontend/{name}/node` with `{"host": "127.0.10.1", "fields": {"type": "web"}}` add new node on name
* DELETE on `/frontend/{name}/node` with `{"host": "127.0.10.1"}` deletes all nodes for this host

**Look into `/res`**, there is a Postman collection of request.

NB. Thanks [@fntlnz](https://github.com/fntlnz) for the awesome name!
