# 适配器模式(Adapter)

由于Golang没有继承，所以，只能实现委托方式。

##程序说明

oldAPI：存放老API程序
adapter：利用委托的方式对老API进行适配。
main.go：直接调用适配器，实现想要的功能。