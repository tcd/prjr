read -p "Arguments to pass: "  args

cd ./cmd/prjr && go run *.go "$args"
