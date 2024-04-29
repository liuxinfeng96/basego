# Basego

## 环境准备

### 配置文件修改

#### 数据库配置

配置目录：conf/config.yaml

配置内容：

```yaml
db_config:
  user: root
  password: 123456
  ip: 127.0.0.1
  port: 33096
  dbname: basego
```

- user：数据库用户名
- password：数据库密码
- ip：数据库IP地址
- port：数据库端口
- dbname：数据库名称


## 部署启动

在项目主目录下，运行部署脚本：

```shell
$ ./deploy.sh [version]
```

-- version: 镜像版本

