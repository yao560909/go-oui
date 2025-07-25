package oui

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// OUI 表示OUI数据库中的一条记录
type OUI struct {
	Registry     string `json:"registry"`
	Assignment   string `json:"assignment"`
	Organization string `json:"organization"`
	Address      string `json:"address"`
}

// Database 表示完整的OUI数据库
type Database struct {
	Rows map[string]OUI `json:"rows"`
}

// NewDatabase 创建并初始化Database实例
func NewDatabase() *Database {
	return &Database{
		Rows: make(map[string]OUI),
	}
}

// Load 从CSV文件加载数据到Database（自动跳过标题行）
func (db *Database) Load() error {
    // 修改为项目根目录的相对路径
    file, err := os.Open("internal/data/oui.csv")
    if err != nil {
        return err
    }
    defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return errors.New("empty CSV file")
	}
	// 跳过标题行（从索引1开始）
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) >= 4 {
			ouiEntry := OUI{
				Registry:     record[0],
				Assignment:   record[1],
				Organization: record[2],
				Address:      strings.Join(record[3:], " "),
			}
			db.Rows[ouiEntry.Assignment] = ouiEntry
		}
	}
	return nil
}

// Lookup 根据MAC地址查询企业信息
// 参数格式支持：08:9E:84、08-9E-84、089E84 等常见格式
func (db *Database) Lookup(mac string) (OUI, error) {
	// 处理空MAC地址
	if mac == "" {
		return OUI{}, errors.New("MAC address cannot be empty")
	}
	// 统一转换为大写并去除分隔符
	normalized := strings.ToUpper(strings.NewReplacer(":", "", "-", "", ".", "").Replace(mac))

	// 提取前6字符作为OUI
	if len(normalized) < 6 {
		return OUI{}, errors.New("invalid MAC address format")
	}
	ouiKey := normalized[:6]

	entry, exists := db.Rows[ouiKey]
	if !exists {
		return OUI{}, errors.New("OUI not found")
	}
	return entry, nil
}
