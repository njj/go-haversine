package haversine

import (
    "math"
)

type Coords struct {
    StartLatitude  float64
    StartLongitude float64

    EndLatitude    float64
    EndLongitude   float64
}

func toRadians(num float64) float64 {
    return num * math.Pi / 180
}

func (coords Coords) Calculate() float64 {

    // Earth's surface to center
    km   := float64(6371)

    dLat := toRadians(coords.EndLatitude - coords.StartLatitude)
    dLon := toRadians(coords.EndLongitude - coords.StartLongitude)

    lat1 := toRadians(coords.StartLatitude)
    lat2 := toRadians(coords.EndLatitude)

    a    := math.Sin(dLat/2) * math.Sin(dLat/2) +
            math.Sin(dLon/2) * math.Sin(dLon/2) *
            math.Cos(lat1)   * math.Cos(lat2)

    c    := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    return km * c
}
