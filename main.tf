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

# resource "aws_lambda_function" "get_recommendation" {
#   function_name = "get_recommendation"
#   handler       = "clothing_recommender"
#   role          = aws_iam_role.iam_for_lambda.arn
#   runtime       = "go1.x"

#   filename = "get_recommendation.zip"
# }

# # This block is used to create an IAM role for Lambda service
# resource "aws_iam_role" "iam_for_lambda" {
#   name = "iam_for_lambda"

#   assume_role_policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Principal": {
#         "Service": "lambda.amazonaws.com"
#       },
#       "Effect": "Allow",
#       "Sid": ""
#     }
#   ]
# }
# EOF
# }

# # This block is used to create an EventBridge rule to trigger the 'get_recommendation' Lambda function
# resource "aws_cloudwatch_event_rule" "every_morning_at_8" {
#   name                = "every-morning-at-8"
#   schedule_expression = "cron(0 8 * * ? *)"
# }

# resource "aws_cloudwatch_event_target" "get_recommendation_target" {
#   rule      = aws_cloudwatch_event_rule.every_morning_at_8.name
#   target_id = "get_recommendation_target"
#   arn       = aws_lambda_function.get_recommendation.arn
# }

# resource "aws_lambda_permission" "allow_cloudwatch" {
#   statement_id  = "AllowExecutionFromCloudWatch"
#   action        = "lambda:InvokeFunction"
#   function_name = aws_lambda_function.get_recommendation.function_name
#   principal     = "events.amazonaws.com"
#   source_arn    = aws_cloudwatch_event_rule.every_morning_at_8.arn
# }
