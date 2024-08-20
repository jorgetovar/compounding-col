# Go Backend

docker build -t compounding/service .

aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin XXX.dkr.ecr.us-west-2.amazonaws.com

docker tag compounding/service:latest XXX.dkr.ecr.us-west-2.amazonaws.com/compounding/service:latest

aws ecr create-repository --repository-name compounding/service --region us-west-2

docker push XXX.dkr.ecr.us-west-2.amazonaws.com/compounding/service:latest