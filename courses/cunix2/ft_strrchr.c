#include "libft.h"


char *ft_strrchr(const char *s, int c){
   int size = 0;
   while(s[size] != 0) 
       size++;

   for (int i = size-1; i>=0; i--){
        if (s[i] == c)
        {
            return (char *)&s[i];
        }
   }
   return NULL;
}
