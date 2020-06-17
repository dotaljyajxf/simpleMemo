
curPwd=`pwd`

rm -rf  $HOME/LittleCai/dist/*

cd $HOME/LittleCai

npm run buildDev

cd $HOME/run/static

ls ./images/* |grep -E "[a-z0-9]{32}.jpeg" | xargs rm 

cp $HOME/LittleCai/dist/images/*  ./images/ 

if [ ! `ls ./js/littleCai-*.js* | wc -l` -eq 0 ];then
    mv ./js/littleCai-*.js* /tmp/  
fi

if [ ! `ls ./js/vendors~littleCai*.js* | wc -l ` -eq 0 ];then 
    mv ./js/vendors~littleCai*.js*  /tmp/
fi

cp $HOME/LittleCai/dist/js/*  ./js/

cp $HOME/LittleCai/dist/index.html  ../views/ 

cd $curPwd
