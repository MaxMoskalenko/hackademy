#include "libft.h"

char *ft_strstr(const char *haystack, const char *needle){
    for (int i = 0; haystack[i] != 0; i++){
        for(int j = 0; needle[j] != 0; j++){
            if(haystack[i+j] != needle[j])
                break;
            if(needle[j+1] == 0)
                return (char *) &haystack[i];
    
        }
    }
    return NULL;
}
