# Go AWS DevOps CI/CD Pipeline

## 1. Project Overview

This project demonstrates the implementation of an automated Continuous Integration and Continuous Delivery (CI/CD) pipeline for a Go application using AWS managed DevOps services.

The objective is to automate source code management, application testing, compilation, artifact generation and pipeline execution using AWS CodeCommit, AWS CodeBuild, AWS CodePipeline, AWS Identity and Access Management (IAM), and Amazon S3.

The project follows DevOps principles, including automation, version control, repeatable builds, least-privilege access and continuous validation of application changes.

The application includes Go source code and automated unit tests. Whenever a developer pushes an approved change to the configured source branch, the pipeline can automatically retrieve the latest source code, execute the build specification, run tests and generate a deployable application artifact.

## 2. Project Objectives

The objectives of this project are to:

- Develop a Go application with automated unit tests.
- Manage source code using Git and AWS CodeCommit.
- Configure AWS CodeBuild to validate, test and compile the application.
- Use a `buildspec.yml` file to define repeatable build instructions.
- Orchestrate the workflow using AWS CodePipeline.
- Configure IAM service roles and permissions for secure communication between AWS services.
- Store build artifacts using Amazon S3.
- Demonstrate feature branch development and controlled code integration.
- Capture evidence of successful local and AWS-based testing.
- Document the configuration, security controls and implementation process.

## 3. Technology Stack

| Technology | Purpose |
|---|---|
| Go (Golang) | Application development |
| Visual Studio Code | Code editing and project management |
| Bash | Command-line operations and automation |
| Git | Version control |
| AWS CodeCommit | Source code repository |
| AWS CodeBuild | Automated testing and compilation |
| AWS CodePipeline | Pipeline orchestration |
| Amazon S3 | Pipeline artifact storage |
| AWS IAM | Identity and access management |
| AWS CloudWatch Logs | Build log monitoring |
| YAML | Build configuration |

## 4. Solution Architecture

The project uses AWS managed services to automate the movement of application code from source control through testing and artifact generation.

### CI/CD workflow

1. A developer creates or modifies Go source code locally.
2. Git records the changes in a commit.
3. The developer pushes the changes to AWS CodeCommit.
4. An AWS CodePipeline source action retrieves the latest committed code.
5. AWS CodeBuild executes the commands defined in `buildspec.yml`.
6. Go formatting and dependency checks are performed.
7. Automated unit tests validate the application.
8. The application is compiled into a Linux executable.
9. The build output is published as a pipeline artifact in Amazon S3.
10. Pipeline execution results and build logs are reviewed in the AWS Console.

**Architecture flow:**

`Developer → VS Code → Git → CodeCommit → CodePipeline → CodeBuild → Amazon S3`

An additional deployment stage can be configured if the assessment requires a complete Source, Build and Deploy pipeline. The deployment target and deployment configuration must be implemented before claiming successful deployment.

## 5. Repository Structure

The repository is organised to separate application code, tests, configuration and evidence.

```text
go-aws-ci-pipeline/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   └── greeting/
│       ├── greeting.go
│       └── greeting_test.go
├── screenshots/
│   ├── 01-local-environment.png
│   ├── 02-go-tests.png
│   ├── 03-git-initial-commit.png
│   ├── 04-codecommit-repository.png
│   ├── 05-git-push.png
│   ├── 06-codecommit-source-files.png
│   ├── 07-iam-service-roles.png
│   ├── 08-codebuild-project.png
│   ├── 09-codebuild-success.png
│   ├── 10-codepipeline-configuration.png
│   ├── 11-codepipeline-succeeded.png
│   ├── 12-s3-artifact.png
│   └── 13-feature-branch-workflow.png
├── .gitignore
├── buildspec.yml
├── go.mod
└── README.md
```

The screenshots directory should contain genuine evidence captured from the local development environment and AWS Management Console. Filenames can be adjusted to match the actual screenshots submitted.

## 6. Application Description

The Go application generates a greeting message and demonstrates the use of a separate internal package.

The application entry point is located at:

`cmd/app/main.go`

The greeting function is implemented in:

`internal/greeting/greeting.go`

Automated tests are defined in:

`internal/greeting/greeting_test.go`

The tests validate the expected output for a supplied name and for an empty name.

This structure demonstrates basic Go modularity, separation of responsibilities and automated unit testing.

## 7. Local Development and Testing

### Prerequisites

The local environment requires:

- Go installed and available on the system PATH.
- Git installed and configured.
- Visual Studio Code.
- Bash terminal, Git Bash or WSL2.
- An AWS account with the necessary service permissions.
- AWS CLI v2 for AWS command-line operations, if required.

### Verify the environment

