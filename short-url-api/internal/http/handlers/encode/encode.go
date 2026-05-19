// Package encode responsible for hashing the id to insert into the db
package encode

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = uint64(len(alphabet))

func Encode(id uint64) string {
	if id == 0 {
		return string(alphabet[0])
	}

	var sb strings.Builder

	for id > 0 {
		rem := id % base
		sb.WriteByte(alphabet[rem])

		id = id / base
	}

	return reverse(sb.String())
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
