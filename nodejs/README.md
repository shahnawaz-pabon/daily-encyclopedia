<div align="center">
  <a href="https://nodejs.org/en/">
    <img alt="nodejs" src="../logos/nodejs.png"/ height="96" width="96">
  </a>
  <h1>Nodejs</h1>
</div>

# Table of Contents

- [Install latest nvm on ubuntu](#install-latest-nvm-on-ubuntu)
- [Third party API call via node-fetch](#third-party-api-call-via-node-fetch)
- [Third party API call via axios](#third-party-api-call-via-axios)

## Install latest nvm on ubuntu

To always download and install the latest version of NVM, you can use the following URL which points to the latest release script on GitHub:

```sh
#!/bin/bash

# Update the package list and upgrade all packages
echo "Updating system packages..."
sudo apt update && sudo apt upgrade -y

# Install curl and build-essential if they are not already installed
echo "Installing required dependencies (curl and build-essential)..."
sudo apt install curl build-essential -y

# Download and run the NVM installation script from the latest release
echo "Installing NVM..."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/master/install.sh | bash

# Load NVM into the current shell session
echo "Loading NVM..."
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# Verify the installation by checking the NVM version
echo "Verifying NVM installation..."
nvm --version

# Optional: Install the latest Node.js version
# Uncomment the following line if you want to install the latest Node.js version automatically
# echo "Installing the latest Node.js version..."
# nvm install node

echo "NVM installation complete. You can now use NVM to manage Node.js versions."
```

To use this script:

1. Save it to a file, for example, install_nvm.sh.
2. Make the script executable:

```sh
chmod +x install_nvm.sh
```

3. Run the script

```sh
./install_nvm.sh
```

## Third party API call via node-fetch

> Used `node-fetch` version `v2.6.1`

```js
const fetch = require("node-fetch");

const method = "POST";
const url = "http://ip_address:port/api/v1/auth/login";
const headers = {
  "Content-Type": "application/json",
};

const body = JSON.stringify({
  username: "username",
  password: "password",
}); // your json data here

const response = await fetch(url, { body, method, headers }).then((res) =>
  res.json()
);

console.log(response);
```

## Third party API call via axios

> Used `axios` version `v1.2.0`

```js
const axios = require("axios");

const body = JSON.stringify({
  username: "username",
  password: "password",
}); // your json data here

let requestHeader = {
  headers: {
    "Content-Type": "application/json;charset=UTF-8",
    "Access-Control-Allow-Origin": "*",
    Authorization: "Bearer YourToken",
  },
};

axios
  .post(
    "http://ip_address:port/api/v1/notification-event/save",
    body,
    requestHeader
  )
  .then((res) => {
    console.log("RESPONSE RECEIVED");
  })
  .catch((err) => {
    console.log("AXIOS ERROR: ", err);
  });
```
