#!/bin/bash

if [ $# -lt 2 ]; then
  echo "Usage: $0 <domain> <query> <filename>"
  exit 1
fi

domain_name=$1
query_name=$2
file_name=${3:-$query_name}

path=application/query/${domain_name}

# 將第一的字母轉小寫
first_letter=$(echo "${query_name:0:1}" | tr '[:upper:]' '[:lower:]')
private_name="${first_letter}${query_name:1}"

cat << EOF > "${path}/${file_name}.go"
package ${domain_name}

import (
    "context"
    "ddd-sample/application/query"
)

type ${query_name}Query query.Query[${query_name}QueryInput, ${query_name}QueryOutput]

type ${private_name}Query struct {
}

type ${query_name}QueryInput struct {
}

type ${query_name}QueryOutput struct {
}

func New${query_name}Query() ${query_name}Query {
    return &${private_name}Query{}
}

func (q *${private_name}Query) Execute(ctx context.Context, input ${query_name}QueryInput) (output ${query_name}QueryOutput, err error) {
    return
}
EOF

echo "${file_name}.go 文件已在 ${path} 中創建完成！"
