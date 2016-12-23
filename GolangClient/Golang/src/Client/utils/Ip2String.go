package utils

import (
    "strconv"

    "strings"
    "bytes"
    "encoding/binary"
)

/**
    字节数组的IP转换为字符串
    例如：[127,0,0,1,0x1b,0x59]
    转为"127.0.0.1:7001"
*/
func Ip2String(ip []byte) string{
    var a1 int = int(ip[0])
    var a2 int = int(ip[1])
    var a3 int = int(ip[2])
    var a4 int = int(ip[3])
    var port1 int = int(ip[4])<<8
    var port2 int = int(ip[5])
    port := port1+port2
    str := strconv.Itoa(a1)+"."+
            strconv.Itoa(a2)+"."+
            strconv.Itoa(a3)+"."+
            strconv.Itoa(a4)+":"+
            strconv.Itoa(port)
    return str
}
/**
    字节数组的IP转换为字符串
    例如："127.0.0.1:7001"
    转为[127,0,0,1,0x1b,0x59]
*/
func String2IP(ip string)[]byte  {
    ipv4:=make([]byte,0)
    strArr :=  strings.Split(ip, ":")
    ipSeg := strings.Split(strArr[0], ".")      //string 字符串 ["127" "0" "0" "1"]
    ip0,_ := strconv.Atoi(ipSeg[0])
    ip1,_ := strconv.Atoi(ipSeg[1])
    ip2,_ := strconv.Atoi(ipSeg[2])
    ip3,_ := strconv.Atoi(ipSeg[3])
    port,_ := strconv.Atoi(strArr[1])           //int   7001
    ipv4 = append(ipv4,
        Int2byteArr(ip0)[3],
        Int2byteArr(ip1)[3],
        Int2byteArr(ip2)[3],
        Int2byteArr(ip3)[3],
        Int2byteArr(port)[2],
        Int2byteArr(port)[3])
    return ipv4
}


func Int2byteArr(x int)[]byte{
    b_buf := bytes.NewBuffer([]byte{})
    binary.Write(b_buf, binary.BigEndian, int32(x))
    return b_buf.Bytes()
}