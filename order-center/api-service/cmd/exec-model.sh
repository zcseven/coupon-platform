#/bin/bash
datasource="root:Zc@123456@tcp(127.0.0.1:3306)/order_center"
table=$1
modelDir=$2
#goctl model mysql datasource -url="$datasource" -table="$table" -c -dir ../internal/model/$modelDir -style=goZero
mkdir -p ../internal/model/$modelDir
goctl model mysql datasource -url="$datasource" -table="$table" -dir ../internal/model/$modelDir -style=goZero