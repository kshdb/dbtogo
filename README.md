# 数据库表 -> go结构体


### 安装 go 1.17+
```shell
go install github.com/kzkzzzz/dbtogo@main
```

```text
Usage:
dbtogo [flags]

Flags:
-d, --dsn string      dsn链接
-h, --help            help for dbtogo
-o, --output string   文件输出目录, 未指定则只打印到终端
-s, --source string   数据源(支持:mysql), 默认mysql (default "mysql")
-t, --table strings   表名, 多个逗号隔开
```

## 在当前目录生成文件

```shell
dbtogo -o . --dsn="remote:remote123@tcp(127.0.0.1:3306)/db" -t user,order
```

### user.go

```go
package main

type User struct {
	Id        int    `json:"id"`         // Id
	Phone     string `json:"phone"`      // 手机
	Username  string `json:"username"`   // 用户名
	Avatar    string `json:"avatar"`     // 头像
	CreatedAt int    `json:"created_at"` // CreatedAt
	UpdatedAt int    `json:"updated_at"` // UpdatedAt
}

func (u *User) TableName() string {
	return "user"
}
```

### order.go

```go
package main

type Order struct {
	Id        int   `json:"id"`         // Id
	OrderId   int64 `json:"order_id"`   // 订单id
	UserId    int   `json:"user_id"`    // 用户id
	CreatedAt int   `json:"created_at"` // 创建时间
}

func (o *Order) TableName() string {
	return "order"
}

```