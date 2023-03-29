package xcrypto

import (
	"testing"
	"time"
)

func TestRsaEncryptionAndDecryption(t *testing.T) {
	data := "hello world"
	timeNow := time.Now()
	t.Log("加密前数据：", data)
	ciphertext, err := RsaEncryptionByPubKey(data, RsaPublicKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("公钥加密：", ciphertext)
	t.Log("加密后长度：", len(ciphertext))
	t.Log("加密耗时 = ", time.Since(timeNow).Microseconds(), "微秒")

	timeNow2 := time.Now()

	result, err := RsaDecryptionByPriKey(ciphertext, RsaPrivateKey)
	if err != nil {
		t.Error(err)
		return
	}

	// 1.264
	t.Log("私钥解密：", result)
	t.Log("解密耗时 = ", time.Since(timeNow2).Microseconds(), "微秒")

}
