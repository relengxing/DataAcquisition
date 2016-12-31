package com.relengxing.server;

import com.relengxing.bean.Date1c20;
import com.relengxing.dao.Date1c20Dao;
import com.relengxing.utils.Bcd2Hex;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;
import java.util.Calendar;
import java.util.Date;

/**
 * Created by relengxing on 2016/12/28.
 */
public class ServerThread implements Runnable {

    private Socket socket;

    public ServerThread(Socket socket) {
        this.socket = socket;
    }

    @Override
    public void run() {
        //获取输入输出流
        try (
                InputStream inputStream = socket.getInputStream();
                OutputStream outputStream = socket.getOutputStream();
                ){
            //帧处理        交给Protocol处理
            byte[] readBuff = new byte[1024];
            int readLength = 0;
            while ((readLength = inputStream.read(readBuff)) != 0){       //读取到数据
                //获取完整的一个数据帧
                for (int i = 0; i < readLength; i++) {
                    if (readBuff[i] == 0x68){
                        if (readBuff[i] == 0x68){
                            int lenH = readBuff[i+9];
                            int lenL = readBuff[i+10];
                            int len = lenH*256 + lenL;
                            if (readBuff[i+13+len] == 0x16){
                                //解析数据帧
                                String address =
                                        String.format("%02x",readBuff[2])+
                                        String.format("%02x",readBuff[3])+
                                        String.format("%02x",readBuff[4])+
                                        String.format("%02x",readBuff[5])+
                                        String.format("%02x",readBuff[6])+
                                        String.format("%02x",readBuff[7]);

                                int year = Bcd2Hex.Bcd2Hex(readBuff[i+14]);
                                int month = Bcd2Hex.Bcd2Hex(readBuff[i+15]);
                                int day = Bcd2Hex.Bcd2Hex(readBuff[i+16]);
                                int hour = Bcd2Hex.Bcd2Hex(readBuff[i+17]);
                                int minue = Bcd2Hex.Bcd2Hex(readBuff[i+18]);
                                System.out.println(""+year+month+day+hour+minue);
                                byte tempH = readBuff[i+19];
                                byte tempL = readBuff[i+20];
                                byte humH = readBuff[i+21];
                                byte humL = readBuff[i+22];
                                byte elect = readBuff[i+23];
                                //时间转Date
                                Calendar cal=Calendar.getInstance();
                                cal.set(Calendar.YEAR, 2000+year);
                                cal.set(Calendar.MONTH, month-1);
                                cal.set(Calendar.DAY_OF_MONTH, day);
                                cal.set(Calendar.HOUR_OF_DAY,hour);
                                cal.set(Calendar.MINUTE,minue);
                                Date date = cal.getTime();
                                System.out.println(date);
                                //温度
                                int temp = Bcd2Hex.Bcd2Hex(tempH)*100 + Bcd2Hex.Bcd2Hex(tempL);
                                //湿度
                                int hum = Bcd2Hex.Bcd2Hex(humH)*100 + Bcd2Hex.Bcd2Hex(humL);
                                //电量
                                int ele = Bcd2Hex.Bcd2Hex(elect);
                                //填充进数据实体
                                Date1c20 date1C20Entity = new Date1c20();
                                date1C20Entity.setAddress(address);
                                date1C20Entity.setTemerature(temp);
                                date1C20Entity.setHumidity(hum);
                                date1C20Entity.setElectricity(ele);
                                date1C20Entity.setTime(date);
                                //传送数据实体给Dao层处理。
                                Date1c20Dao.add(date1C20Entity);
                                //应答
                                System.out.println(date1C20Entity.toString());
                            }
                        }
                    }
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
