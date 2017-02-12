package com.relengxing.utils;

/**
 * Created by relengxing on 2016/12/28.
 */
public class Bcd2Hex {

    /*
    * 只适合8位的
    * */
    public static int Bcd2Hex(int sour){
        int Temp1, Temp2;
        Temp1 = (sour & 0xF0) >> 4;
        Temp2 = sour & 0x0F;
        Temp1 = Temp1 * 10 + Temp2;
        return Temp1;
    }
}
