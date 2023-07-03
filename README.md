# Module upgraded to v14

# AWS Bootstrap Module

This module is to create an S3 bucket to be used for the Terraform State.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| aws\_region | AWS Region where the S3 bucket will be created | string | `""` | no |
| bucket\_name | The name to give the bucket | string | n/a | yes |
| force\_destroy | Whether to allow a forceful destruction of this bucket | string | `"false"` | no |
| versioning | Enables versioning for objects in the S3 bucket | string | `"true"` | no |

## Outputs

| Name | Description |
|------|-------------|
| bucket\_arn | `arn` exported from `aws_s3_bucket` |
| bucket\_id | `id` exported from `aws_s3_bucket` |
| region | `region` exported from `aws_s3_bucket` |
| url | Derived URL to the S3 bucket |


## Sample

```
module "bootstrap" {
  source      = "git::ssh: //git@bitbucket.org/deloittepe/tf-module-aws-bootstrap.git"
  aws_region  = "${var.aws_region}"
  bucket_name = "${var.bucket_name}"
}
```
