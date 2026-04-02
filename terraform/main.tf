provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_namespace" "test" {
  metadata {
    name = "test-namespace"
  }
}
