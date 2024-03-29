THIS LEARNING MODULES BASED ON 
[https://dasarpemrogramangolang.novalagung.com]

INIT GO PROJECT WITH DEPENDENCY MANAGER
> go mod init [project-name]

RUN AND COMPILE GO : 
> go run main.go

BUILD AND COMPILE GO INTO BINARIES : 
-- basic build into same name as golang file within the same directory (root)
> go build main.go
-- change output .exe file name for such [main.go] --> [program.exe]
> go build -o program.exe main.go
-- insert into directory folder such as [bin] (if the folder doesn't exist, it will be created automatically)
> go build -o bin/main.exe main.go