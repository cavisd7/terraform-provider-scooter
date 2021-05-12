terraform {
  required_providers {
    scooter = {
      version = "0.0.1"
      source = "nil.xyz/ns/scooter"
    }
  }
}

provider "scooter" {
  address = "http://localhost"
  port    = 3001
}

resource "scooter_test_item" "my_test_item" {
  name = "mycoolresource"
  description = "Description of my cool item"
}
