MAKEFLAGS += --no-print-directory
EXEC=tree-walker
SRCDIR=./
TST_FILES=$(wildcard $(SRCDIR)/*/*/*_test.go $(SRCDIR)/*/*_test.go $(SRCDIR)/*_test.go)
SRC_FILES=$(wildcard $(SRCDIR)/*/*/*.go $(SRCDIR)/*/*.go $(SRCDIR)/*.go)
SRC1= $(filter-out $(wildcard $(SRCDIR)/*_test.go), $(SRC_FILES))
SRC2= $(filter-out $(wildcard $(SRCDIR)/*/*_test.go), $(SRC1))
SRC= $(filter-out $(wildcard $(SRCDIR)/*/*/*_test.go), $(SRC2))
INTERFACE_FILES=$(wildcard $(SRCDIR)/*/*/i[A-Z]*.go $(SRCDIR)/*/i[A-Z]*.go $(SRCDIR)/i[A-Z]*.go)
TEST_PACKAGE_WITHOUT_COVERAGE=tree-walker/integrationTests
EXCLUDED_PACKAGE_FOR_BENCHMARK=$(shell go list ./... | grep -v generatedMocks | grep -v model/)
EXCLUDED_PACKAGE_FOR_COVERAGE=$(filter-out $(TEST_PACKAGE_WITHOUT_COVERAGE), $(EXCLUDED_PACKAGE_FOR_BENCHMARK))

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

mock-generator: $(INTERFACE_FILES)
	$(foreach interface, $^, mockgen -package=generatedMocks -source=./$(interface) -destination=./internal/generatedMocks/$(shell basename $(shell dirname $(interface)))-$(basename $(notdir $(interface)))Mock.go;)

run-tests: mock-generator $(TST_FILES)
	@echo "${_RED}  --LAUNCH TESTS ${_END}"
	
	go mod tidy
	go test -cover -v $(EXCLUDED_PACKAGE_FOR_COVERAGE)
	go test -v $(TEST_PACKAGE_WITHOUT_COVERAGE)
	go test -bench=. $(EXCLUDED_PACKAGE_FOR_BENCHMARK)

list-todo:
	@echo "${_RED}  --LOOKING FOR TODO ${_END}"
	@grep -I -ri 'todo ' $(SRCDIR) | grep -v 'makefile:' || true
