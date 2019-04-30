#!/bin/bash

problem=$1
lang=${2:-go}

dir=${lang}/${problem/./\/}

echo $problem $(date +"%Y-%m-%dT%H:%M:%S%z") >> "questions/${lang}"
git add $dir "questions/done"
git commit -m "solution to $problem"
