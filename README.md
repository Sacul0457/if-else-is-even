# About
This is just a silly repo that consists of code generators in different languages to see if a number is even.

# Python Version
```python
def is_even(x: int) -> bool:
    # This will print the generated code and return a function of the generated code which you can call
    func = if_else_master_code(x, 'x', print_code=True, return_func=True)
    return func() # Whether x is even or not

print(is_even(10))
