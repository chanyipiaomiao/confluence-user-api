## confluence-user-api

提供confluence用户的增加、删除、禁用功能

## 配置文件

```shell script
{
    "listen": ":11111", // 监听端口
    "api": {
        "prefix": "https://xxx.xxx.xxx",  // confluence访问地址
        "login_index": "/login.action",  // 登录界面
        "do_login": "/dologin.action",   // 登录表单提交接口
        "second_authenticate": "/doauthenticate.action", // 二次验证接口
        "disable_confirm": "/admin/users/deactivateuser-confirm.action", // 禁用确认接口
        "create": "/admin/users/docreateuser.action",   // 创建用户接口
        "delete": "/admin/users/removeuser-confirm.action"  // 删除用户接口
    },
    "admin_user": {
        "username": "xxxx",  // 管理员用户名
        "password": "xxxx"   // 管理员用户密码
    }
}
```

## 接口

创建用户

```golang
POST http://127.0.0.1:11111/createUser
Content-Type: application/json

{"username": "test2", "full_name": "test2", "email": "test2@xxx.com", "password": "xxxx"}
```

删除用户

```golang
POST http://127.0.0.1:11111/deleteUser?username=test2
```

禁用用户

```golang
POST http://127.0.0.1:11111/disableUser?username=test2
```

## 测试通过版本

 Confluence 6.7.1