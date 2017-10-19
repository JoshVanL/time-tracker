FILENAME=time-tracker

help:
	# build - go build targets

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o $(FILENAME)

