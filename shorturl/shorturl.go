package shorturl

import (
	"goshorturl/baidu"
	"goshorturl/factory"
	"goshorturl/sina"
	"log"
)

func Make(name, url string) string {
	var (
		m factory.Factory
	)
	switch name {
	case "baidu":
		m = baidu.Baidu{}
	case "sina":
		m = sina.Sina{}
	}

	if m == nil {
		log.Fatal("not support name:" + name)
	}

	return factory.Create(m, url)
}
