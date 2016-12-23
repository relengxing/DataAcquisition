package utils

import (
    "encoding/json"
    "io"
    "Client/models"
)

type Result struct{
    Ret int
    Reason string
    Data interface{}
}


func OutputJson(w io.Writer, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}


func OutputSimulationMeter(w io.Writer,meter models.SimulationMeter){
    b,err := json.Marshal(meter)
    if err != nil {
        return
    }
    w.Write(b)
}