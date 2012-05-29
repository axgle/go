#include <stdlib.h>
#include <stdio.h>

extern char _binary_hello_txt_start[];

int main (int argc, char *argv[])
{
    char *p;
    p = _binary_hello_txt_start;
    printf("%s",p);
    return 0;
}