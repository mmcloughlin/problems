#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>

uint64_t rev_naive(uint64_t x)
{
  uint64_t y;
  int i;
  for(i = 0; i < 64; i++) {
    y = (y << 1) | (x & 1);
    x >>= 1;
  }
  return y;
}

uint64_t rev_masks(uint64_t x)
{
  x = ((x & 0xffffffff00000000) >> 32) | ((x & 0x00000000ffffffff) << 32);
  x = ((x & 0xffff0000ffff0000) >> 16) | ((x & 0x0000ffff0000ffff) << 16);
  x = ((x & 0xff00ff00ff00ff00) >>  8) | ((x & 0x00ff00ff00ff00ff) <<  8);
  /* we could use an 8-bit lookup table at this point */
  x = ((x & 0xf0f0f0f0f0f0f0f0) >>  4) | ((x & 0x0f0f0f0f0f0f0f0f) <<  4);
  x = ((x & 0xcccccccccccccccc) >>  2) | ((x & 0x3333333333333333) <<  2);
  x = ((x & 0xaaaaaaaaaaaaaaaa) >>  1) | ((x & 0x5555555555555555) <<  1);
  return x;
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


uint64_t rev(uint64_t x)
{
  return rev_masks(x);
}

int main()
{
  while(1) {
    uint64_t x = rand64();
    printf("x=%" PRIx64 " rev(x)=%" PRIx64 " rev(rev(x))=%" PRIx64 "\n", x, rev(x), rev(rev(x)));
    if(rev_naive(x) != rev_masks(x)) {
      printf("FAIL\n");
    }
  }
}
