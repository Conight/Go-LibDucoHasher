package main

import "C"
import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
	"time"
)

func newHashSet(lastH string) hash.Hash {
	h := sha1.New()
	h.Write([]byte(lastH))
	return h
}

//export DUCOS1
func DUCOS1(lastH string, expH string, diff int, eff float64) (int, float64) {
	timeStart := float64(time.Now().UnixNano()/100) / float64(10000000)

	for nonce := int(0); nonce < 100*diff+1; nonce++ {
		h := newHashSet(lastH)
		h.Write([]byte(strconv.Itoa(nonce)))
		dRes := hex.EncodeToString(h.Sum(nil))

		if eff != 0 {
			if nonce%5000 == 0 {
				time.Sleep(time.Duration(eff/100*1000000000) * time.Nanosecond)
			}
		}

		if dRes == expH {
			timeElapsed := float64(time.Now().UnixNano()/100)/float64(10000000) - timeStart
			hashRate := float64(nonce) / timeElapsed
			return nonce, hashRate
		}
	}
	return 0, 0
}

//export DUCOS1Nonce
func DUCOS1Nonce(lastH string, expH string, diff int, eff float64) int {
	for nonce := int(0); nonce < 100*diff+1; nonce++ {
		h := newHashSet(lastH)
		h.Write([]byte(strconv.Itoa(nonce)))
		dRes := hex.EncodeToString(h.Sum(nil))

		if eff != 0 {
			if nonce%5000 == 0 {
				time.Sleep(time.Duration(eff/100*1000000000) * time.Nanosecond)
			}
		}

		if dRes == expH {
			return nonce
		}
	}
	return 0
}

func main() {
	timeElapsed, hashRate := DUCOS1(
		"f316f4cd012371b15da767fa66c6c7478bf9593e",
		"98ce69810af441b69971eec6ba4d87b766bf0213",
		500000,
		0.005,
	)

	fmt.Println(timeElapsed, hashRate)
}
