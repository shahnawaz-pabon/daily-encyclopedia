<div align="center">
  <a href="https://go.dev/">
    <img alt="go" src="../logos/go.png"/ height="96" width="96">
  </a>
  <h1>GO</h1>
</div>

# Table of Contents

- [Install GO on linux](#install-go-on-linux)

## Install GO on linux

```sh
# The -O flag ensures that this outputs to a file, and the L flag instructs HTTPS redirects, since this link was taken from the Go website and will redirect here before the file downloads
curl -OL https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
# To verify the integrity of the file you downloaded by running the sha256sum command and pass it to the filename as an argument
sha256sum go1.17.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.17.7.linux-amd64.tar.gz
```

```sh
sudo vim ~/.bashrc
```

Add the following information to the end of your file:

```sh
export PATH=$PATH:/usr/local/go/bin
```

```sh
# Refresh your profile by running the following command
source ~/.bashrc
# Check go version by running
go version
```

My distro was `Ubuntu 20.04`
