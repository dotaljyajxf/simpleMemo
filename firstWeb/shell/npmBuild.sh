
curPwd=`pwd`

rm -rf  /home/game/LittleCai/dist/* 

cd /home/game/LittleCai

npm run buildDev

cd /home/game/runWeb/static

ls ./images/* |grep -E "[a-z0-9]{32}.jpeg" | xargs rm 

cp /home/game/LittleCai/dist/images/*  ./images/ 

if [ ! `ls ./js/littleCai-*.js* | wc -l` -eq 0 ];then
    mv ./js/littleCai-*.js* /tmp/  
fi

if [ ! `ls ./js/vendors~littleCai*.js* | wc -l ` -eq 0 ];then 
    mv ./js/vendors~littleCai*.js*  /tmp/
fi

cp /home/game/LittleCai/dist/js/*  ./js/

cp /home/game/LittleCai/dist/index.html  ../views/ 

cd $curPwd
