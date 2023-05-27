package main

import "fmt"

func Sum() int {
	//实现的功能：给出五个学生的成绩，求出成绩的总和，平均数：
	//给出五个学生的成绩：
	score1 := 95
	score2 := 91
	score3 := 39
	score4 := 60
	score5 := 21
	//求和：
	sum := score1 + score2 + score3 + score4 + score5
	//平均数：
	avg := sum / 5
	//输出
	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v", sum, avg)
	return 1
}

func ScoresSum() int {
	//实现的功能：给出五个学生的成绩，求出成绩的总和，平均数：
	//给出五个学生的成绩：--->数组存储：
	//定义一个数组：
	var scores [5]int
	//将成绩存入数组：
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21
	//求和：
	//定义一个变量专门接收成绩的和：
	sum := 0
	for i := 0; i < len(scores); i++ { //i: 0,1,2,3,4
		sum += scores[i]
	}
	//平均数：
	avg := sum / len(scores)
	//输出
	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v", sum, avg)
	return 1
}
func main() {
	Sum()
	//缺点：
	//成绩变量定义个数太多，成绩管理费劲，维护费劲  ---》 将成绩找个地方存起来 ---》 数组
	//---》存储相同类型的数据

	ScoresSum()

	//总结：
	//数组定义格式：
	//var 数组名 [数组大小]数据类型
}
