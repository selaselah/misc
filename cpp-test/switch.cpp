#include <stdio.h>

void try_switch(int i) {
  switch (i) {
   case 1:
    printf("1\n");
   case 2:
    printf("2\n");
   case 3:
    printf("3\n");
   default:
    printf("default=%d\n", i);
   //  break;
   // default in the middle need break
   case 4:
    printf("4\n");
    break;
   case 5:
    printf("5\n");
   case 6:
    printf("6\n");
   // case 6: 
   //  printf("6\n");
   // double case 6 cause compile error
  }
}

int main() {
  for (int i = 0; i != 8; ++i) {
    printf("try_switch(%d)\n", i);
    try_switch(i);
  }
  return 0;
}

