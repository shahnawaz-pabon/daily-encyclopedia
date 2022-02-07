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
