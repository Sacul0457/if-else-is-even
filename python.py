from typing import Callable


def if_else_master_code(number: int, variable_name: str) -> Callable:
    """
    Generates python if-else code that allows you to check if a number is even or odd.
    
    Parameters
    ----------
    number: :class:`int`
        The number to check.
    variable_name: :class:`str`
        The variable used for the code generator when comparing.

    Returns
    --------
        :class:`Callable`
    """
    def custom_enum(iterable, start=0):
        count = start -1
        odd = True
        for item in iterable:
            count += 1
            if not odd:
                yield count, item,odd
                odd = True
            else:
                yield count, item, odd
                odd = False

    code = ["def __private_is_even_or_odd_code_generator():", f"  {variable_name} = {number}"]
    for i, item, value in custom_enum(range(number + 1)):
        ret = f'if {variable_name} == {i}: return {value}'
        print(ret)
        code.append(f"  {ret}")
    namespace = {}
    exec("\n".join(code), namespace)
    return namespace['__private_is_even_or_odd_code_generator']
