clean:
	rm -rf output/

auto:
	sh ./shell/auto.sh

npm: build
	sh ./shell/npmBuild.sh

build:
	sh ./shell/webBuild.sh

test: install
	go build -o output/bin/go_qschou_reward ./app

run:
	sh ./shell/webRun.sh