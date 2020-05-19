#/bin/bash

dirRoot="/home/game/LittleCai"
runPath="/home/game/bin"

op=$1

function usage() 
{
    echo "webCtl  autoPb|npmBuild|webRun|webBuild"
    echo " autoPb:   auto generate PB and code"
    echo " npmBuild:  run webpack and move to static"
    echo " webBuild: build goServer web"
    echo " webRun: run or restart web"
} 


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
