package generator

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	Count   int
	Length  int
	Charset string
	Prefix  string
	Postfix string
	Pattern string
}

func Generate(c *Config) ([]string, error) {

	if !isFeasable(c.Charset, c.Pattern, c.Count) {
		return nil, errors.New("Not possible to generate requested number of codes.")
	}

	codes := map[string]bool{}
	count := c.Count

	for count > 0 {
		code := generateOne(c)
		if !codes[code] {
			codes[code] = true
			count--
		}
	}

	result := []string{}

	for k := range codes {
		result = append(result, k)
	}

	return result, nil
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func pickRandomElement(arr []string) string {
	return arr[randomInt(0, len(arr)-1)]
}

func Repeat(str string, count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += str
	}
	return res
}

func generateOne(c *Config) string {
	code := ""
	for _, char := range strings.Split(c.Pattern, "") {
		if char == "#" {
			code += pickRandomElement(strings.Split(c.Charset, ""))
			continue
		}
		code += char
	}
	return fmt.Sprintf("%s%s%s", c.Prefix, code, c.Postfix)
}

func isFeasable(charset, pattern string, count int) bool {
	re := regexp.MustCompile(`#`)
	p := 0
	p = len(re.FindAll([]byte(pattern), -1))
	return math.Pow(float64(len(charset)), float64(p)) >= float64(count)
}

func Charset(name string) string {
	charsets := map[string]string{
		"numbers":      "0123456789",
		"alphabetic":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"alphanumeric": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	return charsets[name]
}
