#!/bin/bash

problem=$1

dir=${problem/./\/}

echo $problem $(date +"%Y-%m-%dT%H:%M:%S%z") >> "questions/done"
git add $dir "questions/done"
git commit -m "solution to $problem"
