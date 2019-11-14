# Consulta CEP - Go AWS Lambda Ready

## Description
This project was made for GO and AWS Lambda study purposes.

It consumes ViaCep service to search for a given CEP (Brazilian Zip code) and returns the result or one of the errors listed below.    

ViaCep: https://viacep.com.br

## Dependencies

* AWS Lambda GO package - github.com/aws/aws-lambda-go/lambda

`go get github.com/aws/aws-lambda-go/lambda`

* AWS Lambda Deployment Package - github.com/aws/aws-lambda-go/cmd/build-lambda-zip

`go get github.com/aws/aws-lambda-go/cmd/build-lambda-zip`

## Steps to deploy

1. Install AWS Lambda GO package 
2. Install AWS Lambda Deployment Package 
3. Execute build-aws-zip.bat on PowerShell or CMD
4. Upload the zip file generated in /target folder to your AWS account
5. Just have fun =)

## Errors expected

* ConsultaIndisponivel - If the ViaCep's service is unavailable for any reason
* CepNaoEncontrado - If the given CEP is not found
* ErroInesperado - If the ViaCep's result could not be parsed for any reason

## Tests

`go test`
