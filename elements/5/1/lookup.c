#include "parity.h"

#include <inttypes.h>

#include "byte_parity.h"

int parity(uint64_t x)
{
  int par = 0;
  for(int i = 0; i < 8; i++) {
    par ^= byte_parity[x&0xff];
    x >>= 8;
  }
  return par;
}
