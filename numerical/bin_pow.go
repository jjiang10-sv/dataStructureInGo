package numerical

// BinPow evaluates (base ^ deg) % rem
func BinPow(base int, deg int, rem int) int {
	var res = 1
	for deg > 0 {
		if (deg & 1) > 0 {
			res = int(int64(res) * int64(base) % int64(rem))
		}
		base = int((int64(base) * int64(base)) % int64(rem))
		deg >>= 1
	}
	return res
}
// bug : int datatype overflow
func BinPow0520(base int, exp int) int {
	res := 1
	for exp > 0 {
		if (exp & 1) > 0 {
			res *= base
		}
		base *= base
		exp >>= 1
	}
	return res
}

func BinPowMod0520(base, exp,mod int) int {
	res := 1
	for exp > 0 {
		if (exp & 1) > 0 {
			res = int(int64(res)*int64(base)%int64(mod))
		}
		base = int(int64(base)*int64(base)%int64(mod))
		exp >>= 1
	}
	return res
}