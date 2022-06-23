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

## Get the Audio Codec for your machine's model

```sh
cat /proc/asound/card*/codec* | grep Codec
```

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

## dpkg - error processing package

```sh
sudo apt --fix-broken install
```
