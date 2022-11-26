<div align="center">
  <a href="https://nodejs.org/en/">
    <img alt="nodejs" src="../logos/nodejs.png"/ height="96" width="96">
  </a>
  <h1>Nodejs</h1>
</div>

# Table of Contents

- [Third party API call via node-fetch](#third-party-api-call-via-node-fetch)
- [Third party API call via axios](#third-party-api-call-via-axios)

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
