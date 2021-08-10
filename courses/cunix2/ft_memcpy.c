#include "libft.h"

void *ft_memcpy(void *restrict dst, const void *restrict src, size_t n){
    char* csrc = (char*)src;
    char* cdst = dst;
    for(int i = 0; i < n; i++){
        cdst[i] = (char)csrc[i];
    }
    return dst;
}
