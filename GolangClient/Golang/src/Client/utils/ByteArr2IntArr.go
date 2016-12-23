package utils

func ByteArr2IntArr(b []byte) []int {
    list :=make([]int,0,1024)
    for _,x:=range b{
        list=append(list,int(x))
    }
    return list[:]
}
