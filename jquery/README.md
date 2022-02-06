<div align="center">
  <a href="https://jquery.com/">
    <img alt="jquery" src="../logos/jquery.png"/>
  </a>
  <h1>jQuery</h1>
</div>

## Miscellaneous codes

Add specific class while removing from others' siblings. Selector can be `HTMLElement` object(performed by `this` keyword), element's `id` or `class`'s name.

```js
$(this).addClass("active").siblings().removeClass("active");
```

Get the value of child element's input field by name when clicked on parent element.

```js
$(this).children('input[name="inputName"]').val();
```

Loop throuh array of objects and add dynamic elements

```js
name_of_array_of_objects.forEach((el) => {
  $("#time-slot").append(
    `
            <div id="` +
      el["id"] +
      `" class="booking-time" style="cursor: pointer;">
                <input type="hidden" id="` +
      date +
      `" name="nameDate" value="` +
      date +
      `">
                <input type="hidden" id="` +
      year +
      `" name="nameYear" value="` +
      year +
      `">
                <input type="hidden" id="` +
      day +
      `" name="nameDay" value="` +
      day +
      `">
                <input type="hidden" name="slotTime" value="` +
      el["start_time"] +
      ` - ` +
      el["end_time"] +
      `">
                <p><i class="far fa-clock mr-2"></i>` +
      el["start_time"] +
      ` - ` +
      el["end_time"] +
      `</p>
            </div>
        `
  );
});
```
