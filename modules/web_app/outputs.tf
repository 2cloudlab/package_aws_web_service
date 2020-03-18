output "public_ip" {
  value       = aws_instance.ec2_instance.public_ip
  description = "The public IP of the Instance"
}

output "listening_port" {
  value = local.listening_port
  description = "port to listen"
}
