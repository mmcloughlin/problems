#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

// reverse reverses the first n characters of s in-place.
void reverse(char *s, int n)
{
  int i;
  for(i = 0; i < n/2; i++) {
    int j = n-1-i;
    char t = s[i];
    s[i] = s[j];
    s[j] = t;
  }
}

// reverse_each_word splits s on whitespace and reverses the characters of
// each word.
void reverse_each_word(char *s)
{
  char *w = s;
  int m = 0;
  while(1) {
    /* collect word */
    while(*s && !isspace(*s)) {
      m++;
      s++;
    }

    if(m > 1) {
      reverse(w, m);
    }

    if(!*s) break;

    /* skip whitespace */
    while(*s && isspace(*s)) {
      s++;
    }

    if(!*s) break;

    w = s;
    m = 0;
  }
}

// reverse_words reverse the word order in s. Operates in place.
void reverse_words(char *s)
{
  int n = strlen(s);
  reverse(s, n);
  reverse_each_word(s);
}

int main()
{
  char *s = (char *)malloc(16*sizeof(char));
  strcpy(s, "Alice likes Bob");

  printf("%s\n", s);
  reverse_words(s);
  printf("%s\n", s);

  free(s);
}
