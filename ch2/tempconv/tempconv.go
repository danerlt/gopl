package tempconv

import "fmt"

type Celsius float64    //摄氏度
type Fahrenheit float64 //华氏度

//摄氏度常量
const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //水结冰温度
	BoilingC      Celsius = 100     //水沸腾温度
)

func (c Celsius) String() string    { return fmt.Sprintf("%g摄氏度", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g华氏度", f) }
