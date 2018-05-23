package main

import (
	"fmt"
)

var (
	Poweroff        = fsmState("关闭")
	FirstGear       = fsmState("1档")
	SecondGear      = fsmState("2档")
	ThirdGear       = fsmState("3档")
	PowerOffEvent   = fsmEvent("按下关闭按钮")
	FirstGearEvent  = fsmEvent("按下1档按钮")
	SecondGearEvent = fsmEvent("按下2档按钮")
	ThirdGearEvent  = fsmEvent("按下3档按钮")
	PowerOffHandler = fsmHandler(func() fsmState {
		fmt.Println("电风扇已关闭")
		return Poweroff
	})
	FirstGearHandler = fsmHandler(func() fsmState {
		fmt.Println("电风扇开启1档，微风徐来！")
		return FirstGear
	})
	SecondGearHandler = fsmHandler(func() fsmState {
		fmt.Println("电风扇开启2档，凉飕飕！")
		return SecondGear
	})
	ThirdGearHandler = fsmHandler(func() fsmState {
		fmt.Println("电风扇开启3档，发型被吹乱了！")
		return ThirdGear
	})
)

// 电风扇
type ElectricFan struct {
	*fsm
}

// 实例化电风扇
func NewElectricFan(initState fsmState) *ElectricFan {
	return &ElectricFan{
		fsm: newFSM(initState),
	}
}

// 入口函数
func main() {

	efan := NewElectricFan(Poweroff) // 初始状态是关闭的
	// 关闭状态
	efan.addHandler(Poweroff, PowerOffEvent, PowerOffHandler)
	efan.addHandler(Poweroff, FirstGearEvent, FirstGearHandler)
	efan.addHandler(Poweroff, SecondGearEvent, SecondGearHandler)
	efan.addHandler(Poweroff, ThirdGearEvent, ThirdGearHandler)
	// 1档状态
	efan.addHandler(FirstGear, PowerOffEvent, PowerOffHandler)
	efan.addHandler(FirstGear, FirstGearEvent, FirstGearHandler)
	efan.addHandler(FirstGear, SecondGearEvent, SecondGearHandler)
	efan.addHandler(FirstGear, ThirdGearEvent, ThirdGearHandler)
	// 2档状态
	efan.addHandler(SecondGear, PowerOffEvent, PowerOffHandler)
	efan.addHandler(SecondGear, FirstGearEvent, FirstGearHandler)
	efan.addHandler(SecondGear, SecondGearEvent, SecondGearHandler)
	efan.addHandler(SecondGear, ThirdGearEvent, ThirdGearHandler)
	// 3档状态
	efan.addHandler(ThirdGear, PowerOffEvent, PowerOffHandler)
	efan.addHandler(ThirdGear, FirstGearEvent, FirstGearHandler)
	efan.addHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	efan.addHandler(ThirdGear, ThirdGearEvent, ThirdGearHandler)

	// 开始测试状态变化
	efan.call(ThirdGearEvent)  // 按下3档按钮
	efan.call(FirstGearEvent)  // 按下1档按钮
	efan.call(PowerOffEvent)   // 按下关闭按钮
	efan.call(SecondGearEvent) // 按下2档按钮
	efan.call(PowerOffEvent)   // 按下关闭按钮
}
