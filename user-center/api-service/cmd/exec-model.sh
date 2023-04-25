#/bin/bash
datasource="root:Zc@123456@tcp(127.0.0.1:3306)/coupon_user"
table=$1
modelDir=$2
#goctl model mysql datasource -url="$datasource" -table="$table" -c -dir ../internal/model/$modelDir -style=goZero

goctl model mysql datasource -url="$datasource" -table="$table" -dir ../internal/model/$modelDir -style=goZero