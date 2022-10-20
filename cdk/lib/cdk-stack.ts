import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as iam from "aws-cdk-lib/aws-iam";
import {Fn} from "aws-cdk-lib";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class CdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const lambdaFunction = new lambda.Function(this, 'main', {
        runtime: lambda.Runtime.GO_1_X,
        code: lambda.Code.fromAsset("cmd/lambda"),
        handler: "main",
        timeout: cdk.Duration.seconds(300),
        memorySize: 128,
        role: iam.Role.fromRoleArn(
            this, 'LambdaRole', Fn.importValue(
                `core-data-common-CoreDataLambdaRole-dev-Arn`
            )
        )
    })
  }
}
