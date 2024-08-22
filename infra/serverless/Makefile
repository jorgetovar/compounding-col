STACK_NAME ?= serverless-go-demo
FUNCTIONS := CompoundingFunction
REGION := us-west-2

# To try different version of Go
GO := go

# Make sure to install aarch64 GCC compilers if you want to compile with GCC.
CC := aarch64-linux-gnu-gcc
GCCGO := aarch64-linux-gnu-gccgo-10

build:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

build-%:
		cd functions/$* && GOOS=linux GOARCH=arm64 CGO_ENABLED=0 ${GO} build -o bootstrap

clean:
	@rm $(foreach function,${FUNCTIONS}, functions/${function}/bootstrap)

deploy:
	if [ -f samconfig.toml ]; \
		then sam deploy --stack-name ${STACK_NAME}; \
		else sam deploy -g --stack-name ${STACK_NAME}; \
  fi

load-testing:
	API_URL=$$(aws cloudformation describe-stacks --stack-name $(STACK_NAME) \
	  --region $(REGION) \
		--query 'Stacks[0].Outputs[?OutputKey==`ApiUrl`].OutputValue' \
		--output text) artillery run load-testing/load-test.yml

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy) &&) true
