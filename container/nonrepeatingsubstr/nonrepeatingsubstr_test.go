package main

import "testing"

// go test -coverprofile=c.out
// go tool cover -html=c.out

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcaaacbb", 3},
		{"abcabffcaaacbb", 4},
		{"abasdcabcaaacbb", 5},
		{"asbcabcaaacbb", 4},
		{"", 0},
		{"bb", 1},
		{"abcd", 4},
	}

	for _, tt := range tests {
		actual := length0fNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("calcTriangle %d ; got %s; expected %d.", actual, tt.s, tt.ans)
		}
	}

}

// go test -bench .
func BenchmarkSubstr(b *testing.B) {
	s := "Yes我爱我的家乡!"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 6

	for i := 0; i < b.N; i++ {
		actual := length0fNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("calcTriangle %d ; got %s; expected %d.", actual, s, ans)
		}
	}
}
