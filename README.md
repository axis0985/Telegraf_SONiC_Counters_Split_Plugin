# Telegraf Processor Plugin for gNMI Plugin Retrieving SONiC gNMI Counters Data

## Build
```bash
go build -o counters_split cmd/main.go
```

## Usage

Add the following to the telegraf configuration
```
[[processors.execd]]
  command=["/path/to/counters_split"]
```
