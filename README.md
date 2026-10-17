# Go AWS DevOps CI/CD Pipeline

![Language](https://img.shields.io/badge/Language-Go-00ADD8?logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-AWS-232F3E?logo=amazonaws&logoColor=white)
![Workflow](https://img.shields.io/badge/Workflow-CI%2FCD-blue)

## 1. Project Overview

This project demonstrates a continuous integration (CI) workflow for a Go application using AWS managed DevOps services. It combines local development in Visual Studio Code, Git version control, AWS CodeCommit, AWS CodeBuild, AWS CodePipeline, AWS Identity and Access Management (IAM), Amazon S3, and Amazon CloudWatch Logs.

The pipeline is designed to retrieve committed source code, validate and test the Go application, compile a Linux executable, and publish build output as a pipeline artifact. The intended workflow is:

**Developer → VS Code/Bash → Git → CodeCommit → CodePipeline → CodeBuild → S3 artifact store**

This README documents the intended architecture and the evidence required to verify the actual AWS implementation. AWS resources, permissions, successful statuses, and screenshots must be confirmed in the account before being reported as completed.

## 2. Project Objectives

- Build and test a small Go application.
- Track source changes with Git and host the remote repository in AWS CodeCommit.
- Define repeatable build steps in `buildspec.yml`.
- Use CodeBuild to run checks, unit tests, and compile the application.
- Use CodePipeline to orchestrate source and build actions.
- Store pipeline artifacts in Amazon S3.
- Apply least-privilege IAM permissions and document service-role trust relationships.
- Verify automatic pipeline execution after a commit is pushed to the configured branch.
- Demonstrate a feature-branch and merge workflow.
- Provide clear AWS Console screenshots as assessment evidence.

## 3. Technology Stack

| Technology | Purpose |
|---|---|
| Go (Golang) | Application and unit tests |
| Visual Studio Code | Editing and project navigation |
| Bash / Git Bash / WSL2 | Local terminal commands |
| Git | Version control and branching |
| AWS CodeCommit | Remote source repository |
| AWS CodeBuild | Automated validation, testing, and compilation |
| AWS CodePipeline | Source-to-build workflow orchestration |
| Amazon S3 | Pipeline artifact storage |
| AWS IAM | Identity, roles, policies, and service trust |
| Amazon CloudWatch Logs | Build logs and troubleshooting |
| YAML | Build specification configuration |

## 4. Architecture and Workflow

1. The developer edits Go source code in VS Code.
2. Git records changes in a commit.
3. The commit is pushed to the configured CodeCommit branch.
4. CodePipeline detects the source change when change detection is enabled and correctly configured.
5. The Source action produces an input artifact for the next stage.
6. CodeBuild reads `buildspec.yml`, runs the configured checks and unit tests, and compiles the application.
7. CodeBuild publishes its configured output artifact.
8. CodePipeline stores and passes artifacts through its S3 artifact store.
9. The developer checks the execution result and logs in the AWS Console.

**Required baseline pipeline:** `Source → Build`

A Deploy stage is optional unless specifically required by the assignment. If a Deploy stage is added, configure a real deployment provider and target; do not add an empty stage or describe an artifact upload as deploying a running Go service.

## 5. Repository Structure

The following is the intended layout. Adjust it to match the files actually present in your repository.

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
│   ├── 02-local-tests.png
│   ├── 03-git-history.png
│   ├── 04-codecommit-repository.png
│   ├── 05-codecommit-branches.png
│   ├── 06-codebuild-project.png
│   ├── 07-codebuild-succeeded.png
│   ├── 08-codebuild-logs.png
│   ├── 09-codepipeline-succeeded.png
│   ├── 10-automatic-trigger.png
│   ├── 11-s3-artifact.png
│   ├── 12-iam-codebuild-role.png
│   └── 13-iam-codepipeline-role.png
├── .gitignore
├── AWS_ENVIRONMENT_SETUP.md
├── buildspec.yml
├── go.mod
└── README.md
```

Only include screenshot filenames once the corresponding screenshots have been captured and committed. Do not create empty or illustrative screenshots to represent AWS results.

## 6. Local Setup and Validation

### Prerequisites

- Go installed and available on `PATH`.
- Git installed.
- Visual Studio Code.
- Bash, Git Bash, or WSL2.
- An AWS account and authorised access to the required services.
- AWS CLI v2 if using the command line to inspect AWS identity or resources.

Check the tools:

```bash
go version
git --version
aws --version
```

Run the tests from the repository root:

```bash
go test ./...
```

Run the application:

```bash
go run ./cmd/app
```

Format the Go files:

```bash
gofmt -w cmd/app/main.go internal/greeting/*.go
```

Compile locally, if the repository follows the structure above:

```bash
mkdir -p dist
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o dist/go-aws-ci-pipeline ./cmd/app
```

If your project uses different file paths, update the commands to match the actual layout. Capture the successful local test output as evidence, but keep it separate from AWS CodeBuild evidence.

## 7. CodeCommit Repository and Git Workflow

The expected repository name is `go-aws-ci-pipeline`, and the expected primary branch is `main`. Verify both in the AWS Console and in the configured Git remote.

Check the remote:

```bash
git remote -v
git branch --show-current
git log --oneline --decorate -5
```

If connecting an existing local repository to CodeCommit for the first time, use the HTTPS clone/push URL copied from the CodeCommit Console. Do not use a guessed URL or commit credentials to the repository.

### Feature branch workflow

Create a feature branch:

```bash
git checkout main
git pull origin main
git checkout -b feature/ci-improvements
```

Make a small, relevant change, run tests, and commit it:

```bash
go test ./...
git add README.md
git commit -m "Document CI pipeline improvements"
git push -u origin feature/ci-improvements
```

If a pull-request workflow is available in the configured environment, open a pull request, review the change, and merge it into `main`. If not, demonstrate a local merge and push the updated branch:

```bash
git checkout main
git pull origin main
git merge --no-ff feature/ci-improvements \
  -m "Merge feature/ci-improvements"
git push origin main
```

Document which method was actually used. Do not claim that a pull request was created if the workflow used a local merge instead.

## 8. AWS CodeCommit Configuration

In the AWS Console, open CodeCommit and verify:

- The repository exists and is in the expected AWS Region.
- The repository contains the application files, `go.mod`, `buildspec.yml`, `.gitignore`, and README.
- The `main` branch exists and contains the latest commit.
- A feature branch and merge history are visible if you completed that workflow.
- The CodePipeline Source action references the correct repository and branch.

**Evidence to capture:** repository overview, branch list, and commit history. The local `git log` output alone does not prove that the remote repository is correctly configured.

## 9. AWS CodeBuild Configuration

The expected CodeBuild project name is `go-aws-ci-build`. Verify the actual name in the Console.

Check these settings:

- Source is configured appropriately for a standalone CodeBuild test or for the CodePipeline-provided input artifact.
- The build environment is a supported Linux image.
- The Go runtime version is compatible with `go.mod` and the selected image.
- The build specification points to `buildspec.yml` at the correct path.
- The CodeBuild service role can write logs and access the configured source and artifacts.
- CloudWatch Logs are enabled.

The build specification should run the checks and tests, compile the Go application, and identify the intended output artifacts. The actual commands and artifact paths must agree with the repository layout.

### Verify a build

1. Open CodeBuild and select the project.
2. Start a build if one is not already available.
3. Open the latest build execution.
4. Confirm the status is `SUCCEEDED`.
5. Inspect the logs for passing tests, successful compilation, and artifact upload.
6. Record the build ID and date for your assessment notes.

**Evidence to capture:** project configuration, build history with a `SUCCEEDED` execution, and logs showing the test/build steps. A README statement alone is not evidence of a successful AWS build.

## 10. AWS CodePipeline Configuration

The expected pipeline name is `go-aws-ci-pipeline`. Verify the actual pipeline name and Region.

### Source stage

- Provider: AWS CodeCommit.
- Repository: `go-aws-ci-pipeline`.
- Branch: `main`.
- Source-change detection: enabled and correctly configured if automatic triggering is required.

### Build stage

- Provider: AWS CodeBuild.
- Build project: `go-aws-ci-build`.
- Input: the artifact produced by the Source stage.
- Output: the configured build artifact, where required by the pipeline.

### Execution verification

1. Open CodePipeline and select the pipeline.
2. Run the pipeline or push a new commit to the configured branch.
3. Open the execution details.
4. Confirm Source and Build succeeded.
5. Check that the execution corresponds to the expected commit.
6. If automatic triggering is required, push a new commit and verify that a new execution starts without manually selecting **Release change** or **Start pipeline**.

**Evidence to capture:** the full pipeline view with Source and Build showing `Succeeded`, and a separate screenshot of the new execution caused by a source change.

## 11. Amazon S3 Artifact Verification

CodePipeline uses an S3 artifact store to pass artifacts between actions. The bucket and object names may be generated automatically.

To verify the artifact:

1. Open the pipeline details/settings and identify the artifact store bucket.
2. Open S3 in the same AWS Region.
3. Select the identified bucket.
4. Locate the relevant pipeline execution artifacts.
5. Confirm that the build action completed and produced an output artifact.
6. If the build output is separately downloadable, inspect it to verify the expected files.

The build output may contain a compiled Linux executable and README, depending on the actual `buildspec.yml`. Do not assume the S3 object will be named `go-aws-ci-pipeline.zip`.

Keep the artifact bucket private. Do not enable public access to make evidence easier to view.

**Evidence to capture:** the bucket name and the relevant artifact object in the S3 Console.

## 12. IAM Permissions and Trust Relationships

Use `AWS_ENVIRONMENT_SETUP.md` to record the **actual** role names, attached policies, relevant actions, resource scope, and trust relationships. The examples below are areas to inspect, not a claim that these permissions are already attached.

### Developer identity

Depending on the tasks performed, the developer identity may need permissions to access CodeCommit, inspect/start CodeBuild and CodePipeline executions, inspect S3 artifacts, and configure project resources. Resource creation and IAM administration permissions should be limited to what is necessary.

### CodeBuild service role

Inspect whether the role has the permissions needed for:

- CloudWatch Logs stream creation and log publishing.
- Reading source or input artifacts.
- Writing output artifacts to the configured S3 location.
- Using the configured encryption key, if applicable.

The CodeBuild service role normally trusts `codebuild.amazonaws.com`.

### CodePipeline service role

Inspect whether the role has the permissions needed for:

- Retrieving source information and source artifacts from CodeCommit.
- Starting the configured CodeBuild project and reading its status.
- Reading and writing pipeline artifacts in S3.
- Using the configured encryption key, if applicable.
- Invoking a deployment provider only if a Deploy stage is configured.

The CodePipeline service role normally trusts `codepipeline.amazonaws.com`.

### IAM evidence

Capture the Permissions tab for the CodeBuild and CodePipeline roles and their Trust relationships tabs. Record attached managed policies and relevant inline policy actions in `AWS_ENVIRONMENT_SETUP.md`. Where practical, note which resources each policy applies to.

Do not assume a managed policy is attached just because it is commonly used. Verify the policy name and contents in the IAM Console. Avoid publishing access keys, credentials, or sensitive account details.

## 13. Resource Tagging Strategy

Apply a consistent tagging scheme where supported and record the actual tags used.

| Tag key | Example value |
|---|---|
| `Project` | `go-aws-ci-pipeline` |
| `Environment` | `Development` |
| `ManagedBy` | `Student` |
| `Purpose` | `CI-CD-Learning` |

Check the resource's Tags section in AWS. These are suggested values, not confirmation that the resources have already been tagged.

## 14. Environment Variables and Secrets

The basic Go greeting application does not require external API credentials. The build uses the runtime and target-platform settings defined in `buildspec.yml`.

Do not store AWS access keys, passwords, or tokens in source files, README content, or committed build configuration. If future extensions require secrets, use an appropriate managed secret store and grant access only to the required role.

Record any non-secret environment variables actually configured for CodeBuild in `AWS_ENVIRONMENT_SETUP.md`. If none are configured, state that explicitly.

## 15. Troubleshooting

| Problem | What to check |
|---|---|
| `go: command not found` locally | Install Go and verify `PATH`. |
| Tests fail | Read the test output and correct code or expectations. |
| `buildspec.yml` not found | Confirm it is committed at the configured path, normally the repository root. |
| CodeCommit authentication fails | Verify the repository URL, Git authentication setup, and repository permissions. |
| CodeBuild `AccessDenied` | Review the CodeBuild service role, S3 permissions, and KMS access if applicable. |
| Pipeline cannot start CodeBuild | Check the pipeline role, project selection, and input artifact configuration. |
| Pipeline does not trigger automatically | Confirm the pushed branch, source-change detection, and latest remote commit. |
| S3 artifact is missing | Check the configured artifact store and the relevant pipeline execution. |
| Build succeeds locally but fails in AWS | Compare Go versions, build image, paths, environment variables, and build logs. |

## 16. Assessment Evidence Checklist

Place genuine screenshots in `screenshots/` and use the actual filenames in the table below.

| Evidence | What it must show | Filename example |
|---|---|---|
| Local environment | Go, Git, and AWS CLI versions | `01-local-environment.png` |
| Local unit tests | Successful `go test ./...` | `02-local-tests.png` |
| Git history | Descriptive commits | `03-git-history.png` |
| CodeCommit repository | Repository and committed files | `04-codecommit-repository.png` |
| CodeCommit branches | `main` and feature branch if created | `05-codecommit-branches.png` |
| CodeBuild project | Source, environment, buildspec | `06-codebuild-project.png` |
| CodeBuild success | Build history with `SUCCEEDED` | `07-codebuild-succeeded.png` |
| Build logs | Passing tests and compilation | `08-codebuild-logs.png` |
| CodePipeline success | Source and Build both `Succeeded` | `09-codepipeline-succeeded.png` |
| Automatic trigger | New commit and resulting execution | `10-automatic-trigger.png` |
| S3 artifact | Bucket and relevant object | `11-s3-artifact.png` |
| IAM roles | Actual role permissions and trust | `12-iam-codebuild-role.png`, `13-iam-codepipeline-role.png` |

If a screenshot has not yet been captured, mark that item as pending in your own progress notes rather than implying the evidence exists.

## 17. Final Validation

Before submission:

- [ ] `go test ./...` passes locally.
- [ ] CodeCommit contains the latest project files.
- [ ] Feature branch and merge workflow is documented with real evidence.
- [ ] CodeBuild has a successful execution.
- [ ] CodePipeline Source and Build stages succeed.
- [ ] A new push triggers a pipeline execution automatically, if configured.
- [ ] The S3 artifact has been located and evidenced.
- [ ] Actual IAM policies and trust relationships are documented.
- [ ] Resource tags are recorded accurately.
- [ ] Screenshots are readable, genuine, and referenced correctly.
- [ ] No credentials or secrets are committed.

## 18. Conclusion

This project demonstrates how a Go application can be validated and compiled through AWS-managed CI services. The combination of Git, CodeCommit, CodeBuild, CodePipeline, IAM, CloudWatch Logs, and S3 provides a repeatable workflow for source control, automated testing, build execution, and artifact storage.

The quality of the submission should be demonstrated through actual AWS configuration and execution evidence, not documentation alone. The repository, successful build, successful pipeline stages, artifact object, IAM permissions, and branch workflow should all be verified before the project is submitted.

## 19. Project Details

- **Project:** Go AWS DevOps CI/CD Pipeline
- **CodeCommit repository:** `go-aws-ci-pipeline` (verify in AWS Console)
- **CodeBuild project:** `go-aws-ci-build` (verify in AWS Console)
- **CodePipeline pipeline:** `go-aws-ci-pipeline` (verify in AWS Console)
- **AWS Region:** Record the actual region
- **Evidence folder:** `screenshots/`
- **Environment documentation:** `AWS_ENVIRONMENT_SETUP.md`

> **Accuracy note:** Replace example names, paths, and screenshot filenames where your actual configuration differs. Only report a stage, artifact, IAM policy, branch, or automatic trigger as completed after verifying it in your environment.
