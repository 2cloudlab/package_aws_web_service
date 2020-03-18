terraform {
  required_version = "= 0.12.19"
}

resource "aws_instance" "ec2_instance" {
  ami                    = var.ami
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.security_group.id]

  user_data = <<-EOF
              #!/bin/bash
              echo "Connect to (${var.db_address},${var.db_port}) with (${var.db_name},${var.db_password})" > index.html
              nohup busybox httpd -f -p ${local.listening_port} &
              EOF

  tags = {
    Name = "2cloudlab"
  }
}

resource "aws_security_group" "security_group" {

  name = var.security_group_name

  ingress {
    from_port   = local.listening_port
    to_port     = local.listening_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

locals {
  listening_port = 8080
}
