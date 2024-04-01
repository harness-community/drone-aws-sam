# drone-aws-sam-deploy

- [Synopsis](#Synopsis)
- [Parameters](#Parameters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin enables the deployment of AWS Serverless Application Model (SAM) applications. It provides various options for authenticating with AWS, including using access keys, session tokens, and assuming roles with or without web identity tokens.

## Parameters

| Parameter                                                                                                                        | Choices/<span style="color:blue;">Defaults</span> | Comments                                         |
| :------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | ------------------------------------------------ |
| AWS_ACCESS_KEY <span style="font-size: 10px"><br/>`string`</span>                                                                |                                                   | The AWS access key for authentication.           |
| AWS_SECRET_KEY <span style="font-size: 10px"><br/>`string`</span>                                                                |                                                   | The AWS secret key for authentication.           |
| AWS_SESSION_TOKEN <span style="font-size: 10px"><br/>`string`</span>                                                             |                                                   | The AWS session token for authentication.        |
| AWS_STS_EXTERNAL_ID <span style="font-size: 10px"><br/>`string`</span>                                                           |                                                   | The external ID for assuming a role with STS.    |
| AWS_ROLE_ARN <span style="font-size: 10px"><br/>`string`</span>                                                                  |                                                   | The ARN of the AWS role to assume.               |
| AWS_REGION <span style="font-size: 10px"><br/>`string`</span>                                                                    |                                                   | The AWS region for deployment.                   |
| TEMPLATE_FILE_PATH <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | The path to the SAM template file.               |
| STACK_NAME <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>         |                                                   | The name of the AWS CloudFormation stack.        |
| S3_BUCKET <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>          |                                                   | The S3 bucket for deployment artifacts.          |
| SESSION_NAME <span style="font-size: 10px"><br/>`string`</span>                                                                  | `sam-deploy-plugin`                               | Session name for AWS.                            |
| DEPLOY_COMMAND_OPTIONS <span style="font-size: 10px"><br/>`string`</span>                                                        |                                                   | Additional options for the `sam deploy` command. |

## Notes

- There are several ways of authenticating with AWS:
  - `AWS_ACCESS_KEY` and `AWS_SECRET_KEY`
  - `AWS_ACCESS_KEY`, `AWS_SECRET_KEY`, and `AWS_SESSION_TOKEN`, 
  - Only `AWS_ROLE_ARN` when EKS Cluster is already configured with required permissions and that `AWS_WEB_IDENTITY_TOKEN_FILE` is already present

## Plugin Image

The plugin `plugins/aws-sam-deploy` is available for the following architectures:

| OS          | Tag      |
| ----------- | -------- |
| linux-amd64 | `latest` |

## Examples

```
    - step:
        type: Plugin
        name: aws-sam-deploy
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: plugins/aws-sam-deploy:latest
                settings:
                    aws_access_key: ACCESS_KEY
                    aws_secret_key: SECRET_KEY
                    aws_region: us-east-1
                    stack_name: aws-sam
                    s3_bucket: sam-plugin
                    template_file_path: template.yaml

    - step:
        type: Plugin
        name: aws-sam-deploy
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: plugins/aws-sam-deploy:latest
                settings:
                    aws_access_key: ACCESS_KEY
                    aws_secret_key: SECRET_KEY
                    aws_session_token: SESSION_TOKEN
                    aws_region: us-east-1
                    stack_name: aws-sam
                    s3_bucket: sam-plugin
                    template_file_path: template.yaml

    - step:
        type: Plugin
        name: aws-sam-deploy
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: plugins/aws-sam-deploy:latest
                settings:
                    aws_access_key: ACCESS_KEY
                    aws_secret_key: SECRET_KEY
                    aws_region: us-east-1
                    stack_name: aws-sam
                    s3_bucket: sam-plugin
                    template_file_path: template.yaml
                    aws_role_arn: arn-role


    - step:
        type: Plugin
        name: aws-sam-deploy
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: plugins/aws-sam-deploy:latest
                settings:
                    aws_region: us-east-1
                    stack_name: aws-sam
                    s3_bucket: sam-plugin
                    template_file_path: template.yaml
                    aws_role_arn: arn-role
```
