#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "liballocate.h"

int main(int argc, char *argv[])
{
    if (argc < 2) {
        exit(1);
    }
    int size = atoi(argv[1]);
    printf("size = %d\n", size);
    FILE* fp = fopen("pid.txt", "w");
    printf("pid = %d\n", getpid());
    fclose(fp);
    for (int i = 0; ; i++) {
        allocate(size);
        printf("%d: returned from Go\n", i);
        sleep(1);
    }
    return 0;
}
