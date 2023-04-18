TARGET=kube-sched
GO=go
BIN_DIR=bin/
GO_MODULE=GO111MODULE=on
FLAGS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: all clean $(TARGET)

all: $(TARGET)

kube-sched:
	$(GO_MODULE) $(FLAGS) $(GO) build -o $(BIN_DIR)$@

clean:
	rm $(BIN_DIR)* 2>/dev/null; exit 0