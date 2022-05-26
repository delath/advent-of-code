#include <stdio.h>
int main(){
    int values[3];
    int newSum=0;
    int oldSum;
    int cont=0;
    int flag=0;
    for(int i=0;i<3;i++){
        scanf("%d", &values[i]);
        newSum+=values[i];
    }
    while(!feof(stdin)){
        oldSum=newSum;
        switch(flag){
            case 0:
                newSum-=values[0];
                scanf("%d", &values[0]);
                newSum+=values[0];
                flag++;
                break;
            case 1:
                newSum-=values[1];
                scanf("%d", &values[1]);
                newSum+=values[1];
                flag++;
                break;
            case 2:
                newSum-=values[2];
                scanf("%d", &values[2]);
                newSum+=values[2];
                flag=0;
                break;
        }
        if(newSum>oldSum){
            cont++;
        }
    }
    printf("%d increases\n", cont);
    return 0;
}