package com.relengxing.server;

import org.springframework.context.annotation.Bean;
import rx.Observable;
import rx.Subscriber;
import rx.schedulers.Schedulers;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.PrintStream;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * Created by relengxing on 2016/12/26.
 */
public class TcpServer implements Runnable{

    public static boolean isRun = false;

    public void run() {
        isRun = true;
        try {
            ServerSocket serverSocket = new ServerSocket(7002);
            while (true){
                Socket socket = serverSocket.accept();          //阻塞线程，监听到有人连接后继续
                //有新连接时，单独开启一个线程用于读写数据
                new Thread(new ServerThread(socket)).start();
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        isRun = false;
    }
}
