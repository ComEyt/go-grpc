package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var (
	cfgFile     string
	projectBase string
	userLicense string
)

//还可以在 init() 函数中定义标志和处理配置，例如 cmd/root.go。
func init() {
	cobra.OnInitialize(initConfig)
	// 持久化标志,这意味着该标志可用于它所分配的命令以及该命令下的每个子命令
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file ")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "project base")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name ")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name user license")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// 将标志绑定到Viper，这样可以使用viper.Get获取标志的值
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
	// 设置标志为必选，如果没有提供该标志，cobra会报错
	//rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)"
	//rootCmd.MarkFlagRequired("region")

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
love by spf13 and friends in Go.
Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


