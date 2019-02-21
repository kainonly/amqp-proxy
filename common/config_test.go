package common

import (
	"github.com/go-ini/ini"
	"testing"
)

func TestLoadCogs(t *testing.T) {
	var err error
	config := new(Cogs)
	// load env
	if err = ini.MapTo(config, "../cogs.ini"); err != nil {
		t.Error(err.Error())
	}

	if config.ValidateArgs() {
		t.Error("please set cogs.ini!")
	}
}
