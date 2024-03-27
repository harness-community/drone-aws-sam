# drone-aws-sam-build

- [Synopsis](#Synopsis)
- [Parameters](#Parameters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin enables building AWS Serverless Application Model (SAM) applications using the `sam build` command. It supports various options for building, including using a specific Docker image, providing build command options, and authenticating with a private Docker registry.

## Parameters

| Parameter                                                                                                                        | Choices/<span style="color:blue;">Defaults</span> | Comments                                                          |
| :------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | ----------------------------------------------------------------- |
| BUILD_IMAGE <span style="font-size: 10px"><br/>`string`</span>                                                                   |                                                   | The Docker image to use for building the SAM application.         |
| TEMPLATE_FILE_PATH <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | The path to the SAM template file.                                |
| BUILD_COMMAND_OPTIONS <span style="font-size: 10px"><br/>`string`</span>                                                         |                                                   | Additional options for the `sam build` command.                   |
| PRIVATE_REGISTRY_URL <span style="font-size: 10px"><br/>`string`</span>                                                          |                                                   | The URL of the private Docker registry.                           |
| PRIVATE_REGISTRY_USERNAME <span style="font-size: 10px"><br/>`string`</span>                                                     |                                                   | The username for authenticating with the private Docker registry. |
| PRIVATE_REGISTRY_PASSWORD <span style="font-size: 10px"><br/>`string`</span>                                                     |                                                   | The password for authenticating with the private Docker registry. |

## Notes

- If `BUILD_IMAGE` is provided, the `sam build` command will use the specified Docker image for building the SAM application.
- If `PRIVATE_REGISTRY_URL`, `PRIVATE_REGISTRY_USERNAME`, and `PRIVATE_REGISTRY_PASSWORD` are provided, the plugin will authenticate with the specified private Docker registry before building the SAM application.

## Plugin Image

The plugin `harnesscommunitytest/aws-sam-build` is available for the following architectures:

| OS          | Tag      |
| ----------- | -------- |
| linux-amd64 | `latest` |

## Examples

```
    - step:
        type: Plugin
        name: aws-sam-build
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: harnesscommunitytest/aws-sam-build
                settings:
                    TEMPLATE_FILE_PATH: template.yaml
                    BUILD_IMAGE: public.ecr.aws/sam/build-python3.9:1.112.0-20240313001230


    - step:
        type: Plugin
        name: aws-sam-build
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: harnesscommunitytest/aws-sam-build
                settings:
                    TEMPLATE_FILE_PATH: template.yaml

    - step:
        type: Plugin
        name: aws-sam-build
        identifier: sam_plugin
        spec:
                connectorRef: <connector>
                image: harnesscommunitytest/aws-sam-build
                settings:
                    TEMPLATE_FILE_PATH: template.yaml
                    BUILD_IMAGE: image
                    PRIVATE_REGISTRY_URL: registry-url
                    PRIVATE_REGISTRY_USERNAME: username
                    PRIVATE_REGISTRY_PASSWORD: password
```
