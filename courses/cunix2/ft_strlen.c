#include "libft.h"

int ft_strlen(char* s){
    int size = 0;
    while (s[size] != 0) 
        size++;
    return size;
}
