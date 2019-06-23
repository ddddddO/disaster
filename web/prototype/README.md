## Go WebAssembly(after go1.11)
1. $HOME/sdk/go1.12/misc/wasm下から、以下ファイルをプロジェクト内へcp(verによってコード異なるので注意)  
   ・wasm_exec.html  
   ・wasm_exec.js  

1. main.goファイル作成  

1. 以下実行し、wasmビルド  
   `GOOS=js GOARCH=wasm go build test.wasm main.go`  

1. 以下実行し、webサーバー起動  
   `goexec 'http.ListenAndServe(":8081", http.FileServer(http.Dir(".
")))'`  
   (`goexec`無い場合、`go get -u github.com/shurcooL/goexec`)  

1. ブラウザから`localhost:8081/wasm_exec.html`へアクセス  
   デベロッパーモードでConsole開き、画面の「Run」ボタン押下で、文字列が出力されればOK  

