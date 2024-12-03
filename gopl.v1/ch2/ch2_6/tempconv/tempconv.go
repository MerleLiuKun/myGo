package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 为 Celsius 类型添加 字符串转换方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// 为 Fahrenheit 类型添加 字符串转换方法
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%gK", k)
}
