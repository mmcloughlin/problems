CC=clang

all: lookup.test

%.test: test.c %.c
	$(CC) $^ -o $@

lookup.c: byte_parity.h

%.h: %.py
	python $< > $@
