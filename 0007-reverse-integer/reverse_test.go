package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "positive",
			x:    123,
			want: 321,
		},
		{
			name: "negative",
			x:    -123,
			want: -321,
		},
		{
			name: "trailing zero dropped",
			x:    120,
			want: 21,
		},
		{
			name: "zero",
			x:    0,
			want: 0,
		},
		{
			name: "single digit positive",
			x:    7,
			want: 7,
		},
		{
			name: "single digit negative",
			x:    -7,
			want: -7,
		},
		{
			name: "overflow positive returns zero",
			x:    1534236469,
			want: 0,
		},
		{
			name: "max int32 overflows",
			x:    2147483647,
			want: 0,
		},
		{
			name: "min int32 overflows",
			x:    -2147483648,
			want: 0,
		},
		{
			name: "near min int32 no overflow",
			x:    -2147483412,
			want: -2143847412,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)
			if got != tt.want {
				t.Errorf("reverse(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
