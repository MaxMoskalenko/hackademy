
#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>
#include "test.h"

<<<<<<< HEAD
<<<<<<< HEAD
void   test_small()
=======
void  test_small()
>>>>>>> e8472f2 (ex03 solution)
=======
void  test_small()
>>>>>>> e8472f2fab3ed6218f023c04a4bd7c9d7487d22b
{
  char buf[32];

  assert(my_strcpy(buf, "") == buf);
  assert(strcmp(buf, "") == 0);

  assert(my_strcpy(buf, "123456789") == buf);
  assert(strcmp(buf, "123456789") == 0);
}

void           test_long()
{
  char        buf[32];
  const char  *ptr;

  ptr = "ooheijiewuudo8wuif5eeKoexie3aih";
  assert(my_strcpy(buf, ptr) == buf);
  assert(strcmp(buf, ptr) == 0);
}

void           test_verif()
{
  char        buf[32];
  const char  *ptr;

  ptr = "lohque7iebo7in3Einga0DumijaeR7";
  assert(strcpy(buf, ptr) == buf);
  assert(strcmp(buf, ptr) == 0);
}

int   main()
{
  test_small();
  test_long();
  test_verif();
  return(0);
}
