#include "libft.h"

int ft_isalpha(int c){
    return (c >= 65 && c <= 90) || (c >= 97 && c <= 122) ? 1 : 0;
}
