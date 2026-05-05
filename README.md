# About
This is just a silly repo that consists of code generators in different languages to see if a number is even.

## Table of contents

### Interpreted languages:

1. [Python](#python-version)
2. [Perl](#perl-version)
3. [Lua](#lua-version)
4. [TypeScript](#typescript-version)

### Compiled languages:
1. [C](#c-version)
2. [Rust](#rust-version)
3. [Pascal](#pascal-version)
4. [Standard ML](#sml-version)
5. [FORTRAN](#fortran-version)

### Esoteric languages:
1. [HTMLish](#htmlish-version)

### Other languages:
1. [Docker](#docker-version)

## Python Version

### Generating the code

```sh
python python.py 10
```

If you want to run the generated script, run these commands instead:
- ```sh
  python python.py 10 > is_even.py
  ```
- ```sh
  python is_even.py
  ```

### Running the function directly

```py
x = 10
func = is_even(x, return_func=True) # This will return a callable than you can call
print(func()) # Call the function which prints 'True'
```

[(GO TO HOME)](#table-of-contents)

## Perl Version

> Make sure you have `perl` installed.

1. Run the script generator:

```sh
perl perl.pl - 5 > is_even.pl
```

> [!NOTE]
> The `-` symbol after the script is to pass the argument to the script.
> If you want to pass negative number. You can do something like this:
>
> ```sh
> perl perl.pl - -10 > is_even.pl
> ```

2. Run the generated script:

```sh
perl is_even.pl
# 0
```

[(GO TO HOME)](#table-of-contents)

## Lua Version

### Run the script generator:

<details>
<summary>Only generate the code</summary>

```sh
lua lua.lua 9
```

</details>

<details>
<summary>Redirect the output to lua to run it directly</summary>

```sh
lua lua.lua 9 | lua
# false
```

</details>

<details>
<summary>Redirect the output to a file to run it later:</summary>

```sh
lua lua.lua 9 > is_9_even.lua
```

Later on, if you forget if 9 is even or not, just run the script:

```sh
lua is_9_even.lua
# false
```

</details>

[(GO TO HOME)](#table-of-contents)

---

## C Version


1. Compile:

```sh
gcc c_code.c -o c_code
```

2. Usage:

```sh
# Unix
./c_code 10 > is_even.c

# Windows
.\c_code.exe 10 > is_even.c
```

3. Compile that newly baked program:

```sh 
gcc is_even.c -o is_even
```

4. Check the result, is 10 even?:

```sh
# Unix
./is_even

# Windows
.\is_even.exe

# true
```


[(GO TO HOME)](#table-of-contents)

## Rust Version

> Make sure you have `rustup` installed.

1. Compile:

```sh
rustc rust.rs
```

2. Usage:

```sh
# Unix
./rust 10 > is_even.rs

# Windows
.\rust.exe 10 > is_even.rs
```

3. Compile that newly baked program:

```sh 
rustc is_even.rs
```

4. Check the result, is 10 even?:

```sh
# Unix
./is_even

# Windows
.\is_even.exe

# true
```

[(GO TO HOME)](#table-of-contents)

## Pascal Version

> I use fpc to compile. Use whatever free pascal compiler you have.

1. Compile:

```sh
fpc pascal.pp
```

2. Usage:

```sh
# Unix:
./pascal 9 > isEven.pp

# Windows:
.\pascal.exe 9 > isEven.pp
```

3. Compile the new isEven program:

```sh
fpc isEven.pp
```

4. Run it:

```sh
# Unix:
./isEven

# Windows:
.\isEven.exe

# FALSE
```

[(GO TO HOME)](#table-of-contents)

---

## Standard ML Version

> Tested with both SML/NJ and MLTon, I will use MLTon for the sake of standalone binaries

1. Compile the generator

```sh
mlton sml.sml
```

2. Generate the "if even program"

```sh
./sml 100 > out.sml
```

3. Compile and run
```sh
mlton out.sml && ./out
```

Good luck have fun


[(GO TO HOME)](#table-of-contents)

---
## HTMLish Version

### Import function

```xml
<head>
  <import name="is even" module="if-else-is-even"/>
</head>
```

### Usage

```xml
<body>
  <comment>False</comment>
  <call function="is even">
    <integer>1</integer>
  </call>
  <output>
    <comment>True</comment>
    <call function="is even">
      <integer>2</integer>
    </call>
  </output>
</body>
```

### Execution

```bash
$ htmlish htmlish.html
```

[(GO TO HOME)](#table-of-contents)

---

## TypeScript Version

### Import

Copy **typescript.ts** into your repository and import it directly. In this example, we have name the file "ifElseIsEven.ts":

```ts
import { evaluate } from "ifElseIsEven.ts";
```

### Usage

```ts
import { evaluate } from "ifElseIsEvent.ts";
 
function fn() {
    const parityOf3 = evaluate(3); // "Odd"
    const parityOf8 = evaluate(8); // "Even"
    const output = `3 is ${parityOf3} and 8 is ${parityOf8}`;
    console.log(output); // Prints "3 is Odd and 8 is Even"
}
```

### Notes
> [!CAUTION]
> Calling `evaluate` for numbers with an absolute value > 100,000,000 will likely take at least 10 seconds to process and consume a **very high amount of RAM** (estimated between 40 and 80 bytes per number). The program is structured in a way that does not overflow the stack; it will continue running until it has exhausted your system's memory and may cause your runtime or operating system to crash.

- Supports negative inputs
- Does not support `Infinity` or `-Infinity` (an error message will display)
- Does not support `NaN` (an error message will display)
- The return type of `evaluate` is a string literal - "Even" or "Odd"

[(GO TO HOME)](#table-of-contents)

## FORTRAN Version

### Usage
1. Compile:
```sh
gfortran fortran.f95 -o a.out
```
2. Usage:
```sh
./a.out 10 > is_even.f95
```


[(GO TO HOME)](#table-of-contents)

## Docker Version

> Make sure you have `docker` installed.

1. Build to an image:

```sh
docker build -t is-even .
```

2. Run the image:

```sh
docker run --rm is-even ./is-even 4
```

[(GO TO HOME)](#table-of-contents)