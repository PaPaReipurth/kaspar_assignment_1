package server

import (
	"math"
	"testing"

	"golang.org/x/net/context"

	"github.com/papareipurth/kaspar_assignment_1/link_station"
)

const (
	MARGIN float32 = .2
)

func TestGetNearest(t *testing.T) {
	s := LinkStationServer{}

	tests := []struct {
		point *link_station.Point
		want  float32
	}{
		{
			point: &link_station.Point{
				X: 0, Y: 0,
			},
			want: 100,
		},
		{
			point: &link_station.Point{
				X: 100, Y: 100,
			},
			want: 0,
		},
		{
			point: &link_station.Point{
				X: 15, Y: 10,
			},
			want: 0.67,
		},
		{
			point: &link_station.Point{
				X: 18, Y: 18,
			},
			want: 4.71,
		},
	}

	for _, tt := range tests {
		resp, _ := s.GetNearest(context.Background(), tt.point)
		if math.Abs(float64(resp.GetP()-tt.want)) > float64(MARGIN) {
			t.Errorf("GetNearest(%v) got %v expected %v", tt.point, resp.GetP(), tt.want)
		}
	}
}
