package main

import (
	"fmt"
	"log"

	"github.com/yao560909/go-oui/pkg/oui"
)

func main() {
	// 初始化OUI数据库
	db := oui.NewDatabase()
	
	// 加载内置数据
	if err := db.Load(); err != nil {
		log.Fatalf("数据加载失败: %v", err)
	}

	// 测试不同格式的MAC地址
	macAddresses := []string{
		"08:EA:44:00:00:00",    // 冒号分隔
		"08-EA-44-11-22-33",   // 连字符分隔
		"08EA44AABBCC",         // 无分隔符
		"invalid-mac",          // 无效格式
	}

	for _, mac := range macAddresses {
		result, err := db.Lookup(mac)
		if err != nil {
			fmt.Printf("查询 %-15s 失败: %v\n", mac, err)
			continue
		}
		
		fmt.Printf("MAC: %-15s → 企业: %s\n", mac, result.Organization)
		fmt.Printf("     Registry: %s\n", result.Registry)
		fmt.Printf("     Address: %s\n\n", result.Address)
	}
}