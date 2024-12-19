package testmarker

import "testing"

func TestMarkerString(t *testing.T) {
	tests := []struct {
		name string
		m    Marker
		want string
	}{
		{
			name: "Case 1",
			m:    Mark("Case 1"),
			want: "Case 1(testmarker_test.go:13)",
		},
		{
			name: "Case 2",
			m:    Mark("Case 2"),
			want: "Case 2(testmarker_test.go:18)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Marker.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
