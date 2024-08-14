#!/bin/bash
script_path=$(readlink -f "$0")
script_dir=$(dirname "$script_path")

cd $script_dir
root_path=$(builtin cd $script_dir/..; pwd)
cd $root_path

cd module1

# stack を作る（初回以降はエラーなるので結果は破棄する）
pulumi stack init $STACK_NAME 2>/dev/null

pulumi up -s $STACK_NAME
