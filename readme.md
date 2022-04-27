# golang project admin

## init project

### 1 安装 adm

``` shell
go install github.com/GoAdminGroup/adm
```

### 2 修改 config.json

### 3 配置gorm

``` go
 // 本代码中配置多数据源 sqlite数据库为Admin框架所需,mysql为业务数据库
 orm, err = gorm.Open(sqlite.Open(s.GetConfig("default").File), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Info),
 })
 orm.Use(dbresolver.Register(
  dbresolver.Config{
   Replicas: []gorm.Dialector{mysql.Open(m.GetConfig("main").GetDSN())},
  }, "blind"),
 )
```

### 4 运行

```shell
go run main.go
```
