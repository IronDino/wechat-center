package models

const (
	// AccessTokenURL 获取AccessToken URL地址
	AccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

// GetAccessToken 获取公众号AccessToken
func GetAccessToken(appID, appSecret string) (string, error) {
	return "", nil
}
