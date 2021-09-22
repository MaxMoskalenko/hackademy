#include "libft.h"

char* ft_strsub(char const *s, unsigned int start, size_t len){
    char* sub = (char*)malloc(len * sizeof(char));
    
    if(sub == NULL)
        return NULL;
    int i = 0;
    for (; i < len; i++){
        sub[i] = s[start + i];
    }
    sub[i] = '\0';

    return sub;
}
