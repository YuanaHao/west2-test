package main

import "fmt"

func main() {
	var a [10]int // 定义一个长度为10的数组
	var height int
	var sum int
	//写入数组内的十个元素，即十个苹果高度
	_, err := fmt.Scan(&a[0], &a[1], &a[2], &a[3], &a[4], &a[5], &a[6], &a[7], &a[8], &a[9])
	//写入人的高度
	//此处不选择scanf的原因是其要求string类型
	//不选择两行都用scanln的原因是对于scanln会自动在缓存区读取换行符替换下一行第一个字符
	_, err = fmt.Scan(&height)
	if err == nil { //即上行函数不报错的情况下
		for i := 0; i < len(a); i++ {
			if a[i] <= height { //比较人和苹果的高度
				sum++
			} else {
				m := height + 30 //加上板凳的高度
				if a[i] <= m {
					sum++
				}
			}
		}
		fmt.Println(sum)
	}

}
