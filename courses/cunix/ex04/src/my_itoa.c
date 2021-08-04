#include <stdlib.h>

char* my_itoa(int nmb){
    int string_size = 0;
    long digit = 1;
    int nmb2 = nmb;

    if (nmb == 0){
        static char result[2];
        result[0] = '0';
        result[1] = '\0';
        return result;
    }

    for (int i = 0; nmb2 != 0; i++){
         nmb2 /= 10;
         digit *= 10;
         string_size++;
    }

    static char result[12];

    digit /= 10; 

    int i = 0;
    if(nmb < 0){
        i = 1;
        result[0] = '-';
        nmb = -nmb;
        string_size++;
    }
    for(; i < string_size; i++){
        result[i] = (nmb / digit) + '0';
        nmb %= digit;
        digit /= 10;
    }
    
    result[i] = '\0';

    return result;
}

