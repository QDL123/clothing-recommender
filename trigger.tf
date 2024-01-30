resource "aws_cloudwatch_event_rule" "every_morning" {
  name                = "every-morning"
  description         = "Fires every morning at 8am Pacific"
  schedule_expression = "cron(0 16 ? * * *)"
}

resource "aws_cloudwatch_event_target" "run_lambda_every_morning" {
  rule      = aws_cloudwatch_event_rule.every_morning.name
  target_id = "clothing_recommender_lambda"
  arn       = aws_lambda_function.clothing_recommender_function.arn
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.clothing_recommender_function.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.every_morning.arn
}