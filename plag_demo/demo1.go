package plag_demo

import "github.com/spf13/pflag"

var version bool

func init() {
	// 自定义FlagSet
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.BoolVar(&version, "version", true, "Print version information and quit")

	// 使用全局的FlagSet,底层调用了CommandLine，CommandLine 是一个包级别的变量。
	pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit")
}
