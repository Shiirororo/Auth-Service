package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/user_service/global"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
