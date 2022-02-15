<div align="center">
  <a href="https://www.javascript.com/">
    <img alt="javascript" src="../logos/javascript.gif"/ height="96" width="96">
  </a>
  <h1>JavaScript</h1>
</div>

## Miscellaneous codes

### Get the index of the object inside an array, matching a condition

```js
a = [
  { prop1: "abc", prop2: "qwe" },
  { prop1: "bnmb", prop2: "yutu" },
  { prop1: "zxvz", prop2: "qwrq" },
];

index = a.findIndex((x) => x.prop2 === "yutu");

console.log(index);
```

### Make unique array of objects from the duplicate one

```js
duplicate_array_of_objects = [
  { id: "1", prop2: "qwe" },
  { id: "2", prop2: "yutu" },
  { id: "1", prop2: "qwe" },
  { id: "2", prop2: "yutu" },
];

var unique_array_of_objects = duplicate_array_of_objects.filter(function (
  elem
) {
  return !this[elem.id] && (this[elem.id] = true);
},
Object.create(null));

console.log("Unique array");
console.log(unique_array_of_objects);
```

### Remove a specific object from the array of objects, matching a condition

```js
array = [
  { id: 1, prop2: "qwe" },
  { id: 2, prop2: "yutu" },
];

var len = array.length;

for (var i = 0; i < len; i++) {
  let index = array.findIndex((x) => x.id === 1);
  if (index > -1) {
    array.splice(index, 1); // 2nd parameter means remove one item only
  }
}

console.log(array);
```
