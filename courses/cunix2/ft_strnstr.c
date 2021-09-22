#include "libft.h"

char *ft_strnstr(const char *haystack, const char *needle, size_t len){
    for (int i = 0; haystack[i] != 0 && i < len; i++){
        for (int j = 0; needle[j] != 0 && i + j < len; j++){
            if(haystack[i+j] != needle[j])
                 break;
             if(needle[j+1] == 0)
                 return (char *) &haystack[i];
        }
    }
    return NULL;
}
