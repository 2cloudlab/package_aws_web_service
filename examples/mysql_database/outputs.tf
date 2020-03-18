output "address" {
  value       = module.mysql_database.address
  description = "Connect to the database at this endpoint"
}

output "port" {
  value       = module.mysql_database.port
  description = "The port the database is listening on"
}