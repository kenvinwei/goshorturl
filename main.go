package main

import (
	"fmt"
	"goshorturl/shorturl"
)

//Usage
func main() {
	fmt.Println(shorturl.Make("baidu", "https://996.icu/#/zh_CN"))
	fmt.Println(shorturl.Make("sina", "https://996.icu/#/zh_CN"))
}
