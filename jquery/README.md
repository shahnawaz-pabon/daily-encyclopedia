<div align="center">
  <a href="https://jquery.com/">
    <img alt="jquery" src="../logos/jquery.png"/>
  </a>
  <h1>jQuery</h1>
</div>

## Miscellaneous codes

### Add specific class while removing from others' siblings. Selector can be `HTMLElement` object(performed by `this` keyword), element's `id` or `class`'s name.

```js
$(this).addClass("active").siblings().removeClass("active");
```

### Get the value of child element's input field by name when clicked on parent element.

```js
$(this).children('input[name="inputName"]').val();
```

### Make custom tooltip

**HTML**

```html
<h2 class="tool-tip">Hello tool-tip</h2>
```

**CSS**

```css
.tool-tip {
  cursor: pointer;
}
.tool-tip[number-tooltip]:hover:after {
  content: attr(number-tooltip);
  border: #c0c0c0 1px dotted;
  padding: 10px;
  display: block;
  background-color: #000000;
  color: #ffffff;
  max-width: 300px;
  text-align: center;
}
```

**Javascript**

```js
$(".tool-tip").attr("number-tooltip", "This is a custom tool-tip");
```
