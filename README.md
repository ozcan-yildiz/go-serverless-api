# Go Serverless API: CI/CD Pipeline with AWS Lambda & ECR

This project demonstrates a modern, scalable, and cost-effective **Serverless** architecture using Go. Unlike traditional server-based deployments, this application runs as a functional unit that scales automatically and only incurs costs when executed.

### Project Overview
The goal was to transition from an EC2-based (Server-full) architecture to a **Serverless (Lambda)** approach. The application is containerized using Docker and deployed to AWS Lambda, providing a high-availability environment without the overhead of managing underlying servers.

### Tech Stack & Tools
* **Language:** Go (Golang)
* **Architecture:** Serverless
* **Containerization:** Docker (Multi-stage builds)
* **Registry:** AWS ECR (Elastic Container Registry)
* **Compute:** AWS Lambda
* **CI/CD:** GitHub Actions

### How the Pipeline Works
The deployment is fully automated to ensure "Zero-Touch" delivery:

1.  **Code Push:** Developers push code to the `main` branch.
2.  **Container Build:** GitHub Actions builds a specialized Docker image optimized for the **AWS Lambda Provided AL2** runtime.
3.  **Image Push:** The container image is pushed to **AWS ECR** with the `latest` tag.
4.  **Lambda Update:** The pipeline uses the AWS CLI to trigger a `UpdateFunctionCode` command, telling Lambda to pull and deploy the new image from ECR.

### Security & IAM Governance
Infrastructure security was prioritized by implementing granular access controls:
* **GitHub Secrets:** All AWS credentials are encrypted and never hardcoded.
* **Resource-Based Policies:** Implemented specific ECR Repository Policies to allow Lambda to securely pull images.
* **Least Privilege:** The IAM user and Lambda execution role are restricted to only the necessary actions (`ecr:BatchGetImage`, `lambda:UpdateFunctionCode`, etc.).

### Key Features
* **Cost Efficiency:** No idle server costs; you only pay per request.
* **Scalability:** AWS Lambda handles thousands of concurrent requests automatically.
* **Containerized Portability:** Using Docker for Lambda ensures the local development environment matches production exactly.
* **Automated Recovery:** AWS manages the infrastructure; if a function fails, the system is self-healing.