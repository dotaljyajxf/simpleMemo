#!/bin/bash

dirRoot="$HOME/LittleCai"
runPath="$HOME/LittleCai/firstWeb/shell"

op=$1

function usage() 
{
    echo "webCtl  autoPb|npmBuild|webRun|webBuild"
    echo " autoPb:   auto generate PB and code"
    echo " npmBuild:  run webpack and move to static"
    echo " webBuild: build goServer web"
    echo " webRun: run or restart web"
} 

function makeRunDir()
{
  if [ ! -d $HOME/run ];then
    mkdir $HOME/run
    mkdir $HOME/run/views
    mkdir -p $HOME/run/static
    mkdir -p $HOME/run/js
    mkdir -p $HOME/run/images
    cp $HOME/LittleCai/firstWeb/bin/run.sh $HOME/run
  fi
}

makeRunDir
case $op in
    autoPb)
    sh $runPath/auto.sh 
    ;;
    npmBuild)
    sh $runPath/npmBuild.sh
    ;;
    webRun)
    sh $runPath/webRun.sh 
    ;;
    webBuild)
    sh $runPath/webBuild.sh 
    ;;
    *)
    usage
    ;;
esac
