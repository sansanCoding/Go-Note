package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
)

//加密

type encrypt struct {

}

var Encrypt *encrypt

func init(){
	Encrypt = NewEncrypt()
}

func NewEncrypt() *encrypt {
	return &encrypt{

	}
}

//md5加密
func (thisObj *encrypt) MD5(message string) string {
	res := md5.Sum([]byte(message))
	return fmt.Sprintf("%x", res)
}

//md5加密16
func (thisObj *encrypt) MD5By16(message string) string {
	res := md5.Sum([]byte(message))
	return fmt.Sprintf("%x", res)[8:24]
}

//sha1加密
func (thisObj *encrypt) SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//获取字符串HashCode
func (thisObj *encrypt) HashCode(str string) int {
	v := int(crc32.ChecksumIEEE([]byte(str)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}

	return 0
}

//格式化Rsa的秘钥
//@params string key 需要格式化的秘钥内容体
// keyType=1,为key增加头尾，并每间隔64位换行
// ifPublic true 为公钥， false为私钥
// keyType=0,不变
func (thisObj *encrypt) FormatRSAKey(key string, keyType int, ifPublic bool) string {
	if keyType == 0 {
		return key
	}
	if ifPublic {
		publicBegin 	:= "-----BEGIN PUBLIC KEY-----\r\n"
		publicEnd 		:= "-----END PUBLIC KEY-----"
		temp 			:= ""
		//Rsa秘钥每64个字符拼接\r\n
		thisObj.RSA64Split(key, &temp)
		return publicBegin + temp + publicEnd
	} else {
		privateBegin 	:= "-----BEGIN RSA PRIVATE KEY-----\r\n"
		privateEnd 		:= "-----END RSA PRIVATE KEY-----"
		temp := ""
		//Rsa秘钥每64个字符拼接\r\n
		thisObj.RSA64Split(key, &temp)
		return privateBegin + temp + privateEnd
	}
}
//RSA每64位分割换行处理
func (thisObj *encrypt) RSA64Split(key string, temp *string) {
	if len(key) <= 64 {
		*temp = *temp + key + "\r\n"
	}
	for i := 0; i < len(key); i++ {
		if (i+1)%64 == 0 {
			*temp = *temp + key[:i+1] + "\r\n"
			key = key[i+1:]
			thisObj.RSA64Split(key, temp)
			break
		}
	}
}
