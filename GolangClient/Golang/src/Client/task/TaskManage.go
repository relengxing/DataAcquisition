package task

import (
    "sync"
    "time"
)


//任务管理器
type TaskManage struct {
    Enqueue  map[string]*BaseTask//任务队列
    mt       sync.Mutex
}

//初始化
func NewTaskManage() *TaskManage {
    taskManage := &TaskManage{}
    taskManage.Enqueue = make(map[string]*BaseTask)
    return taskManage
}
/**
    开始运行任务管理器
    执行任务管理器任务队列中的任务
*/
func (taskManage *TaskManage)Run()  {
    go func() {
        timer1 := time.NewTicker(1 * time.Second)       //tick为1秒
        for {
            select {
            case <-timer1.C:
                for name, task := range taskManage.Enqueue {
                    if task.TaskTime == 0 {                 //task时间到
                        task.TaskMethod(task.TaskParam...)     //运行任务
                        if task.TaskCycle == 0 {            //周期是0，代表只运行一次
                            delete(taskManage.Enqueue,name) //删除任务
                        }else {
                            taskManage.Enqueue[name].TaskTime = task.TaskCycle  //重新定时
                        }
                    }else {
                        taskManage.Enqueue[name].TaskTime--
                    }
                }
            }
        }
    }()
}
/**
    添加任务
*/
func (taskManage *TaskManage)RegisterTask(name string,basetask *BaseTask)  {
    taskManage.mt.Lock()
    defer taskManage.mt.Unlock()
    taskManage.Enqueue[name] = basetask
}

/**
    删除任务
*/
func (taskManage *TaskManage)DeleteTask(name string)  {
    taskManage.mt.Lock()
    defer taskManage.mt.Unlock()
    delete(taskManage.Enqueue,name) //删除任务
}