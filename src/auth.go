package main

import (
	"crypto/des"
	"fmt"
	"strconv"
)

//
// GetVNCAuthenticationBytes 获取VNCAuthentication挑战响应值
//  @Description: 获取VNCAuthentication挑战响应值
//   加密方式参考: RFB协议: https://vncdotool.readthedocs.io/en/0.8.0/rfbproto.html#vnc-authentication
//   加密方式参考: https://blog.csdn.net/counsellor/article/details/106838216
//  @param password 密码
//  @param challenge 挑战码
//  @return []byte 挑战响应值
//
func GetVNCAuthenticationBytes(password, challenge []byte) []byte {
	// 取8位密码,如果超出则进行截取,如果不够8位则在最后进行空字符串补位
	var pwd8 []byte
	if len(password) > 8 {
		pwd8 = password[:8]
	} else {
		pwd8 = []byte(fmt.Sprintf("%-8s", password))
	}
	pwdKeyByte := reverseBinaryBlock(pwd8)
	return encryptDES(challenge, pwdKeyByte)
}

//
// reverseBinaryBlock 对每个二进制块进行反转
//  @Description: 对每个二进制块进行反转,采用的笨办法逐个反转,应该会有更好的方式
//   原始内容:      12345678
//   二进制块:      00110001 00110010 00110011 00110100 00110101 00110110 00110111 00111000
//   反转后二进制块: 10001100 01001100 11001100 00101100 10101100 01101100 11101100 00011100
//  @param n 要反转的bytes
//  @return rev 反转后的bytes
//
func reverseBinaryBlock(n []uint8) (rev []uint8) {
	for i := 0; i < len(n); i++ {
		uint8Str := fmt.Sprintf("%08b", n[i])
		retUint8Bytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			retUint8Bytes[i] = uint8Str[8-1-i]
		}
		retUint8, _ := strconv.ParseUint(string(retUint8Bytes), 2, 8)
		rev = append(rev, uint8(retUint8))
	}
	return
}

//
// encryptDES DES-ECB方式加密
//  @Description: DES-ECB方式加密
//  @param src 加密数据
//  @param key 密钥
//  @return []byte 加密结果
//
func encryptDES(src, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		// 对明文按照blockSize进行分块加密
		// 必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out
}
