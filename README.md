# simple-admin-example-api-single v1.0.11 gen by tools v1.5.10
Simple Admin 单体API 服务例子。Simple Admin single API service example.

# 生成命令 | Command

```shell
goctls api new example -i -c -a -m github.com/suyuan32/simple-admin-example-api -p 8081 -g -e

cd example
# 修改 ent/schema/example 为 student | Modify ent/schema/example to student

# 整理下依赖 | Optimize Dependencies
go mod tidy

# 添加 student schema, 直接复制本项目的到本地， 位置在 ./ent/schema | Add student schema, directly copy this project to the local, the location is ./ent/schema

# 生成 Ent 代码 | Generate Ent Code
make gen-ent

# 生成逻辑代码
make gen-api-ent-logic model=Student group=student

# 生成必要文件
make gen-api
```
