<div align="center">
  <h1>Python</h1>
</div>

# Table of Contents

- [\*args and \*\*kwargs](https://book.pythontips.com/en/latest/args_and_kwargs.html)
- [Handling exceptions by using decorators](#handling-exceptions-by-using-decorators)

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
