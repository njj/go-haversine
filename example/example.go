package main

import ".."
import "fmt"

func main() {
    coords := haversine.Coords{ 30.849635, -83.24559, 27.950575, -82.457178 }

    fmt.Println(coords.Calculate())
}