```bash
git --version
go version
aws --version
```

### Run the unit tests

From the project root:

```bash
go test ./...
```

### Run the application

```bash
go run ./cmd/app
```

### Compile the application

```bash
mkdir -p dist

go build -o dist/go-aws-ci-pipeline ./cmd/app
```

### Validate formatting

```bash
gofmt -w cmd/app/main.go internal/greeting/*.go
```

These commands provide a repeatable local process for validating the application before changes are pushed to CodeCommit.

## 8. Git Version Control and CodeCommit

Git is used to track application changes, maintain commit history and manage development branches.

AWS CodeCommit hosts the remote repository and provides the source used by CodePipeline.

### Repository configuration

The repository name is:

`go-aws-ci-pipeline`

The primary branch is:

`main`

The remote repository URL should be obtained from the CodeCommit Console rather than copied from an unrelated example.

### Initial Git configuration

```bash
git init
git branch -M main

git add .
git commit -m "Initial Go application and CI configuration"

git remote add origin YOUR_CODECOMMIT_HTTPS_URL
git push -u origin main
```

Replace the remote URL placeholder with the actual repository URL.

### Feature branch workflow

Feature branches allow developers to make changes independently before integrating them into the main branch.

Create a feature branch:

```bash
git checkout main
git pull origin main

git checkout -b feature/improve-greeting
```

Make the application changes, run the tests and commit the work:

```bash
go test ./...

git add .
git commit -m "Improve greeting application and tests"

git push -u origin feature/improve-greeting
```

Where the repository and available tooling support pull requests, open a pull request from `feature/improve-greeting` into `main`, review the changes and merge the branch after approval.

If pull requests are not available in the configured CodeCommit environment, demonstrate the branch merge locally:

```bash
git checkout main
git merge --no-ff feature/improve-greeting \
  -m "Merge feature/improve-greeting"

git push origin main
```

A merge into `main` can trigger the configured pipeline when source-change detection is enabled.

**Evidence required:** Include a screenshot of the feature branch, commit history, pull request where available, or successful merge in the Git history. Do not claim that a pull request was completed unless one was actually created and reviewed.

## 9. AWS CodeCommit Configuration

AWS CodeCommit is the source control service used by the pipeline.

The repository contains:

- Go application source files.
- Automated unit tests.
- `go.mod`.
- `buildspec.yml`.
- `.gitignore`.
- `README.md`.

The configured pipeline source should point to the correct CodeCommit repository and branch.

### Verification

1. Open the AWS Management Console.
2. Navigate to CodeCommit.
3. Open the `go-aws-ci-pipeline` repository.
4. Confirm that the source files are present.
5. Confirm that the latest commit is visible.
6. Verify that the pipeline is configured to use the intended branch.

**Evidence:** Add a screenshot showing the repository name, branch and committed files.

## 10. AWS CodeBuild Configuration

AWS CodeBuild provides the managed build environment used to validate and compile the Go application.

### Build project

The example project name is:

`go-aws-ci-build`

The project should use:

- AWS CodeCommit as the source provider for an independent build test.
- A supported managed Linux build image.
- A Go runtime compatible with `go.mod`.
- `buildspec.yml` as the build specification.
- A CodeBuild service role with the necessary permissions.
- CloudWatch Logs for build output.

The exact image version and available Go runtimes should be verified in the AWS Console.

### Build phases

The build specification defines the following phases:

**Install:** Prepares the Go runtime.

**Pre-build:** Checks Go formatting and validates module dependencies.

**Build:** Runs unit tests and compiles the application for Linux AMD64.

**Post-build:** Confirms the build output is available for packaging.

**Artifacts:** Publishes the compiled executable and README file.

### Verification

After starting a build:

1. Open the CodeBuild project.
2. Navigate to the build history.
3. Select the latest execution.
4. Confirm that the build status is `SUCCEEDED`.
5. Review the logs to verify that tests passed and compilation completed.
6. Confirm that the expected output artifact was generated.

**Evidence:** Include a screenshot of the CodeBuild build history showing a successful execution and a screenshot of the relevant build logs.

## 11. AWS CodePipeline Configuration

AWS CodePipeline orchestrates the source and build actions.

The example pipeline name is:

`go-aws-ci-pipeline`

### Source stage

The Source stage retrieves the committed code from AWS CodeCommit.

Configuration:

- Source provider: AWS CodeCommit.
- Repository: `go-aws-ci-pipeline`.
- Branch: `main`.
- Change detection: Enabled where supported and configured.

### Build stage

The Build stage invokes AWS CodeBuild using the project:

`go-aws-ci-build`

CodeBuild receives the source artifact, executes `buildspec.yml` and produces the build output.

