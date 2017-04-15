
export BIN 		:= $(abspath ./bin)
export SRC 		:= $(abspath ./src)
export PKG 		:= $(abspath ./pkg)
export BUILD 	:= $(abspath ./build)
export GOPATH	:= ${GOPATH}:$(abspath .)

TAG 			:= entry_app

.PHONY : all

all : 
	$(MAKE) -C $(SRC)
	$(MAKE) make-image

make-image:
	cp build/Dockerfile $(BIN)
	cd $(BIN) ; docker build --rm -t $(TAG) . 

clean:
	-@rm -fv $(BIN)/*
	-@rm -fvr $(PKG)/*
	-@docker rmi -f $(TAG)
