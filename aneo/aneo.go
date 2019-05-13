package aneo

import (
	"fmt"
	"io"
	"math"
)

type trafficLight struct {
	dstFromStartM, lightDurationSec int
}

func solve(in io.Reader, out io.Writer) {
	var maxSpeedKmPerHour, trafficLightsCount int
	trafficLights := []trafficLight{}

	fmt.Fscanf(in, "%d\n", &maxSpeedKmPerHour)
	fmt.Fscanf(in, "%d\n", &trafficLightsCount)
	for i := 0; i < trafficLightsCount; i++ {
		var dstFromStartMIth, lightDurationSecIth int
		fmt.Fscanf(in, "%d %d\n", &dstFromStartMIth, &lightDurationSecIth)
		trafficLights = append(trafficLights, trafficLight{
			dstFromStartMIth, lightDurationSecIth,
		})
	}

	speedKmPerHour := maxSpeedKmPerHour
outer:
	for ; speedKmPerHour > 0; speedKmPerHour-- {
		speedMPerSec := float64(speedKmPerHour*1000) / float64(3600)
		for _, tl := range trafficLights {
			t := float64(tl.dstFromStartM) / speedMPerSec
			t = math.Round(t*100000) / 100000
			n := t / float64(tl.lightDurationSec)
			n = math.Floor(n)
			if int(n)%2 == 1 {
				continue outer
			}
		}
		break
	}

	fmt.Fprint(out, speedKmPerHour)
}
