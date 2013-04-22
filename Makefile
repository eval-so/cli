TARGET= evalso

all: $(TARGET)

evalso: src/eval.so/evalso/evalso.go
	GOPATH=${PWD} go build $<

clean:
	rm -f $(TARGET)

distclean: clean
	rm -rf .gostuff
