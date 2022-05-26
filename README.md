<<<<<<< HEAD
# paopao-ce
PaoPao Community Edition
=======
# PaoPao CE

官网 https://www.paopao.info 

## 1. 代码说明

- paopao-api 泡泡后端
- paopao-web 泡泡前端

## 2. 安装说明

- 需要提前安装Go(>=1.17)、Node.js(>=14)、MySQL(>=5.7)、Redis环境

- 需要安装 Zinc （轻量级全文搜索引擎），可以查阅 https://zincsearch.com/ 安装

### 2.1 paopao-api 后端服务

- 导入 `paopao-api` 目录下的`paopao.sql`文件至数据库

- 拷贝 config.yaml.sample 文件至 config.yaml，然后按照注释完成配置修改

```
cd 项目根目录/paopao-api
go mod download
go build .
```
编译成功后，将产出 `paopao-api` 二进制文件，运行：

```
chmod +x paopao-api
./paopao-api
```

即可正常启动


### 2.2 paopao-web 前端服务

编译前需要修改 .env里面的后端地址

```
cd 项目根目录/paopao-web
vim .env
yarn
yarn build
```
build完成后，可以在dist目录获取编译产出，配置nginx指向至该目录即可


## 3. 其他

后端服务建议使用 supervisor 守护进程，然后通过nginx反向代理

短信通道使用的juhe数据，如果申请不下来，可以考虑替换其他服务商的

代码结构比较简单，喜欢的朋友欢迎贡献PR。


>>>>>>> 76a0b2b (update yarn build command.)
