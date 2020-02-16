package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/yuchenfw/gocrypt"
)

//@todo RAS加解密,只作参考,不作实际生产环境使用!

type rsaCrypt struct {
	secretInfo RSASecret
}

type RSASecret struct {
	PublicKey          string
	PublicKeyDataType  gocrypt.Encode
	PrivateKey         string
	PrivateKeyDataType gocrypt.Encode
	PrivateKeyType     gocrypt.Secret
}

func NewRSACrypt(secretInfo RSASecret) *rsaCrypt {
	return &rsaCrypt{
		secretInfo: secretInfo,
	}
}

//调用示例:
//	//签名信息,值如UserName=Test&Age=30
//	signInfo := strings.TrimRight("UserName=Test&Age=30", "&")
//	secretInfo := util.RSASecret{
//		//公钥
//		PublicKey:          params["PublicKey"].(string),
//		PublicKeyDataType:  gocrypt.Base64,
//		//私钥
//		PrivateKey:         params["PrivateKey"].(string),
//		PrivateKeyType:     gocrypt.PKCS8,
//		PrivateKeyDataType: gocrypt.Base64,
//	}
//	rasCryptStr,rasCryptErr = util.NewRSACrypt(secretInfo).Sign(signInfo, gocrypt.MD5, gocrypt.Base64)
//	//输出调试:
//	fmt.Println("rasCryptStr:",rasCryptStr)
//	fmt.Println("rasCryptErr:",rasCryptErr)

//RSA签名
func (thisObj *rsaCrypt) Sign(src string, hashType gocrypt.Hash, outputDataType gocrypt.Encode) (dst string, err error) {
	secretInfo := thisObj.secretInfo
	if secretInfo.PrivateKey == "" {
		return "", fmt.Errorf("secretInfo PrivateKey can't be empty")
	}
	privateKeyDecoded, err := gocrypt.DecodeString(secretInfo.PrivateKey, secretInfo.PrivateKeyDataType)
	if err != nil {
		return
	}
	prvKey, err := gocrypt.ParsePrivateKey(privateKeyDecoded, secretInfo.PrivateKeyType)
	if err != nil {
		return
	}
	cryptoHash, hashed, err := gocrypt.GetHash([]byte(src), hashType)
	if err != nil {
		return
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, prvKey, cryptoHash, hashed)
	if err != nil {
		return
	}
	return gocrypt.EncodeToString(signature, outputDataType)
}

//RSA验签
func (thisObj *rsaCrypt) VerifySign(src string, hashType gocrypt.Hash, signedData string, signDataType gocrypt.Encode) (bool, error) {
	secretInfo := thisObj.secretInfo
	if secretInfo.PublicKey == "" {
		return false, fmt.Errorf("secretInfo PublicKey can't be empty")
	}
	publicKeyDecoded, err := gocrypt.DecodeString(secretInfo.PublicKey, secretInfo.PublicKeyDataType)
	if err != nil {
		return false, err
	}
	pubKey, err := x509.ParsePKIXPublicKey(publicKeyDecoded)
	if err != nil {
		return false, err
	}
	cryptoHash, hashed, err := gocrypt.GetHash([]byte(src), hashType)
	if err != nil {
		return false, err
	}
	signDecoded, err := gocrypt.DecodeString(signedData, signDataType)
	if err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), cryptoHash, hashed, signDecoded); err != nil {
		return false, err
	}
	return true, nil
}
