package oui

import (
    "os"
    "path/filepath"
    "testing"
)

func TestLoadAndLookup(t *testing.T) {
    // 初始化测试数据库
    db := NewDatabase()
    
    // 获取测试文件绝对路径
    testCSVPath := filepath.Join("..", "..", "internal", "data", "oui.csv")
    if _, err := os.Stat(testCSVPath); os.IsNotExist(err) {
        t.Fatalf("测试文件不存在: %s", testCSVPath)
    }

    // 测试加载功能
    if err := db.Load(); err != nil {
        t.Fatalf("加载失败: %v", err)
    }

    // 更新测试用例表
    tests := []struct {
        name    string
        mac     string
        wantOrg string
        wantErr bool
    }{
        {
            name:    "标准冒号格式",
            mac:     "08:EA:44:05:07:80", // 改为Extreme Networks的OUI
            wantOrg: "Extreme Networks Headquarters",
            wantErr: false,
        },
        {
            name:    "连字符格式",
            mac:     "08-EA-44-05-07-80", // 同步修改OUI部分
            wantOrg: "Extreme Networks Headquarters",
            wantErr: false,
        },
        {
            name:    "纯数字格式",
            mac:     "08EA4405780", // 前6位对应08EA44
            wantOrg: "Extreme Networks Headquarters",
            wantErr: false,
        },
        {
            name:    "空MAC地址",
            mac:     "",
            wantErr: true,
        },
        {
            name:    "无效长度",
            mac:     "089E8",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := db.Lookup(tt.mac)
            if (err != nil) != tt.wantErr {
                t.Errorf("Lookup() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && got.Organization != tt.wantOrg {
                t.Errorf("Lookup() = %v, want %v", got.Organization, tt.wantOrg)
            }
        })
    }
}

func TestEmptyDatabase(t *testing.T) {
    db := NewDatabase()
    _, err := db.Lookup("08:9E:84:05:07:80")
    if err == nil {
        t.Error("空数据库应返回错误")
    }
}