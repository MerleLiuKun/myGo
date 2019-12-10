package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var shaStr = flag.String("sha", "sha256", "输入哈希值算法：")
	var myStr = flag.String("str", "x", "输入需要哈希的字符串")

	flag.Parse()
	fmt.Printf("Parse: method: %v, string: %v\n", *shaStr, *myStr)
	ParseSha(*shaStr, *myStr)
}

func ParseSha(flag, str string) {
	switch flag {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	}
}

/*
Run:
 go run flagSha.go -sha sha512 -str liukun
Parse: method: sha512, string: liukun
3afb1e6d5177393678787b96171feccf22aec73c73cf61aed015f804cba9dd2bdb16f2ad9cf8738d4a00302d747cf5861dd820989b2562bb60f97710dc6b06af
 go run flagSha.go -str liukun
Parse: method: sha256, string: liukun
a818d8f7fe7ecd48a136900a91cc34b66efc0bbb31ad514342280afda6f8da47
*/
