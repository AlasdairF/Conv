package conv

const (
 digits01 = "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
 digits10 = "0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999"
)

type Buf struct {
 buf []byte
}

func NewConverter() *Buf {
	obj := new(Buf)
	obj.buf = make([]byte, 20)
	return obj
}

func (obj *Buf) String(u int) string {
	return string(obj.format(u, 0))
}

func (obj *Buf) StringPad(u int, p int) string {
	return string(obj.format(u, p))
}

func (obj *Buf) Bytes(u int) []byte {
	return obj.format(u, 0)
}

func (obj *Buf) BytesPad(u int, p int) []byte {
	return obj.format(u, p)
}

func (obj *Buf) format(u int, padding int) []byte {

	var neg bool
	if u < 0 {
		neg = true
		u = -u
	}

	var q int
	var j uintptr
	a := obj.buf
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
		return a[i:]
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
	
	return a[i:]
}

func String(u int) string {
	return string(format(u, 0))
}

func StringPad(u int, p int) string {
	return string(format(u, p))
}

func Bytes(u int) []byte {
	return format(u, 0)
}

func BytesPad(u int, p int) []byte {
	return format(u, p)
}

func format(u int, padding int) []byte {

	var neg bool
	if u < 0 {
		neg = true
		u = -u
	}

	var q int
	var j uintptr
	var a [20]byte
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
		return a[i:]
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
	
	return a[i:]
}

func Int(a []byte) (result int) {
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
	var m uint = 1
	for i:=len(a)-1; i>=0; i-- {
		result += uint(a[i]-48) * m
		m *= 10
	}
	return result
}
