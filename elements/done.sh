#!/bin/bash

problem=$1

dir=${problem/./\/}

echo $problem >> "questions/done"
git add $dir "questions/done"
git commit -m "solution to $problem"
