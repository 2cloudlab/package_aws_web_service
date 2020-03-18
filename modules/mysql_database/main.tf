terraform {
  required_version = "= 0.12.19"
}

resource "aws_db_instance" "db_instance" {
  identifier_prefix   = "2cloudlab.com"
  engine              = "mysql"
  allocated_storage   = 10
  instance_class      = "db.t2.micro"
  name                = var.db_name
  username            = var.db_username
  password            = var.db_password
  skip_final_snapshot = true
}