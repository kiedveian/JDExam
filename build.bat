set MAIN_PATH=github.com/kiedveian/JDExam
set FOPS_PATH=%MAIN_PATH%/fops
set VERSION_STRING=v0.0.1
go get %MAIN_PATH%
go test %FOPS_PATH%
go build -v -ldflags "-X github.com/kiedveian/JDExam/fops.Version=%VERSION_STRING%" -o fops.exe  %MAIN_PATH%
