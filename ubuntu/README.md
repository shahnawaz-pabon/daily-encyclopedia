<div align="center">
  <a href="https://ubuntu.com/">
    <img alt="ubuntu" src="../logos/ubuntu.png"/>
  </a>
  <h1>Ubuntu</h1>
</div>

# Table of Contents

- [Get the Audio Codec for your machine's model](#get-the-audio-codec-for-your-machine's-model)

## Get the Audio Codec for your machine's model

```sh
cat /proc/asound/card*/codec* | grep Codec
```
