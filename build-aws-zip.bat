set GOOS=linux
go build
mkdir target
move .\awslambda-go-viacep-consultacep .\target\
build-lambda-zip --output .\target\viacep-consultacep.zip .\target\awslambda-go-viacep-consultacep
