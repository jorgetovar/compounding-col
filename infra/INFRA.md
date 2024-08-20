# Infrastructure

aws cloudformation create-stack \
--stack-name CompoundingCoreStack \
--capabilities CAPABILITY_NAMED_IAM \
--template-body file://core.yml

aws cloudformation describe-stacks --stack-name CompoundingCoreStack

aws cloudformation describe-stacks --stack-name CompoundingCoreStack > ./cloudformation-core-output.json

cfn-lint core.yml

aws cloudformation update-stack \
--stack-name CompoundingCoreStack \
--capabilities CAPABILITY_NAMED_IAM \
--template-body file://core.yml
