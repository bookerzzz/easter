/*
Package easter is tiny library that returns Easter Sunday date for a year
*/
package easter

import (
	"time"
)

// GregorianEaster returns date for Easter
// https://en.wikipedia.org/wiki/Computus#Software
func GregorianEaster(year int) (date time.Time) {
	a := year % 19
	b := year >> 2
	c := b/25 + 1
	d := (c * 3) >> 2
	e := ((a * 19) - ((c*8 + 5) / 25) + d + 15) % 30
	e += (29578 - a - e*32) >> 10
	e -= ((year % 7) + b - d + e + 2) % 7
	d = e >> 5
	day := e - d*31
	month := d + 3
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

// GregorianAscension returns 40th day of easter
func GregorianAscension(year int) time.Time {
	return GregorianEaster(year).AddDate(0, 0, 39)
}

// GregorianPentecost returns 50th day of easter
func GregorianPentecost(year int) time.Time {
	return GregorianEaster(year).AddDate(0, 0, 49)
}
