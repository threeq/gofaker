package gofaker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	assertion := assert.New(t)
	trueNum := 0
	falseNum := 0
	batch(100, func() {
		if Bool() {
			trueNum ++
		} else {
			falseNum++
		}
	})

	assertion.True(trueNum > 0)
	assertion.True(falseNum > 0)
}

func TestBoolean(t *testing.T) {
	assertion := assert.New(t)
	trueNum := 0
	falseNum := 0
	batch(100, func() {
		if Boolean(1, 9, true) {
			trueNum ++
		} else {
			falseNum++
		}
	})

	assertion.True(trueNum > 0)
	assertion.True(falseNum > 0)

	trueNum = 0
	falseNum = 0
	batch(100, func() {
		if Boolean(0, 9, true) {
			trueNum ++
		} else {
			falseNum++
		}
	})

	assertion.True(trueNum > 0)
	assertion.True(falseNum > 0)

	trueNum = 0
	falseNum = 0
	batch(100, func() {
		if Boolean(-1, 1, true) {
			trueNum ++
		} else {
			falseNum++
		}
	})

	assertion.True(trueNum > 0)
	assertion.True(falseNum > 0)
}

func TestNatural(t *testing.T) {
	assertion := assert.New(t)
	batch(100, func() {
		assertion.True(Natural() > 0)
	})
}

func TestNaturalN(t *testing.T) {
	assertion := assert.New(t)
	batch(100, func() {
		num := NaturalN(10, 20)
		assertion.True(num >= 10)
		assertion.True(num <= 20)
	})
}

func TestInt(t *testing.T) {
	assertion := assert.New(t)
	set := map[int]bool{}
	batch(100, func() {
		set[Int()] = true
	})

	assertion.True(len(set) > 1)
}

func TestIntN(t *testing.T) {
	assertion := assert.New(t)
	set := map[int]bool{}
	batch(100, func() {
		num := IntN(1, 2)
		assertion.True(num == 1 || num == 2)

		if num < 1 || num > 2 {
			// Error
			assertion.Equal(1, num)
		}

		set[num] = true
	})

	assertion.Equal(2, len(set))

	assertion.Equal(4, IntN(4,4))
}

func TestFloat(t *testing.T) {
	assertion := assert.New(t)
	batch(100, func() {
		assertion.True(Float() > 0)
		assertion.True(Float() < 1)
	})
}

func TestFloatN(t *testing.T) {
	assertion := assert.New(t)
	batch(100, func() {
		assertion.True(FloatN(0.1, 2) > 0.1)
		assertion.True(FloatN(1.2, 2.1) < 2.1)
	})
}

func TestChar(t *testing.T) {
	assertion := assert.New(t)

	var pools = map[string]string{
		"lower":  "abcdefghijklmnopqrstuvwxyz",
		"upper":  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"number": "0123456789",
		"symbol": "!@#$%^&*()[]",
	}

	pools["alpha"] = pools["lower"] + pools["upper"]
	all := pools["lower"] + pools["upper"] + pools["number"] + pools["symbol"]

	batch(100, func() {
		c := Char("")
		assertion.Equal(1, len(c))
		assertion.Contains(all, c)
	})

	batch(100, func() {
		c := Char("alpha")
		assertion.Equal(1, len(c))
		assertion.Contains(pools["alpha"], c)
	})

	batch(100, func() {
		c := Char("number")
		assertion.Equal(1, len(c))
		assertion.Contains(pools["number"], c)
	})
}

func TestStr(t *testing.T) {
	assertion := assert.New(t)
	batch(100, func() {
		s := Str("")
		assertion.True(len(s) > 0)
	})
}

func TestStrN(t *testing.T) {
	assertion := assert.New(t)
	batch(100, func() {
		s := StrN("", 6, 6)
		assertion.Equal(6, len(s))
	})
}
