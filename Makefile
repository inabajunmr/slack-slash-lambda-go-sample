.PHONY: clean build deploy

clean: 
	rm -rf ./app/app
	
build:
	GOOS=linux GOARCH=amd64 go build -o app/app ./app

deploy: build
	sam package --template-file template.yaml --s3-bucket ${BUCKET} --output-template-file packaged.yaml
	aws cloudformation deploy --template-file packaged.yaml --stack-name ${STACK_NAME}  --capabilities CAPABILITY_IAM --profile ${PROFILE}
