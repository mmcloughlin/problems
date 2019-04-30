#!/usr/bin/env python3

import random
import sys
import os.path

QUESTIONS_DIR = './questions'
LANGUAGE = 'go'


def read_questions_file(name):
    questions_file = os.path.join(QUESTIONS_DIR, name)
    with open(questions_file) as f:
        return filter(None, [q.split()[0] for q in f])


# arguments
sources = sys.argv[1:] if len(sys.argv) > 1 else ['all']
sources.append('-' + LANGUAGE)

# form set of possible questions
questions = set()
for source in sources:
    if source[0] == '+':
        questions |= set(read_questions_file(source[1:]))
    elif source[0] == '-':
        questions -= set(read_questions_file(source[1:]))
    else:
        questions |= set(read_questions_file(source))

# pick random one
print(random.choice(list(questions)))
