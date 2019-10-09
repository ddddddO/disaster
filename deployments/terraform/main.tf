provider "google" {
  #project = "disaster"
  project = "disaster-245318"
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_compute_instance" "vm_instance" {
  name = "terra-instance"
  machine_type = "f1-micro"

  boot_disk {
	initialize_params {
	  image = "debian-cloud/debian-9"
	}
  }

  network_interface {
	#network = "default"
	network = "${google_compute_network.vpc_network.self_link}"
	access_config {
	}
  }
}

resource "google_compute_network" "vpc_network" {
  name = "terra-network"
  auto_create_subnetworks = "true"  
}
