# almond

A simple glue solution for simplifying the configuration and use of the 
prometheus-alertmanager-grafana stack for service monitoring and alerting.

## Architecture

### Components

* Consul
  * Service registration for Prometheus to discover and scrape.
  * KV store for essential information.
* Prometheus
  * Where metrics are collected.
  * Where __alerting rules__ are defined.
* Prometheus Alertmanager
  * Where alerts are received and distributed.
* Grafana
  * Where metrics are exhibited.
* Almond
  * Providing deployment template to setup the stack easily.
  * Providing APIs to setup the rules and dashboards quickly.

### Diagram 

TODO

## Deployment

### Use with docker
```
docker run -p 8080:8080 --name=almond joshuakwan/almond
```

### Use with docker-compose

Use the compose file in the repo to launch the almond application along with dependent services
```
docker-compose up
```