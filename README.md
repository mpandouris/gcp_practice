# gcp_practice
This repo is a simple service that demonstrates the following:

1. Simple Golang HTTP server
1. Full CI/CD pipeline:

    1. Deployment to DEV, UAT namespaces
    1. Canary deployments for isolated testing based on feature branches.
    1. Necessary cleanup to avoid leftover stale deployments
    1. Self-Hosted github runner.

1. Terraform for building the following infrastructure:

    1. GKE cluster
    1. Artifact Registry for service images
    1. VPC Networks for communication between services
    1. Terraform state sharing via a GCP bucket
1. Helm Chart
1. Prometheus
1. Grafana for dashboard visualisation of the prometheus logs and metrics

---

