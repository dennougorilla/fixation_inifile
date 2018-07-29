import os

path = './GameUserSettings.ini'

with open(path) as f:
    l = f.readlines()
    if 'bDisableMouseAcceleration=False\n' in l:
        index = l.index('bDisableMouseAcceleration=False\n')
        del l[index]
        l.insert(index, 'bDisableMouseAcceleration=True\n')

with open(path, mode='w') as f:
    f.writelines(l)
