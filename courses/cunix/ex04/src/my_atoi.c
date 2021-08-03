int my_atoi(const char *nptr){
    int sign = nptr[0] == '-' ? -1 : 1;
    int number = 0;
    int string_size = my_strlen(nptr);
    int digits_number = 0;

    for (int i = 0; i < string_size; i++){
        if( nptr[i] - '0' < 0 || nptr[i] - '0' > 9 ) continue;
        number += (nptr[i] - '0') * pow(10, digits_number);
        digits_number++;
    }
    return sign * number;
}

unsigned int my_strlen(char *str)
{
    int index = 0;
    while (str[index++] != '\0'){}

    return index - 1;
}

int pow(int base, int power){
    int result = 1;
    for (int i = 0; i < power; i++){
        result *= base;
    }
    return base;
}
