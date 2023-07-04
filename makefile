api-gen:
	goctl api go --api ./service/user/api/user.api --dir ./service/user

model-gen:
	goctl model mysql datasource \
	--url "root:123456@tcp(127.0.0.1:3306)/book" \
	-table="*" \
	-dir ./service/user/internal/model