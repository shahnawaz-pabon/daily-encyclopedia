<div align="center">
  <a href="https://www.javascript.com/">
    <img alt="javascript" src="../logos/javascript.gif"/ height="96" width="96">
  </a>
  <h1>JavaScript</h1>
</div>

## Miscellaneous codes

### Get the index of the object inside an array, matching a condition

```js
let a = [
  { prop1: "abc", prop2: "qwe" },
  { prop1: "bnmb", prop2: "yutu" },
  { prop1: "zxvz", prop2: "qwrq" },
];

let index = a.findIndex((x) => x.prop2 === "yutu");

console.log(index);
```

<br>

### Make unique array of objects from the duplicate one

```js
let duplicate_array_of_objects = [
  { id: "1", prop2: "qwe" },
  { id: "2", prop2: "yutu" },
  { id: "1", prop2: "qwe" },
  { id: "2", prop2: "yutu" },
];

let unique_array_of_objects = duplicate_array_of_objects.filter(function (
  elem
) {
  return !this[elem.id] && (this[elem.id] = true);
},
Object.create(null));

console.log("Unique array");
console.log(unique_array_of_objects);
```

<br>

### Remove a specific object from the array of objects, matching a condition

```js
let array = [
  { id: 1, prop2: "qwe" },
  { id: 2, prop2: "yutu" },
];

let len = array.length;

for (let i = 0; i < len; i++) {
  let index = array.findIndex((x) => x.id === 1);
  if (index > -1) {
    array.splice(index, 1); // 2nd parameter means remove one item only
  }
}

console.log(array);
```

<br>

## Keep only specific key-value pairs from an array of objects

```js
// We want to keep only prop1 and prop2
let a = [
  { prop1: "abc", prop2: "qwe", prop3: "abcd" },
  { prop1: "bnmb", prop2: "yutu", prop3: "abcd" },
  { prop1: "zxvz", prop2: "qwrq", prop3: "abcd" },
];

const result = a.map((object) =>
  Object.fromEntries(
    Object.entries(object).filter(
      ([key]) => key.includes("prop1") || key.includes("prop2")
    )
  )
);

console.log(result);
```

<br>

###

```js
var arr = [];
var len = oFullResponse.results.length;
for (var i = 0; i < len; i++) {
  arr.push({
    key: oFullResponse.results[i].label,
    sortable: true,
    resizeable: true,
  });
}
```
