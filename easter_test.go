package easter

import (
	"testing"
	"time"
)

// http://tlarsen2.tripod.com/thomaslarsen/easterdates.html#List20
var (
	easterCases = [][3]int{
		{23, 4, 2000}, {20, 4, 2025}, {10, 4, 2050}, {7, 4, 2075},
		{15, 4, 2001}, {5, 4, 2026}, {2, 4, 2051}, {19, 4, 2076},
		{31, 3, 2002}, {28, 3, 2027}, {21, 4, 2052}, {11, 4, 2077},
		{20, 4, 2003}, {16, 4, 2028}, {6, 4, 2053}, {3, 4, 2078},
		{11, 4, 2004}, {1, 4, 2029}, {29, 3, 2054}, {23, 4, 2079},
		{27, 3, 2005}, {21, 4, 2030}, {18, 4, 2055}, {7, 4, 2080},
		{16, 4, 2006}, {13, 4, 2031}, {2, 4, 2056}, {30, 3, 2081},
		{8, 4, 2007}, {28, 3, 2032}, {22, 4, 2057}, {19, 4, 2082},
		{23, 3, 2008}, {17, 4, 2033}, {14, 4, 2058}, {4, 4, 2083},
		{12, 4, 2009}, {9, 4, 2034}, {30, 3, 2059}, {26, 3, 2084},
		{4, 4, 2010}, {25, 3, 2035}, {18, 4, 2060}, {15, 4, 2085},
		{24, 4, 2011}, {13, 4, 2036}, {10, 4, 2061}, {31, 3, 2086},
		{8, 4, 2012}, {5, 4, 2037}, {26, 3, 2062}, {20, 4, 2087},
		{31, 3, 2013}, {25, 4, 2038}, {15, 4, 2063}, {11, 4, 2088},
		{20, 4, 2014}, {10, 4, 2039}, {6, 4, 2064}, {3, 4, 2089},
		{5, 4, 2015}, {1, 4, 2040}, {29, 3, 2065}, {16, 4, 2090},
		{27, 3, 2016}, {21, 4, 2041}, {11, 4, 2066}, {8, 4, 2091},
		{16, 4, 2017}, {6, 4, 2042}, {3, 4, 2067}, {30, 3, 2092},
		{1, 4, 2018}, {29, 3, 2043}, {22, 4, 2068}, {12, 4, 2093},
		{21, 4, 2019}, {17, 4, 2044}, {14, 4, 2069}, {4, 4, 2094},
		{12, 4, 2020}, {9, 4, 2045}, {30, 3, 2070}, {24, 4, 2095},
		{4, 4, 2021}, {25, 3, 2046}, {19, 4, 2071}, {15, 4, 2096},
		{17, 4, 2022}, {14, 4, 2047}, {10, 4, 2072}, {31, 3, 2097},
		{9, 4, 2023}, {5, 4, 2048}, {26, 3, 2073}, {20, 4, 2098},
		{31, 3, 2024}, {18, 4, 2049}, {15, 4, 2074}, {12, 4, 2099},

		// Dates out of 2000 - 2099
		{11, 4, 1700}, {4, 4, 1999}, {30, 3, 2206}, {26, 3, 2282},

		// Dates way off anything known
		{11, 4, 1902010}, {25, 4, 302010},
	}
	ascensionCases = [][3]int{
		{10, 5, 2018}, {30, 5, 2019},
	}

	pentecostCases = [][3]int{
		{20, 5, 2018}, {9, 6, 2019},
	}
)

func TestGregorianEaster(t *testing.T) {
	for _, ec := range easterCases {
		check := time.Date(ec[2], time.Month(ec[1]), ec[0], 0, 0, 0, 0, time.Local)
		if d := GregorianEaster(ec[2]); d != check {
			t.Errorf("Easter Sunday does not match for year %d: %#v != %#v", ec[2], d, check)
		}
	}
}

func TestGregorianAscension(t *testing.T) {
	for _, ac := range ascensionCases {
		for _, ec := range easterCases {
			if ec[2] != ac[2] {
				continue
			}
			check := time.Date(ac[2], time.Month(ac[1]), ac[0], 0, 0, 0, 0, time.Local)
			if d := GregorianAscension(ec[2]); d != check {
				t.Errorf("Ascension does not match for year %d: %#v != %#v", ac[2], d, check)
			}
		}
	}
}

func TestGregorianPentecost(t *testing.T) {
	for _, pc := range pentecostCases {
		for _, ec := range easterCases {
			if ec[2] != pc[2] {
				continue
			}
			check := time.Date(pc[2], time.Month(pc[1]), pc[0], 0, 0, 0, 0, time.Local)
			if d := GregorianPentecost(ec[2]); d != check {
				t.Errorf("Pentecost does not match for year %d: %#v != %#v", pc[2], d, check)
			}
		}
	}
}
