package config

import "amqp-proxy/config/options"

type Config struct {
	Debug    string                 `yaml:"debug"`
	Listen   string                 `yaml:"listen"`
	Gateway  string                 `yaml:"gateway"`
	Amqp     string                 `yaml:"amqp"`
	Transfer options.TransferOption `yaml:"transfer"`
}
