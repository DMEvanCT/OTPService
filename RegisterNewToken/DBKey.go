package RegisterNewToken

import (
	"OTPService/Logging"
	"OTPService/ReadInConfig"
	"github.com/DMEvanCT/GoBase/Database"
)
// This is used to store the keys for the OTP in thee database




// KeyType Ex: Yubikey
func KeyToDatbase(keytype, keyid string, Userid int)  (string) {
	if keytype == "Yubikey" {
		db := Database.DatabaseInitAuth()
		databasename := ReadInConfig.GetDatabaseName()
		tablename :=   ReadInConfig.GetTableName()
		setkey, err := db.Begin()
		if err != nil {
			Logging.LogForMSInfo(err)
		}

		defer setkey.Rollback()
		createkeyid, err := setkey.Prepare("UPDATE " + databasename + "." + tablename + "SET Yubikey = " + keyid + " WHERE UserID = (?)")
		if err != nil {
			Logging.LogForMSInfo(err)
		}
		defer createkeyid.Close()
		_, err = createkeyid.Exec(&Userid)
		if err != nil {
			Logging.LogForMSInfo(err)
		}
		setkey.Commit()

	}

	return "Result: Key created successfully"

}