package xcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

// RsaEncryptionByPubKeyFile RSA加密(公钥加密)
func RsaEncryptionByPubKeyFile(data string, publicKeyFile string) (string, error) {
	bytes, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		return "", errors.New("public key errors")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err != nil {
		return "", err
	}
	encodeStr := base64.StdEncoding.EncodeToString(ciphertext)
	return encodeStr, nil
}

// RsaEncryptionByPubKey RSA加密(公钥加密)
func RsaEncryptionByPubKey(data string, publicKey string) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key errors")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err != nil {
		return "", err
	}
	encodeStr := base64.StdEncoding.EncodeToString(ciphertext)
	return encodeStr, nil
}

// RsaDecryptionByPriKeyFile RSA解密(私钥解密)
func RsaDecryptionByPriKeyFile(ciphertext string, privateKeyFile string) (string, error) {
	decodeStr, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return "", err
	}
	//获取私钥
	block, _ := pem.Decode(bytes)
	if block == nil {
		return "", errors.New("private key errors")
	}
	//解析PKCS1格式的私钥
	pkcs1PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, pkcs1PrivateKey, []byte(decodeStr))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RsaDecryptionByPriKey RSA解密(私钥解密)
func RsaDecryptionByPriKey(ciphertext string, privateKey string) (string, error) {
	decodeStr, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	//获取私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key errors")
	}
	//解析PKCS1格式的私钥
	pkcs1PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, pkcs1PrivateKey, []byte(decodeStr))
	if err != nil {
		return "", err
	}
	return string(data), nil
}
