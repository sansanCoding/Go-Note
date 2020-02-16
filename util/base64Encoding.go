package util

import (
	"encoding/base64"
	"strings"
)

//Base64编码

type base64Encoding struct {

}

var Base64Encoding *base64Encoding

func init(){
	Base64Encoding = NewBase64Encrypt()
}

func NewBase64Encrypt() *base64Encoding {
	return &base64Encoding {

	}
}

//推荐: base64.RawURLEncoding.EncodeToString编码 + base64.RawURLEncoding.DecodeString解码
//	适用于url传输的,且base64.RawURLEncoding.EncodeToString编码后,是不携带=,也不用针对base64编码后再进行 / + =的处理
//	同样base64.RawURLEncoding.DecodeString解码,也不需要替换=为空字符串的处理
//参考示例:
//	qrcodeContent := "http://www.baidu.com?name=~!@#$%^&*()_+&b=1/2/3+4+5+6"
//	stdBytes := base64.RawURLEncoding.EncodeToString([]byte(qrcodeContent))
//	fmt.Println(string(stdBytes))
//	stdStr,stdStrErr := base64.RawURLEncoding.DecodeString(string(stdBytes))
//	fmt.Println(string(stdStr),stdStrErr)
//输出结果:
//	aHR0cDovL3d3dy5iYWlkdS5jb20_bmFtZT1-IUAjJCVeJiooKV8rJmI9MS8yLzMrNCs1KzY
//	http://www.baidu.com?name=~!@#$%^&*()_+&b=1/2/3+4+5+6 <nil>
//------------------------------------------------------------
//注:若是PHP之类的语言收到 aHR0cDovL3d3dy5iYWlkdS5jb20_bmFtZT1-IUAjJCVeJiooKV8rJmI9MS8yLzMrNCs1KzY 这样的base64编码字符串,
//	 需要第一步针对 _ - 这2个字符串进行替换,将 _ 替换成 /,将 - 替换成 +;
//	 PHP之类的语言才可以base64解码成功.
//	 以PHP代码示例:
//		//未实事先针对 _ - 进行替换,是无法解码成功的!
//		php -r "
//			echo base64_decode('aHR0cDovL3d3dy5iYWlkdS5jb20_bmFtZT1-IUAjJCVeJiooKV8rJmI9MS8yLzMrNCs1KzY');
//		"
//		//输出:
//		http://www.baidu.comOR2BUᢢᤣђ㱲³B³R³
//		//针对 _ - 进行替换后,解码成功!
//		php -r "
//			echo base64_decode('aHR0cDovL3d3dy5iYWlkdS5jb20/bmFtZT1+IUAjJCVeJiooKV8rJmI9MS8yLzMrNCs1KzY');
//		"
//		//输出:
//		http://www.baidu.com?name=~!@#$%^&*()_+&b=1/2/3+4+5+6
//------------------------------------------------------------
func (thisObj *base64Encoding) RawURLEncoding(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
func (thisObj *base64Encoding) RawURLDecoding(data string) (string,error) {
	decodeByte,err := base64.RawURLEncoding.DecodeString(data)
	if err!=nil {
		return "",err
	}
	return string(decodeByte),nil
}

//这个是针对标准base64编码的一种处理,将url传输的字符进行替换处理!
//	这个跟base64.RawURLEncoding.EncodeToString()效果差不多!
func (thisObj *base64Encoding) UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	byteArr := base64.StdEncoding.EncodeToString(source)
	safeUrl := strings.Replace(string(byteArr), "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return safeUrl
}

//关于base64的坑参考文章地址:
//https://blog.csdn.net/u014270740/article/details/91038606

//------------------------------ 针对标准的base64和适用于url的base64比较 start ------------------------------
/*
//qrcodeContent := "http://www.baidu.com/test=1+1"

//StdEncoding 输出:
//#######################################################
//	StdEncoding是标准的base64编码，如
//	RFC 4648号。
//#######################################################
//@todo 可以适用于文件,但不适用于url传输,因为不会对/和+字符进行处理
//#######################################################
//stdBytes := base64.StdEncoding.EncodeToString([]byte(qrcodeContent))
//fmt.Println(string(stdBytes))
//stdStr,stdStrErr := base64.StdEncoding.DecodeString(string(stdBytes))
//fmt.Println(string(stdStr),stdStrErr)
//输出结果:
//aHR0cDovL3d3dy5iYWlkdS5jb20vdGVzdD0xKzE=
//http://www.baidu.com/test=1+1 <nil>

//UrlEncoding 输出:
//#######################################################
//	URLEncoding是RFC 4648中定义的备用base64编码。
//	它通常用于url和文件名中。
//#######################################################
//@todo 可以适用于文件,也可以用于url传输,但是有填充字符=,一旦url传参多了,也不大适用于url传输
//#######################################################
//stdBytes := base64.URLEncoding.EncodeToString([]byte(qrcodeContent))
//fmt.Println(string(stdBytes))
//stdStr,stdStrErr := base64.URLEncoding.DecodeString(string(stdBytes))
//fmt.Println(string(stdStr),stdStrErr)
//输出结果:
//aHR0cDovL3d3dy5iYWlkdS5jb20vdGVzdD0xKzE=
//http://www.baidu.com/test=1+1 <nil>

//RawStdEncoding 输出:
//#######################################################
//	RawStdEncoding是标准的原始、未添加的base64编码，
//	如RFC 4648第3.2节所定义。
//	这与StdEncoding相同，但省略了填充字符。
//#######################################################
//@todo 与StdEncoding一样,虽然省略了填充字符,但是也不适合用于url传输
//#######################################################
//stdBytes := base64.RawStdEncoding.EncodeToString([]byte(qrcodeContent))
//fmt.Println(string(stdBytes))
//stdStr,stdStrErr := base64.RawStdEncoding.DecodeString(string(stdBytes))
//fmt.Println(string(stdStr),stdStrErr)
//输出结果:
//aHR0cDovL3d3dy5iYWlkdS5jb20vdGVzdD0xKzE
//http://www.baidu.com/test=1+1 <nil>

//RawURLEncoding 输出:
//#######################################################
//	RawURLEncoding是RFC 4648中定义的未添加的备用base64编码。
//	它通常用于url和文件名中。
//	这与URLEncoding相同，但省略了填充字符。
//#######################################################
//@todo 与UrlEncoding一样,省略了填充字符,适合用于url传输;
//@todo base64.RawURLEncoding.EncodeToString也不用特地针对/ + =进行处理
//@todo base64.RawURLEncoding.DecodeString也适合进行无填充=的解码处理
//#######################################################
//stdBytes := base64.RawURLEncoding.EncodeToString([]byte(qrcodeContent))
//fmt.Println(string(stdBytes))
//stdStr,stdStrErr := base64.RawURLEncoding.DecodeString(string(stdBytes))
//fmt.Println(string(stdStr),stdStrErr)
//输出结果:
//aHR0cDovL3d3dy5iYWlkdS5jb20vdGVzdD0xKzE
//http://www.baidu.com/test=1+1 <nil>

*/
//------------------------------ 针对标准的base64和适用于url的base64比较 end ------------------------------

//------------------------------ 针对 一段RSA加密后base64编码的字符串 解码测试 start ------------------------------
/*
//qrcodeContent := "Cf1WA2nBMo3H9G2UPhlLBBVBsMDl4udWr7__e6Iy93eIqLKi3EOjGhk8TkHujL1Uj6aGfZJNBzIbVE2NfNaz4pob8uiQvGaeTZdWP-8lFmAm6J1sz8N15xQkO7ADa5bNLCCqtlQbN2z7JcNenvFuID_rZGqb_1gmr-BGubGRMiMSK7RdjQYrMHaBcHLPB0UteakzcQwgKxCW7u0ECHqPJ39ne9JUG22JBWRo1ORuX5r30J_XrW3SQcdPSxfe0kvd61y12QOYh8VlOBBdBeDNnyDXefI_tDJDBFeqTXCgKu9wFkkWIZiM7WwqogaY-bvjUisbrPO4_fjJ1c0nWDOqRA"
//StdEncoding-base64解析:
//stdStr,stdStrErr := base64.StdEncoding.DecodeString(qrcodeContent)
//fmt.Println(string(stdStr),stdStrErr)
//输出结果(报错):
//	�Vi�2���m�>KA�����V illegal base64 data at input byte 34

//URLEncoding-base64解析:
//stdStr,stdStrErr := base64.URLEncoding.DecodeString(qrcodeContent)
//fmt.Println(string(stdStr),stdStrErr)
//输出结果(报错):
//	�Vi�2���m�>KA�����V���{�2�w�����C�<NAT���}�M2TM�|ֳ��萼f�M�V?�%`&�l��u�$;�k��, ��T7l�%�^��n ?�dj��X&��F���2#+�]�+0v�pr�E-y�3q +���z�'g{�Tm�dh��n_��П׭m�A�OK��K��\�����e8]�͟ �y�?�2CW�Mp�*�pI!���l*�����R+�������'X3� illegal base64 data at input byte 340

//RawStdEncoding-base64解析:
//stdStr,stdStrErr := base64.RawStdEncoding.DecodeString(qrcodeContent)
//fmt.Println(string(stdStr),stdStrErr)
//输出结果(报错):
//	�Vi�2���m�>KA�����V illegal base64 data at input byte 34

//RawURLEncoding-base64解析:
//stdStr,stdStrErr := base64.RawURLEncoding.DecodeString(qrcodeContent)
//fmt.Println(string(stdStr),stdStrErr)
//输出结果(正确):
//	�Vi�2���m�>KA�����V���{�2�w�����C�<NAT���}�M2TM�|ֳ��萼f�M�V?�%`&�l��u�$;�k��, ��T7l�%�^��n ?�dj��X&��F���2#+�]�+0v�pr�E-y�3q +���z�'g{�Tm�dh��n_��П׭m�A�OK��K��\�����e8]�͟ �y�?�2CW�Mp�*�pI!���l*�����R+�������'X3�D <nil>
*/
//------------------------------ 针对 一段RSA加密后base64编码的字符串 解码测试 end ------------------------------