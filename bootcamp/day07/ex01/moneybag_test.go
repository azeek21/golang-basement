package ex01_test

import (
	"moneybag/ex00"
	"testing"
)

// AAA arrange act assert

type TestCase struct {
	Desc    string
	Target  int
	Options []int
	Exp     []int
}

func getCases() TestCase {
	return TestCase{
		Desc:   "duplicate unordered options",
		Target: 13,
		Options: []int{27, 52,
			4,
			97,
			58,
			33,
			7,
			62,
			57,
			8,
			85,
			12,
			92,
			78,
			66,
			74,
			89,
			6,
			3,
			20,
			10,
			90,
			53,
			71,
			84,
			5,
			25,
			79,
			23,
			87,
			40,
			21,
			86,
			17,
			2,
			70,
			41,
			16,
			88,
			59,
			73,
			38,
			93,
			49,
			95,
			37,
			35,
			9,
			36,
			55,
			72,
			34,
			22,
			67,
			15,
			18,
			46,
			61,
			31,
			45,
			50,
			13,
			11,
			91,
			99,
			83,
			24,
			1,
			43,
			32,
			30,
			65,
			75,
			82,
			19,
			28,
			64,
			51,
			81,
			76,
			63,
			14,
			96,
			29,
			68,
			44,
			54,
			39,
			77,
			80,
			48,
			60,
			69,
			56,
			100,
			98,
			42,
			26, 47, 94},
	}

}

func BenchmarkMinCoins(b *testing.B) {
	tcase := getCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ex00.MinCoins(tcase.Target, tcase.Options)
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	tcase := getCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ex00.MinCoins2(tcase.Target, tcase.Options)
	}

}

func BenchmarkMinCoins2Optimized(b *testing.B) {
	tcase := getCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ex00.MinCoins2Optimized(tcase.Target, tcase.Options)
	}

}
