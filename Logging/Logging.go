package Logging

import (
	"OTPService/ReadInConfig"
	"fmt"
	"log"
	"github.com/getsentry/sentry-go"
	"time"
)


func InitSentry() (bool) {
	sentrydsn, sentryenabled := ReadInConfig.SentryDSN()
	if sentryenabled == false {
		log.Println("Sentry is noto enabled")
	} else {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: sentrydsn,
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
		defer sentry.Flush(2 * time.Second)
	}
	return sentryenabled
}

func LogForMSInfo(err error) {
	st := InitSentry()
	if st == false {
		log.Println(err)
	} else {
		errorcap := fmt.Sprintf("%v", err)
		sentry.CaptureMessage(errorcap)
	}

}
