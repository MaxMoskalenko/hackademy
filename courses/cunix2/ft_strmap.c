#include "libft.h"

char* ft_strmap(char const *s, char (*f)(char)){
    int size = 0;
    while(s[size] != 0) 
        size++;
    char* s2 = (char*) malloc(size * sizeof(char));
    
    int i = 0;
    for(; s[i] != 0; i++){
        s2[i] = f(s[i]);
    }
    s2[i] = 0;
    return s2;

}
