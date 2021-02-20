# Build Linux

env GOOS=linux GOARCH=amd64 go build -o ./out/ ./cmd/sourcefiles

# Build Windows

env GOOS=windows GOARCH=amd64 go build -o ./out/ ./cmd/sourcefiles