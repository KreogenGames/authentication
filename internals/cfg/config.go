package cfg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Cfg struct {
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("SERV") //Задание префикса для переменных окружения сервера
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "143018h")
	v.SetDefault("DBHOST", "") //localhost
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "postgres")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}

	return cfg
}

func (cfg *Cfg) GetDBString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
