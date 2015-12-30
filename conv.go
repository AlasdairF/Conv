package conv

import (
	"sync"
	"io"
)

const (
 digits01 = "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
 digits10 = "0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999"
)

var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 20)
    },
}

func DecimalMark(num []byte, mark byte) []byte {
	l := len(num)
	l2 := l + ((l - 1) / 3)
	newar := make([]byte, l2)
	var i int
	l2--
	for l--; l>=0; l-- {
		if i++; i == 4 {
			newar[l2] = mark
			l2--
			i = 1
		}
		newar[l2] = num[l]
		l2--
	}
	return newar
}

func String(u int) string {
	return formatString(u, 0)
}

func StringPad(u int, p int) string {
	return formatString(u, p)
}

func Bytes(u int) []byte {
	return format(u, 0)
}

func BytesPad(u int, p int) []byte {
	return format(u, p)
}

func FloatString(f float64, prec int) string {
	return string(FloatBytes(f, prec))
}

func FloatBytes(f float64, prec int) []byte {
	if prec == 0 {
	  return format(int(f), 0)
	}
	u := int(f)
	save := u
	var neg bool
	if u < 0 {
		neg = true
		u = -u
	}

	var q int
	var j uintptr
	a := pool.Get().([]byte)
	defer pool.Put(a)
	i := 19 - prec

	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	if neg {
		i--
		a[i] = '-'
	}
	
	a[19 - prec] = '.'
	switch prec {
		case 1: u = int(f * 10) - (save * 10)
		case 2: u = int(f * 100) - (save * 100)
		case 3: u = int(f * 1000) - (save * 1000)
		case 4: u = int(f * 10000) - (save * 10000)
		case 5: u = int(f * 100000) - (save * 100000)
		case 6: u = int(f * 1000000) - (save * 1000000)
		case 7: u = int(f * 10000000) - (save * 10000000)
		case 8: u = int(f * 100000000) - (save * 100000000)
		case 9: u = int(f * 1000000000) - (save * 1000000000)
	}
	if neg {
		u = -u
	}
	save = i
	
	i = 20
	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	cpy := make([]byte, 20 - save)
	copy(cpy, a[save:])
	return cpy
}

func format(u int, padding int) []byte {
	var neg bool
	if u < 0 {
		neg = true
		u = -u
	} else {
		if u < 10 && padding == 0 {
			return []byte{byte(u) + 48}
		}
	}

	var q int
	var j uintptr
	a := pool.Get().([]byte)
	defer pool.Put(a)
	i := 20

	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	if padding == 0 {
		if neg {
			i--
			a[i] = '-'
		}
		cpy := make([]byte, 20 - i)
		copy(cpy, a[i:])
		return cpy
	}
	
	if neg {
		padding = 21 - padding
	} else {
		padding = 20 - padding
	}
	for i > padding {
		i--
		a[i] = '0'
	}
	if neg {
		i--
		a[i] = '-'
	}
	
	cpy := make([]byte, 20 - i)
	copy(cpy, a[i:])
	return cpy
}

func formatString(u int, padding int) string {
	var neg bool
	if u < 0 {
		neg = true
		u = -u
	} else {
		if u < 10 && padding == 0 {
			switch u {
				case 0: return `0`
				case 1: return `1`
				case 2: return `2`
				case 3: return `3`
				case 4: return `4`
				case 5: return `5`
				case 6: return `6`
				case 7: return `7`
				case 8: return `8`
				case 9: return `9`
			}
		}
	}

	var q int
	var j uintptr
	a := pool.Get().([]byte)
	defer pool.Put(a)
	i := 20

	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	if padding == 0 {
		if neg {
			i--
			a[i] = '-'
		}
		return string(a[i:])
	}
	
	if neg {
		padding = 21 - padding
	} else {
		padding = 20 - padding
	}
	for i > padding {
		i--
		a[i] = '0'
	}
	if neg {
		i--
		a[i] = '-'
	}
	return string(a[i:])
}

