<h1 align="center"> GUSH </h1>
<p align="center"> &lt;gush deploy&gt; </p>
简易的全平台项目自动化部署工具。

## 特性
- 简单易上手
- 使用灵活，支持绝大部分项目部署场景
- 全平台支持
- 无需在服务器端部署


## 安装

## 快速入门

### 在项目目录下创建 gush.yml 文件
```
gush init
```
执行上述命令会生成 gush.yml 模板文件，一个简单的 gush.yml 文件内容如下：
```
servers:
  default:
    host: site.com
    user: root
    port: 22
    password: ""
    ssh_key: ""

tasks:
  default:
    - local_shell: |
       cd /path/to/
       git push origin master
    - remote_shell: |
        cd /www/
        git pull origin master
   
```
上面配置定义了一个名称为`default`的 server(服务器) 和一个名称为`default`的 task(任务)，任务下面定义的是多个连贯的 action(动作)。

### 执行部署
在项目目录下执行:
```
gush deploy
```
上面命令会读取当前目录下的 gush.yml 文件配置，然后通过SSH连接`default` server 并执行`default` task。

也可以指定连接的 server 名称和执行的 task 名称：
```
gush deploy <task> <server>
```

如果 gush.yml 文件不在当前目录，可以指定 gush.yml 文件路径:
```
gush deploy -c /path/to/gush.yml <task> <server>
```

### 部署执行过程
以下面 `gush.yml` 配置内容为例,下面配置简单演示前端 vue.js 项目部署过程:
```
servers:
  default:
    host: site.com
    user: root
    port: 22
    password: ""
    ssh_key: ""

tasks:
  default:
    - local_shell: |
       cd /www/vue-project/
       yarn build
       tar -czf dist.tar.gz dist/
    - upload:
        local: ./dist.tar.gz
        remote: /www/wwwroot/dist.tar.gz
    - remote_shell: |
        cd /www/wwwroot/
        tar -xzf dist.tar.gz
        rm dist.tar.gz
    - local_shell: |
        rm dist.tar.gz
```

执行`gush deploy`过程如下：
- SSH 连接 `default` server
- 执行 `default` task,依次执行 action:
    - 在本地编译 vue 项目，并打包编译后的文件为一个压缩包文件
    - 上传压缩包到服务器
    - 在服务器上解压压缩文件，并清理压缩包文件
    - 清理本地压缩包文件

## gush.yml 文件详解



