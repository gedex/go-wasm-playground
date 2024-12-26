hello: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./hello/main.go
	cp ${GOROOT}/misc/wasm/wasm_exec.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

goroutines: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./goroutines/main.go
	cp ${GOROOT}/misc/wasm/wasm_exec.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

channels: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./channels/main.go
	cp ${GOROOT}/misc/wasm/wasm_exec.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

js: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./js/main.go
	cp ${GOROOT}/misc/wasm/wasm_exec.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

js-timer: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./js-timer/main.go
	cp ${GOROOT}/misc/wasm/wasm_exec.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

func: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./func/main.go
	cp ./func/index.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

todolist: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./todolist/main.go
	cp ./todolist/index.html ./build/index.html
	cp ./todolist/index.js ./build/index.js

request: clean
	GOOS=js GOARCH=wasm go build -o ./build/test.wasm ./request/main.go
	cp ./request/index.html ./build/index.html
	cp ${GOROOT}/misc/wasm/wasm_exec.js ./build/wasm_exec.js

clean:
	rm -rf ./build/*
