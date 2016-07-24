# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.network "forwarded_port", guest: 80, host: 8080
  config.vm.provider "virtualbox" do |vb|
    vb.memory = 1024
    vb.cpus = 2
  end

  config.vm.provision "go", type: "shell", inline: <<-SHELL
    # Install Go 1.6
    wget --quiet https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.6.linux-amd64.tar.gz
    rm go1.6.linux-amd64.tar.gz
  SHELL
  config.vm.provision "enable_go", type: "shell", privileged: false, inline: <<-SHELL
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
  SHELL

  config.ssh.forward_agent = true
  config.vm.synced_folder ".", "/opt/gopath/src/github.com/bostongolang/golang-lab-slack"
end
