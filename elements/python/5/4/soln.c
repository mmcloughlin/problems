#include <stdio.h>
#include <inttypes.h>
#include <stdlib.h>
#include <time.h>

uint64_t closest_same_weight(uint64_t x)
{
  uint64_t flip;
  if(x&1) {
    flip = (x+1)&~x;
  } else {
    flip = (~x+1)&x;
  }
  flip |= flip >> 1;
  return x^flip;
}

void printb(uint64_t x)
{
  int i;
  for(i = 63; i >= 0; i--) {
    printf("%d", (int)((x>>i)&1));
  }
}

uint64_t rand64()
{
  uint64_t x;
  x = (x << 16) | (rand() & 0xffff);
  x = (x << 16) | (rand() & 0xffff);
  x = (x << 16) | (rand() & 0xffff);
  x = (x << 16) | (rand() & 0xffff);
  return x;
}

int main()
{
  srand(time(NULL));
  uint64_t x = rand64();
  printb(x); printf("\n");
  uint64_t y = closest_same_weight(x);
  printb(y); printf("\n");
  printb(x^y); printf("\n");
}
