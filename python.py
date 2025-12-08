from typing import Callable


def if_else_master_code(
    number: int,
    variable_name: str | None = None,
    *,
    print_code: bool = True,
    return_func: bool = False,
) -> Callable | None:
    """
    Generates python if-else code that allows you to check if a number is even or odd.

    Parameters
    ----------
    number: :class:`int`
        The number to check.
    variable_name: :class:`str` | `None`
        The variable used for the code generator when comparing. If `None`, the variable name `x` is used.
    print_code: :class:`bool`
        Whether to print the generated code. Defaults to `True`.
    return_func: :class:`bool`
        Whether to return a :class:`Callable` of the generated code. Defaults to `False`.

    Raises
    -------
    ValueError
        When both `print_code` and `return_func` is set to `False` as the function would esentially do nothing.

    Returns
    --------
        The :class:`Callable` which you can call and returns a :class:`bool` if `return_func` is set to `True`, else `None`
    """

    def custom_enum(iterable, start=0):
        count = start - 1
        odd = True
        for item in iterable:
            count += 1
            if not odd:
                yield count, item, odd
                odd = True
            else:
                yield count, item, odd
                odd = False

    if not print_code and not return_func:
        raise ValueError(
            "You have to either print or return a function, else it esentially does nothing."
        )

    variable_name = variable_name or "x"
    code = [
        "def __private_is_even_or_odd_code_generator():",
        f"  {variable_name} = {number}",
    ]

    for i, item, value in custom_enum(range(number + 1)):
        ret = f"if {variable_name} == {i}: return {value}"
        code.append(f"  {ret}")
    if return_func:
        namespace = {}
        exec("\n".join(code), namespace)
        return namespace["__private_is_even_or_odd_code_generator"]
    if print_code:
        print("\n".join(code))


x = 10
if_else_master_code(x, "k")
