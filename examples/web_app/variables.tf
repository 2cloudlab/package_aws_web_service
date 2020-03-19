variable "security_group_name" {
  description = "The name of the security group"
  type        = string
  default     = "web_app_security_group_name"
}

variable "ami" {
  description = "The id of Amazon Machine Image"
  type = string
  default = "ami-0fc20dd1da406780b"
}

variable "instance_type" {
  description = "Instance type"
  type = string
  default = "t2.micro"
}

variable "db_address" {
  description = "db address"
  type = string
  default = "127.0.0.1"
}

variable "db_port" {
  description = "db port"
  type = string
  default = "3306"
}

variable "db_name" {
    description = "db name"
    type = string
    default = "self_name"
}

variable "db_password" {
    description = "db password"
    type = string
    default = "self_password"
}
