#include "libft.h"

int ft_memcmp(const void *s1, const void *s2, size_t n){
    char* cs1 = (char *) s1;
    char* cs2 = (char *) s2;

    for (int i = 0; i < n; i++){
        if (cs1[i] != cs2[i])
            return cs1[i] > cs2[i] ? 1 : -1;
    }
    return 0;
}
