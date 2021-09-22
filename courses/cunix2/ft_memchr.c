#include "libft.h"

void *ft_memchr(const void *s, int c, size_t n){
    char* cs = (char *) s;
    for (int i = 0; i < n; i++){
        if(cs[i] == c) 
            return &cs[i];
    }
    return NULL;
}
