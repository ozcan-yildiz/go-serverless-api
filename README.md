# Go Serverless API: End-to-End Production Architecture

This project demonstrates a fully automated, production-grade **Serverless** architecture. It covers the entire lifecycle of a Go application, from containerization and CI/CD to custom domain management with SSL/TLS termination.

### 🌐 Live Endpoint
The API is globally accessible at: `https://api.atlascon.co.uk`

### 🏗️ Advanced Architecture
Beyond simple function execution, this project implements a complete cloud networking stack:

* **API Gateway (HTTP API):** Acts as the managed entry point, handling request routing to the Lambda function.
* **Route 53 (DNS):** Managed DNS records using **Alias** records to point a professional subdomain (`api.atlascon.co.uk`) to the AWS infrastructure.
* **AWS Certificate Manager (ACM):** Provisioned and validated SSL/TLS certificates to ensure all traffic is encrypted via HTTPS.
* **Lambda & ECR:** Container-based serverless execution using Go 1.x runtime on Amazon Linux 2.



### 🚀 CI/CD Pipeline Flow
The automation remains the heart of the project:
1.  **Build:** GitHub Actions triggers on every push to `main`.
2.  **Containerize:** Multi-stage Docker build creates a minimal footprint image.
3.  **Deploy:** The image is pushed to **AWS ECR**, and the **Lambda** function is updated via AWS CLI.
4.  **Instant Availability:** Changes are immediately reflected behind the **API Gateway** and custom domain.

### 🔐 Security & Governance
* **Granular IAM Policies:** Implemented least-privilege access for the GitHub deployer and the Lambda execution role.
* **Resource-Based Policies:** Configured ECR to allow cross-service access specifically for the Lambda service.
* **Zero-Downtime Updates:** Lambda's versioning ensures the API stays online during deployments.

### 🛠️ Key Technical Skills Demonstrated
* **Cloud Networking:** Configuring DNS A-records, Alias targets, and CNAMEs.
* **Security:** Managing SSL certificates and HTTPS enforcement.
* **Serverless Orchestration:** Connecting multiple AWS services (Route 53 -> API Gateway -> Lambda -> ECR) into a cohesive system.
* **DevOps Best Practices:** Decoupling infrastructure from code while maintaining full automation.

---
*Developed as part of a professional cloud engineering portfolio.*