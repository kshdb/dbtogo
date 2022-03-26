package cmd

import (
	"database/sql"
	"github.com/kzkzzzz/dbtogo/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type (
	MysqlGen struct {
	}
	MysqlColumnInfo struct {
		Table   string `gorm:"column:TABLE_NAME"`
		Name    string `gorm:"column:COLUMN_NAME"`
		Comment string `gorm:"column:COLUMN_COMMENT"`
		Type    string `gorm:"column:COLUMN_TYPE"`
	}
)

func (m *MysqlGen) GetColumns() []ColumnInfo {
	gm, err := gorm.Open(mysql.Open(cmdParam.Dsn))
	if err != nil {
		common.Log.Fatalf("连接失败: %s", err)
	}
	gm = gm.Debug()

	var currentDb sql.NullString
	err = gm.Raw("select database() as db").Scan(&currentDb).Error
	if err != nil {
		common.Log.Fatal(err)
	}

	if !currentDb.Valid {
		common.Log.Fatal("未指定数据库")
	}

	columns := make([]MysqlColumnInfo, 0)
	gm.Table("information_schema.COLUMNS").
		Select([]string{"TABLE_NAME", "COLUMN_NAME", "COLUMN_TYPE", "COLUMN_COMMENT"}).
		Where("TABLE_SCHEMA = ?", currentDb).
		Where("TABLE_NAME in ?", cmdParam.Table).
		Find(&columns)

	result := make([]ColumnInfo, 0)
	for _, v := range columns {
		result = append(result, ColumnInfo{
			Table:   v.Table,
			Name:    v.Name,
			Comment: v.Comment,
			Type:    v.Type,
			GoName:  common.StrToCamelCase(v.Name),
			GoType:  m.convertTypeToGo(v.Type),
		})
	}
	//fmt.Println(conf.Mysql)
	return result
}

func (m *MysqlGen) convertTypeToGo(srcType string) (dstType string) {
	srcType = strings.ToLower(strings.TrimSpace(srcType))

	switch {
	case strings.HasPrefix(srcType, "bigint"):
		dstType = "int64"
	case strings.Contains(srcType, "int"):
		dstType = "int"
	case strings.HasPrefix(srcType, "decimal") || strings.HasPrefix(srcType, "dobule"):
		dstType = "float64"
	case strings.HasPrefix(srcType, "date") || strings.HasPrefix(srcType, "time") ||
		strings.HasPrefix(srcType, "year") || strings.HasPrefix(srcType, "datetime") ||
		strings.HasPrefix(srcType, "timestamp"):
		dstType = "time.Time"

	default:
		dstType = "string"
	}
	//
	//if strings.Contains(dstType, "int") && strings.Contains(srcType, "unsigned") {
	//	dstType = fmt.Sprintf("u%s", dstType)
	//}

	return dstType
}