### Deploy stage

If the assessment requires Source, Build and Deploy stages, a valid deployment provider and deployment target must be configured.

The deployment target depends on the chosen architecture. For example, an S3 deployment action can publish static application files, but it does not by itself execute a compiled Go server application.

For a deployable Go service, select an appropriate deployment target and configure its permissions and deployment process before adding the stage.

Do not add an empty or misleading deployment stage merely to make the pipeline appear complete.

### Verification

A successful pipeline execution should demonstrate:

- Source action completed successfully.
- Build action completed successfully.
- Deploy action completed successfully, if deployment has been configured.
- The latest execution corresponds to the intended commit.
- The resulting artifact is available at the configured destination.

**Evidence:** Include a screenshot showing the complete pipeline execution and the status of every configured stage. If all three stages are required, capture the full Source, Build and Deploy pipeline with each stage showing `Succeeded`.

## 12. IAM Security and Permissions

AWS IAM controls the permissions used by developers and AWS services.

This project follows the principle of least privilege: each identity should receive only the permissions needed to perform its task.

### Developer identity

The developer identity needs permissions to create and manage the project resources and to access the CodeCommit repository.

Depending on how the project is configured, required permissions can include:

- `codecommit:CreateRepository`
- `codecommit:GetRepository`
- `codecommit:GitPull`
- `codecommit:GitPush`
- `codebuild:CreateProject`
- `codebuild:StartBuild`
- `codebuild:BatchGetBuilds`
- `codepipeline:CreatePipeline`
- `codepipeline:GetPipeline`
- `codepipeline:GetPipelineState`
- `codepipeline:StartPipelineExecution`
- Relevant IAM role and policy management permissions.
- Relevant S3 permissions for artifact inspection.

These actions are examples of relevant permissions, not a claim that all are attached to the developer identity. Resource creation permissions should be restricted to the required resources where practical.

### CodeBuild service role

The CodeBuild service role allows CodeBuild to execute builds.

Depending on the configured source and artifact handling, permissions can include:

- CloudWatch Logs permissions to create log streams and publish build logs.
- S3 permissions to read input artifacts and write output artifacts.
- Any additional permissions required by the chosen source provider.

For CodePipeline-integrated builds, ensure the CodeBuild role can access the configured artifact objects and encryption keys, if applicable.

### CodePipeline service role

The CodePipeline service role allows the pipeline to coordinate the source, build and deployment actions.

Depending on the configured stages, permissions can include:

- CodeCommit permissions to retrieve repository information and source revisions.
- CodeBuild permissions to start builds and retrieve build status.
- S3 permissions to read and write pipeline artifacts.
- Relevant deployment-service permissions if a Deploy stage is implemented.
- AWS KMS permissions if the artifact bucket uses a customer-managed encryption key.

### Trust relationships

Service roles should trust the appropriate AWS service principal. For example, CodeBuild roles normally trust `codebuild.amazonaws.com`, while CodePipeline roles normally trust `codepipeline.amazonaws.com`.

### Documenting the actual policies

The final submission must identify the actual IAM role names and policies attached in the AWS account.

| Identity | Actual role or user name | Attached policies |
|---|---|---|
| Developer | Complete from AWS Console | Record actual policies |
| CodeBuild | Complete from AWS Console | Record actual policies |
| CodePipeline | Complete from AWS Console | Record actual policies |
| Deployment service, if used | Complete from AWS Console | Record actual policies |

Review the IAM Console and record the actual policies rather than assuming that example actions are already attached.

Avoid using root credentials for routine work, committing access keys or granting unrestricted administrator access to service roles.

## 13. AWS Resource Tagging Strategy

Consistent tags make resources easier to identify, manage and audit.

Where supported, apply tags such as:

| Tag key | Example value |
|---|---|
| `Project` | `go-aws-ci-pipeline` |
| `Environment` | `Development` |
| `ManagedBy` | `Student` |
| `Purpose` | `CI-CD-Learning` |
| `Owner` | `Project-Owner` |

Apply appropriate tags to supported resources, including CodeBuild projects, CodePipeline resources, S3 buckets and IAM resources where tagging is supported.

These are recommended tag values. Document the actual values used in your AWS account.

## 14. Amazon S3 Artifact Storage

Amazon S3 is used as the pipeline artifact store.

The pipeline artifact bucket holds the files exchanged between pipeline stages. The build output should include the compiled Go executable and README file.

The exact artifact object name and location may be generated by CodePipeline.

### Verification

1. Open the CodePipeline configuration.
2. Identify the configured artifact store bucket.
3. Open the corresponding S3 bucket.
4. Locate the relevant pipeline execution artifacts.
5. Confirm that the build completed successfully and produced an output artifact.

