variable "vm_project_id" {
  default = "testing"
}

variable "vm_name" {
  default = "auto"
}

variable "vm_size" {}

variable "vm_image" {}

variable "vm_bootstrap_runlist" {
  default = []
  type = "list"
}

variable "vm_http_proxy" {}

variable "vm_no_proxy" {
  default = []
  type = "list"
}

variable "vm_env" {}

variable "vm_ip" {
  default = ""
}

variable "vm_sg" {
  default = ["ALL"]

  type = "list"
}

variable "vm_keypair" {
  default = "mykeypair"
}

variable "vm_count" {
  default = 1
}

variable "vm_chefuser" {
  default = "mydefaultusername"
}

varialbe "vm_chef_server" {
   default = https://mychefserver.com/organizations/myorg"	 
}


resource "openstack_compute_instance_v2" "openstack_vm" {
  count           = "${var.vm_count}"
  name            = "${format("%s-%s", var.vm_project_id, var.vm_name)}"
  image_id        = "${var.vm_image}"
  key_pair        = "${var.vm_keypair}"
  security_groups = "${var.vm_sg}"
  flavor_name     = "${var.vm_size}"
  floating_ip     = "${var.vm_ip}"

  metadata {
    role = "docker_testing"
  }

  provisioner "chef" {
    node_name       = "${self.name}"
    server_url      = "${var.vm_chef_server}"
    secret_key      = "${file("encrypted_data_bag_secret")}"
    environment     = "${var.vm_env}"
    user_name       = "jhu"
    user_key        = "${file("../chef-repo/.chef/mykey.pem")}"
    recreate_client = true
    run_list        = "${var.vm_bootstrap_runlist}"
    http_proxy      = "${var.vm_http_proxy}"
    https_proxy     = "${var.vm_http_proxy}"
    no_proxy	    = "${var.vm_no_proxy}"
    version         = "12.13.37"

    connection {
      user        = "${var.vm_chefuser}"
      private_key = "${file("~/.ssh/id_rsa")}"
    }
  }
}


