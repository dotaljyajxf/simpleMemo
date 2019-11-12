#/bin/bash


function op_build()
{
    go build -o web main.go
    echo "build ok!"
}

function op_show()
{
     ps -ef | grep web | grep -v grep
}

function op_start()
{
    nohup ./web $args  > ./log/nohup.log 2>&1 &
    op_show
}


function op_stop()
{
    pid=`ps -ef | grep web | grep -v grep | awk '{print $2}'`
    if [ "$pid" != "" ];then
        kill -9 $pid
    fi
}

function op_restart()
{
    op_stop
    sleep 5s
    op_start
}

function usage()
{
    echo "sh $0 start/stop/restart/build/show"
}

op=$1
if [ ! -f ./conf/app.ini ];then
   echo "./conf/app.ini not exist"
   exit
fi

if [ ! -d ./static ];then
   echo "./static not exist"
   exit
fi

if [ ! -d ./views ];then
   echo "./views not exist"
   exit
fi

args="-c ./conf/app.ini -s ./static -v ./views"

case $op in
    start)
    op_start
    ;;
    stop)
    op_stop
    ;;
    restart)
    op_restart
    ;;
    build)
    op_build
    ;;
    show)
    op_show
    ;;
    *)
    usage
    ;;
esac
