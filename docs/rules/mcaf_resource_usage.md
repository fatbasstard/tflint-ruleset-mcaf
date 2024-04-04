# mcaf_resource_usage

This rule ensures that an MCAF module is used instead of a basic Terraform resource.

## Example

```hcl
resource "aws_s3_bucket" "example" {
  bucket = "my-s3-bucket"
}
```

```console
$ tflint
1 issue(s) found:

Warning: resource aws_s3_bucket.example should be replaced with MCAF module terraform-aws-mcaf-s3 (mcaf_resource_usage)

  on test.tf line 1:
  1: resource "aws_s3_bucket" "example" {
```

## Why


## How To Fix

```hcl
module "s3_bucket" {
  source              = "github.com/schubergphilis/terraform-aws-mcaf-s3?ref=v0.11.0"
  name                = "my-s3-bucket"
}```
