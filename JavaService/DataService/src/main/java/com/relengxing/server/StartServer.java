package com.relengxing.server;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.stereotype.Service;

/**
 * Created by relengxing on 2016/12/26.
 */
@Service
public class StartServer implements ApplicationListener<ContextRefreshedEvent> {

    public void onApplicationEvent(ContextRefreshedEvent contextRefreshedEvent) {
        System.out.println("=====服务器启动，开始执行任务=====");
        startMethod();
    }
    /**
     * 启动时调用该方法
     * */
    private void startMethod() {
        if (tcpServer.isRun){
            return;
        }else {
            new Thread(tcpServer,"TcpServer").start();
        }
    }
    private TcpServer tcpServer;
    public void setTcpServer(TcpServer server) {
        this.tcpServer = server;
    }
}
