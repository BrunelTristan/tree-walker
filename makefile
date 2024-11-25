MAKEFLAGS += --no-print-directory
EXEC=tree-walker
SRCDIR=./
TST_FILES=$(wildcard $(SRCDIR)/*/*/*_test.go $(SRCDIR)/*/*_test.go $(SRCDIR)/*_test.go)
SRC_FILES=$(wildcard $(SRCDIR)/*/*/*.go $(SRCDIR)/*/*.go $(SRCDIR)/*.go)
SRC1= $(filter-out $(wildcard $(SRCDIR)/*_test.go), $(SRC_FILES))
SRC2= $(filter-out $(wildcard $(SRCDIR)/*/*_test.go), $(SRC1))
SRC= $(filter-out $(wildcard $(SRCDIR)/*/*/*_test.go), $(SRC2))

.PHONY: all clear 

all: displayCompilation $(EXEC) run-tests code-analysis list-todo 

$(EXEC): $(SRC) makefile
	go mod tidy
	go fmt ./...
	go build -o . ./...

displayCompilation:
	@echo "${_RED}  --COMPILATION ${_END}"

code-analysis:
	go vet ./...
	golangci-lint run -D exportloopref

clear: clean
	rm -rf $(EXEC)
	
run-tests: $(TST_FILES)
	@echo "${_RED}  --LAUNCH TESTS ${_END}"
	go test -cover -v ./...
	go test -bench=. ./...

list-todo:
	@echo "${_RED}  --LOOKING FOR TODO ${_END}"
	@grep -I -ri 'todo ' $(SRCDIR) | grep -v 'makefile:' || true
