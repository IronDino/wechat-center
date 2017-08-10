package models

const (
	// APITicketURL 获取APITicket URL地址
	APITicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

// GetAPITicket 获取公众号APITicket
func GetAPITicket(appID, appSecret string) (string, error) {
	return "", nil
}
