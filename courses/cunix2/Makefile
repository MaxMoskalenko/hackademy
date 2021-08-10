CC :=gcc
CFLAGS := -std=c99 -c
 
library: compile
	@ar -rcs libft.a $(wildcard ft_*.o) 
	@echo "libft.a Created/Updated"

compile:
	@$(CC) $< $(CFLAGS) $(wildcard ft_*.c)
	@echo "Compiled"


