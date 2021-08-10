#include "libft.h"

char** ft_strsplit(char const *s, char c){
    int possible_word_num = 1;

    for (int i = 0; s[i] != 0; i++)
        if (s[i] == c)
            possible_word_num++;

    
    char** arr = (char**)malloc(possible_word_num * sizeof(char*));
    
    int s_idx = 0, size = 0, word_count = 0;

    for (int i = 0; s[i] != 0; i++){
        if(s[i] == c && size != 0){
            char* word = (char*)malloc(size*sizeof(char));
            for (int j = s_idx, k = 0; k < size; j++, k++){
                word[k] = s[j];
            }
            arr[word_count] = word;
            s_idx = i+1;
            size = 0;
            word_count++;
            continue;
        }
        if(s[i] == c && size == 0){
            s_idx++;
            continue;
        }
        size++;
    }
    
    if(size != 0){
        char* word = (char*)malloc(size*sizeof(char));
        for (int j = s_idx, k = 0; k < size; j++, k++){
            word[k] = s[j];
        }
        arr[word_count] = word;
    }


    return arr;
}
