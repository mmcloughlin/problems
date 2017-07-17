#!/bin/bash

problem=$1

dir=${problem/./\/}

mkdir -p $dir

touch $dir/README.md
cat > $dir/soln.py <<EOF
def f():
  pass


def test(trials=1000):
  for trial in xrange(trials):
    f()
  print 'pass'


def main():
  test()


if __name__ == '__main__':
  main()
EOF
