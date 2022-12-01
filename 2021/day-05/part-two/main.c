#include <stdio.h>
#include <stdlib.h>
#define MAX 1000
int main(){
    int x1, y1;
    int x2, y2;
    int overlap=0;
    int temp;
    int i, j;
    int** field=(int**)malloc(MAX*sizeof(int*));
    for (i=0;i<MAX;i++){
        field[i]=(int*)malloc(MAX*sizeof(int));
    }
    for(i=0;i<MAX;i++){
        for(j=0;j<MAX;j++){
            field[i][j]=0;
        }
    }
    while(!feof(stdin)){
        scanf("%d,%d -> %d,%d", &x1, &y1, &x2, &y2);
        if(x1==x2){
            if(y1>y2){
                temp=y1-y2;
                temp++;
                for(i=0;i<temp;i++){
                    field[x1][y2+i]++;
                    if(field[x1][y2+i]==2)
                        overlap++;
                }
            }else{
                temp=y2-y1;
                temp++;
                for(i=0;i<temp;i++){
                    field[x1][y1+i]++;
                    if(field[x1][y1+i]==2)
                        overlap++;
                }
            }
        }else if(y1==y2){
            if(x1>x2){
                temp=x1-x2;
                temp++;
                for(i=0;i<temp;i++){
                    field[x2+i][y1]++;
                    if(field[x2+i][y1]==2)
                        overlap++;
                }
            }else{
                temp=x2-x1;
                temp++;
                for(i=0;i<temp;i++){
                    field[x1+i][y1]++;
                    if(field[x1+i][y1]==2)
                        overlap++;
                }
            }
        }else{
            if(x1>x2 && y1>y2){
                temp=x1-x2;
                temp++;
                for(i=0;i<temp;i++){
                    field[x2+i][y2+i]++;
                    if(field[x2+i][y2+i]==2)
                        overlap++;
                }
            }else if(x1>x2 && y1<y2){
                temp=x1-x2;
                temp++;
                for(i=0;i<temp;i++){
                    field[x2+i][y2-i]++;
                    if(field[x2+i][y2-i]==2)
                        overlap++;
                }
            }else if(x1<x2 && y1<y2){
                temp=x2-x1;
                temp++;
                for(i=0;i<temp;i++){
                    field[x1+i][y1+i]++;
                    if(field[x1+i][y1+i]==2)
                        overlap++;
                }
            }else if(x1<x2 && y1>y2){
                temp=x2-x1;
                temp++;
                for(i=0;i<temp;i++){
                    field[x1+i][y1-i]++;
                    if(field[x1+i][y1-i]==2)
                        overlap++;
                }
            }
        }
    }
    printf("%d\n", overlap);
    return 0;
}
