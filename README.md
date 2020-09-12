<h1 align="center"> GUSH </h1>
<p align="center"> &lt;gush deploy&gt; </p>

适用于个人项目的简易的全平台项目部署工具。

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

`gush deploy` 命令等价于执行 `gush deploy default default`.


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

gush.yml 文件用于描述服务器信息(servers)和部署任务(tasks)，结构如下：
```
servers:
  server1:
    ....
  server2:
  	....

tasks:
  task1:
    ...
  task2:
  	...

```
`servers`配置项下用于配置多套 server(服务器) 信息，`tasks`配置项下用于配置多个 task(任务)，在执行`gush deploy`部署命令时可以指定 server 和 task,如:
```
gush deploy task1 server2
```
上述命令将在 server2 服务器上执行 task1 部署任务。

### server 配置
server 配置项：
- **host**：服务器地址
- **user**: 服务器SSH登录用户名
- **port**: 服务器SSH登录端口号
- **password**：可选，服务器SSH登录密码，如果不配置则在执行部署命令时会要求从命令行输入密码
- **ssh_key**：可选，用于免密登录服务器的SSH私钥

### task 配置

task 是由多个 action(动作) 组成的。

以下是全部可用的 action ：
- **local_shell**：字符串，在本地执行 shell 指令
- **remote_shell**：字符串，在远程服务器上执行 shell 指令
- **upload**：对象，上传本地文件到远程服务器
	- **local**：字符串，本地文件路径
	- **remote**：字符串，远程文件路径
- **download**：对象，下载远程服务器上的文件到本地
	- **remote**：字符串，远程文件路径
    - **local**：字符串，本地文件路径
- **run**：数组，执行其它多个 task

#### run action详解
为了可以更加细致的控制部署任务，可以将任务进行细化拆分，并可以组合多个任务去执行，如：
```
...

tasks:
    task_push:
      - local_shell: |
         cd /path/to/
         git push origin master
    task_pull:
        - remote_shell: |
           cd /www/
           git pull origin master
    default:
       -run:
        - task_push
        - task_pull
```
- 只执行 `task_pull` 部署任务:
```
gush deploy task_pull
```
- 执行 `default` 部署任务，它会先执行任务`task_push`，然后执行`task_pull`:
```
gush deploy default
```

## 相似项目
- [Envoy](https://laravel.com/docs/7.x/envoy)

## License
Apache
