package main

func init() {
	// 初始化  包被引入的时执行，只执行一次
	// 因为我们不能确定 init 方法的执行顺序(目前是按照所在文件名排序)
	// 只能曲线救国
	initBeforeSomething()
	initSomething()
	initAfterSomething()
}

func initBeforeSomething() {

}

func initSomething() {

}

func initAfterSomething() {

}
