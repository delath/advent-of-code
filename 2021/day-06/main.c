#include <stdio.h>
#define DAYS 256
int main(){
    long fishes[9]; 
    long temp;
    int i;
    long sum=0;
    char input;
    for (i = 0; i < 9; i++) {
        fishes[i]=0;
    }
    scanf("%c", &input);
    while (input != '\n') {
        if (input == '0')
            fishes[0]++;
        if (input == '1')
            fishes[1]++;
        if (input == '2')
            fishes[2]++;
        if (input == '3')
            fishes[3]++;
        if (input == '4')
            fishes[4]++;
        if (input == '5')
            fishes[5]++;
        if (input == '6')
            fishes[6]++;
        if (input == '7')
            fishes[7]++;
        if (input == '8')
            fishes[8]++;

        scanf("%c", &input);
    }
    for(i=0;i<DAYS;i++){
        temp=fishes[0];
        fishes[0]=fishes[1];
        fishes[1]=fishes[2];
        fishes[2]=fishes[3];
        fishes[3]=fishes[4];
        fishes[4]=fishes[5];
        fishes[5]=fishes[6];
        fishes[6]=fishes[7]+temp;
        fishes[7]=fishes[8];
        fishes[8]=temp;
    }
    for(i=0;i<9;i++){
        sum+=fishes[i];
    }
    printf("%ld\n", sum);
    return 0;
}
