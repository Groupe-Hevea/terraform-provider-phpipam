terraform {
  required_providers {
    phpipam = {
      version = "1.3.0"
      source  = "hashicorp.com/Groupe-Hevea/phpipam"
    }
  }
}

provider "phpipam" {
  app_id = "tf_rw"
  endpoint = "https://phpipam.lan"
  username = ""
  password = ""
}

data "phpipam_subnet" "my_subnet" {
  subnet_address = "192.168.1.0"
  subnet_mask    = 24
}

data "phpipam_nth_free_addresses" "ten_addresses" {
  number    = 10
  subnet_id = data.phpipam_subnet.my_subnet.subnet_id
}

output "some_free_addresses" {
  value = data.phpipam_nth_free_addresses.ten_addresses
}
