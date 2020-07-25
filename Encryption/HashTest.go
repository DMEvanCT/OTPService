package Encryption


import (
	"testing"
)

func TestMD5(t *testing.T) {
	hash := Md5Hash("super_Secret_value", "super_secret_salt")
	if string(hash) != "413589c502421ab3bf41d8139223e545" {
		t.Fail()
	}


}
