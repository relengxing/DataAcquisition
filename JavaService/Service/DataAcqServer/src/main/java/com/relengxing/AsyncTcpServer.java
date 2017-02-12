package com.relengxing;

import com.relengxing.entity.Date1C20;
import com.relengxing.mapper.Date1c20Mapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Date;

/**
 * Created by relengxing on 2017/2/3.
 */
@Service
public class AsyncTcpServer {
    public static boolean isRun = false;

    @Autowired
    Date1c20Mapper date1C20Mapper;

    @Autowired
    AsyncDataServer asyncDataServer;
    @Async
    public void executeAsyncTask(){
        isRun = true;
        try {
            ServerSocket serverSocket = new ServerSocket(7002);
            while (true){
                Socket socket = serverSocket.accept();          //阻塞线程，监听到有人连接后继续
                //有新连接时，单独开启一个线程用于读写数据
                asyncDataServer.executeAsyncTask(socket);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        isRun = false;
    }


}
