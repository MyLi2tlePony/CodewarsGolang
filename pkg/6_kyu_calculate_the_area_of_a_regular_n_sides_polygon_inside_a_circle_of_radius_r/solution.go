package calculate_the_area_of_a_regular_n_sides_polygon_inside_a_circle_of_radius_r

import (
	"math"
)

func solution(circleRadius float64, numberOfSides int) float64 {
	angle := 2 * math.Pi / float64(numberOfSides)
	angleSin := math.Sin(angle)
	triangleS := circleRadius * circleRadius * angleSin / 2
	area := float64(numberOfSides) * triangleS
	return math.Round(area*1000) / 1000
}
