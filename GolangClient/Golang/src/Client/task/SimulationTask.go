package task

import (
    "fmt"
    "Client/server"
    "Client/models"
    "time"
)
/**
    模拟表记任务
*/
func SimulationMeterTask(arg ...interface{}) error {
    var addr []byte
    var ip []byte
    switch v:=arg[0].(type) {
    case []byte: addr = v
    default:
        fmt.Println("类型错误")
    }
    switch v:=arg[1].(type) {
    case []byte: ip = v
    default:
        fmt.Println("类型错误")
    }
    server.SimulationMeterClient(addr,ip,10)

    return nil
}

/**
    轮训表记数据库任务
*/
func PollingTask(arg ...interface{}) error {
    //1.轮训本身是一个任务，每个小时（测试阶段可以修改）执行一次，
    //执行后把所有表记从数据库取出，判断周期和当前时间是否整除，整除则执行
    meterList := models.GetAllSimulationMeter()
    for _,meter:= range meterList {
        runFlag := int(time.Hour)%meter.Cycle
        if runFlag == 0 {
            fmt.Println("一个模拟表记运行......",meter,runFlag)
            go server.SimulationMeterClient(meter.Address,meter.Ipv4,10)
        }
    }
    return nil
}
