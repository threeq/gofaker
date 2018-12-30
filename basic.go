package gofaker

import "math/rand"

// 返回一个随机的布尔值。
func Boolean(min, max int, cur bool) bool {
	if min == 0.0 || (min+max) == 0.0 {
		return Bool()
	}
	temp := 1.0 / float32((min+max)*min)
	if rand.Float32() >= temp {
		return !cur
	}
	return cur
}

func Bool() bool {
	return rand.Float32() >= 0.5
}

// 返回一个随机的自然数（大于等于 0 的整数）
func Natural() int {
	return int(rand.Uint32())
}

func NaturalN(min, max int) int {
	return int(Float()*float32(max-min)) + min
}

// 返回一个随机的整数
func Int() int {
	return rand.Int()
}

// 区间 [min, max]
func IntN(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min+1) + min
}

// 返回一个随机的浮点数
func Float() float32 {
	return rand.Float32()
}

func FloatN(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

// 返回一个随机字符
func Char(pool string) string {
	var pools = map[string]string{
		"lower":  "abcdefghijklmnopqrstuvwxyz",
		"upper":  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"number": "0123456789",
		"symbol": "!@#$%^&*()[]",
	}

	pools["alpha"] = pools["lower"] + pools["upper"]

	if pool == "" {
		pool = pools["lower"] + pools["upper"] + pools["number"] + pools["symbol"]
	} else if pools[pool] != "" {
		pool = pools[pool]
	}
	at := rand.Intn(len(pool))
	return pool[at : at+1]
}

func Str(pool string) string {
	size := int(NaturalN(1, 1000))
	var text = ""
	for i := 0; i < size; i++ {
		text += Char(pool)
	}
	return text
}

func StrN(pool string, min, max int) string {
	size := int(NaturalN(min, max))
	var text = ""
	for i := 0; i < size; i++ {
		text += Char(pool)
	}
	return text
}
