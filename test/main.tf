# tf-module upgrade complete #

variable "aws_region" {
  default = "us-east-1"
  type    = string
}

variable "bucket_name" {
  default = "test-bucket-tf-state"
}

module "bootstrap" {
  source      = "../"
  aws_region  = var.aws_region
  bucket_name = var.bucket_name
}

output "bucket_id" {
  value = module.bootstrap.bucket_id
}

output "bucket_region" {
  value = module.bootstrap.region
}

