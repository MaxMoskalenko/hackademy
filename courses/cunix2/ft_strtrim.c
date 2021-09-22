#include "libft.h"

char* ft_strtrim(char const *s){
    int size = 0;
    while (s[size] != 0)
        size++;
    
    char* ns = (char*)malloc(size * sizeof(char));
    if(ns == NULL)
        return NULL;
    
    int delete_spaces = 1;
    int j = 0;
    for (int i = 0; s[i] != 0; i++){
        if((delete_spaces == 1) && (s[i] == ' ' || s[i] == '\n' || s[i] == '\t'))
            continue;
        
        delete_spaces = 0;
        ns[j] = s[i];
        j++;
    }

    ns[j] = 0;

    
    for (int i = j - 1; i >= 0; i--){
        if (ns[i] == ' ' || ns[i] == '\n' || ns[i] == '\t')
               ns[i] = 0;
        else
            break;
    }

    return ns;


} 
