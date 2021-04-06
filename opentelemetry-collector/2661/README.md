## Introduction
Adapted from [Brian Brazil's DropWizard tutorial], this code is meant to help debug and dissect [OpenTelemetry-Prometheus Receiver Dropwizard related bug](https://github.com/open-telemetry/opentelemetry-collector/issues/2661)

## Running it
```shell
$ mvn install && mvn exec:java -Dexec.mainClass=io.robustperception.exhibit.Exhibit
```

and the Prometheus scrape endpoint will be available at localhost:1234/metrics
