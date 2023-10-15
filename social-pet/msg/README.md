

### 调用方法

消息体的返回结构

```
code:状态码，0为正常
msg:当code为非0时，传递error信息
data:正常业务返回实体
```
```golang
{
  "code": 0,
  "msg": "",
  "data": {
    "title": "xxx",
    "notice": "xxx"
  }
}
```