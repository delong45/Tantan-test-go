package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func Md5(text string) string {
	m := md5.New()
	io.WriteString(m, text)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func GetUid() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	n := rand.Int63()
	id := Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(n, 10)))
	return id
}
