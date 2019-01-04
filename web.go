package gofaker

import (
	"strconv"
)

func Ip() string {
	return strconv.Itoa(NaturalN(0, 255)) + "." +
		strconv.Itoa(NaturalN(0, 255)) + "." +
		strconv.Itoa(NaturalN(0, 255)) + "." +
		strconv.Itoa(NaturalN(0, 255))
}

func Protocol() string {
	return protocols[IntN(0, len(protocols)-1)]
}

/*
	随机生成一个顶级域名。
	国际顶级域名 international top-level domain-names, iTLDs
	国家顶级域名 national top-level domainnames, nTLDs
	[域名后缀大全](http://www.163ns.com/zixun/post/4417.html)
*/
func Tld() string {
	return tlds[IntN(0, len(tlds)-1)]
}

// 随机生成一个域名
func Domain(tld string) string {
	if tld == "" {
		tld = Tld()
	}
	return Word() + "." + tld
}

// 随机生成一个邮件地址
func Email(domain string) string {
	if domain == "" {
		domain = Domain("")
	}
	pool := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"._-"
	return StrN(pool, 3, 16) + "@" + domain
}

func Url(protocol, host string) string {
	if protocol=="" {
		protocol = Protocol()
	}
	if host =="" {
		host = Domain("")
	}
	return  protocol + "://" + host + "/" + Word()
}