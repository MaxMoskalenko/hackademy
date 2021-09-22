int my_strcmp(char *str1, char *str2){
    int index = 0;
    
    //Until null-symbol compares each symbol of strings
    while(str1[index] != '\0' && str2[index] != '\0'){
        if(str1[index] != str2[index])
            return (str1[index] > str2[index]) ? 1 : -1;
        index++;
    }

    //If strings have the equal sizes then the are the same, otherwise the longest is bigger 
    if(str1[index] == str2[index])
        return 0;
    if(str1[index] == '\0')
        return -1;
    else
        return 1;
}    
