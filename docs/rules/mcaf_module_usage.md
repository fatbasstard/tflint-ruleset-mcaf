# mcaf_module_usage

This rule ensures that an MCAF module is used instead of any other module.

## Example

```hcl
module "s3_bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "4.1.1"

  bucket = "my-s3-bucket"

  versioning = {
    enabled = true
  }
}```

```console
$ tflint
1 issue(s) found:

Warning: module s3_bucket should be replaced with MCAF module terraform-aws-mcaf-s3 (mcaf_module_usage)

  on test.tf line 1:
  79:   source  = "terraform-aws-modules/s3-bucket/aws"
```

## Why


## How To Fix

```hcl
module "s3_bucket" {
  source              = "github.com/schubergphilis/terraform-aws-mcaf-s3?ref=v0.11.0"
  name                = "my-s3-bucket"
  versioning          = true
}```
