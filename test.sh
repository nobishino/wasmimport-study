
GOOS=wasip1 GOARCH=wasm go build -o main
WASMTIME_LOG=wasmtime_wasi=trace wasmtime main