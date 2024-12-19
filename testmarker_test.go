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
			m:    Marker{caseName: "Case 1", File: "testmarker_test.go", Line: 1},
			want: "Case 1(testmarker_test.go:1)",
		},
		{
			name: "Case 2",
			m:    Marker{caseName: "Case 2", File: "testmarker_test.go", Line: 2},
			want: "Case 2(testmarker_test.go:2)",
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
