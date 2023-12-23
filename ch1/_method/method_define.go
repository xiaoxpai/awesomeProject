package main

//单个返回值

//多个返回值

// Fun2 返回值有名字，可以在内部复制，最后直接return
func Fun2(userName string, password string) (u string, p string) {
	//return userName, password //这种写法不好，不如下面的写法
	u = userName
	p = password
	return
}
