import os

if not os.path.isfile(path_w):
    with open(path_w, mode='w') as f:
        f.write(s)

path = './test.txt'
with open(path) as f:
    while True:
        s_line = f.readline()
        if not s_line:

with open(path_w, mode='w') as f:
    f.write(s)

with open(path_w, mode='r+') as f:
    f.seek(3)
    f.write('-')

with open(path_w) as f:
    print(f.read())
# 123-56o
# ThreeFour
# Four
    
