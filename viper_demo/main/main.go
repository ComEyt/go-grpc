package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func main() {
	ReadNestConfig()
}

var (
	cfg = pflag.StringP("config","c","","Configuration file")
	help = pflag.BoolP("help","h",false,"Show this help message")
)

func ViperReadConfig() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 从配置文件中读取配置
	if *cfg == "" {
		viper.SetConfigFile(*cfg)  //指定了配置文件名
		viper.SetConfigType("yaml") //如果配置文件名 也就是config后面的参数没有
	}else{
		viper.AddConfigPath(".")  // 把当前文件夹加入到配置文件的搜索路径中
		viper.AddConfigPath("../") //配置文件搜索路径
		viper.SetConfigName("config") //配置文件名称
	}
	if err := viper.ReadInConfig();err != nil{
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}

// HandleViperHotReload viper 热加载
func HandleViperHotReload() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件发生更改之后会调用的回调函数
		fmt.Println("Config file Changed",e.Name)
	})
	viper.Set("user.username","abc")
	username := viper.GetString("user.username")
	fmt.Println(username)
}

// HandleViperReadEnv 读取环境变量
func HandleViperReadEnv() {
	os.Setenv("apiversion_ViperReadEnv","viper_test")
	viper.AutomaticEnv()
	viper.BindEnv()
	viper.SetEnvPrefix("apiversion")
	get := viper.Get("ViperReadEnv")
	get2 := viper.Get("apiversion")
	fmt.Println(get)
	fmt.Println(get2)
}

// ReadNestConfig 读取嵌套的键
func ReadNestConfig() {
	viper.SetConfigName("configJson")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./viper_demo")

	if err := viper.ReadInConfig();err != nil{
		fmt.Println(err)
	}
	address := viper.GetString("host.address")
	sb := viper.GetString("sb.lbj")
	fmt.Println(address)
	fmt.Println(sb)
}