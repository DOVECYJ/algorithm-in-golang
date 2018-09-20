// big_int_product project main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 991086334737101014141447823434124179111084831929311089596731075933473557543639836
// 73378335716510109111963597814326771110829569262663543213269127839633
func main() {
	x := "991086334737101014141447823434124179111084831929311089596731075933473557543639836"
	y := "73378335716510109111963597814326771110829569262663543213269127839633"
	//fmt.Scanf("%s %s\n", &x, &y)
	fmt.Println(x, y)
	fmt.Println("[main]:", sproduct(x, y))
}

func big_int_product(x, y string) string {
	return sproduct(x, y)
}

func split(x string) (a, b string, halfn int) {
	n := len(x)
	halfn = n / 2
	a, b = x[0:halfn], x[halfn:]
	halfn = n - halfn
	return
}

func split_by_n(x string, n int) (string, string) {
	n = len(x) - n
	if n <= 0 {
		return "", x
	}
	return x[0:n], x[n:]
}

func sadd(x, y string) (sum string) {
	/**
	* 0~9 : 48 49 59 51 52 53 54 55 56 57
	* 0+1+9=48+49+57=154
	* 1+9=49+57=106
	**/
	lx, ly := len(x), len(y)
	a, b, carry := byte(48), byte(48), byte(48)
	for i := 1; i <= lx || i <= ly; i++ {
		if i > lx {
			a, b = 48, y[ly-i]
		} else if i > ly {
			a, b = x[lx-i], 48
		} else {
			a, b = x[lx-i], y[ly-i]
		}
		s := a + b + carry
		if s >= 154 {
			sum = string(s-106) + sum
		} else {
			sum = string(s-96) + sum
		}
		carry = s/154 + 48
	}
	if carry != 48 {
		sum = string(carry) + sum
	}
	return
}

func sproduct(x, y string) string {
	if len(x) < 10 && len(y) < 10 {
		x_int, _ := strconv.Atoi(x)
		y_int, _ := strconv.Atoi(y)
		return fmt.Sprintf("%d", x_int*y_int)
	}

	halfn := max(len(x), len(y)) / 2
	a, b := split_by_n(x, halfn)
	c, d := split_by_n(y, halfn)

	z2 := sproduct(a, c)                       //a×c
	z0 := sproduct(b, d)                       //b×d
	z1 := sadd(sproduct(a, d), sproduct(b, c)) //a×d+b×c

	for i := 0; i < halfn; i++ {
		z2 += "00"
		z1 += "0"
	}
	//fmt.Println("[z-210]", z2, z1, z0)
	return strings.TrimLeft(sadd(sadd(z2, z1), z0), "0")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
