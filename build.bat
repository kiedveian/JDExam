set FOPS_PATH=github.com/kiedveian/JDExam/fops
go get %FOPS_PATH%
go test %FOPS_PATH%
go build -o fops %FOPS_PATH%
