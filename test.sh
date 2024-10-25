
GOOS=wasip1 GOARCH=wasm go build
WASMTIME_LOG=wasmtime_wasi=trace wasmtime wasmimport-study 