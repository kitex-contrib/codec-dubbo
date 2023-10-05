# Benchmark

Make use of deploy.sh to deploy. Benchmark results would be stored in **stress_log**, client
and server logs would be stored in **client_log** and **server_log**

## Sub Commands

### Start the dubbo server process

```bash
# dubbo_server process would listen on :20001
sh deploy.sh dubbo_server

# specify listening port
sh deploy.sh dubbo_server -p 21001
```

### Start the kitex server process

```bash
# kitex_server process would listen on :20001
sh deploy.sh kitex_server

# specify listening port
sh deploy.sh kitex_server -p 21001
```

### Start the dubbo client process

```bash
# dubbo_client process would connect to dubbo_server with "127.0.0.1:20001"
# and listen on :20000
sh deploy.sh dubbo_client

# specify listening port
sh deploy.sh dubbo_client -p 21000

# specify dubbo_server address
sh deploy.sh dubbo_client -addr "192.168.0.2:20001"
```

### Start the kitex client process

```bash
# kitex_client process would connect to kitex_server with "127.0.0.1:20001"
# and listen on :20000
sh deploy.sh kitex_client

# specify listening port
sh deploy.sh kitex_client -p 21000

# specify kitex_server address
sh deploy.sh kitex_client -addr "192.168.0.2:20001"
```

### Start the stress process

```bash
# stress process would connect to dubbo_server with "127.0.0.1:20000"
# default tps: 100, parallel: 10, payloadLen: 10
sh deploy.sh stress

# specify dubbo_client address
sh deploy.sh stress -addr "192.168.0.3:20000"

# specify tps upper bound
sh deploy.sh stress -t 200

# specify parallel clients
sh deploy.sh stress -p 100

# specify payloadLen with byte unit
sh deploy.sh stress -l 256
```

### Close processes

```bash
# close all related processes
sh deploy.sh close all

sh deploy.sh close stress

sh deploy.sh close dubbo_client

sh deploy.sh close kitex_client

sh deploy.sh close dubbo_server

sh deploy.sh close kitex_server
```

## Benchmark kitex2kitex

```bash
# stress, kitex_client, kitex_server are on the same machine
# startup order: server -> client -> stress
sh deploy.sh kitex_server
sh deploy.sh kitex_client
sh deploy.sh stress -t 50000
sh deploy.sh close all
```

## Benchmark dubbo2dubbo

```bash
# assuming stress, dubbo_client, dubbo_server are on the same machine
# startup order: server -> client -> stress
sh deploy.sh dubbo_server
sh deploy.sh dubbo_client
sh deploy.sh stress -t 50000
sh deploy.sh close all
```

## Benchmark kitex2dubbo

```bash
# assuming stress, kitex_client, dubbo_server are on the same machine
# startup order: server -> client -> stress
sh deploy.sh dubbo_server
sh deploy.sh kitex_client
sh deploy.sh stress -t 50000
sh deploy.sh close all
```

## Benchmark dubbo2kitex

```bash
# assuming stress, dubbo_client, kitex_server are on the same machine
# startup order: server -> client -> stress
sh deploy.sh kitex_server
sh deploy.sh dubbo_client
sh deploy.sh stress -t 50000
sh deploy.sh close all
```

