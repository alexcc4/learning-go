第一周作业
=======

#### 问题描述
 > 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
 

### 问题解答
- 是否需要 `Wrap` 这个 `error`, 抛给上层？
    * 需要，因为 `DAO` (类似 `golang DDD` 的 `infrastructure` 层)只为上面各层提供通用的技术能力，如数据库交互，数据持久化等基础功能，与具体的业务无关
    * 在 `DAO` 不需要处理遇到的此类错误，只需要加必要的信息 `Wrap` 到上一层
    * 在 `DAO` 上一层，可以根据业务需要来处理这类 `error`,注意处理完 `error`， 这个函数就是一个正常返回的函数
    
- 参考：
   * `Packages that are reusable across many projects only return root error values.`
   * `If the error is not going to be handled, wrap and return up the call stack.`
   * `Once an error is handled, it is not allowed to be passed up the call stack any longer.`
  

- 代码参照 `service.go`

### 代码本地测试
#### 依赖
- `mysql`
  
#### 运行
- 导入 `schema.sql`
- 修改 `dao.go` 的相关数据库配置 `{user}:{password}@tcp(127.0.0.1:3306)/{database}`
- `go run service.go`

