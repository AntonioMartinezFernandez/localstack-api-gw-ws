provider "aws" {
  access_key                  = "mock_access_key"
  region                      = "eu-central-1"
  s3_use_path_style           = true
  secret_key                  = "mock_secret_key"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    apigateway     = "http://localhost:4566"
    apigatewayv2   = "http://localhost:4566"
    cloudformation = "http://localhost:4566"
    cloudwatch     = "http://localhost:4566"
    dynamodb       = "http://localhost:4566"
    es             = "http://localhost:4566"
    firehose       = "http://localhost:4566"
    iam            = "http://localhost:4566"
    kinesis        = "http://localhost:4566"
    lambda         = "http://localhost:4566"
    route53        = "http://localhost:4566"
    redshift       = "http://localhost:4566"
    s3             = "http://localhost:4566"
    secretsmanager = "http://localhost:4566"
    ses            = "http://localhost:4566"
    sns            = "http://localhost:4566"
    sqs            = "http://localhost:4566"
    ssm            = "http://localhost:4566"
    stepfunctions  = "http://localhost:4566"
    sts            = "http://localhost:4566"
    iot            = "http://localhost:4566"
  }
}

# Create SNS
resource "aws_sns_topic" "sns_topic" {
  name = "custom-sns-topic"
}

# Create SQS
resource "aws_sqs_queue" "sqs_queue" {
  name = "custom-sqs-queue"
}

# Subscribe SQS queue to SNS topic
resource "aws_sns_topic_subscription" "subscription_sqs_to_sns" {
  topic_arn = aws_sns_topic.sns_topic.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.sqs_queue.arn
}

# API Gateway Websockets

// API GW
resource "aws_apigatewayv2_api" "api_gw_ws" {
  name                       = "custom-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

// default
resource "aws_apigatewayv2_integration" "api_gw_integration_default" {
  api_id           = aws_apigatewayv2_api.api_gw_ws.id
  integration_type = "HTTP"

  request_templates = {
    "application/json" = <<EOF
  {
     "input": $input.json('$'),
     "context": $context
  }
  EOF
  }

  integration_method = "POST"
  integration_uri    = "http://host.docker.internal:8000/default"
}

resource "aws_apigatewayv2_route" "api_gw_route_default" {
  api_id    = aws_apigatewayv2_api.api_gw_ws.id
  route_key = "$default"

  target = "integrations/${aws_apigatewayv2_integration.api_gw_integration_default.id}"
}

// connect
resource "aws_apigatewayv2_integration" "api_gw_integration_connect" {
  api_id           = aws_apigatewayv2_api.api_gw_ws.id
  integration_type = "HTTP_PROXY"

  integration_method = "POST"
  integration_uri    = "http://host.docker.internal:8000/connect"
}

resource "aws_apigatewayv2_route" "api_gw_route_connect" {
  api_id    = aws_apigatewayv2_api.api_gw_ws.id
  route_key = "$connect"

  target = "integrations/${aws_apigatewayv2_integration.api_gw_integration_connect.id}"
}

// disconnect
resource "aws_apigatewayv2_integration" "api_gw_integration_disconnect" {
  api_id           = aws_apigatewayv2_api.api_gw_ws.id
  integration_type = "HTTP_PROXY"

  integration_method = "POST"
  integration_uri    = "http://host.docker.internal:8000/disconnect"
}

resource "aws_apigatewayv2_route" "api_gw_route_disconnect" {
  api_id    = aws_apigatewayv2_api.api_gw_ws.id
  route_key = "$disconnect"

  target = "integrations/${aws_apigatewayv2_integration.api_gw_integration_disconnect.id}"
}
