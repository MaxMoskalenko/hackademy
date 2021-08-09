#include "libft.h"

char *ft_strdup(const char *s1){
    int size = 0;
    while(s1[size] != '\0') 
        size++;
    char* s2 = (char* )malloc(size*sizeof(char));
    for(int i = 0; i < size; i++) 
        s2[i] = s1[i];
    return s2;
}
