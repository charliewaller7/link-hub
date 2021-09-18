# link-hub
A simple UI for viewing multiple web pages, built for managing multiple apps running on different localhost ports.

## Usage

Apps are defined in YAML config in the following format

```yaml
apps:
  - name: Elasticsearch
    url: localhost:9200
  - name: Kibana
    url: localhost:5601
  - name: My Web Server
    url: localhost:8080
```

To run Link hub with a specific configuration, pass in the config at runtime with the `-file` flag.

```bash
./linkhub -file path/to/config.yaml
```

The default port is `10101` (obscure to avoid clashing with common default ports), but can 
be overwritten with the `PORT` environment variable, of `-port` flag.

The command line `-port` flag takes precedence over the `PORT` env var.

```bash
> ./linkhub 
...running on 10101

> ./linkhub -port 7777
...running on 7777

> PORT=1111 ./linkhub
...running on 1111

>  PORT=1111 ./linkhub -port 7777
...running on 7777
```

### Binaries

The `./linkhub` binary is compiled for MacOS. Clone and re-compile for another OS.