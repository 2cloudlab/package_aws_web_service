variable "security_group_name" {
  description = "The name of the security group"
  type        = string
  default     = "terraform-example-instance"
}

variable "ami" {
  description = "The id of"
  type        = string
}

variable "instance_type" {
  description = "Instance type"
  type        = string
}

variable "db_address" {
  description = "db address"
  type        = string
}

variable "db_port" {
  description = "db port"
  type        = string
}

variable "db_name" {
  description = "db name"
  type        = string
}

variable "db_password" {
  description = "db password"
  type        = string
}
