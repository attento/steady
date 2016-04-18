# Redundancy

Redundancy is a HTTP balancer and reverse proxy, it helps you to configure and
manage dynamic HTTP application.

This is a typical configuration file when `gianarb.it` is a frontend, you can
have more of one, each contains `port`, `bind` and a list of `nodes`.

```json
{
  "my.gianarb.it": {
      "frontend": {
          "port": 8089,
          "bind": "my.gianarb.it"
      },
      "nodes": [
          {"host":"www.google.com"},
          {"host":  "www.corley.it"}
      ]
  }
}
```

Each `node` is describe with an `host`.
