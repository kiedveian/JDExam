set FOPS_PATH=github.com/kiedveian/JDExam/fops
set VERSION_STRING=v0.0.1
go get %FOPS_PATH%
go test %FOPS_PATH%
go build -ldflags "-X main.Version=%VERSION_STRING%" -o fops  %FOPS_PATH%
