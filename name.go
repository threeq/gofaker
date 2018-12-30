package gofaker

import "strings"

var firstNames = []string{
	// male
	"James", "John", "Robert", "Michael", "William",
	"David", "Richard", "Charles", "Joseph", "Thomas",
	"Christopher", "Daniel", "Paul", "Mark", "Donald",
	"George", "Kenneth", "Steven", "Edward", "Brian",
	"Ronald", "Anthony", "Kevin", "Jason", "Matthew",
	"Gary", "Timothy", "Jose", "Larry", "Jeffrey",
	"Frank", "Scott", "Eric",

	// female
	"Mary", "Patricia", "Linda", "Barbara", "Elizabeth",
	"Jennifer", "Maria", "Susan", "Margaret", "Dorothy",
	"Lisa", "Nancy", "Karen", "Betty", "Helen",
	"Sandra", "Donna", "Carol", "Ruth", "Sharon",
	"Michelle", "Laura", "Sarah", "Kimberly", "Deborah",
	"Jessica", "Shirley", "Cynthia", "Angela", "Melissa",
	"Brenda", "Amy", "Anna",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones",
	"Miller", "Davis", "Garcia", "Rodriguez", "Wilson",
	"Martinez", "Anderson", "Taylor", "Thomas", "Hernandez",
	"Moore", "Martin", "Jackson", "Thompson", "White",
	"Lopez", "Lee", "Gonzalez", "Harris", "Clark",
	"Lewis", "Robinson", "Walker", "Perez", "Hall",
	"Young", "Allen",
}

// 随机生成一个常见的英文名
func NameFirst() string {
	return firstNames[IntN(0, len(firstNames)-1)]
}

func NameLast() string {
	return lastNames[IntN(0, len(lastNames)-1)]
}

func Name(middle bool) string {
	mid := ""
	if middle {
		mid = NameFirst() + " "
	}

	return NameFirst() + " " + mid + NameLast()
}

/*
	随机生成一个常见的中文姓。
	[世界常用姓氏排行](http://baike.baidu.com/view/1719115.htm)
	[玄派网 - 网络小说创作辅助平台](http://xuanpai.sinaapp.com/)
*/
var cnFirstNames = strings.Split(
	"王 李 张 刘 陈 杨 赵 黄 周 吴 "+
		"徐 孙 胡 朱 高 林 何 郭 马 罗 "+
		"梁 宋 郑 谢 韩 唐 冯 于 董 萧 "+
		"程 曹 袁 邓 许 傅 沈 曾 彭 吕 "+
		"苏 卢 蒋 蔡 贾 丁 魏 薛 叶 阎 "+
		"余 潘 杜 戴 夏 锺 汪 田 任 姜 "+
		"范 方 石 姚 谭 廖 邹 熊 金 陆 "+
		"郝 孔 白 崔 康 毛 邱 秦 江 史 "+
		"顾 侯 邵 孟 龙 万 段 雷 钱 汤 "+
		"尹 黎 易 常 武 乔 贺 赖 龚 文",
	" ")
var cnLastNames = strings.Split(
	"伟 芳 娜 秀英 敏 静 丽 强 磊 军 "+
		"洋 勇 艳 杰 娟 涛 明 超 秀兰 霞 "+
		"平 刚 桂英 树 树军 俊 元纬 正伟 "+
		"皓轩",
	" ")

func NameFirstCN() string {
	return cnFirstNames[IntN(0, len(cnFirstNames)-1)]
}

func NameLastCN() string {
	return cnLastNames[IntN(0, len(cnLastNames)-1)]
}

func NameCN() string {
	return NameFirstCN() + NameLastCN()
}
