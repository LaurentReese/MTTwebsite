set -x
./before_compile_backend.sh
go build serverGOLANG.go sendmailGOLANG.go databaseGOLANG.go
./after_compile_backend.sh
set +x