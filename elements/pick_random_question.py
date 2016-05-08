import random

# number of questions per chapter
QUESTIONS_PER_CHAPTER = {
    5 : 15,
    6 : 26,
    7 : 14,
    8 : 19,
    9 : 15,
    10: 17,
    11: 10,
    12: 16,
    13: 18,
    14: 17,
    15: 22,
    16: 16,
    17: 23,
    18: 14,
    19: 12,
}

# generate full list of questions
questions = []
for ch, num_questions in QUESTIONS_PER_CHAPTER.items():
    for n in range(num_questions):
        q = '{chapter}.{number}'.format(chapter=ch, number=n)
        questions.append(q)

# pick a random one
picked = random.choice(questions)
print picked
