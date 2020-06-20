
curPwd=`pwd`

staticPath=$HOME/LittleCai/bin

rm -rf  $staticPath/static

if [ ! -d $staticPath/static ];then
   mkdir -p $staticPath/static/js
   mkdir -p $staticPath/static/image
   mkdir -p $staticPath/static/css
fi

if [ ! -d $staticPath/views ];then
   mkdir -p $staticPath/views
fi

npmPath=$HOME/LittleCai/frontend

rm -rf $npmPath/dist/*

cd $npmPath && npm run buildDev

cp $npmPath/dist/images/*  $staticPath/static/images/

#if [ ! `ls ./js/littleCai-*.js* | wc -l` -eq 0 ];then
#    mv ./js/littleCai-*.js* /tmp/
#fi
#
#if [ ! `ls ./js/vendors~littleCai*.js* | wc -l ` -eq 0 ];then
#    mv ./js/vendors~littleCai*.js*  /tmp/
#fi

cp $npmPath/dist/js/*  $staticPath/static/js/

cp $npmPath/dist/index.html  ../views/

cd $curPwd
