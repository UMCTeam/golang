#! /bin/bash
cmd="docker run -d --name goland_$1 -p 3501:8081 $1"
$cmd