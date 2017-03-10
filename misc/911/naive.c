// Experiment to see how long the naive approach takes in a performant
// language.
//
// $ clang -O3 naive.c
// $ time ./a.out
//
// 8927139961
//
// real	17m33.587s
// user	17m19.576s
// sys	0m3.984s

#include <stdio.h>
#include <inttypes.h>
#include <math.h>
#include <string.h>

int64_t count_non_emergency_numbers(int n) {
  int64_t count = 0;
  int64_t i;
  char s[100];
  for(i = powl(10, n-1); i < powl(10, n); i++) {
    sprintf(s, "%" PRIi64, i);
    count += (strstr(s, "911") == NULL);
  }
  return count;
}

int main(int argc, char *argv[]) {
  int64_t count = count_non_emergency_numbers(10);
  printf("%" PRIi64 "\n", count);
}
