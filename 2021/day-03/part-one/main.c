#include <stdio.h>
#define BITS 12
int main(){
    int ones[BITS];
    int zeros[BITS];
    int sol[BITS];
    int gamma=0, epsilon=0;
    int i, j;
    int temp;
    char bit;
    for(i=0;i<BITS;i++){
        ones[i]=0;
        zeros[i]=0;
    }
    scanf("%c", &bit);
    while(!feof(stdin)){
        for(i=0;i<BITS;i++){
            if(bit=='1')
                ones[i]++;
            if(bit=='0')
                zeros[i]++;
            scanf("%c", &bit);
        }
        scanf("%c", &bit);
    }
    for(i=0;i<BITS;i++){
        if(ones[i]>zeros[i]){
            sol[i]=1;
        }else{
            sol[i]=0;
        }
    }
    for(i=0;i<BITS;i++){
        if(i==0 && sol[BITS-1-i]==1){
            gamma++;
        }else if(i==0 && sol[BITS-1-i]==0){
            epsilon++;
        }
        if(i!=0){
            temp=2;
            for(j=1;j<i;j++){
                temp=temp*2;
            }
            if(sol[BITS-1-i]==1){
                gamma=gamma+temp;
            }else{
                epsilon=epsilon+temp;
            }
        }
    }
    printf("%d\n", gamma*epsilon);
    return 0;
}
