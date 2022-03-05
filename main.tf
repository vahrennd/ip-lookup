terraform {
  required_providers {
    aws = {
      version = "~> 2.0"
    }
    archive = {
      version = "~> 1.3.0"
    }
  }
}

provider "aws" {
  region  = "us-east-1"
}

// Create a zip of the binary
data "archive_file" "zip" {
  type = "zip"
  source_file = "bin/ip-lookup"
  output_path = "ip-lookup.zip"
}

// Create lambda function
resource "aws_lambda_function" "ip-lookup" {
  function_name    = "ip-lookup"
  filename         = "ip-lookup.zip"
  handler          = "ip-lookup"
  source_code_hash = "data.archive_file.zip.output_base64sha256"
  role             = "${aws_iam_role.iam_for_lambda.arn}"
  runtime          = "go1.x"
  memory_size      = 128
}

// Create gateway API
resource "aws_api_gateway_rest_api" "api" {
  name = "ip_lookup_api"
}

resource "aws_api_gateway_resource" "resource" {
  path_part   = "ip-lookup"
  parent_id   = aws_api_gateway_rest_api.api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_resource.resource.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_resource.resource.id
  http_method             = aws_api_gateway_method.method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.ip-lookup.invoke_arn
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.ip-lookup.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
}

resource "aws_api_gateway_deployment" "ip_lookup_deploy" {
  depends_on = [aws_api_gateway_integration.integration]

  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = "v1"
}

output "url" {
  value = "${aws_api_gateway_deployment.ip_lookup_deploy.invoke_url}${aws_api_gateway_resource.resource.path}"
}

// Create access control role
resource "aws_iam_role" "iam_for_lambda" {
  name = "ip_lookup_iam"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}