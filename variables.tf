# Module upgraded to v14

variable "bucket_name" {
  description = "The name to give the bucket"
}

variable "versioning" {
  default     = "true"
  description = "Enables versioning for objects in the S3 bucket"
}

variable "aws_region" {
  default     = ""
  description = "AWS Region where the S3 bucket will be created"
  type        = string
}

variable "force_destroy" {
  description = "Whether to allow a forceful destruction of this bucket"
  default     = false
}

