package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/user_service/pkg/settings"
)

func LoadConfig() settings.Config {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../../configs")

	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read the configuration %w \n", err))
	}
	fmt.Println("Server Port:: ", viper.GetInt("server.port"))
	var config settings.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
	return config
}
