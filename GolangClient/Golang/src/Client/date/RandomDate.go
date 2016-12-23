package date

import (
    "time"
    "math/rand"
    "Client/utils"
)
/**
    0x100~0x300
    代表温度范围10~30度，精确到1位小数
*/
func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}
/**
    返回的是BCD码
*/
func GetRandomBcd(start int,end int) byte  {
    date := utils.ByteHexToBcd(rand.Intn(end)+start)
    return date
}



