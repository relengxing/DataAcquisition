package controllers

import (
    "fmt"
    "html/template"
)

type HomeController struct {
    controller
}

func (this HomeController)Get(){
    //meterList:=models.GetAllSimulationMeter()
    //fmt.Println(meterList)
    //utils.OutputJson(this.W,1,"访问地址成功",meterList)
    t,err := template.ParseFiles("static/html/home.html")
    if err != nil {
        fmt.Println("Home",err)
    }
    t.Execute(this.W,nil)
}
