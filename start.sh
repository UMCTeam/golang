#! /bin/bash
cmd="docker run -d --name goland_$1 -p 8082:8081 $1"
$cmd