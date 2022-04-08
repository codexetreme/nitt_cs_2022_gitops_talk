terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

# https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/resources/database_cluster
resource "digitalocean_database_cluster" "postgres" {
  name       = "gitops-postgres-cluster-1"
  engine     = "pg"
  version    = "14"
  size       = "db-s-1vcpu-1gb"
  region     = "blr1"
  node_count = 1
}

resource "digitalocean_database_db" "students_db" {
  cluster_id = digitalocean_database_cluster.postgres.id
  name       = "students"
}

locals {
#  https://www.terraform.io/language/functions/replace
  private-conn_str = replace(digitalocean_database_cluster.postgres.private_uri, digitalocean_database_cluster.postgres.database, digitalocean_database_db.students_db.name)
  public-conn_str = replace(digitalocean_database_cluster.postgres.uri, digitalocean_database_cluster.postgres.database, digitalocean_database_db.students_db.name)
}


output "public-connection-string" {
  value     = local.public-conn_str
  sensitive = true
}

# https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/resources/app
resource "digitalocean_app" "gitops-app" {
  spec {
    name   = "gitops-app"
    region = "blr"

    service {
      name               = "students-app"
      environment_slug   = "go"
      instance_count     = 1
      instance_size_slug = "professional-xs"
      image {
        registry_type = "DOCR"
        repository    = "students-app"
        tag = "v-8a730eb"
      }
      env {
        key = "MODE"
        value = "staging"
        type = "GENERAL" # SECRET, GENERAL
      }
      env {
        key = "DATABASE_URL"
        value = local.public-conn_str
        type = "SECRET" # SECRET, GENERAL
      }
    }
    database {
      name = "db"
      production = "true"
      db_name = digitalocean_database_db.students_db.name
      db_user = digitalocean_database_cluster.postgres.user
      cluster_name = digitalocean_database_cluster.postgres.name
      version = digitalocean_database_cluster.postgres.version
      engine = "PG"
    }
  }
}

output "app-live-url" {
  value = digitalocean_app.gitops-app.live_url
}

resource "digitalocean_database_firewall" "allow-only-app" {
  cluster_id = digitalocean_database_cluster.postgres.id

  rule {
    type  = "app"
    value = digitalocean_app.gitops-app.id
  }
  rule {
    type  = "ip_addr"
    value = "49.207.210.156"
  }
}