data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

variable "database_url" {
  type    = string
  default = getenv("DATABASE_URL")
}


env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/16/dev"
  url = var.database_url
  migration {
    dir = "file://src/db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}