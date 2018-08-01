import os
from time import sleep

path = os.getenv('LOCALAPPDATA')+'/FortniteGame/Saved/Config/WindowsClient/GameUserSettings.ini'

with open(path) as f:
    l = f.readlines()
    if 'bDisableMouseAcceleration=False\n' in l:
        print('setting...')
        index = l.index('bDisableMouseAcceleration=False\n')
        print(str(index)+': '+l[index]),
        del l[index]
        l.insert(index, 'bDisableMouseAcceleration=True\n')
        print(str(index)+': '+l[index]),
    else:
        print(l[24]),

with open(path, mode='w') as f:
        print('writing...')
        f.writelines(l)
        print('Done')
        sleep(10)
