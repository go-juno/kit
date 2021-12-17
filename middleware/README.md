# 中间件使用

## Auth

## 使用 Auth

```go
r.GET("/ping", middleware.Auth(),
  func(c *gin.Context) {
   // 获取绑定上下文用户域控
   auth := middleware.GetBindAuth(c)
   if middleware.GetBindDataAccess(c) {
    c.JSON(200, gin.H{
     "message": auth,
    })
    return
   }
   c.JSON(403, nil)
  })
```

## Role

### 创建 Role 实例

```go
 config := &middleware.RoleConfig{
  Host:    "test-auth.yupaopao.com", // 权限系统地址
  Port:    8088, // 端口
  SiteKey: "53c9014f9afc4b378ed3bda86876d67a", // 权限系统分配的站点地址
 }

 roleMiddleware, err := middleware.NewRole(*config)

```

### 使用 Role

```go
 r.GET("/ping", roleMiddleware.Handler(), func(c *gin.Context) {
  // 获取中间件绑定的角色上下文
  roles := middleware.MustGetBindRole(c)
  // 判断为某一角色
  if roles.Has("880e0a7c37fc41229e6505b535dd757f") {
   c.JSON(200, gin.H{
    "message": "pong 角色 A ",
   })
   return
  }
  // 判断为多个角色之一
  if roles.OneOf("880e0a7c37fc41229e6505b535dd757f",
   "b6de126a99504a78be9bc90480d84b52",
   "cd26497a22b4400b90ab8533ede63537") {
   c.JSON(200, gin.H{
    "message": "角色 A|B|C",
   })
   return
  }
  // 判断同时为多个角色
  if roles.Both("880e0a7c37fc41229e6505b535dd757f",
   "b6de126a99504a78be9bc90480d84b52",
   "cd26497a22b4400b90ab8533ede63537") {
   c.JSON(200, gin.H{
    "message": "角色 A&B&C",
   })
   return
  }
  c.JSON(403, nil)
 })
```

## Privilege

### 创建 Privilege 实例

```go
 config := &middleware.PrivilegeConfig{
  Host:    "test-auth.yupaopao.com",
  Port:    8088,
  SiteKey: "53c9014f9afc4b378ed3bda86876d67a",
 }

 privilegeMiddleware, err := middleware.NewPrivilege(*config)
```

### 使用 Privilege

```go
r.GET("/ping", privilegeMiddleware.Handler("category", "", ""), func(c *gin.Context) {
  c.JSON(200, gin.H{
   "message": "pong",
  })
 })
```

## DataAccess

### 创建 DataAccess 实例

```go
config := &middleware.DataAccessConfig{
  Host:    "test-auth.yupaopao.com",
  Port:    8088,
  SiteKey: "53c9014f9afc4b378ed3bda86876d67a",
 }

 dataAccessMiddleware, err := middleware.NewDataAccess(*config)
```

### 使用 DataAccess

```go
 r.GET("/ping", dataAccessMiddleware.Handler(
  "880e0a7c37fc41229e6505b535dd757f",
  "880e0a7c37fc41229e6505b535dd757f"),
  func(c *gin.Context) {
   // 获取用户数据权限
   if middleware.GetBindDataAccess(c) {
    c.JSON(200, gin.H{
     "message": "pong",
    })
    return
   }
   c.JSON(403, nil)
  })
```
