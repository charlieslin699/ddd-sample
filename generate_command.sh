#!/bin/bash

if [ $# -lt 2 ]; then
  echo "Usage: $0 <domain> <command> <filename>"
  exit 1
fi

domain_name=$1
command_name=$2
file_name=${3:-$command_name}

path=application/command/${domain_name}

# 將第一的字母轉小寫
first_letter=$(echo "${command_name:0:1}" | tr '[:upper:]' '[:lower:]')
private_name="${first_letter}${command_name:1}"

cat << EOF > "${path}/${file_name}.go"
package ${domain_name}

import (
    "context"
    "ddd-sample/application/command"
)

type ${command_name}Command command.Command[${command_name}CommandInput, ${command_name}CommandOutput]

type ${private_name}Command struct {
}

type ${command_name}CommandInput struct {
}

type ${command_name}CommandOutput struct {
}

func New${command_name}Command() ${command_name}Command {
    return &${private_name}Command{}
}

func (c *${private_name}Command) Execute(ctx context.Context, input ${command_name}CommandInput) (output ${command_name}CommandOutput, err error) {
    return
}
EOF

echo "${file_name}.go 文件已在 ${path} 中創建完成！"
