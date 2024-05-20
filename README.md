### 项目介绍

>  [1Panel](https://1panel.cn) 多主机管理


### 使用说明

```bash
# 克隆项目
git clone https://github.com/baiiiii/many-panel.git
# 构建镜像
docker build -t many-panel:1.0 .
# 新建容器
docker run -d -p 9999:9999 --name many-panel many-panel:1.0
```
```
用户名：admin
密码：admin123
```