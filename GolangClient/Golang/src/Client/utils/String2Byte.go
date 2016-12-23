package utils


/**
    String2Byte
    这个是一个专用的，不具备通用性
    使用在地址处理那个部分
    "1b-59-e3-85-23-45"这个字符串先拆分成["1b"......]这个字符串数组
    然后把1b还原成byte数字,只看前两位
*/
func String2Byte(str string) byte {
    dic := make(map[byte]byte)
    dic['0'] = 0
    dic['1'] = 1
    dic['2'] = 2
    dic['3'] = 3
    dic['4'] = 4
    dic['5'] = 5
    dic['6'] = 6
    dic['7'] = 7
    dic['8'] = 8
    dic['9'] = 9
    dic['a'] = 10
    dic['b'] = 11
    dic['c'] = 12
    dic['d'] = 13
    dic['e'] = 14
    dic['f'] = 15
    dic['A'] = 10
    dic['B'] = 11
    dic['C'] = 12
    dic['D'] = 13
    dic['E'] = 14
    dic['F'] = 15
    if len(str) == 0 {
        return 0
    }else if len(str)==1 {
        return dic[str[0]]
    }else {
        return (dic[str[0]]<<4+dic[str[1]])
    }
}