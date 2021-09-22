#include "libft.h"

void *ft_memmove(void *dst, const void *src, size_t len){
    char* cdst = dst;
    char* csrc = (char *) src;

    char* tmp = (char*) malloc(len * sizeof(char));

    for (int i = 0; i < len; i++){
        tmp[i] = csrc[i];
    }

    for (int i = 0; i < len; i++){
        cdst[i] = tmp[i];
    }
    free(tmp);
    return dst;
}
