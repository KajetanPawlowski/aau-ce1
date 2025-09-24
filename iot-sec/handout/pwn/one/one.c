#include <stdio.h>
#include <stdlib.h>
#include <string.h>



int main(int argc, char const *argv[])
{

    if (argc < 2) {
        printf("Usage: %s <username>\n", argv[0]);
        return 1;
    }

    puts("Admin users have privileged access. Contact your admin to receive the role");

    long admin = 0x4141414141414141;
    char username[32];
    strcpy(username, argv[1]);

    if (admin == 0x3131313131313131) {
        puts("Extra privileges added!\n");
        puts("aau{buff3r_0v3rfl0ws_4r3_fun}\n");
    }

    return 0;
}
