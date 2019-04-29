CC=clang

all: lookup.test bitmasks.test

%.test: test.c %.c
	$(CC) $^ -o $@

lookup.c: byte_parity.h

%.h: %.py
	python $< > $@

clean:
	$(RM) *.test byte_parity.h

.PHONY: all clean
