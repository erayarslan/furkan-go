EXECUTABLE := furkan
TARGET := main.go

make:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-linux $(TARGET)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-windows.exe $(TARGET)
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-darwin $(TARGET)