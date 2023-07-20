package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Addr string
}

func GetConfigs() (*Config, error) {
	cfg := Config{}

	rootCmd := &cobra.Command{
		Use: "Blockchain Node",
	}
	cobraInit(rootCmd, &cfg)
	if err := rootCmd.Execute(); err != nil {
		return nil, err
	}
	viper.Unmarshal(&cfg)

	return &cfg, nil
}

func cobraInit(rootCmd *cobra.Command, cfg *Config) {
	rootCmd.Flags().StringP("addr", "a", ":8080", "Address to listen on")
	viper.BindPFlag("addr", rootCmd.Flags().Lookup("addr"))
}
