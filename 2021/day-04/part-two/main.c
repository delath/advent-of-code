#include <stdio.h>
#include <stdbool.h>
#define SIZE 5
#define LOT 100
int main(){
    int extracted[LOT];
    bool flag[SIZE][SIZE];
    int board[SIZE][SIZE];
    int output;
    int rounds=0;
    int temp;
    bool win=0;
    int sum=0;
    int i, j, k;
    for(i=0;i<LOT;i++){
        scanf("%d", &extracted[i]);
        if(i!=99){
            getchar();
        }
    }
    while(!feof(stdin)){
        for(i=0;i<SIZE;i++){
            for(j=0;j<SIZE;j++){
                flag[i][j]=0;
            }
        }
        for(i=0;i<SIZE;i++){
            for(j=0;j<SIZE;j++){
                scanf("%d", &board[i][j]);
            }
        }
        for(k=0;k<LOT;k++){
            for(i=0;i<SIZE;i++){
                for(j=0;j<SIZE;j++){
                    if(board[i][j]==extracted[k]){
                        flag[i][j]=1;
                    }
                }
            }
            for(i=0;i<SIZE;i++){
                for(j=0;j<SIZE;j++){
                    if(flag[i][j]==0){
                        break;
                    }else if(j==4 && flag[i][4]==1){
                        win=1;
                        temp=k;
                    }
                }
            }
            for(i=0;i<SIZE;i++){
                for(j=0;j<SIZE;j++){
                    if(flag[j][i]==0){
                        break;
                    }else if(j==4 && flag[4][i]==1){
                        win=1;
                        temp=k;
                    }
                }
            }
            if(win){
                break;   
            }
        }
        if(temp>rounds){
            rounds=temp;
            for(i=0;i<SIZE;i++){
                for(j=0;j<SIZE;j++){
                    if(flag[i][j]==0){
                        sum+=board[i][j];
                    }
                }
            }
            printf("sum %d e ultimo %d", sum, extracted[rounds]);
            output=sum*extracted[rounds];
        }
        sum=0;
        win=0;
    }
    printf("%d\n", output);
    return 0;
}
