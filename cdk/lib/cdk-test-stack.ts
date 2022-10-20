import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from "aws-cdk-lib/aws-lambda";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class CdkTestStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'CdkQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });

    new lambda.Function(this, 'main', {
        runtime: lambda.Runtime.GO_1_X,
        code: lambda.Code.fromAsset("cmd/lambda/main"),
        handler: "main",
        timeout: cdk.Duration.seconds(300),
        memorySize: 128
    })
  }
}
