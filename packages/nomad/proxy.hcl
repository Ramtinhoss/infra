variable "gcp_zone" {
  type    = string
}

variable "gcp_project" {
  type    = string
}

variable "gcp_region" {
  type    = string
}

variable "port" {
  type    = number
  default = 5009
}

variable "docker_registry" {
  type    = string
  default = ""
}

variable "api_secret" {
  type    = string
  default = ""
}

variable "otel_tracing_print" {
  type    = string
  default = ""
}

variable "environment" {
  type    = string
  default = ""
}

variable "proxy_checksum" {
  type    = string
  default = ""
}

variable "google_service_account_key" {
  type    = string
  default = ""
}

variable "bucket_name" {
  type    = string
  default = ""
}

variable "template_bucket_name" {
  type    = string
  default = ""
}

variable "otel_collector_grpc_endpoint" {
  type    = string
  default = ""
}

job "proxy" {
  datacenters = [var.gcp_zone]
  node_pool = "api"

  priority = 70

  group "proxy" {
    network {
      port "proxy" {
        static = var.port
      }
    }

    service {
      name = "proxy"
      port = var.port
    }

    task "start" {
      driver = "raw_exec"

      resources {
        memory     = 1024
        cpu        = 256
      }

      env {
        GOOGLE_SERVICE_ACCOUNT_BASE64 = var.google_service_account_key
        GCP_PROJECT_ID                = var.gcp_project
        GCP_REGION                    = var.gcp_region
        GCP_DOCKER_REPOSITORY_NAME    = var.docker_registry
        API_SECRET                    = var.api_secret
        OTEL_TRACING_PRINT            = var.otel_tracing_print
        ENVIRONMENT                   = var.environment
        TEMPLATE_BUCKET_NAME          = var.template_bucket_name
        OTEL_COLLECTOR_GRPC_ENDPOINT  = var.otel_collector_grpc_endpoint
      }

      config {
        command = "/bin/bash"
        args    = ["-c", " chmod +x local/proxy && local/proxy --port ${var.port}"]
      }

      artifact {
        source      = "gcs::https://www.googleapis.com/storage/v1/${var.bucket_name}/proxy"
        options {
          checksum    = "md5:${var.proxy_checksum}"
        }
      }
    }
  }
}
