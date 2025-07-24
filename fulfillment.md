# AutoRest for Marketplace SaaS Fulfillment API code generation

This markdown file can be passed as `autorest fulfillment.md`, and the block below will be automagically parsed as configuration for the AutoRest code generation step. 

See [here for Go-specifc AutoRest options](https://github.com/Azure/autorest/blob/main/docs/generate/flags.md#go-flags).

``` yaml
go: true
output-folder: './fulfillment'
file-prefix: 'gen_'
module: github.com/fastah/azuremarketplacesaas/fulfillment
module-version: '0.11.0'
openapi-type: data-plane
azure-arm: false
license-header: MICROSOFT_MIT_NO_VERSION
credential-scope: "20e940b3-4c77-4b0b-9a53-9e16a1b010a7/.default"
```
