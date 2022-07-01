<div align="center">
  <a href="https://www.nginx.com/">
    <img alt="nginx" src="../logos/nginx.png"/>
  </a>
  <h1>NGINX</h1>
</div>

# Table of Contents

- [Setup nginx on Ubuntu](#setup-nginx-on-ubuntu)
- [Configure a Server Block](#configure-a-server-block)

## Setup nginx on Ubuntu

```sh
# update software repositories
sudo apt-get update

# install nginx from ubuntu repositories
sudo apt-get install nginx

#  verify the installation
nginx -v

# controlling the nginx Service
sudo systemctl status nginx
sudo systemctl start nginx
sudo systemctl enable nginx
sudo systemctl stop nginx
# preventing nginx from loading when the system boots
sudo systemctl disable nginx
sudo systemctl reload nginx
sudo systemctl restart nginx

# system’s firewalls
sudo ufw app list
# granting nginx access through the default Ubuntu firewall
sudo ufw allow 'nginx http'
# encrypted (https) traffic
sudo ufw allow 'nginx https'
# refresh the firewall settings
sudo ufw reload
# to allow both
sudo ufw allow 'nginx full'

# test nginx
http://127.0.0.1
```

## Configure a Server Block

```sh
# create a directory for the test domain
sudo mkdir -p /var/www/test_domain.com/html

# configure ownership and permissions
sudo chown -R $USER:$USER /var/www/test_domain.com
sudo chmod -R 755 /var/www/test_domain.com

# create an index.html file for the server block
sudo vim /var/www/test_domain.com/html/index.html
```
