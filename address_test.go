package gofaker

import (
	Assert "github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestPickAddr(t *testing.T) {
	assert := Assert.New(t)
	assert.Nil(pickAddr([]*Addr{}))
	assert.Nil(pickAddr(nil))
	assert.NotNil(pickAddr([]*Addr{{}}))
}

func TestBuildAddrTree(t *testing.T) {
	assert := Assert.New(t)
	addrTree := buildAddrTree()

	assert.Equal("", addrTree[0].pid)
	assert.NotEqual("", addrTree[0].name)
	assert.Equal("0000", addrTree[0].id[2:])
}

func TestRegion(t *testing.T) {
	assert := Assert.New(t)

	batch(100, func() {
		r := Region()
		assert.Contains(addrRegion, r)
	})
}

func TestProvince(t *testing.T) {
	assert := Assert.New(t)
	batch(100, func() {
		name := Province()
		assert.True(len(name)>0)
	})
}

func TestCity(t *testing.T) {
	assert := Assert.New(t)
	batch(100, func() {
		name := City(false)
		assert.True(len(name)>0)
		assert.NotContains(name, " ")
	})
	batch(100, func() {
		name := City(true)
		assert.True(len(name)>0)
		assert.Contains(name, " ")
	})
}

func TestCounty(t *testing.T) {
	assert := Assert.New(t)
	batch(100, func() {
		name := County(false)
		assert.True(len(name)>0)
		assert.NotContains(name, " ")
	})
	reg := regexp.MustCompile( `^\S+ \S+ \S+$`)
	batch(100, func() {
		name := County(true)
		assert.True(len(name)>0)
		assert.Contains(name, " ")
		assert.Regexp(reg, name)
		assert.Len(strings.Split(name, " "), 3)
	})
}

func TestZip(t *testing.T) {
	assert := Assert.New(t)

	reg := regexp.MustCompile( `^\d{6}$`)
	batch(100, func() {
		zipCode := Zip()
		assert.Len(zipCode, 6)
		assert.Regexp(reg, zipCode)
	})
}

