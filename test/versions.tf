
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
  required_version = ">= 0.14"
}

provider "aws" {
  region  = var.aws_region
  version = "~> 2.0"
}

provider "null" {
  version = "~> 2.0"
}

provider "random" {
  version = "~> 2.0"
}