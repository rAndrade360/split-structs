run:
	echo "Splitting structs..."
	@cd cmd/cli && go run main.go && cd ../../

serve:
	echo "Running server"
	@cd cmd/server && go run main.go && cd ../../


