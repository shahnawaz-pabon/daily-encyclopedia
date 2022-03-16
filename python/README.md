<div align="center">
  <a href="https://ubuntu.com/">
    <img alt="ubuntu" src="../logos/python.png"/>
  </a>
  <h1>Python</h1>
</div>

# Table of Contents

- [\*args and \*\*kwargs](https://book.pythontips.com/en/latest/args_and_kwargs.html)
- [Handling exceptions by using decorators](#handling-exceptions-by-using-decorators)
- [Keep specific keys from a list of dictionaries](#keep-specific-keys-from-a-list-of-dictionaries)
- [Convert a query string to valid json](#convert-a-query-string-to-valid-json)
- [Custom date formatter](#custom-date-formatter)

<br>

## Handling exceptions by using decorators

[Decorators in Python](https://www.geeksforgeeks.org/decorators-in-python/)

Here is an example without using decorator

```python
def mean(a,b):
    try:
        print((a+b)/2)
    except TypeError:
        print("wrong data types. enter numeric")


def square(sq):
    try:
        print(sq*sq)
    except TypeError:
        print("wrong data types. enter numeric")


def divide(l,b):
    try:
        print(b/l)
    except TypeError:
                print("wrong data types. enter numeric")
mean(4,5)
square(21)
divide(8,4)
divide("two","one")
```

So with the use of decorators, the above code can be more presentable and understandable:

```python
def Error_Handler(func):
	def Inner_Function(*args, **kwargs):
		try:
			func(*args, **kwargs)
		except TypeError:
			print(f"{func.__name__} wrong data types. enter numeric")
	return Inner_Function
@Error_Handler
def Mean(a,b):
		print((a+b)/2)

@Error_Handler
def Square(sq):
		print(sq*sq)

@Error_Handler
def Divide(l,b):
		print(b/l)

Mean(4,5)
Square(14)
Divide(8,4)
Square("three")
Divide("two","one")
Mean("six","five")
```

<br>

## Keep specific keys from a list of dictionaries

If we have a dictionary like this:

```python
data = {
    "list_of_dictionaries": [
        {
            "title":"Title",
            "description":"description",
            "name":"Pabon",
            "another_list":[
                {
                    "key1":1,
                },
                {
                    "key2":2,
                },
                {
                    "key3":3,
                }
            ]
        }
    ],
    "another_key": "value",
}
```

We want specific keys(e.g: name, another_list) keeping into the dictionary, then we can do this:

```python

keys_to_preserve = ["name", "another_list"]

res = [
    {key: item[key] for key in keys_to_preserve}
    for item in data["list_of_dictionaries"]
]

data["list_of_dictionaries"] = res

print(data)
```

<br>

## Convert a query string to valid json

I faced this issue in `django`. The issue was, I was receiving `application/x-www-form-urlencoded` data from the client and I had to sent this to another third party API. Also I had some processing with this data and that's why I had to convert it to `json` format. So I sent the request.body directly to the third party API and do these for my processing.

```python
from urllib import parse
import requests

# Parse a query string given as a string argument (data of type application/x-www-form-urlencoded). Data are returned as a list of name, value pairs.
parsed_data = parse.parse_qs(request.body.decode("utf-8"))

# dictionary comprehension to convert it to a single item
parsed_data_to_single_item = {
    k: v for k, vs in parsed_data.items() for v in vs
}
print(parsed_data_to_single_item)


# For third party API call
headers = {
    "Content-Type": "application/x-www-form-urlencoded",
}
url = "my_url"
response = requests.post(url, data=request.body, headers=headers)
print(response.json())

```

## Custom date formatter

```python
def get_month_year(custom_date):
    return custom_date.strftime("%d %B, %Y")


def get_am_pm(custom_date):
    # date_time = datetime.strftime(custom_date)
    return custom_date.strftime("%I:%M %p")


def get_weekday(custom_date):
    return custom_date.strftime("%A")


def get_abbreviated_weekday(custom_date):
    return custom_date.strftime("%a")
```
