<div align="center">
  <a href="https://nodejs.org/en/">
    <img alt="nodejs" src="../logos/nodejs.png"/ height="96" width="96">
  </a>
  <h1>Nodejs</h1>
</div>

# Table of Contents

- [Third party API call via node-fetch](#third-party-api-call-via-node-fetch)

## Third party API call via node-fetch

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
