#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <time.h>
#include <assert.h>

#include "parity.h"

#define TRIALS (1<<20)

uint64_t rand64()
{
  uint64_t x = 0;
  for(int i = 0; i < 8; i++) {
    x <<= 8;
    x |= rand()&0xff;
  }
  return x;
}

int parity_naive(uint64_t x)
{
  uint64_t par = 0;
  for( ; x; x >>= 1) {
    par ^= x&1;
  }
  return (int)par;
}

int main()
{
  srand(time(NULL));
  for(uint64_t t = 0; t < TRIALS; t++) {
    uint64_t x = rand64();
    int expect = parity_naive(x);
    int got = parity(x);
    assert(expect == got);
  }
  printf("pass\n");
  return 0;
}
