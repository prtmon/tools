package utils

import (
	"bytes"
	"github.com/pquerna/otp/totp"
	"image/png"
)

// GenerateTotpKey 生成TOTP密钥及QRCode图片
func GenerateTotpKey(issuer string, accountName string) (string, []byte, error) {
	secret := ""
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,      // 发行者名称
		AccountName: accountName, // 可以是邮箱地址或电话号码等唯一用户标识
	})
	if err != nil {
		return secret, nil, err
	}
	secret = key.Secret()
	// Convert TOTP key into a PNG
	img, err := key.Image(200, 200)
	if err != nil {
		return secret, nil, err
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return secret, nil, err
	}
	return secret, buf.Bytes(), err
}

// ValidTotpKey 检验totpKey是否有效
func ValidTotpKey(passCode string, secret string) bool {
	return totp.Validate(passCode, secret)
}
