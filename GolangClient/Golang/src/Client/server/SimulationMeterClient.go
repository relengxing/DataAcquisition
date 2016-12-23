package server

import (
    "fmt"
    "os"
    "net"
    "time"
    "Client/date"
    "Client/utils"
)

func checkError(err error)  {
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
}
/**
    根据IP地址建立TCP连接，
    ip格式："127.0.0.1:7001"
*/
func SimulationMeterClient(addr []byte,ip []byte,keepTime int)  {
    go func() {
        ch := make(chan []byte, 1)
        var str string = utils.Ip2String(ip)
        tcpAddr,err := net.ResolveTCPAddr("tcp",str)
        checkError(err)
        conn,err := net.DialTCP("tcp",nil,tcpAddr)
        checkError(err)
        //获取一个随机数据，然后修改一些东西
        //addr := []byte{0x11,0x11,0x11,0x22,0x22,0x22}
        date := getRandomDate(addr)
        //开一个线程上报数据
        go writeData(conn,date)
        //开一个线程读数据
        go readData(conn,ch)
        for {
            select {
            case <-time.After(time.Duration(keepTime) * time.Second):
                fmt.Println("长时间未接收到信息，断开连接")
                goto EndClient
            case date:=<-ch:
                fmt.Printf("读取到数据%x",date)
            //对数据帧进行处理

            }
        }
        EndClient:
        conn.Close()
    }()
}
/**
    获取随机数据
*/
func getRandomDate(addr []byte) []byte {
    buf := make([]byte,0,1024)
    buf = append(buf,0x68)      //帧头
    buf = append(buf,0x10)      //类型
    buf = append(buf,addr[0])   //地址域
    buf = append(buf,addr[1])   //地址域
    buf = append(buf,addr[2])   //地址域
    buf = append(buf,addr[3])   //地址域
    buf = append(buf,addr[4])   //地址域
    buf = append(buf,addr[5])   //地址域
    buf = append(buf,0x00)      //控制码
    buf = append(buf,0x00)      //数据长度
    buf = append(buf,0x10)      //数据长度
    buf = append(buf,0x00)      //SEQ
    buf = append(buf,0x1c)      //数据标识
    buf = append(buf,0x20)      //数据标识
    buf = append(buf,utils.ByteHexToBcd(time.Now().Year()))      //时间
    buf = append(buf,utils.ByteHexToBcd(int(time.Now().Month())))      //时间
    buf = append(buf,utils.ByteHexToBcd(time.Now().Day()))      //时间
    buf = append(buf,utils.ByteHexToBcd(time.Now().Hour()))      //时间
    buf = append(buf,utils.ByteHexToBcd(time.Now().Minute()))      //时间
    buf = append(buf,date.GetRandomBcd(10,20))      //温度
    buf = append(buf,date.GetRandomBcd(0,99))      //温度
    buf = append(buf,date.GetRandomBcd(0,99))      //湿度
    buf = append(buf,date.GetRandomBcd(0,99))      //湿度
    buf = append(buf,date.GetRandomBcd(0,99))      //电压
    buf = append(buf,date.GetRandomBcd(0,99))      //电压
    buf = append(buf,0x55)      //状态码
    buf = append(buf,0xaa)      //状态码
    buf = append(buf,0x12)      //CRC
    buf = append(buf,0x34)      //CRC
    buf = append(buf,0x16)      //结束符
    return buf[:]
}

/**
    发送数据
*/
func writeData(conn *net.TCPConn,b []byte)  {
    _,err := conn.Write(b)
    if err != nil {
        fmt.Println("发送失败")
        fmt.Println(err)
        conn.Close()
    }
}
/**
    读取数据，并且进行数据帧的预处理，数据处理好了以后再发给SimulationMeterClient
    发送给SimulationMeterClient的是完整、正确并经过验证了的数据帧
*/
func readData(conn *net.TCPConn,readchan chan []byte){
    buf := make([]byte, 1024)
    for  {
        c, err := conn.Read(buf)
        if err != nil {
            fmt.Println("读取服务器数据异常:", err.Error())
            conn.Close()
            break
        }
        //数据帧解析
        //fmt.Println("读取服务器数据:",buf[0:c])
        for i,x := range buf[0:c]{
            if x == 0x68 {      //判断帧头
                var lenH int = int(buf[i+9])        //获取数据域长度
                var lenL int = int(buf[i+10])
                len := lenH<<8 + lenL
                if buf[i+13+len] == 0x16 {      //判断帧尾
                    //只有帧格式正确才给上级代码，不去管帧的具体内容
                    readchan <- buf[i:i+13+len+1]
                    break
                }
            }
        }
    }
}

