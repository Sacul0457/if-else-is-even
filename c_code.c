#include <stdio.h>
#include <stdlib.h>


void print_is_even(int n)
{
    // 1 for true, 0 for false
    printf("    int i = %d;\n", n);
    printf("    int is_even = 0;\n");
    int is_even = 1;
    int i = 0;
    if (n >= 0){
        for (; i < n + 1; i++) {
            printf("    if (i == %d){is_even = %d;}\n", i, is_even);
            is_even = !is_even;
        }
    }
    else {
        for (; i > n - 1; i--) {
            printf("    if (i == %d){is_even = %d;}\n", i, is_even);
            is_even = !is_even;
        }
    }
}


static inline int
parse_number(const char *str_number)
{
    char *end;
    long num = strtol(str_number, &end, 10);
    
    // conversion failed
    if (end == str_number || *end != '\0') {
        printf("Expected an integer value, got '%s' instead.\n", str_number);
        exit(-1);
    }

    if (num < INT_MIN || num > INT_MAX) {
        printf("Number '%s' too big/small for 'int' type, try a smaller/bigger number.\n", str_number);
        exit(-1);
    }

    return (int)num;
}


int main(int argc, char *argv[])
{
    if (argc != 2) {
        printf("Expected 1 argument, got %d instead.\n", argc - 2);
        printf("Example command: '.\\is_even 10'\n");
        return -1;
    }

    int num = parse_number(argv[1]);
    
    printf("int main()\n{\n");
    print_is_even(num);
    printf("    printf(\"%%s\", is_even? \"true\" : \"false\");\n");
    printf("}\n");
    return 0;
}