Do not make the pipeline artifact bucket public. Keep the bucket permissions restricted to the relevant AWS services and authorised users.

**Evidence:** Include a screenshot showing the S3 bucket and the relevant artifact object.

## 15. Environment Variables and Build Configuration

The application does not require custom secrets or external API credentials.

The build process uses the runtime and target platform settings in `buildspec.yml`.

The compilation command is:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o dist/go-aws-ci-pipeline ./cmd/app
```

This produces a Linux AMD64 executable without requiring a C toolchain for this application.

No passwords, access keys or other secrets should be placed in `buildspec.yml` or committed to the repository.

If additional configuration becomes necessary, use appropriate AWS-managed secret storage and access controls.

## 16. Testing and Validation

Testing takes place locally and in AWS CodeBuild.

### Local testing

```bash
go test ./...
go run ./cmd/app
```

### AWS testing

The AWS pipeline validates that:

1. Source code is retrieved successfully.
2. The Go runtime is available.
3. Formatting and dependency checks pass.
4. Unit tests complete successfully.
5. The Go executable compiles successfully.
6. The output artifact is generated.
7. Pipeline execution results are available for review.

### Automatic trigger testing

To verify automatic change detection:

1. Create a small code change on a feature branch.
2. Run the local tests.
3. Commit and push the feature branch.
4. Merge the change into `main` using the documented workflow.
5. Confirm that the new commit reaches CodeCommit.
6. Open CodePipeline and inspect the new execution.
7. Confirm that the configured stages complete successfully.

This demonstrates how version control and pipeline automation support continuous integration.

## 17. Evidence and Screenshots

The following screenshots should be included in the `screenshots/` directory.

| Evidence | Required content |
|---|---|
| Local environment | Go, Git and AWS CLI version output |
| Local tests | Successful `go test ./...` output |
| Git history | Descriptive commits and branch workflow |
| CodeCommit | Repository, branch and source files in AWS Console |
| IAM | Actual service role names and attached policies |
| CodeBuild configuration | Source, environment and build specification |
| CodeBuild history | Successful build showing `SUCCEEDED` |
| CodeBuild logs | Passing tests and successful compilation |
| CodePipeline | Full configured pipeline execution |
| Deployment stage | Successful deployment status, if required |
| S3 artifact | Artifact bucket and output object |
| Automatic trigger | New commit followed by a pipeline execution |

Screenshots should be genuine, readable and consistent with the resource names documented in this README. Remove or obscure sensitive account details where appropriate.

## 18. Troubleshooting

| Issue | Recommended action |
|---|---|
| Go tests fail | Inspect the test output and correct the application or test expectations. |
| Build specification not found | Confirm that `buildspec.yml` is committed at the repository root. |
| CodeCommit authentication fails | Verify Git credentials and repository access permissions. |
| CodeBuild reports `AccessDenied` | Review the CodeBuild service role and artifact permissions. |
| CodePipeline fails to start a build | Check the pipeline role, selected CodeBuild project and source configuration. |
| Automatic execution does not start | Verify the configured branch, source-change detection and latest remote commit. |
| Artifact cannot be found | Confirm the pipeline artifact store bucket and inspect the correct execution's output. |
| Deployment stage fails | Review the selected provider, deployment configuration, target permissions and deployment logs. |

## 19. Project Outcomes

This project demonstrates practical knowledge of:

- Go application development and unit testing.
- Git version control and feature branch workflows.
- AWS CodeCommit repository management.
- AWS CodeBuild configuration and build automation.
- AWS CodePipeline orchestration.
- IAM roles, service trust relationships and least-privilege access.
- Amazon S3 artifact management.
- Automated change detection and pipeline validation.
- Technical documentation and evidence-based assessment.

## 20. Conclusion

The Go AWS DevOps CI/CD Pipeline demonstrates how source control, automated testing, compilation and artifact storage can be integrated using managed AWS services.

The implementation provides a foundation for extending the workflow with additional quality checks, pull request controls, security scanning and a suitable deployment target.

Successful completion should be demonstrated through the actual AWS Console configuration, passing build logs, pipeline execution history, IAM policy documentation, source branch workflow and artifact evidence.

## 21. Project Links

**GitHub repository:** Add your own public repository URL here, if applicable.

**AWS CodeCommit repository:** Add the repository name and console reference used for your project.

**AWS CodeBuild project:** `go-aws-ci-build`

**AWS CodePipeline pipeline:** `go-aws-ci-pipeline`

**Evidence directory:** `screenshots/`

*Note: Replace example names and descriptions with the actual configuration from your AWS account. Only mark stages as successful after verifying the corresponding AWS Console execution.*


