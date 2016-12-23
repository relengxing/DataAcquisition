package models

import (
    "fmt"
    "Client/utils/MysqlORM"
    "errors"
    "log"
)
/**
    模拟表记设置
    DAO层
    增删改查全部测试完毕
    时间：2016年12月19日
*/

type SimulationMeter struct {
    Uid int
    Location string
    Cycle int
    MeterType []byte
    Address []byte
    Ipv4 []byte
}



/**
    获取所有存入数据库的模拟表记设置信息
*/
func GetAllSimulationMeter() []SimulationMeter{
    //MysqlORM.Connect("root","1234","data_acquisition")
    list := make([]SimulationMeter,0)
    rows, err := MysqlORM.Db.Query("SELECT * FROM simulation_meter;")
    if err != nil {
        fmt.Println("fetech data failed:", err.Error())
    }
    //defer rows.Close()
    for rows.Next() {
        var uid int
        var location string
        var cycle int
        var meterType []byte
        var address []byte
        var ipv4 []byte
        rows.Scan(&uid, &location, &cycle,&meterType,&address,&ipv4)
        list = append(list,SimulationMeter{
            Uid:uid,Location:location,Cycle:cycle,MeterType:meterType,Address:address,Ipv4:ipv4})
    }
    return list
}

/**
    增加一个模拟表记
*/
func AddOneSimulationMeter(location string,cycle int,meter_type []byte,address []byte,ipv4 []byte) bool {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Runtime error caught: %v", r)
        }
    }()
    result, err := MysqlORM.Db.Prepare(
        "insert into simulation_meter(location,cycle,meter_type,address,ipv4)values(?,?,?,?,?)")
    if err != nil {
        log.Println(err)
    }
    rs, err := result.Exec(location, cycle,meter_type,address,ipv4)
    if err != nil {
        log.Println(err)
    }
    ////我们可以获得插入的id
    //id, err := rs.LastInsertId()
    ////可以获得影响行数
    affect, err := rs.RowsAffected()
    if err != nil {
        fmt.Println("数据库插入错误",err)
        return false
    }
    if affect == 1{
        return true
    }
    return false
}

/**
    删除一个模拟表记
*/
func DeleteOneSimulationMeter(id int) bool {
    result, err := MysqlORM.Db.Prepare(
        "DELETE FROM `simulation_meter` WHERE id = ?")
    if err != nil {
        log.Println(err)
    }
    rs, err := result.Exec(id)
    if err != nil {
        log.Println(err)
    }
    ////我们可以获得插入的id
    //id, err := rs.LastInsertId()
    ////可以获得影响行数
    affect, err := rs.RowsAffected()
    if affect == 1{
        return true
    }
    return false
}
/**
    修改一个模拟标记
*/
func UpdateOneSimulationMeter(meter SimulationMeter) bool {
    result, err := MysqlORM.Db.Prepare(
        "UPDATE simulation_meter SET location=?,cycle=?,meter_type=?,address=?,ipv4=? WHERE id=?")
    if err != nil {
        log.Println(err)
    }
    rs, err := result.Exec(meter.Location, meter.Cycle,meter.MeterType,meter.Address,meter.Ipv4,meter.Uid)
    if err != nil {
        log.Println(err)
    }
    ////我们可以获得插入的id
    //id, err := rs.LastInsertId()
    ////可以获得影响行数
    affect, err := rs.RowsAffected()
    if affect == 1{
        return true
    }
    return false
}

/**
    查询一个模拟表记
*/
func SelectOneSimulationMeter(id int)  (meter SimulationMeter,err error){
    MysqlORM.Connect("root","1234","data_acquisition")
    rows, err := MysqlORM.Db.Query("SELECT * FROM simulation_meter where id = ?;",id)
    if err != nil {
        fmt.Println("fetech data failed:", err.Error())
    }
    //defer rows.Close()
    for rows.Next() {
        var uid int
        var location string
        var cycle int
        var meterType []byte
        var address []byte
        var ipv4 []byte
        rows.Scan(&uid, &location, &cycle,&meterType,&address,&ipv4)
        meter = SimulationMeter{
            Uid:uid,Location:location,Cycle:cycle,MeterType:meterType,Address:address,Ipv4:ipv4}
        return meter,nil
    }
    return meter,errors.New("数据库中无此数据")
}

