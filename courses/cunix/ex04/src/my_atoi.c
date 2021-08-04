int my_atoi(const char *nptr){
    int sign = nptr[0] == '-' ? -1 : 1;
    int number = 0;

    for (int i = 0; nptr[i] != '\0'; i++){
        if(nptr[i] - '0' < 0 || nptr[i] - '0' > 9) continue;
        
        number = number * 10 + (nptr[i] - '0');
    }
    number = sign * number;
    return sign < 0 ? number % 2147483649 : number % 2147483648;
}
