package controllers

import (
    "Client/models"
    "Client/utils"
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

type MeterController struct {
    controller
}


/**
    下面的代码是为了曲线救国。。
    委曲求全啊。。
    具体遇到的问题就是byte数组会转换为base64，这样给前面又不太方便了
*/
type SimulationMeterForJson struct {
    Uid int             `json:"id"`
    Location string     `json:"location"`
    Cycle int           `json:"cycle"`
    MeterType []int   `json:"meterType"`
    Address []int      `json:"address"`
    Ipv4 []int         `json:"ipv4"`
}

func GetAllSimulationMeterForJson(meterList[]models.SimulationMeter) []SimulationMeterForJson {
    jsonList := make([]SimulationMeterForJson,0,1024)
    for _,meter := range meterList{
        jsonMeter := &SimulationMeterForJson{
            meter.Uid,meter.Location,meter.Cycle,
            utils.ByteArr2IntArr(meter.MeterType),
            utils.ByteArr2IntArr(meter.Address),
            utils.ByteArr2IntArr(meter.Ipv4)}
        jsonList = append(jsonList,*jsonMeter)
    }
    return jsonList
}

func (this MeterController)Get(){
    meterList:=models.GetAllSimulationMeter()
    jsonList := GetAllSimulationMeterForJson(meterList)
    utils.OutputJson(this.W,1,"访问地址成功",jsonList)
}

func (this MeterController)Put(){
    fmt.Println("有人用PUT方法访问Meter")

}
func (this MeterController)Post(){
    id := this.R.FormValue("id")
    location := this.R.FormValue("location")
    actType := this.R.FormValue("type")
    address := this.R.FormValue("address")
    ipv4 := this.R.FormValue("ipv4")
    cycle,_ := strconv.Atoi(this.R.FormValue("cycle"))
    meterType := []byte{0x10}
    if  strings.EqualFold("create",actType){    //如果是创建
        fmt.Println("接收到创建请求")
        if location!=""&& address!=""&&ipv4!=""{
            ok := models.AddOneSimulationMeter(location,cycle,meterType,DealAddress(address), DealIpv4(ipv4))
            if ok {
                fmt.Println("创建成功")
            }else {
                fmt.Println("创建失败")
            }
        }else {
            fmt.Println("参数不满足要求")
        }
    } else if strings.EqualFold("update",actType) {
        fmt.Println("接收到修改请求",id)
        idInt,_ := strconv.Atoi(id)
        meter:= &models.SimulationMeter{Uid:idInt,Location:location,MeterType:meterType,
        Address:DealAddress(address),Ipv4:DealIpv4(ipv4),Cycle:cycle}
        models.UpdateOneSimulationMeter(*meter)
    } else if strings.EqualFold("delete",actType) {
        fmt.Println("接收到删除请求",id)
        idInt,_ := strconv.Atoi(id)
        models.DeleteOneSimulationMeter(idInt)
    }
    http.Redirect(this.W,this.R,"/home",http.StatusFound)

}
/**
    识别前端过来的IPV4
*/
func DealAddress(addr string)[]byte {
    address:=make([]byte,0)
    addrSeg :=  strings.Split(addr, "-") //string 字符串 ["11" "11" "11" "00" "00" "00"]
    address = append(address,utils.String2Byte(addrSeg[0]))
    address = append(address,utils.String2Byte(addrSeg[1]))
    address = append(address,utils.String2Byte(addrSeg[2]))
    address = append(address,utils.String2Byte(addrSeg[3]))
    address = append(address,utils.String2Byte(addrSeg[4]))
    address = append(address,utils.String2Byte(addrSeg[5]))
    return address
}

/**
    识别前端过来的address
*/
func DealIpv4(ipv4 string) []byte {
    return utils.String2IP(ipv4)
}

