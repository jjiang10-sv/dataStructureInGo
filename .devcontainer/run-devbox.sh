#!/usr/bin/env bash 

SOURCE="$( readlink "$SOURCE")" 

while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink 
    bin="$( cd -P"$( dirname "$SOURCE" )" && pwd)" 
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /* ]] && S0URCE="$bin/$SOURCE" #if $SOURCE was a relative symlink, we need to resolve 
    #it relative to the path where the symlink file was located
done
bin="$( cd -P "$( dirname "$SOURCE" )" && pwd )" 
cd $bin

# Setting ssh config 
mkdir -p /root/.ssh && cp -r /root/.ssh-localhost/* root/.ssh && chmod 700 /root/.ssh && chmod 600 /root/.ssh/*

