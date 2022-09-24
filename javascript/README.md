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

### Make an array of object literals in a loop

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

<br>

### Remove array of objects from another array of objects

```js
// Here we will remove array b from array a
var a = [
  {
    id: "1",
    name: "a1",
  },
  {
    id: "2",
    name: "a2",
  },
  {
    id: "3",
    name: "a3",
  },
];

var b = [
  {
    id: "2",
    name: "a2",
  },
];

c = a.filter((x) => !b.filter((y) => y.id === x.id).length);
console.log(c);
```

<br>

### Create dynamic rows and columns for material ui data-grid

```js
const response_got_from_api = [
  { id: 1, name: "Foo", age: 20 },
  { id: 2, name: "Bar", age: 21 },
  { id: 3, name: "MyFoo", age: 20 },
];

const columns = Object.keys(response_got_from_api[0]);

const rows = response_got_from_api.map((item, index) => {
  const returnValue = {};
  returnValue["id"] = index + 1;
  columns.forEach((key) => {
    returnValue[key] = item[key];
  });
  return returnValue;
});

console.log(columns);
console.log(rows);
```

<br>

### Callback function

A `callback` function is a function that is passed as a parameter to another function and is performed when some action has been finished. The following is an example of a basic callback function that logs to the console once some actions have been performed.

```js
function modifyArray(arr, callback) {
  // do something to arr here
  arr.push(100);
  // then execute the callback function that was passed
  callback();
}

var arr = [1, 2, 3, 4, 5];

modifyArray(arr, function () {
  console.log("array has been modified", arr);
});
```

<br>

### Enque and deque using two stacks

```js
var inputStack = []; // First stack
var outputStack = []; // Second stack

// For enqueue, just push the item into the first stack
function enqueue(stackInput, item) {
  return stackInput.push(item);
}

function dequeue(stackInput, stackOutput) {
  // Reverse the stack such that the first element of the output stack is the
  // last element of the input stack. After that, pop the top of the output to
  // get the first element that was ever pushed into the input stack
  if (stackOutput.length <= 0) {
    while (stackInput.length > 0) {
      var elementToOutput = stackInput.pop();
      stackOutput.push(elementToOutput);
    }
  }

  return stackOutput.pop();
}
```
