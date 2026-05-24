package zigzag

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "three rows",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "four rows",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "single row returns input unchanged",
			s:       "ABCDE",
			numRows: 1,
			want:    "ABCDE",
		},
		{
			name:    "single character",
			s:       "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "more rows than characters",
			s:       "AB",
			numRows: 5,
			want:    "AB",
		},
		{
			name:    "empty string",
			s:       "",
			numRows: 3,
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.s, tt.numRows)
			if got != tt.want {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}
