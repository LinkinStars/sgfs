#!/usr/bin/env bash

appId=`ps -ef |grep sgfs |head -n1 |awk '{print $2}'`

if [ -z $appId ];then
    echo "Maybe sgfs not running, please check it..."
else
    echo "The sgfs is stopping..."
    kill $appId
fi