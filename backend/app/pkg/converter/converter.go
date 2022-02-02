package converter

import (
	"math"
	"strings"
)

type Converter struct {
	alphabetCapacity  int64
	alphabetMap       map[int64]string
	reversAlphabetMap map[string]int64
}

const Base62Alphabet = "UwJHRu6IOmEvxlVjZTe8h9oLYi0PfkF5GbzsqAD7KNyX3ar41cd2QCtMpSWngB"

func NewDefaultDecimalConverter() *Converter {
	return NewDecimalConverter([]byte(Base62Alphabet))
}
func NewDecimalConverter(alphabet []byte) *Converter {
	alphabetCapacity := len(alphabet)
	alphabetMap := make(map[int64]string, alphabetCapacity)
	reversAlphabetMap := make(map[string]int64, alphabetCapacity)
	for id, symbol := range alphabet {
		idInt64, letter := int64(id), string(symbol)
		alphabetMap[idInt64], reversAlphabetMap[letter] = letter, idInt64
	}
	return &Converter{alphabetMap: alphabetMap, reversAlphabetMap: reversAlphabetMap, alphabetCapacity: int64(alphabetCapacity)}
}
func (c *Converter) Encode(n int64) string {
	var out []string
	for n > c.alphabetCapacity {
		out = append(out, c.alphabetMap[n%c.alphabetCapacity])
		n /= c.alphabetCapacity
	}
	out = append(out, c.alphabetMap[n])

	out = reverseStringSlice(out)
	sBuilder := strings.Builder{}
	for _, letter := range out {
		sBuilder.WriteString(letter)
	}

	return sBuilder.String()
}
func (c *Converter) Decode(in string) int64 {
	set := reverseStringSlice(strings.Split(in, ""))
	var sum int64
	for i, k := range set {
		sum += c.reversAlphabetMap[k] * int64(math.Pow(float64(c.alphabetCapacity), float64(i)))
	}
	return sum
}
func reverseStringSlice(in []string) []string {
	inLength := len(in)
	halfLength := inLength / 2
	inLength--
	for i := 0; i < halfLength; i++ {
		in[i], in[inLength-i] = in[inLength-i], in[i]
	}
	return in
}
