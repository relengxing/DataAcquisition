package main

import (
    "Client/utils/MysqlORM"
    "Client/route"
    "Client/controllers"
    "Client/task"
    "fmt"

    "Client/utils"
    "os"
    "strings"
    "reflect"
)

func main() {
    //打开数据库连接
    MysqlORM.Connect("root","1234","data_acquisition")
    //fmt.Println(models.GetAllSimulationMeter())
    //go server.SimulationMeterClient("127.0.0.1:7001",10)
    //go server.SimulationMeterClient("127.0.0.1:7001",10)
    //fmt.Printf("0x%x\n",date.GetRandomBcd(0,10))
    //
    test2()
    test1()
    //test3()
}
func test1()  {
    route.StaticRoute.SetStaticRouteDir("/css/","static")
    route.StaticRoute.SetStaticRouteDir("/js/","static")
    route.StaticRoute.SetStaticRouteDir("/html/","static")
    route.StaticRoute.SetStaticRouteDir("/img/","static")
    homeCtrl := &controllers.HomeController{}
    homeCtrl.Child= homeCtrl
    route.RESTfulRoute.SetRESTfulRouteDir("/home", homeCtrl)
    meterCtrl := &controllers.MeterController{}
    meterCtrl.Child = meterCtrl
    route.RESTfulRoute.SetRESTfulRouteDir("/meter",meterCtrl)
    route.FixedRoute.SetFixedRouteDir("/404",route.NotfoundHandler)
    route.FixedRoute.SetFixedRouteDir("/",route.WelcomeHandler)
    route.Route.Start(":9090")
}

func test2()  {
    task1 := &task.BaseTask{TaskMethod: task.PollingTask,
        TaskParam: nil,
        TaskTime: 0,
        TaskCycle: 60}
    taskManage := task.NewTaskManage()
    taskManage.RegisterTask("任务1",task1)
    taskManage.Run()
    fmt.Println("上报任务启动")
}

type ByteTest struct {
    ByteList []uint8 `json:"date"`
}
func test3()  {
    bt := &ByteTest{}
    bt.ByteList = []uint8{129,1,2,3}
    fmt.Println(*bt)
    utils.OutputJson(os.Stdout,1,"访问地址成功",bt)
    //meterList := models.GetAllSimulationMeter()
    //fmt.Println(meterList)
    //jsonList := controllers.GetAllSimulationMeterForJson(meterList)
    //utils.OutputJson(os.Stdout,1,"访问地址成功",jsonList)
}

func test4()  {
    //b := []byte{0x00, 0x00, 0x03, 0xe8}
    //b_buf := bytes.NewBuffer(b)
    //var x int32
    //binary.Read(b_buf, binary.BigEndian, &x)
    //fmt.Println(x)
    //
    //fmt.Println(strings.Repeat("-", 100))
    str := "127.0.0.1:7001"
    fmt.Println(utils.String2IP(str))

    str1 := "1b-19-10-11-10-11"
    addrSeg :=  strings.Split(str1, "-") //string 字符串 ["11" "11" "11" "00" "00" "00"]
    bytes := []byte(addrSeg[0])
    fmt.Println(reflect.TypeOf(addrSeg[0]))
    fmt.Printf("%x",bytes)
    fmt.Println("================")
    fmt.Println(controllers.DealAddress(str1))

}