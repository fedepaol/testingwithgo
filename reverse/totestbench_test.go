package reverse

//START OMIT
import (
	"math/rand"
	"testing"
)

func benchReverse(length int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(RandStringBytes(length))
	}
}

func BenchmarkReverse4(b *testing.B) { benchReverse(4, b)}
func BenchmarkReverse8(b *testing.B) { benchReverse(8, b)}
func BenchmarkReverse16(b *testing.B) { benchReverse(16, b)}
func BenchmarkReverse32(b *testing.B) { benchReverse(32, b)}
func BenchmarkReverse64(b *testing.B) { benchReverse(64, b)}

// END OMIT

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
