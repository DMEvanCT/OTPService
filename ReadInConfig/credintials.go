package ReadInConfig

import (
	"log"
	"github.com/spf13/viper"
)

//  @todo Remove this when we move to AWS instead of file
func init() {
	const ConfigPath = "/etc/dm"
	const ConfigName = "OTPService"

	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName(ConfigName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Looks like there was a problem reading %v/%v  Please check the config and try again.", ConfigPath, ConfigName)
	}

}
//  @todo Get from AWS instead of file
func SentryDSN () (string, bool) {
	EnableSentry := false
	err :=  viper.GetString("Sentry.dsn")
	if err != " " {
		log.Println("There was a problem reading sentry DSN all Sentry logging will be disabled")
	} else {
		EnableSentry = true
	}


	return err, EnableSentry
}
//  @todo Get from AWS instead of file
func GetSalt() (string) {
	err := viper.GetString("OTP.standardsalt")
	if err != " " {
		log.Println("There was an error getting the salt for the hash")
	}
	return err
}

// Used for Set Key ID
//  @todo Get from AWS instead of file
func GetDatabaseName() (string) {
		err := viper.GetString("OTP.databasename")
	if err != " " {
		log.Println("There was an error getting the salt for the hash")
	}
	return err
}
//  @todo Get from AWS instead of file
func GetTableName() (string) {
	err := viper.GetString("OTP.tablename")
	if err != " " {
		log.Println("There was an error getting the salt for the hash")
	}
	return err
}