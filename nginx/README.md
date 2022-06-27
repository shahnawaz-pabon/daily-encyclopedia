<div align="center">
  <a href="https://www.nginx.com/">
    <img alt="nginx" src="../logos/nginx.png"/>
  </a>
  <h1>NGINX</h1>
</div>

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

```
