package aes

import (
	"appointed-registration/global"
	"crypto/aes"
	"encoding/base64"
)

func aesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func aesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

/*
    origData := []byte("17796761085") // 待加密的数据
	key := []byte("imed2019imed2019") // 加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ ECB模式 --------------------")
	encrypted := AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
*/
// 同时加密手机号码和手机验证码
func AesECBPass(mobile, code string) (string, string) {

	encryptedMobile, encryptedCode := aesEncryptECB([]byte(mobile), []byte(global.LoginKey)), aesEncryptECB([]byte(code), []byte(global.LoginKey))

	return base64.StdEncoding.EncodeToString(encryptedMobile), base64.StdEncoding.EncodeToString(encryptedCode)

}

// 只加密手机号码
func AesMobile(mobile string) string {

	encryptedMobile := aesEncryptECB([]byte(mobile), []byte(global.LoginKey))

	return base64.StdEncoding.EncodeToString(encryptedMobile)
}
