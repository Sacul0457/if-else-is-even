#include <stdio.h>


void printfIsEven(int n){
    // 1 for true, 0 for false
    printf("int i = %d\n", n);
    for (int i = 0; i < n + 1; i++){
        if (i % 2 == 0){
            printf("if (i == %d){return 1;}\n", i);
        }
        else{
            printf("if(i == %d){return 0;}\n", i);
        }

    }
}


int main(){
    printfIsEven(10);
    return 0;
}
