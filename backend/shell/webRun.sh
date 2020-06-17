

cd $HOME/run

rm ./web
cp $HOME/LittleCai/firstWeb/bin/web .

sh run.sh restart 

cd - 

ps -ef |grep -v grep |grep web 
