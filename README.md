
# Go AWS DevOps CI Pipeline

## Project Overview

This project implements a continuous integration pipeline for a Go
application using AWS CodeCommit, AWS CodeBuild, AWS CodePipeline,
IAM and Amazon S3.

## Architecture

Developer -> Git -> CodeCommit -> CodePipeline -> CodeBuild -> S3 artifact

## Application

The application prints a greeting and includes automated unit tests.

## Local Setup

Requirements:
- Go 1.24 or compatible
- Git
- Visual Studio Code

Run the tests:

    go test ./...

Run the application:

    go run ./cmd/app

## Build Process

The buildspec.yml file defines the build process:
1. Select the Go runtime.
2. Check formatting and dependencies.
3. Run unit tests.
4. Compile the application for Linux AMD64.
5. Publish the build artifact.

## AWS Resources

- CodeCommit repository: go-aws-ci-pipeline
- CodeBuild project: go-aws-ci-build
- CodePipeline pipeline: go-aws-ci-pipeline
- S3: Pipeline artifact storage
- IAM: Service roles with least-privilege access

## IAM Roles

Document the actual role names created in your AWS account.

## Environment Variables

No custom environment variables or secrets are required for this
application. The Go runtime and target platform are configured
in buildspec.yml.

## Evidence

Screenshots of local tests, source repository, build logs,
pipeline execution and artifact storage are saved in screenshots/.

