#include "libft.h"

void *ft_memset(void *b, int c, size_t len){

    char* cb = b;
    for (int i = 0; i < len; i++){
        cb[i] = (char) c;
    }
    return cb;
}
