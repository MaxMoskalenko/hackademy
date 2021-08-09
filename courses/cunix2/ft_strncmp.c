#include "libft.h"

int ft_strncmp(const char *s1, const char *s2, size_t n){
    int i;
    for(i = 0; i < n; i++){
        if(s1[i] != s2[i])
            return s1[i] > s2[i] ? 1: -1;
        if(s1[i] == 0 || s2[i] == 0)
            break;
    }
    if(s1[i] == 0 && s2[i] == 0)
        return 0;
    if(s2[i] == 0)
        return 1;
    if(s1[i] == 0)
        return -1;
    return 0;
}
