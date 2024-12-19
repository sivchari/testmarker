package testmarker

import "testing"

// NOTE: This example is not working because the Example test can't use the testing.T
// Thus, I don't use Example test correctly here. This is just for demonstration.

func ExampleMarker() {
	t := &testing.T{}
	tests := []struct {
		name Marker
		arg  int
		want int
	}{
		{
			name: Mark("Case 1"),
			arg:  1,
			want: 1,
		},
		{
			name: Mark("Case 2"),
			arg:  2,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name.String(), func(t *testing.T) {
			t.Parallel()
			if got := tt.arg + 1; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
