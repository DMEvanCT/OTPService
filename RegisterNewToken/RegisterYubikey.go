package RegisterNewToken

import (
	"OTPService/Encryption"
	"OTPService/ReadInConfig"
)



// This will take the  12 charecters from  the token to generate a key. Then store that key in the database.
func RegisterNewYubikey(key_otp string, userid int, grpc YubikeyServer) {
	// This is the entire OTP it will  use the first 12 chars for key id
	key_id_to_transform  := []rune(key_otp)
	// This is the key id for the Yubikey
	key_to_store := string(key_id_to_transform[0:12])
	//Converts the key to MD5 to store
	Keys := Encryption.Md5Hash(key_to_store, ReadInConfig.GetSalt())
	KeyToDatbase("Yubikey", string(Keys), userid)

}
