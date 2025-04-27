location               = "East US"
resource_group_name    = "twaut-resource-group"
vnet_name              = "twaut-vnet-prod-aze"
subnet_name            = "twaut-subnet-vm-aze"
address_space          = ["10.0.0.0/16"]
address_prefixes       = ["10.0.2.0/24"]
network_interface_name = "twaut-nic-vm-aze-api"
vm_name                = "twaut-vm-api-prod-aze"
