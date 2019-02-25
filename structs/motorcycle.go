package main

import "fmt"

type motorcycle struct {
    throttleResponse uint16
    brakePressure uint16
    rearWheelPSI float64
    topSpeedMph float64
}

func main() {
    panigale := motorcycle{throttleResponse: 22345,
                brakePressure: 934,
                rearWheelPSI: 36.2,
                topSpeedMph: 202}

    fmt.Println(panigale.brakePressure)
}
