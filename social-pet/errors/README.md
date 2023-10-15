### 主要用途
对系统errors包的封装，方便对错误信息包裹和分类

### 错误码查询表
```
NotFound         ErrorType = 404 //资源不存在
ParamInvalid     ErrorType = 400 //参数错误
UnAuthorize      ErrorType = 401 //没有用户信息
PermissionDenied ErrorType = 403 //没有权限
SystemError      ErrorType = 500 //系统业务错误
```

### 调用方法
```
调试xxx_test.go文件
```
