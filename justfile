


test:
	go build -o testdata/renderer/renderer.so -buildmode=plugin ./testdata/renderer
	go run . \
		--input-directory ./testdata/input/service \
		--names CloudServiceServer \
		--renderer-plugin-path ./testdata/renderer/renderer.so \
		--output-path ./testdata/renderer/output/handler.go

	go test ./testdata/renderer/...
