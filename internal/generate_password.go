package internal

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	lowers   = "abcdefghijklmnopqrstuvxyz"
	uppers   = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	numbers  = "0123456789"
	specials = "+_-?.@#$%!"
	charSet  = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789+_-?.@#$%!"
)

// if length is less then noUpperCase + noLowerCase + noNumber + noSpecials
// length will be ignored
func GeneratePassword(length, noUpperCase, noLowerCase, noNumber, noSpecials int, ignore string) string {
	// build regex pattern of chars to exclude from password
	pattern := fmt.Sprintf("[%s]", ignore)
	rand.Seed(time.Now().Unix())

	var password strings.Builder

	fillWith := length - (noUpperCase + noLowerCase + noNumber + noSpecials)
	if fillWith >= 0 { // char specification is less then length
		noLowerCase += fillWith
	}

	// get random chars from each lower, upper, number and
	// special string vars
	for i := 0; i < noUpperCase; i++ {
		subSet := regexp.MustCompile(pattern).ReplaceAllString(uppers, "")
		index := rand.Intn(len(subSet))
		password.WriteByte(subSet[index])
	}
	for i := 0; i < noNumber; i++ {
		subSet := regexp.MustCompile(pattern).ReplaceAllString(numbers, "")
		index := rand.Intn(len(subSet))
		password.WriteByte(subSet[index])
	}
	for i := 0; i < noSpecials; i++ {
		subSet := regexp.MustCompile(pattern).ReplaceAllString(specials, "")
		index := rand.Intn(len(subSet))
		password.WriteByte(subSet[index])
	}

	for i := 0; i < noLowerCase; i++ {
		subSet := regexp.MustCompile(pattern).ReplaceAllString(lowers, "")
		index := rand.Intn(len(subSet))
		password.WriteByte(subSet[index])
	}
	return password.String()

}
