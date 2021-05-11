import { LambdaRestApi } from "@aws-cdk/aws-apigateway";
import { Code, Function, Runtime } from "@aws-cdk/aws-lambda";
import { Bucket } from "@aws-cdk/aws-s3";
import * as cdk from "@aws-cdk/core";

export class InfrastructureStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const codeBucket = Bucket.fromBucketName(
      this,
      `${id}-code-bucket`,
      process.env["CODE_BUCKET_NAME"] ?? "bad time"
    );

    const dataBucket = new Bucket(this, `${id}-data-bucket`);

    const lambda = new Function(this, `${id}-lambda`, {
      runtime: Runtime.GO_1_X,
      handler: "oh-hell-backend",
      code: Code.fromBucket(
        codeBucket,
        `${process.env["GITHUB_SHA"]}.zip` ?? "bad time"
      ),
      environment: {
        BUCKET: dataBucket.bucketName,
        GIN_MODE: "release",
      },
    });

    new LambdaRestApi(this, `${id}-api`, {
      handler: lambda,
      deployOptions: {
        stageName: "dev",
      },
    });
  }
}
