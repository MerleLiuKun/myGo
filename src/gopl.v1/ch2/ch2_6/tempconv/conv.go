package tempconv

// 摄氏温度 转化为 华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 摄氏温度 转化为 开尔文温度
func CtoK(c Celsius) Kelvins {
	return Kelvins(c - AbsoluteZeroC)
}

// 华氏温度 转化为 摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 华氏温度 转化为 开尔文温度
func FToK(f Fahrenheit) Kelvins {
	return Kelvins((f-32)*5/9 - Fahrenheit(AbsoluteZeroC))
}

// 开尔文温度 转化为 摄氏温度
func KToC(k Kelvins) Celsius {
	return Celsius(k + Kelvins(AbsoluteZeroC))
}

// 开尔文温度 转化为 华氏温度
func KToF(k Kelvins) Fahrenheit {
	return Fahrenheit((k+Kelvins(AbsoluteZeroC))*9/5 + 32)
}
