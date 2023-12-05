<div align="center">
  <a href="https://www.r-project.org/">
    <img alt="r" src="../logos/r-logo.png" height="100" width="100"/>
  </a>
</div>

# Table of Contents

- [Install R](#install-r)

## Install R

### Install R from CRAN repo on Ubuntu

`CRAN` operates as a network of FTP and web servers, providing current releases of R code and documentation for diverse platforms. It includes a dedicated repository specifically for Ubuntu, ensuring access to the latest R version.

```sh
# update indices
sudo apt update -qq
# install two helper packages we need
sudo apt install --no-install-recommends software-properties-common dirmngr
# add the signing key (by Michael Rutter) for these repos
# To verify key, run gpg --show-keys /etc/apt/trusted.gpg.d/cran_ubuntu_key.asc 
# Fingerprint: E298A3A825C0D65DFD57CBB651716619E084DAB9
wget -qO- https://cloud.r-project.org/bin/linux/ubuntu/marutter_pubkey.asc | sudo tee -a /etc/apt/trusted.gpg.d/cran_ubuntu_key.asc
# add the R 4.0 repo from CRAN -- adjust 'focal' to 'groovy' or 'bionic' as needed
sudo add-apt-repository "deb https://cloud.r-project.org/bin/linux/ubuntu $(lsb_release -cs)-cran40/"
```