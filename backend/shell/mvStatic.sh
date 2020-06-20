
curPwd=`pwd`

distPath=$HOME/LittleCai/frontend/dist
staticPath=$HOME/LittleCai/bin

if [ ! -d $distPath ];then
   echo "no dist dir"
   exit 1
fi

if [ ! -d $staticPath/static ];then
   mkdir -p $staticPath/static/js
fi

rm -rf  $staticPath/static/*
cp -r $distPath/* $staticPath/static/

#if [ ! -d $staticPath/static ];then
#   mkdir -p $staticPath/static/js
#fi
#
#if [ ! -d $staticPath/views ];then
#   mkdir -p $staticPath/views
#fi


#rm -rf $npmPath/dist/*
#cd $npmPath && npm run buildDev

#cp $distPath/images/*  $staticPath/static/images/

#if [ ! `ls ./js/littleCai-*.js* | wc -l` -eq 0 ];then
#    mv ./js/littleCai-*.js* /tmp/
#fi
#
#if [ ! `ls ./js/vendors~littleCai*.js* | wc -l ` -eq 0 ];then
#    mv ./js/vendors~littleCai*.js*  /tmp/
#fi

#cp $distPath/js/*  $staticPath/static/js/
#
#cp $distPath/index.html  $staticPath/views/

cd $curPwd
