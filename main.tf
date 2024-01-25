terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    archive = {
      source = "hashicorp/archive"
    }
    null = {
      source = "hashicorp/null"
    }
  }

  required_version = ">= 1.3.7"
}

# Provider block is used to configure the access to AWS
provider "aws" {
  region = "us-west-2"
  profile = "clothing-recommender"

  default_tags {
    tags = {
      app = "clothing-recommender-terraform"
    }
  }
}
