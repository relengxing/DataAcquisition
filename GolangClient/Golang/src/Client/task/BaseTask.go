package task




//标准任务
type workerFunc func(...interface{}) error

type BaseTask struct {
    TaskMethod      workerFunc          //任务方法
    TaskParam       []interface{}      //任务参数
    TaskTime        int64       //任务延时时间
    TaskCycle       int64       //任务周期
    //TaskType        TaskType            //任务类型
}
//任务类型，并行还是串行，暂时都做并行的
//type TaskType int
//const (
//    PARALLEL = 1 << iota        //并行
//    ORDER                       //顺序执行
//)
/**
    一个工具函数，方便转换
*/
func ChangeParam(arg ...interface{})[]interface{}  {
    return arg
}
