# tf-module upgrade complete #
resource "aws_s3_bucket" "remote-state" {
  bucket        = var.bucket_name
  acl           = "private"
  force_destroy = var.force_destroy

  versioning {
    enabled = var.versioning
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }
}

resource "aws_s3_bucket_public_access_block" "public_access_block" {
  bucket                  = aws_s3_bucket.remote-state.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

