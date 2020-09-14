# File Ops

* how to build 
  * run commands:
  <pre> go get github.com/kiedveian/JDExam/fops 
   go build -o fops github.com/kiedveian/JDExam/fops </pre>

* Usage:
  *  fops [flags]
  *  fops [command]

* Available Commands
  * linecount Print line count of file
  * checksum  Print checksum of file
  * version     Show the version info
  * help         Help about commands
* Flags
  * -h, --help   help for fops

* known issues
  * error message inconsistent with example

* TODO 
  * command "linecount" return a error if input is a binary file
  * unit test
  * accurate test cases of error handle
