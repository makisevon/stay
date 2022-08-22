.PHONY: build run test clean

output := stay

build:
	protoc --go_out=. --go-grpc_out=. pb/*.proto && \
	go build -o $(output)

http  := :8080
addr0 := :8081
addr1 := :8082
addr2 := :8083

peers := $(addr0),$(addr1),$(addr2)

run: $(output)
	./$(output) -addr=$(addr0) -http=$(http) -peers=$(peers) lru rds & \
	./$(output) -addr=$(addr1) -peers=$(peers) exp rds & \
	./$(output) -addr=$(addr2) -peers=$(peers) exp lru rds

key := hello
val := world # d29ybGQ=

test:
	@ \
 	echo curl localhost$(http)/set/$(key) -w '\n' -X POST -d val=$(val) && \
	curl localhost$(http)/set/$(key) -w '\n' -X POST -d val=$(val) && \
	echo && \
	echo curl localhost$(http)/get/$(key) -w '\n' && \
	curl localhost$(http)/get/$(key) -w '\n' && \
	echo && \
	echo curl localhost$(http)/del/$(key) -w '\n' -X POST && \
	curl localhost$(http)/del/$(key) -w '\n' -X POST && \
	echo && \
	echo curl localhost$(http)/get/$(key) -w '\n' && \
	curl localhost$(http)/get/$(key) -w '\n'

clean:
	rm pb/*.pb.go
	go clean
