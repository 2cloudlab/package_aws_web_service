output "public_ip" {
  value       = module.web_app.public_ip
  description = "The public IP of the Instance"
}

output "listening_port" {
  value = module.web_app.listening_port
  description = "port to listen"
}