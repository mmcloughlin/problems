#!/bin/bash

problem=$1
lang=${2:-go}

dir=${lang}/${problem/./\/}
done="questions/${lang}"

echo ${problem} $(date +"%Y-%m-%dT%H:%M:%S%z") >> ${done}
git add ${dir} ${done}
git commit -m "elements/${lang}: solution to ${problem}"
