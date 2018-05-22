# hardware-agent
hardware agent 

基于ipmitool，硬件监控专用agent。

##采集的metric列表：
  Hardware.agent.alive

  
##配置说明
配置文件请参照cfg.example.json，修改该文件名为cfg.json，将该文件里的IP换成实际使用的IP。
```
{
  "debug": true,
  "hostname": "zhangt-custom-data",
  "smartapi": "https://devsmarteye.anchnet.com/api/hardware/info",
  "plugin": {
    "enabled": false,
    "dir": "./plugin",
    "git": "https://github.com/open-falcon/plugin.git",
    "logs": "./logs"
  },
  "heartbeat": {
    "enabled": true,
    "addr": "127.0.0.1:6030",
    "interval": 60,
    "timeout": 1000
  },
  "transfer": {
    "enabled": true,
    "addrs": [
      "127.0.0.1:8433",
      "127.0.0.1:8433"
    ],
    "interval": 60,
    "timeout": 1000
  },
  "http": {
    "enabled": false,
    "listen": ":1988",
    "backdoor": false
  },
  "exectimeout": 0,
  "cycle": 0
}

```