func Write(w io.Writer, u int, padding int) (int, error) {
	a := pool.Get().([]byte)
	defer pool.Put(a)
	var neg bool
	if u < 0 {
		neg = true
		u = -u
	} else {
		if u < 10 && padding == 0 {
			a[0] = byte(u) + 48
			return w.Write(a[0:1])
		}
	}

	var q int
	var j uintptr
	i := 20

	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	if padding == 0 {
		if neg {
			i--
			a[i] = '-'
		}
		return w.Write(a[i:])
	}
	
	if neg {
		padding = 21 - padding
	} else {
		padding = 20 - padding
	}
	for i > padding {
		i--
		a[i] = '0'
	}
	if neg {
		i--
		a[i] = '-'
	}
	
	return w.Write(a[i:])
}

func WriteFloat(w io.Writer, f float64, prec int) (int, error) {
	if prec == 0 {
	  return Write(w, int(f), 0)
	}
	u := int(f)
	save := u
	var neg bool
	if u < 0 {
		neg = true
		u = -u
	}

	var q int
	var j uintptr
	a := pool.Get().([]byte)
	defer pool.Put(a)
	i := 19 - prec

	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	if neg {
		i--
		a[i] = '-'
	}
	
	a[19 - prec] = '.'
	switch prec {
		case 1: u = int(f * 10) - (save * 10)
		case 2: u = int(f * 100) - (save * 100)
		case 3: u = int(f * 1000) - (save * 1000)
		case 4: u = int(f * 10000) - (save * 10000)
		case 5: u = int(f * 100000) - (save * 100000)
		case 6: u = int(f * 1000000) - (save * 1000000)
		case 7: u = int(f * 10000000) - (save * 10000000)
		case 8: u = int(f * 100000000) - (save * 100000000)
		case 9: u = int(f * 1000000000) - (save * 1000000000)
	}
	if neg {
		u = -u
	}
	save = i
	
	i = 20
	for u >= 100 {
		i -= 2
		q = u / 100
		j = uintptr(u - q*100)
		a[i+1] = digits01[j]
		a[i] = digits10[j]
		u = q
	}
	if u >= 10 {
		i--
		q = u / 10
		a[i] = digits01[uintptr(u-q*10)]
		u = q
	}
	i--
	a[i] = digits01[uintptr(u)]
	
	return w.Write(a[save:])
}

func Int(a []byte) (result int) {
	if len(a) == 0 {
		return 0
	}
	var neg bool
	if a[0] == '-' {
		neg = true
		a[0] = 48
	}
	var m int = 1
	for i:=len(a)-1; i>=0; i-- {
		result += int(a[i]-48) * m
		m *= 10
	}
	if neg {
		return -result
	}
	return result
}

func Uint(a []byte) (result uint) {
	if len(a) == 0 {
		return 0
	}
	var m uint = 1
	for i:=len(a)-1; i>=0; i-- {
		result += uint(a[i]-48) * m
		m *= 10
	}
	return result
}

func Ints(a []byte) []int {
	pages := make([]int, 0, 3)
	var in bool
	var last int
	for i, b := range a {
		if (b >= '0' && b <= '9') || b == '-' {
			if !in {
				in = true
				last = i
			}
		} else {
			if in {
				pages = append(pages, Int(a[last:i]))
				in = false
			}
		}
	}
	if in {
		pages = append(pages, Int(a[last:]))
	}
	return pages
}

func Uints(a []byte) []uint {
	pages := make([]uint, 0, 3)
	var in bool
	var last int
	for i, b := range a {
		if b >= '0' && b <= '9' {
			if !in {
				in = true
				last = i
			}
		} else {
			if in {
				pages = append(pages, Uint(a[last:i]))
				in = false
			}
		}
	}
	if in {
		pages = append(pages, Uint(a[last:]))
	}
	return pages
}
