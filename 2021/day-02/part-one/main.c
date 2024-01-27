#include <stdio.h>
int main(){
    char input;
    int depth=0;
    int position=0;
    int temp;
    int i;
    
    input=getchar();

    while(!feof(stdin)){
        switch(input){
            case 'f':
                for(i=0;i<7;i++){
                    getchar();
                }
                scanf("%d", &temp);
                position+=temp;
                getchar();
                break;
            case 'd':
                for(i=0;i<4;i++){
                    getchar();
                }
                scanf("%d", &temp);
                depth+=temp;
                getchar();
                break;
            case 'u':
                for(i=0;i<2;i++){
                    getchar();
                }
                scanf("%d", &temp);
                depth-=temp;
                getchar();
                break;
        }
        input=getchar();
    }

    printf("%d\n", depth*position);
    return 0;
}
