<div align="center">
  <a href="https://ubuntu.com/">
    <img alt="ubuntu" src="../logos/ubuntu.png"/>
  </a>
  <h1>Ubuntu</h1>
</div>

# Table of Contents

- [Get the Audio Codec for your machine's model](#get-the-audio-codec-for-your-machine's-model)
- [Fix - Package system is broken](#fix---package-system-is-broken)
- [dpkg - error processing package](#dpkg---error-processing-package)
- [Monitor system processes with htop](#monitor-system-processes-with-htop)
- [Copy a file from remote to local and vice versa via scp](#copy-a-file-from-remote-to-local-and-vice-versa-via-scp)
- [Operate mysql at boot time](#operate-mysql-at-boot-time)
- [Checking open ports](#checking-open-ports)

## Get the Audio Codec for your machine's model

```sh
cat /proc/asound/card*/codec* | grep Codec
```

<br>

## Fix - Package system is broken

```sh
# to force an install of the files that didn't get loaded because of the error
sudo apt-get -f install
sudo apt-get update
# run back and forth until only the package that has the error is left.
sudo apt-get -f install
sudo dpkg --configure -a
# clean the cache
sudo apt-get clean
```

<br>

## dpkg - error processing package

```sh
sudo apt --fix-broken install
```

<br>

## Monitor system processes with htop

```sh
sudo apt install htop
htop
```

<br>

## Copy a file from remote to local and vice versa via scp

### Copy from remote to local

```sh
scp -r -P 80 username@remote_ip:/remote/directory/ /local/directory
```

### Copy from local to remote

```sh
scp -P 80 /local/directory/drive-download-20220720T120659Z-001.zip username@remote_ip:/remote/directory
```

<br>

## Operate mysql at boot time

### Stop running at boot time

```sh
sudo systemctl disable mysql
```

### Check whether enabled currently

```sh
sudo systemctl is-enabled mysql
```

### Start running at boot time

```sh
sudo systemctl enable mysql
```

<br>

## Checking open ports

```sh
sudo lsof -i -P -n | grep LISTEN
sudo netstat -tulpn | grep LISTEN
sudo ss -tulpn | grep LISTEN
sudo lsof -i:22 ## see a specific port such as 22 ##
sudo nmap -sTU -O IP-address-Here
```
