#include <stdio.h>
// 99 compile error
// 11 compile error: list-initializer only usabel in init not assign
// struct A {
//   A(): a() { a = {0, 1}; }
//   int a[2];
// };

// 99 compile error
// 11 compile error: list-initializer should not use in ()
// struct A {
//   A(): a({0, 1}) {}
//   int a[2];
// };

// 99 compile error
// 11 OK
// struct A {
//   A(): a{0, 1} {}
//   int a[2];
// };

// 99 compile error
// 11 OK
// struct A {
//   A() {}
//   int a[2] = {0, 1};
// };

// 99 OK but ugly
struct A {
  A() { a[0] = 0; a[1] = 1; }
  int a[2];
};

int main() {
  A a;
  printf("should 1: %d\n", a.a[1]);
}
