package gofaker

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestIp(t *testing.T) {
	assertion := assert.New(t)
	var segs []int
	for i := 0; i <= 255; i++ {
		segs = append(segs, i)
	}

	batch(1000, func() {
		ip := Ip()
		ipSeg := strings.Split(ip, ".")
		assertion.Len(ipSeg, 4)

		s, err := strconv.Atoi(ipSeg[0])
		assertion.Nil(err)
		assertion.Contains(segs, s)

		s, err = strconv.Atoi(ipSeg[1])
		assertion.Nil(err)
		assertion.Contains(segs, s)

		s, err = strconv.Atoi(ipSeg[2])
		assertion.Nil(err)
		assertion.Contains(segs, s)

		s, err = strconv.Atoi(ipSeg[3])
		assertion.Nil(err)
		assertion.Contains(segs, s)
	})
}

func TestProtocol(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		assertion.Contains(protocols, Protocol())
	})
}

func TestTld(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		assertion.Contains(tlds, Tld())
	})
}

func TestDomain(t *testing.T) {
	assertion := assert.New(t)

	batch(1000, func() {
		c := strings.Count(Domain(""), ".")
		assertion.True(c == 1 || c == 2)
	})
}

func TestEmail(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		assertion.Contains(Email(""), "@")
	})
}

func TestUrl(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		url := Url("","")
		assertion.Contains(url, "://")
		protocol := strings.Split(url, "://")[0]
		assertion.True(len(protocol)>0)
	})
}