#ifndef _LIBFT_H_
#define _LIBFT_H_

#include <stddef.h>
#include <stdlib.h>

void ft_bzero(void *s, size_t n);
char *ft_strdup(const char *s1);
int ft_strncmp(const char *s1, const char *s2, size_t n);
char *ft_strchr(const char *s, int c);
char *ft_strrchr(const char *s, int c);
int ft_isalpha(int c);
int ft_isdigit(int c);
int ft_isascii(int c);
int ft_toupper(int c);
int ft_tolower(int c);

#endif

