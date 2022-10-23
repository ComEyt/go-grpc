package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
	version bool
	flagVar = pflag.Int("flagName", 1234, "help message for flagName")
	ip      = pflag.IntP("flagName2", "f", 1234, "helpMessage")
	port    = pflag.IntP("flagName2", "f", 1234, "helpMessage")
)

func main() {
	SetNoOptDefValue()
}

// Get<Type>
func getType() {
	// 自定义FlagSet
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.BoolVar(&version, "version", true, "Print version information and quit")

	flagType, err := flagSet.GetBool("version")
	fmt.Println(flagType, err)
}

// GetNonOptionParam 获取非选项参数
func GetNonOptionParam() {
	pflag.Parse()

	// 参数数量
	fmt.Printf("argument number is: %v\n", pflag.NArg())
	// 参数列表
	fmt.Printf("argument list is: %v\n", pflag.Args())
	// 第一个参数是
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))
}

// SetNoOptDefValue 指定了选项但是没有指定选项值时的默认值。
func SetNoOptDefValue() {
	pflag.Parse()
	pflag.Lookup("flagName2").NoOptDefVal = "4321"
}

func QuitSymbol() {
	pflag.CommandLine.MarkDeprecated("logmode", "please use --log-mode instead")
}

// QuitEasyLine 保留名为 port 的标志，但是弃用它的简写形式。也就是弃用下列例子中的--P
func QuitEasyLine() {
	pflag.IntVarP(port, "port", "P", 3306, "MySQL service host port.")
	pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port only")
}
