

cd /home/game/runWeb

rm ./web
cp /home/game/LittleCai/src/firstWeb/bin/web .

sh run.sh restart 

cd - 

ps -ef |grep -v grep |grep web 
