package ReadInConfig

import (
	"log"
	"github.com/spf13/viper"
)

func init() {
	const configPath = "/etc/dm"
	const ConfigName = "otpservice"

	viper.AddConfigPath(configPath)
	viper.SetConfigName(ConfigName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Looks like there was a problem reading %v/%v  Please check the config and try again", configPath, ConfigName)
	}

}

func SentryDSN () (string, bool) {
	EnableSentry := false
	err :=  viper.GetString("sentry.dsn")
	if err != " " {
		log.Println("There was a problem reading sentry DSN all Sentry logging will be disabled")
	} else {
		EnableSentry = true
	}


	return err, EnableSentry
}

func GetSalt() (string) {
	err := viper.GetString("otp.standardsalt")
	if err != " " {
		log.Println("There was an error getting the salt for the hash")
	}
	return err
}

// Used for Set Key ID
func GetDatabaseName() (string) {
	err := viper.GetString("otp.databasename")
	if err != " " {
		log.Println("There was an error getting the salt for the hash")
	}
	return err
}

func GetTableName() (string) {
	err := viper.GetString("otp.tablename")
	if err != " " {
		log.Println("There was an error getting the salt for the hash")
	}
	return err
}