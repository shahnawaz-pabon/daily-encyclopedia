<div align="center">
  <a href="https://www.r-project.org/">
    <img alt="r" src="../logos/r-logo.png" height="100" width="100"/>
  </a>
</div>

# Table of Contents

1. [Installing R from the CRAN Repository](#installing-r-from-the-cran-repository)
2. [Installing Jupyter Notebook](#installing-jupyter-notebook)
3. [Setting Up Jupyter to Use the R Kernel](#setting-up-jupyter-to-use-the-r-kernel)
4. [Installing DataExplorer Package in R on Ubuntu](#installing-dataexplorer-package-in-r-on-ubuntu)

## Installing R from the CRAN Repository


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
# update apt cache
sudo apt update
# install r
sudo apt install --no-install-recommends r-base
# launching the R shell as the root user
sudo -i R
```

## Installing Jupyter Notebook

### Steps to Install Jupyter Notebook

1. **Update your package list and upgrade existing packages:**

    ```bash
    sudo apt update
    sudo apt upgrade
    ```

2. **Install Python and pip:**

    ```bash
    sudo apt install python3-pip
    ```

3. **Upgrade pip to the latest version:**

    ```bash
    pip3 install --upgrade pip
    ```

4. **Install Jupyter Notebook:**

    ```bash
    pip3 install jupyter
    ```

5. **Start Jupyter Notebook:**

    ```bash
    jupyter notebook
    ```

## Setting Up Jupyter to Use the R Kernel

### Steps to Install and Register IRKernel

1. **Open an R session:**

    ```bash
    R
    ```

2. **Within the R session, run the following commands to install IRKernel and register it with Jupyter:**

    ```r
    install.packages('IRkernel')
    IRkernel::installspec()
    ```

    **Note:** If you encounter a permission denied error, run R with `sudo`:

    ```bash
    sudo R
    ```

    Then, run the IRKernel installation commands again within the R session.

3. **Verify that the R kernel is available in Jupyter:**

    ```bash
    jupyter kernelspec list
    ```

    You should see an entry for `ir` among the listed kernels.

## Installing DataExplorer Package in R on Ubuntu

### Instructions

Follow the steps below to install the `DataExplorer` package in R along with its dependencies on an Ubuntu system.

### Step 1: Install System Libraries

First, install the necessary system libraries using the following commands:

```bash
sudo apt install build-essential libxml2-dev libglpk40 libglpk-dev libcgal-dev libcurl4-openssl-dev libssl-dev
sudo apt install gfortran
```

### Step 2: Install Required R Packages

Next, open your R console or RStudio and run the following commands to install the required R packages:

```r
# Install required system libraries (if not already installed)
install.packages("curl")
install.packages("openssl")

# Install required R packages
install.packages("igraph")
install.packages("httr")
install.packages("networkD3")
install.packages("covr")

# Install DataExplorer package
install.packages("DataExplorer")
```

### Verification

Finally, you can load the DataExplorer package to ensure it was installed correctly:

```r
library(DataExplorer)
```