# mackerel-plugin-switchbot

A Mackerel plugin to collect metrics from various SwitchBot devices.

A list of supported devices can be found [here](https://github.com/SlashNephy/mackerel-plugin-switchbot/blob/master/metrics.go#L145).

## Example Metrics

```console
$ mackerel-plugin-switchbot
switchbot.battery.XXXXX-battery         95        1699895326
switchbot.battery.YYYYY-battery         97        1699895326
switchbot.battery.ZZZZZ-battery         100       1699895326
switchbot.temperature.XXXXX-temperature 27.800000 1699895326
switchbot.humidity.XXXXX-humidity       42        1699895326
```

## Usage

1. Install via mkr / or download releases directly

```console
$ mkr plugin install SlashNephy/mackerel-plugin-switchbot
```

2. Append following configuration to `mackerel-agent.conf`

```conf
[plugin.metrics.switchbot]
command = "/opt/mackerel-agent/plugins/bin/mackerel-plugin-switchbot --open-token XXX --secret-key XXX"
```

mackerel-plugin-switchbot has some command-line options. Check the help for details.

```console
$ /opt/mackerel-agent/plugins/bin/mackerel-plugin-switchbot --help
```
