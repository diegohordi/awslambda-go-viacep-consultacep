set GOOS=linux
go build
mkdir target
move .\awslambda-go-viacep-consultacep .\target\viacep-consultacep
build-lambda-zip --output .\target\viacep-consultacep.zip .\target\viacep-consultacep
