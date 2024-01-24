package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeatedCharacter := Repeat("a", 3)
	fmt.Println(repeatedCharacter)
	// Output: aaa
}
