package utils

/**
    二进制数转BCD码
*/
func ByteHexToBcd(Source int) byte{
    var Temp1, Temp2 int;
    Temp1 = (Source %100)/ 10;		//修改，舍去百位数字的值
    Temp2 = Source % 10;
    Temp1 = (Temp1 << 4) + Temp2;
    return byte(Temp1);
}
