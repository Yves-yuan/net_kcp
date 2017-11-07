package process

import (
	"os"
	"runtime/pprof"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func initConfig(ctx *cli.Context) {
	viper.SetConfigType("toml")
	c := ctx.String("config")
	if c == "" {
		panic("configuration file argument missing")
	}

	viper.SetConfigFile(c)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func MakeServe(fn func(*cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		initConfig(ctx)
		if cpuprofile := ctx.String("cpuprofile"); cpuprofile != "" {
			f, err := os.Create(cpuprofile)
			if err != nil {
				panic(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}

		return fn(ctx)
	}
}
