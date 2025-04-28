# Redka 客户端

一个 Redka 客户端，使用体验更像在使用 Redis 客户端。

## 特征

- [x] [Strings](https://github.com/nalgeon/redka/blob/main/docs/commands/strings.md) 命令
- [x] [Lists](https://github.com/nalgeon/redka/blob/main/docs/commands/lists.md) 命令
- [x] [Sets](https://github.com/nalgeon/redka/blob/main/docs/commands/sets.md) 命令
- [x] [Hashes](https://github.com/nalgeon/redka/blob/main/docs/commands/hashes.md) 命令
- [x] [Sorted sets](https://github.com/nalgeon/redka/blob/main/docs/commands/sorted-sets.md) 命令
- [x] [Key](https://github.com/nalgeon/redka/blob/main/docs/commands/keys.md) 命令
- [x] [Server/Connection](https://github.com/nalgeon/redka/blob/main/docs/commands/server.md) 命令
- [x] [Transations](https://github.com/nalgeon/redka/blob/main/docs/commands/transactions.md) 命令

## 快速开始

### 先决条件

Redka Client 要求 Go 版本 1.22 或以上。

### 获取 Redka Client

```go
import "github.com/hugiot/redkacli"
```

或者

```bash
go get -u github.com/hugiot/redkacli
```

### 运行

```go
package main

import (
	"fmt"
	"github.com/hugiot/redkacli"
	"log"
)

func main() {
	options := redkacli.NewOptions()
	options.SetPath("./test.db")
	client, err := redkacli.New(options)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// string commands
	_ = client.Set("name", "Tom")
	fmt.Println(client.Get("name"))

	// hash commands
	_, _ = client.HSet("object", map[string]interface{}{
		"name": "Tom",
		"age":  10,
		"city": "New York",
	})
	fmt.Println(client.HGetAll("object"))
	
	// other commands
	// ……
}
```

## 感谢

- [Redka](https://github.com/nalgeon/redka)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)

## 许可证

2025 © [hugiot](https://github.com/hugiot). This project is [MIT](https://github.com/hugiot/redkacli/blob/main/LICENSE) licensed.

