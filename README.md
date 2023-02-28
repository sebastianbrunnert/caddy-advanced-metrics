<p align="center">
  <h3 align="center">caddy-advanced-metrics</h3>

  <p align="center">
    Metric your Caddy servers using Prometheus and Grafana with this simple plugin. It provided per-host metrics for your servers and more useful metrics.
    <br/>
    <br/>
    <a href="https://github.com/sebastianbrunnert/caddy-advanced-metrics/issues">Report Bug</a>
    -
    <a href="https://github.com/sebastianbrunnert/caddy-advanced-metrics/issues">Request Feature</a>
  </p>
</p>

![Downloads](https://img.shields.io/github/downloads/sebastianbrunnert/caddy-advanced-metrics/total) ![Contributors](https://img.shields.io/github/contributors/sebastianbrunnert/caddy-advanced-metrics?color=dark-green) ![Stargazers](https://img.shields.io/github/stars/sebastianbrunnert/caddy-advanced-metrics?style=social) ![Issues](https://img.shields.io/github/issues/sebastianbrunnert/caddy-advanced-metrics) ![License](https://img.shields.io/github/license/sebastianbrunnert/caddy-advanced-metrics) 

## About The Project

With this caddy module built on bases of [xcaddy](https://github.com/caddyserver/xcaddy) you can easily monitor your caddy servers using Prometheus and Grafana. It provides per-host metrics for your servers and more useful metrics. It is possible to monitor your hosts on the same Prometheus port.

## Built With

* xcaddy
* Prometheus Go client library

## Getting Started

### Usage

```
{
	order advanced_metrics before file_server
  order advanced_metrics before reverse_proxy
}

:8080 {
    # Enable advanced metrics on port 1234 (http://localhost:1234/metrics)
    advanced_metrics {
      port 1234
    }

    # Enable file server
    file_server
}

:8081 {
    # Enable advanced metrics (default port 6611)
    advanced_metrics

    # Enable file server
    file_server
}
```

### Download build

You can download the latest build from the [releases page](https://github.com/sebastianbrunnert/caddy-advanced-metrics/releases/).

### Prerequisites

* Golang
* xcaddy

### Build and run

1. Clone the repository

```sh
git clone https://github.com/sebastianbrunnert/caddy-advanced-metrics.git
```

2. Install Golang dependencies

```sh
go mod download
```

3. Build custom caddy binary using xcaddy 

```sh
make build
```

4. Run custom caddy binary

Now you can run your custom caddy binary with the advanced metrics module. You can use the following command to run the binary:

```sh
./caddy run --config ./examples/Caddyfile
```

or you can setup the caddy service using systemd.

```sh
nano /usr/lib/systemd/system/caddy.service
```

```sh
[Service]
ExecStart=/path/to/caddy run --config /path/to/Caddyfile
```

## Roadmap

See the [open issues](https://github.com/sebastianbrunnert/caddy-advanced-metrics/issues) for a list of proposed features (and known issues).


## License

Distributed under the MIT License. See [LICENSE](https://github.com/sebastianbrunnert/caddy-advanced-metrics/blob/main/LICENSE) for more information.