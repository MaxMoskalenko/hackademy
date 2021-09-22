#include "libft.h"

char* ft_strjoin(char const *s1, char const *s2){
  
    int s1_size = 0, s2_size = 0;

    while(s1[s1_size] != 0 && s2[s2_size] != 0){
        
        if(s1[s1_size] != 0) 
            s1_size++;
        if(s2[s2_size] != 0)
            s2_size++;
    }

    char* conc = (char *)malloc( (s1_size + s2_size) * sizeof(char));
    
    if(conc == NULL)
        return NULL;

    int i = 0;
    for (; s1[i] != 0; i++){
        conc[i] = s1[i];
    }
    for (int j = 0; s2[j] != 0; j++, i++){
        conc[i] = s2[j];
    }
    conc[i] = 0;
    return conc;


}
