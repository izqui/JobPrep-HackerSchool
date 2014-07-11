package main

import (
	"fmt"
	"strconv"
)

func main() {

	v, _ := decode([]byte("d5:hellold2:hi3:bye2:wili30ei19eeeee"))

	fmt.Println(v)
}

func decode(b []byte) (interface{}, []byte) {

	i := 0
	switch b[i] {
	case 'i':

		p := []byte{}
		for ; b[i+1] != 'e'; i++ {
			p = append(p, b[i+1])
		}

		n, _ := strconv.Atoi(string(p))

		return n, b[i+2:] //Don't return the 'e'

	case 'l':

		l := []interface{}{}
		r := b[1:]

		for r[0] != 'e' {

			var item interface{}
			item, r = decode(r)
			l = append(l, item)
		}

		return l, r[1:]

	case 'd':

		d := map[string]interface{}{}

		r := b[1:]

		for r[0] != 'e' {

			var key, value interface{}
			key, r = decode(r)
			value, r = decode(r)

			d[key.(string)] = value
		}

		return d, r[1:]

	default:

		i := 0

		p := []byte{}
		for ; b[i] != ':'; i++ {
			p = append(p, b[i])
		}

		i += 1 //Ignore the colon

		n, _ := strconv.Atoi(string(p))
		n += i

		s := []byte{}
		for ; i < n; i++ {
			s = append(s, b[i])
		}

		return string(s), b[i:]

	}

	panic("Unknown")
}
