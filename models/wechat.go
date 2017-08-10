package models

import (
	"time"

	"github.com/gotit/go-net/net"
)

// Wechat 微信配置信息
type Wechat struct {
	AppID                 string    `json:"" db:""` // 公众号AppID
	AppSecret             string    `json:"" db:""` // 公众号AppSecret
	FreshAccessToken      bool      `json:"" db:""` // 公众号AccessToken是否去刷新
	AccessToken           string    `json:"" db:""` // 公众号AccessToken
	AccessTokenExpireTime time.Time `json:"" db:""` // 公众号AccessToken过期时间
	FreshAPITicket        bool      `json:"" db:""` // 公众号APITicket是否去刷新
	APITicket             string    `json:"" db:""` // 公众号ApiTicket
	APITicketExpireTime   time.Time `json:"" db:""` // 公众号APITicket过期时间
	MchID                 string    `json:"" db:""` // 商户MchID
	MchSecret             string    `json:"" db:""` // 商户MchSecret
	NativeAppID           string    `json:"" db:""` // 应用APPID
	PayNotityURL          string    `json:"" db:""` // 支付通知地址
	CertP12               string    `json:"" db:""` // 退款CertP12文件地址
	CertPem               string    `json:"" db:""` // 退款CertPem文件地址
	CertPassword          string    `json:"" db:""` // 退款Cert文件密码
	Net                   *net.Net  // 网络请求包
	CertNet               *net.Net  // 退款Cert网络请求包
}
