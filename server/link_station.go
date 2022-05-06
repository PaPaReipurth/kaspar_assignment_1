package server

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/papareipurth/kaspar_assignment_1/link_station"
	"golang.org/x/net/context"
)

type LinkStationServer struct {
	link_station.UnimplementedLinkStationServiceServer
}

type LinkStationAttr struct {
	X int
	Y int
	R int
}

func (lsa *LinkStationAttr) GetDistance(x, y float64) float64 {
	// d = sqrt((x2-x1)^2+(y2-y1)^2)
	return math.Sqrt(math.Pow(x-float64(lsa.X), 2) + math.Pow(y-float64(lsa.Y), 2))
}

var (
	linkStationsAttr = []*LinkStationAttr{
		{
			X: 0, Y: 0, R: 10,
		},
		{
			X: 20, Y: 20, R: 5,
		},
		{
			X: 10, Y: 0, R: 12,
		},
	}
)

func (s *LinkStationServer) GetNearest(ctx context.Context, point *link_station.Point) (*link_station.LinkStation, error) {
	bestLinkStation := &link_station.LinkStation{}

	for _, linkStationAttr := range linkStationsAttr {
		var p float32

		if d := linkStationAttr.GetDistance(float64(point.GetX()), float64(point.GetY())); d < float64(linkStationAttr.R) {
			// p = (reach-distance)^2
			p = float32(math.Pow(float64(linkStationAttr.R)-d, 2))
		} else {
			p = 0.0
		}

		if p > bestLinkStation.P {

			bestLinkStation = &link_station.LinkStation{
				X: int32(linkStationAttr.X),
				Y: int32(linkStationAttr.Y),
				R: int32(linkStationAttr.R),
				P: p,
			}

		}

	}

	// Not out of reach
	if bestLinkStation.GetP() != 0 {
		log.Printf("Best link station within reach for point %d, %d if %d, %d with power %v", point.GetX(), point.GetY(), bestLinkStation.GetX(), bestLinkStation.GetY(), bestLinkStation.GetP())
		return bestLinkStation, nil
	} else {
		log.Printf("No link station within reach for point %d, %d", point.GetX(), point.GetY())
		return &link_station.LinkStation{}, errors.New(fmt.Sprintf("No link station within reach for point %d, %d", point.GetX(), point.GetY()))
	}
}
