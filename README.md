# simple-admin-example-api-single
Simple Admin 单体API 服务例子。Simple Admin single API service example.

# 生成命令 | Command

```shell
goctls api new example --i18n=true --casbin=true --go_zero_version=v1.5.0 --tool_version=v0.2.8 --trans_err=true --module_name=github.com/suyuan32/simple-admin-example-api --port=8081 --gitlab=true --ent=true

cd example
# 修改 ent/schema/example 为 student | Modify ent/schema/example to student

# 整理下依赖 | Optimize Dependencies
go mod tidy

# 生成 Ent 代码 | Generate Ent Code
make gen-ent

# 生成逻辑代码
make gen-api-ent-logic model=Student group=student
```
