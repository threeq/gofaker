package gofaker

import (
	"math/rand"
	"strconv"
	"strings"
)

type Addr struct {
	id       string
	pid      string
	name     string
	children []*Addr
}

func tree(addrs []*Addr) []*Addr {
	mapped := map[string]*Addr{}
	for _, addr := range addrs {
		mapped[addr.id] = addr
	}

	var result []*Addr
	for _, addr := range addrs {
		if addr.pid == "" {
			result = append(result, addr)
			continue
		}
		pAddr := mapped[addr.pid]
		if pAddr == nil {
			continue
		}
		if pAddr.children == nil {
			pAddr.children = []*Addr{addr}
		} else {
			pAddr.children = append(pAddr.children, addr)
		}

	}
	return result
}

func buildAddrTree() []*Addr {
	var fixed []*Addr
	for id, name := range addressDict {

		getPid(id)

		fixed = append(fixed, &Addr{
			id:   id,
			pid:  getPid(id),
			name: name,
		})
	}
	return tree(fixed)
}

func getPid(id string) string {
	if id[2:] == "0000" {
		return ""
	} else if id[4:] == "00" {
		return id[0:2] + "0000"
	} else {
		return id[0:4] + "00"
	}
}

var addrDict = buildAddrTree()
var addrRegion = []string{"东北", "华北", "华东", "华中", "华南", "西南", "西北"}

// 随机生成一个大区
func Region() string {
	return addrRegion[rand.Intn(len(addrRegion))]
}

// 随机生成一个（中国）省（或直辖市、自治区、特别行政区）
func Province() string {
	return pickAddr(addrDict).name
}

// 随机生成一个（中国）市
func City(prefix bool) string {
	var province = pickAddr(addrDict)
	var city = pickAddr(province.children)

	if prefix {
		return strings.Join([]string{province.name, city.name}, " ")
	}
	return city.name
}

// 随机生成一个（中国）县
func County(prefix bool) string {
	var province = pickAddr(addrDict)
	var city = pickAddr(province.children)
	var county = pickAddr(city.children)

	if county == nil {
		county = &Addr{
			name: "-",
		}
	}

	if prefix {
		return strings.Join([]string{province.name, city.name, county.name}, " ")
	}
	return county.name
}

// 随机生成一个邮政编码（六位数字）
func Zip() string {
	var zip = ""
	for i := 0; i < 6; i++ {
		zip += strconv.Itoa(rand.Intn(10))
	}
	return zip
}

func pickAddr(addrs []*Addr) *Addr {
	if addrs == nil || len(addrs) == 0 {
		return nil
	}
	return addrs[rand.Intn(len(addrs))]
}
