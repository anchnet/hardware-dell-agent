package g

import (
	"encoding/json"
	log "github.com/cihub/seelog"
	"os"
	"sync"

	"github.com/toolkits/file"
	"time"
)

type PluginConfig struct {
	Enabled bool   `json:"enabled"`
	Dir     string `json:"dir"`
	Git     string `json:"git"`
	LogDir  string `json:"logs"`
}

type HeartbeatConfig struct {
	Enabled  bool   `json:"enabled"`
	Addr     string `json:"addr"`
	Interval int    `json:"interval"`
	Timeout  int    `json:"timeout"`
}

type TransferConfig struct {
	Enabled  bool     `json:"enabled"`
	Addrs    []string `json:"addrs"`
	Interval int      `json:"interval"`
	Timeout  int      `json:"timeout"`
}

type HttpConfig struct {
	Enabled  bool   `json:"enabled"`
	Listen   string `json:"listen"`
	Backdoor bool   `json:"backdoor"`
}

type GlobalConfig struct {
	Debug       bool             `json:"debug"`
	Hostname    string           `json:"hostname"`
	Plugin      *PluginConfig    `json:"plugin"`
	Heartbeat   *HeartbeatConfig `json:"heartbeat"`
	Transfer    *TransferConfig  `json:"transfer"`
	Http        *HttpConfig      `json:"http"`
	SmartAPI    string           `json:"smartapi"`
	ExecTimeout time.Duration    `json:"exectimeout"`
	Cycle       time.Duration    `json:"cycle"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Info("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Error("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Error("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Error("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Error("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Info("read config file:", cfg, "successfully")
}
