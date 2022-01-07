# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/focal64"

    config.vm.provider :virtualbox do |vb|
          vb.customize ["modifyvm", :id, "--memory", "4096", "--cpus", "2"]
    end

    config.vm.provision "shell", inline: <<-SHELL
        apt update
        apt install -y gnupg software-properties-common curl
        curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add -
        apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
        apt update && apt install terraform make gcc

        wget https://dl.google.com/go/go1.17.5.linux-amd64.tar.gz
        tar -C /usr/local/ -xzf go1.17.5.linux-amd64.tar.gz
        rm go1.17.5.linux-amd64.tar.gz
        echo 'export PATH="$PATH:/usr/local/go/bin"' >> /home/vagrant/.profile

        echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list
        apt update
        apt install goreleaser
    SHELL
end
