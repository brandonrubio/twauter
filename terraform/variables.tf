variable "subscription_id" {
  type = string
}

variable "tenant_id" {
  type = string
}

variable "client_id" {
  type        = string
  sensitive   = true
  description = "ID for service principal"
}

variable "client_secret" {
  type        = string
  sensitive   = true
  description = "Secret for authenticating with service principal"
}

variable "location" {
  type        = string
  default     = "East US"
  description = "Region where the datacenter that hosts the vm is located"
}

variable "resource_group_name" {
  type        = string
  default     = "default_resource_group"
  description = "Name of the container that holds all of the related resources (vm, vnet, subnt, etc.)"
}

variable "vnet_name" {
  type        = string
  description = "Name of the Azure Virtual Network"
}

variable "address_space" {
  type        = list(string)
  description = "Range of IP addresses for a VNet"
}

variable "subnet_name" {
  type        = string
  description = "Name of the Azure Subnet"
}

variable "address_prefixes" {
  type        = list(string)
  description = "Range of IP addresses for a subnet within a VNet's address space"
}

variable "network_interface_name" {
  type        = string
  description = "Name of the virtual network adapter used to connect the vm and the network"
}

variable "vm_name" {
  type        = string
  description = "Name of the Azure Virtual Machine"
}

variable "vm_username" {
  type      = string
  sensitive = true
}
