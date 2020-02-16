package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/yuchenfw/gocrypt"
)

type aesCrypt struct{

}

//单例helper对象
var AesCrypt *aesCrypt

func init() {
	AesCrypt = NewAesCrypt()
}

func NewAesCrypt() *aesCrypt {
	return &aesCrypt{

	}
}

//--------------------------- php的AES加解密參考 start ---------------------------
//    /**
//     * AES加密
//     * @param $plaintext 需要加密内容
//     * @param $key       加密key,解密时保持key一致
//     * @return string    返回base64_encode的加密内容
//     */
//    function encryptByAES($plaintext,$key){
//        //$key = 'test_openssl_encrypt';
//        //$key previously generated safely, ie: openssl_random_pseudo_bytes
//        //$plaintext = "message to be encrypted";
//        $ivlen = openssl_cipher_iv_length($cipher="AES-128-CBC");
//        $iv = openssl_random_pseudo_bytes($ivlen);
//        $ciphertext_raw = openssl_encrypt($plaintext, $cipher, $key, $options=OPENSSL_RAW_DATA, $iv);
//        $hmac = hash_hmac('sha256', $ciphertext_raw, $key, $as_binary=true);
//        $ciphertext = base64_encode( $iv.$hmac.$ciphertext_raw );
//        return $ciphertext;
//    }
//
//    /**
//     * AES解密
//     * @param $ciphertext   已被encryptByAES()加密过的内容
//     * @param $key          解密key,加密时保持key一致
//     * @return string       返回解密后的内容
//     * @throws Exception
//     */
//    function decryptByAES($ciphertext,$key){
//        //$ciphertext = $argv[1];
//        //$key = 'test_openssl_encrypt';
//        //decrypt later....
//        $c = base64_decode($ciphertext);
//        $ivlen = openssl_cipher_iv_length($cipher="AES-128-CBC");
//        $iv = substr($c, 0, $ivlen);
//        $hmac = substr($c, $ivlen, $sha2len=32);
//        $ciphertext_raw = substr($c, $ivlen+$sha2len);
//        $original_plaintext = openssl_decrypt($ciphertext_raw, $cipher, $key, $options=OPENSSL_RAW_DATA, $iv);
//        $calcmac = hash_hmac('sha256', $ciphertext_raw, $key, $as_binary=true);
//        if (hash_equals($hmac, $calcmac)) { //PHP 5.6+ timing attack safe comparison
//            return $original_plaintext;
//        }else{
//            throw new \Exception('AesCrypt.decryptByAES.hash_equals.result is error!');
//        }
//    }
//
//    //************************ AES加密 end ************************
//
//
//    //************************ AES-128-ECB 加密 start ************************
//    /**
//     * 解密字符串
//     * @param $str 要解密的字符串
//     * @param $key 加密时的key
//     * @param bool $isBase64Decode 是否base64_decode true是 false不是
//     * @return string
//     * @desc 客户端加密-适用范围:
//             $key要保持16位
//                安卓 AES-128-ECB Pkcs5padding   先加密后base_encode输出,由php先base64_decode再decryptByAES128ECB解密
//                IOS AES-128-ECB Pkcs7padding   先加密后base_encode输出,由php先base64_decode再decryptByAES128ECB解密
//                H5  AES-128-ECB Pkcs7padding   先加密后base_encode输出,由php先base64_decode再decryptByAES128ECB解密
//     */
//     function decryptByAES128ECB($str,$key,$isBase64Decode=true){
//        if( $isBase64Decode ){
//            $str = base64_decode($str);
//        }
//        return openssl_decrypt($str,'AES-128-ECB',$key,OPENSSL_RAW_DATA);
//    }
//
//    /**
//     * 加密字符串
//     * @param string $str 要加密的字符串
//     * @return string
//     */
//    function encryptByAES128ECB($str,$key){
//        return base64_encode(openssl_encrypt($str,'AES-128-ECB',$key,OPENSSL_RAW_DATA));
//    }
//    //************************ AES-128-ECB 加密 end ************************
//--------------------------- php的AES加解密參考 end ---------------------------

//************************ AES-128-pkcs7padding-iv-CBC start ************************
//加密
func (cthis *aesCrypt) AES128Pkcs7PaddingIvCBCEncrypt(origData, key []byte,IV []byte) (string, error) {
	if key == nil || len(key) != 16 {
		return "", nil
	}
	if IV != nil && len(IV) != 16 {
		return "", nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = gocrypt.PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData))
	//根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

//解密
func (cthis *aesCrypt) AES128Pkcs7PaddingIvCBCDecrypt(crypteData string, key []byte,IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}

	//base64解密
	crypted, cryptedErr := base64.StdEncoding.DecodeString(crypteData)
	if cryptedErr!=nil {
		return nil,cryptedErr
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block,IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = gocrypt.UnPaddingPKCS7(origData)
	return origData, nil
}
//************************ AES-128-pkcs7padding-iv-CBC end ************************