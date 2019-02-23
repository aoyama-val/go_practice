main: main.c liballocate.so
	gcc -W -Wall -L. -lallocate main.c -o main

liballocate.so: liballocate.go
	go build -buildmode=c-shared -o liballocate.so liballocate.go
