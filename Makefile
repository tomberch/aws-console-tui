NAME 		:= aws-console-tui
OUTPUT_DIR 	:= bin
OUTPUT_PATH := ${OUTPUT_DIR}/${NAME}
PACKAGE    	:= github.com/tomberch/$(NAME)
GIT_REV    	:= $(shell git rev-parse --short HEAD)
VERSION    	:= v0.0.1

build: 
	go build \
		-ldflags "-X ${PACKAGE}/cmd.version=${VERSION} -X ${PACKAGE}/cmd.commitHash=${GIT_REV}" \
		-o ${OUTPUT_PATH} main.go

clean:
	go clean
	rm ${OUTPUT_DIR} -rf

