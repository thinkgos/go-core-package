package extcert

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"strings"
)

// 文件后缀
const base64Prefix = "base64://"

// ParseCrt 解析根证书
func ParseCrt(b []byte) (*x509.Certificate, error) {
	caBlock, _ := pem.Decode(b)
	return x509.ParseCertificate(caBlock.Bytes)
}

// LoadCrtFile 解析根证书文件
func LoadCrtFile(filename string) (*x509.Certificate, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseCrt(b)
}

// ParseKey 解析私钥
func ParseKey(b []byte) (*rsa.PrivateKey, error) {
	keyBlock, _ := pem.Decode(b)
	return x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
}

// LoadKeyFile 解析私钥文件
func LoadKeyFile(filename string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseKey(b)
}

// ParseCrtAndKey 解析根证书和私钥
func ParseCrtAndKey(crt, key []byte) (ca *x509.Certificate, privateKey *rsa.PrivateKey, err error) {
	ca, err = ParseCrt(crt)
	if err != nil {
		return
	}
	privateKey, err = ParseKey(key)
	return
}

// ParseCrtAndKeyFile 解析根证书文件和私钥文件
func ParseCrtAndKeyFile(crtFilename, keyFilename string) (ca *x509.Certificate, key *rsa.PrivateKey, err error) {
	ca, err = LoadCrtFile(crtFilename)
	if err != nil {
		return
	}
	key, err = LoadKeyFile(keyFilename)
	return
}

// LoadCrtAndKeyFile 读取根证书文件和私钥文件
func LoadCrtAndKeyFile(crtFilename, keyFilename string) (crt []byte, key []byte, err error) {
	crt, err = ioutil.ReadFile(crtFilename)
	if err != nil {
		return
	}
	key, err = ioutil.ReadFile(keyFilename)
	return
}

// LoadPair 加载tls cert key
// 如果cert有"base64://"前缀,直接解析后面的字符串,否则认为这是个cert文件名
// 如果key有"base64://"前缀,直接解析后面的字符串,否则认为这是个key文件名
func LoadPair(cert, key string) (certBytes, keyBytes []byte, err error) {
	certBytes, err = LoadCrt(cert)
	if err != nil {
		return
	}
	keyBytes, err = LoadKey(key)
	return
}

// LoadCrt 加载tls cert
// 如果cert有"base64://"前缀,直接解析后面的字符串,否则认为这是个cert文件名
func LoadCrt(cert string) ([]byte, error) {
	if strings.HasPrefix(cert, base64Prefix) {
		return base64.StdEncoding.DecodeString(cert[len(base64Prefix):])
	}
	return ioutil.ReadFile(cert)
}

// LoadKey 加载tls key
// 如果key有"base64://"前缀,直接解析后面的字符串,否则认为这是个key文件名
func LoadKey(key string) ([]byte, error) {
	if strings.HasPrefix(key, base64Prefix) {
		return base64.StdEncoding.DecodeString(key[len(base64Prefix):])
	}
	return ioutil.ReadFile(key)
}
