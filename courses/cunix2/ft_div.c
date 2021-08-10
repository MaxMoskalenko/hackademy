#include "libft.h"

div_t ft_div(int numer, int denom){
    div_t result = {numer / denom, numer % denom};
    return result;
}
