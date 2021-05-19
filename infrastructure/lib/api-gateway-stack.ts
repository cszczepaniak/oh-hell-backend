import {
  BasePathMapping,
  DomainName,
  EndpointType,
  LambdaRestApi,
} from "@aws-cdk/aws-apigateway";
import { Function } from "@aws-cdk/aws-lambda";
import { Certificate } from "@aws-cdk/aws-certificatemanager";
import * as cdk from "@aws-cdk/core";
import { HostedZone, CnameRecord } from "@aws-cdk/aws-route53";

interface Props extends cdk.StackProps {
  lambda: Function;
}

export class ApiGatewayStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: Props) {
    super(scope, id, props);

    if (!props?.lambda) {
      throw new Error(`lambda is a required property`);
    }

    const cert = Certificate.fromCertificateArn(
      this,
      `${id}-cert`,
      process.env["SSL_CERT_ARN"] || "bad time!"
    );

    const customDomain = new DomainName(this, "customDomain", {
      domainName: "api.oh-heck.com",
      certificate: cert,
      endpointType: EndpointType.REGIONAL,
    });

    const api = new LambdaRestApi(this, `${id}-api`, {
      handler: props.lambda,
      deployOptions: {
        stageName: "dev",
      },
    });

    new BasePathMapping(this, `${id}-basepath`, {
      domainName: customDomain,
      restApi: api,
    });

    const hostedZone = HostedZone.fromHostedZoneAttributes(
      this,
      `${id}-hosted-zone`,
      {
        hostedZoneId: process.env["HOSTED_ZONE_ID"] || "bad time!",
        zoneName: "oh-heck.com",
      }
    );

    new CnameRecord(this, `${id}-api-gateway-record`, {
      zone: hostedZone,
      recordName: "api",
      domainName: customDomain.domainNameAliasDomainName,
    });
  }
}
