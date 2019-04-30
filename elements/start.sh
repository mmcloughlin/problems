#!/bin/bash

problem=$1
lang=${2:-go}

dir=${lang}/${problem/./\/}

mkdir -p $dir

touch $dir/README.md
