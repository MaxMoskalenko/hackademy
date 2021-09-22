#include "libft.h"

void *ft_memccpy(void *restrict dst, const void *restrict src, int c, size_t n){
    char* cdst = dst;
    char* csrc = (char*) src;
    for (int i = 0; i < n; i++){ 
        cdst[i] = csrc[i];
        
        if (csrc[i] == c) {
             return &csrc[i+1];
        }
    }
    return NULL;

}
