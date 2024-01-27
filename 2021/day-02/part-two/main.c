#include <stdio.h>
int main(){
    char input;
    long aim=0;
    long position=0;
    long depth=0;
    long temp;
    long sol;
    long i;
    
    input=getchar();

    while(!feof(stdin)){
        switch(input){
            case 'f':
                for(i=0;i<7;i++){
                    getchar();
                }
                scanf("%ld", &temp);
                position+=temp;
                depth+=(aim*temp);
                getchar();
                break;
            case 'd':
                for(i=0;i<4;i++){
                    getchar();
                }
                scanf("%ld", &temp);
                aim+=temp;
                getchar();
                break;
            case 'u':
                for(i=0;i<2;i++){
                    getchar();
                }
                scanf("%ld", &temp);
                aim-=temp;
                getchar();
                break;
        }
        input=getchar();
    }

    printf("%ld\n", depth*position);
    return 0;
}
