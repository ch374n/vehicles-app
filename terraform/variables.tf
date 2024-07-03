variable "kube_config_path" {
  type        = string
  description = "Path to the kubeconfig file"
  default     = "~/.kube/config"
}

variable "kube_context" {
  type        = string
  description = "Kubernetes context to use"
  default     = "minikube"
}

variable "app_name" {
  type        = string
  description = "Name of the application"
  default     = "vehicles-app"
}

variable "app_image" {
  type        = string
  description = "Docker image for the application"
}

variable "app_replicas" {
  type        = number
  description = "Number of application replicas"
  default     = 1
}

variable "container_port" {
  type        = number
  description = "Port the container listens on"
  default     = 8080
}

variable "service_port" {
  type        = number
  description = "Port the service listens on"
  default     = 80
}
