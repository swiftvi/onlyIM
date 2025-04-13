# 生成rpc代码
goctl rpc protoc ./app/user/rpc/user.proto --go_out=./app/user/rpc --go-grpc_out=./app/user/rpc --zrpc_out=./app/user/rpc
# 启动rpc服务
go run ./app/user/rpc/user.go -f ./app/user/rpc/etc/user.yaml
# 生成model代码
goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./app/user/model/" -c
# 生成http API代码
goctl api go -api ./app/user/api/user.api -dir ./app/user/api
# 启动http服务
go run ./app/user/api/user.go -f ./app/user/api/etc/user.yaml