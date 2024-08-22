# Build Go Serverless REST APIs and Deploy to AWS using the SAM framework (Amazon Linux 2 Runtime)

## Why Another Go Tutorial

AWS has been deprecating several services and runtimes recently. As we’ve seen with the discontinuation of our beloved CodeCommit and other crucial services, Go1.x is no longer supported for AWS Lambda functions.

If you try to deploy most of the outdated tutorials, you might encounter errors like this:

```
Resource creation Initiated    
CREATE_FAILED                    AWS::Lambda::Function            DemoFunction                     
                                   Resource handler returned message: 
                                   "The runtime parameter of go1.x is no longer supported for 
                                   creating or updating AWS Lambda functions. We recommend you 
                                   use a supported runtime while creating or updating functions. 
                                   (Service: Lambda, Status Code: 400, Request ID:  
                                   81f1f708-0a7a-40d0-8442-b9c16510d01f)" 
ROLLBACK_IN_PROGRESS             AWS::CloudFormation::Stack       lambda-go-gorilla                
                                   The following resource(s) failed to create: 
                                   [DemoFunction]. Rollback requested by user.
```

**The key takeaway is that the only constant in software is change. However, there are some *timeless principles* that we should always keep in mind**:

To address this issue, I decided to create an up-to-date repository with all the infrastructure needed to deploy a Go application. There are two options available:

1. Deploying with Fargate using Docker containers.
2. Deploying using the SAM framework on AWS.

> You can GitHub find the repository [here](https://github.com/jorgetovar/compounding-col).

### Timeless principles in Software Development

- Infrastructure as Code is essential.
- Good naming conventions in software are crucial.
- Always test your logic.
- Availability & scalability
- Deployment Pipeline as a mechanism to automate the software delivery process. 
- Observability Is Mandatory.
- Security is a first-class citizen in cloud-native applications.
- Go is an excellent option for building APIs.

## Infrastructure as Code is Essential

Immutable infrastructure allows us to declare what we want at a higher level and ensures that development and production environments remain as close as possible. For example:

```yaml
CompoundingFunction:
  Type: AWS::Serverless::Function
  Metadata:
    BuildMethod: makefile
  Properties:
    FunctionName: CompoundingFunction
    Architectures: ["arm64"]
    Handler: bootstrap
    Runtime: provided.al2
    CodeUri: ./functions/CompoundingFunction/
    MemorySize: 512
    Timeout: 10
    Environment:
      Variables:
        COMPOUNDING_TABLE_NAME: !Ref CompoundingTable
    Policies:
      - DynamoDBCrudPolicy:
          TableName: !Ref CompoundingTable
    Events:
      ApiGatewayPost:
        Type: Api
        Properties:
          RestApiId: !Ref ApiGateway
          Path: /compounding
          Method: POST
```

## Good Naming Conventions in Software Are Key

Don’t be afraid to refactor if you have a good suite of tests. Refactoring is an essential activity in software development. Names are important as they appear everywhere in modules, functions, packages, variables, etc.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is the structure for the response JSON
type Response struct {
	Message      string    `json:"message"`
	GainsPerYear []float64 `json:"gainsPerYear"`
}

type Request struct {
	Principal  float64 `json:"principal"`
	AnnualRate float64 `json:"annualRate"`
	Years      int     `json:"years"`
}

func HelloHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return createResponse(400, "Invalid request body")
	}
	fmt.Println("Request", req)
	gainsPerYear := CalculateCompoundInterest(req.Principal, req.AnnualRate, req.Years)
	fmt.Println(gainsPerYear)
	response := Response{
		Message:      "Calculation successful",
		GainsPerYear: gainsPerYear,
	}

	body, err := json.Marshal(response)
	if err != nil {
		return createResponse(500, "Error marshalling response")
	}

	return createResponse(200, string(body))
}

func createResponse(statusCode int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(HelloHandler)
}
```

## Always Test Your Logic

In serverless applications, unit tests are important, but don’t forget to also include integration tests, as most of these applications rely on integrations and policies to solve business problems.

```go
func TestCalculateCompoundInterest(t *testing.T) {
	principal := 100000000.0
	annualRate := 10.0
	years := 10

	result := CalculateCompoundInterest(principal, annualRate, years)
	lastElement := round(result[len(result)-1], 2)

	expected := round(259374246.01, 2)
	if !reflect.DeepEqual(lastElement, expected) {
		t.Errorf("Expected %v, but got %v", expected, lastElement)
	}
}
```
## Availability & Scalability
Serverless architectures are highly available by default and are event-driven, removing most operational tasks. However, if you choose to rely on ECS and containers, it’s important to include a load balancer to distribute traffic among your servers, ensuring both availability and scalability.

```
  CompoundingLoadBalancer:
    Type: AWS::ElasticLoadBalancingV2::LoadBalancer
    Properties:
      Name: compounding-nlb
      Scheme: internet-facing
      Type: network
      Subnets:
        - !Ref PublicSubnetOne
        - !Ref PublicSubnetTwo
```

## Deployment Pipeline

A deployment pipeline automates the software delivery process. We created a Makefile to simplify this process, making it easy to deploy and execute repetitive tasks with a single command. This approach enhances efficiency and consistency in your deployment workflow.


![CICD](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/cwycsr379esizpf47twe.png)



## Observability Is Mandatory

Ensure you have tracing, logging, and metrics in place. With serverless applications, enabling these features is as simple as adding `Tracing: Active`. The ability to see all logs in a central place like CloudWatch and monitor the interactions of the service is invaluable.


![Observability](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/iisqf5oq05f0x2ycype5.png)



## Security Is a First-Class Citizen in Cloud-Native Applications

Security is paramount in all the application. Using Amazon Cognito provides robust user authentication, while API keys add an additional layer of control and authorization, ensuring that only authorized clients can access your APIs.

```yaml
Auth:
  DefaultAuthorizer: CompoundingAuthorizer
  Authorizers:
    CompoundingAuthorizer:
      UserPoolArn:  XXXX
    LambdaTokenAuthorizer:
      FunctionArn: !GetAtt LambdaTokenAuthorizerFunction.Arn
      FunctionPayloadType: REQUEST
      Identity:
        Headers:
          - Authorization
        ReauthorizeEvery: 100
  AddDefaultAuthorizerToCorsPreflight: false
```

Assign the minimal necessary permissions to each service, user, and component to reduce the attack surface and prevent unauthorized access. **Least Privilege Principle**:

```
      Policies:
        - DynamoDBCrudPolicy:
            TableName: !Ref CompoundingTable
```

## References

1. [Terraform in Action](https://learning.oreilly.com/library/view/terraform-in-action/9781617296895/OEBPS/Text/Ch-09.htm#sigil_toc_id_148) - Practical uses and strategies for implementing Terraform, a tool for building, changing, and managing infrastructure.
2. [Continuous Delivery Pipelines](https://leanpub.com/cd-pipelines/)

## Conclusion

Software is constantly evolving, and while some tools and practices will change, the foundational principles remain the same. We need immutable infrastructure, CI/CD, good naming conventions, a robust testing strategy, security in our APIs, and efficiency in our applications. That’s why I decided to recreate this project in a serverless way.

**There has never been a better time to be an engineer and create value in society through software.**

- [LinkedIn](https://www.linkedin.com/in/jorgetovar-sa)
- [Twitter](https://twitter.com/jorgetovar621)
- [GitHub](https://github.com/jorgetovar)

If you enjoyed the articles, visit my blog [jorgetovar.dev](jorgetovar.dev)
