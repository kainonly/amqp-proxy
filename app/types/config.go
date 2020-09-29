package types

type Config struct {
	Debug    string         `yaml:"debug"`
	Listen   string         `yaml:"listen"`
	Amqp     string         `yaml:"amqp"`
	Transfer TransferOption `yaml:"transfer"`
}
