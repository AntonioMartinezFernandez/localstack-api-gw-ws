output "api_gw_endpoint" {
  value = aws_apigatewayv2_api.api_gw_ws.api_endpoint
}

output "sqs_url" {
  value = aws_sqs_queue.sqs_queue.url
}
