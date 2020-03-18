terraform {
  required_version = "= 0.12.19"
}

provider "aws" {
  version = "= 2.46"
  region  = "us-east-2"
}

//create mysql database
module "web_app" {
  source                                         = "../../modules/web_app"
  security_group_name = var.security_group_name
  ami = var.ami
  instance_type = var.instance_type
  db_address = var.db_address
  db_port = var.db_port
  db_name = var.db_name
  db_password = var.db_password
}