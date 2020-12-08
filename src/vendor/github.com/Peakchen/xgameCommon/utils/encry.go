// add by stefan

package utils

import (
	"bytes"
	"github.com/Peakchen/xgameCommon/aktime"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	mrand "math/rand"
	"net"

	uuid "github.com/satori/go.uuid"
)

//get new uuid
func GetUUID() string {
	return uuid.NewV4().String()
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}

// get current computer mac addr.
func GetMacAddrs() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) > 0 {
			return macAddr
		}
	}

	return ""
}

// hash output uint32
func HashNonEncryption(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// add mac and time.now() as seed
func GetHashSeed() int64 {
	mac_adr := GetMacAddrs()
	if len(mac_adr) == 0 {
		return aktime.Now().UnixNano()
	}

	t := aktime.Now().UnixNano() // int64
	return int64(HashNonEncryption(fmt.Sprintf("%d %s", t, mac_adr)))
}

// ECB PKCS5Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// ECB PKCS5UnPadding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// ECB Des encrypt
func ECBEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	origData = PKCS5Padding(origData, bs)
	if len(origData)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

// ECB Des decrypt
func ECBDecrypt(crypted, key []byte) ([]byte, error) {
	if len(crypted) < 1 || len(key) < 1 {
		return nil, errors.New("wrong data or key")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(crypted)%bs != 0 {
		return nil, errors.New("DecryptDES crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(crypted))
	dst := out
	for len(crypted) > 0 {
		block.Decrypt(dst, crypted[:bs])
		crypted = crypted[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

// [golang ECB 3DES Encrypt]
func TripleEcbDesEncrypt(origData, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]

	block, err := des.NewCipher(k1)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	origData = PKCS5Padding(origData, bs)

	buf1, err := ECBEncrypt(origData, k1)
	if err != nil {
		return nil, err
	}
	buf2, err := ECBDecrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := ECBEncrypt(buf2, k3)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// [golang ECB 3DES Decrypt]
func TripleEcbDesDecrypt(crypted, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]
	buf1, err := ECBDecrypt(crypted, k3)
	if err != nil {
		return nil, err
	}
	buf2, err := ECBEncrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := ECBDecrypt(buf2, k1)
	if err != nil {
		return nil, err
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

// GetMd5String compute the md5 sum as string
func GetMd5String_V2(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GetRandString returns randominzed string with given length
func GetRandString(length int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	result := []byte{}
	r := mrand.New(mrand.NewSource(aktime.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// EncodeString encode string use base64 StdEncoding
func EncodeString(source string) string {
	return base64.StdEncoding.EncodeToString([]byte(source))
}

// DecodeString deencode string use base64 StdEncoding
func DecodeString(source string) (string, error) {
	dst, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return "", err
	} else {
		return string(dst), nil
	}
}

// EncodeMap encode the map to gob
func EncodeMap(obj map[interface{}]interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(obj)
	if err != nil {
		return []byte(""), err
	}
	return buf.Bytes(), nil
}

// DecodeMap decode data to map
func DecodeMap(encoded []byte) (map[interface{}]interface{}, error) {
	buf := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(buf)
	var out map[interface{}]interface{}
	err := dec.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func init() {
	mrand.Seed(GetHashSeed())

	gob.Register([]interface{}{})
	gob.Register(map[int]interface{}{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[interface{}]interface{}{})
	gob.Register(map[string]string{})
	gob.Register(map[int]string{})
	gob.Register(map[int]int{})
	gob.Register(map[int]int64{})
}
