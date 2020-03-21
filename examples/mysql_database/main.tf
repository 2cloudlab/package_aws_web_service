terraform {
  required_version = "= 0.12.19"
}

provider "aws" {
  version = "= 2.46"
  region  = "us-east-2"
}

//create mysql database
module "mysql_database" {
  source      = "../../modules/mysql_database"
  db_name     = var.db_name
  db_username = var.db_username
  db_password = var.db_password
}