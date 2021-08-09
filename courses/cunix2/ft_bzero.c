#include "libft.h"

void ft_bzero(void *s, size_t n){
    char* cs = s;
    for(int i=0; i < n; i++){
        cs[i] = 0;
    }

}
