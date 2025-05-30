package model

import (
	"crypto/md5"
	"encoding/hex"
)

// Encrypting user password with hash
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}