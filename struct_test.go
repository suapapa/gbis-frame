package main

import(
	"sort"
	"testing"
)

func TestSortBusArrivalList(t *testing.T) {
	l := busArrivalList{
		busArrival{
			RouteID: "r2",
			PredictTime1: "3",
			PredictTime2: "5",
		},
		busArrival{
			RouteID: "r1",
			PredictTime1: "2",
			PredictTime2: "5",
		},
	}
	sort.Sort(l)
	if l[0].RouteID != "r1" || l[1].RouteID != "r2" {
		t.Error("sort not working!")
	}
}