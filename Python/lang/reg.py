import re
phone = "2004-959-559 # 这是一个电话号码"

# 删除注释
num = re.sub(r'#.*$', "", phone)
print("电话号码 : ", num)

# 移除非数字的内容
num = re.sub(r'\D', "", phone)
print("电话号码 : ", num)

filename = "\n           file = coc1111os2d-js-min.e0380.js           \n"
num = re.sub('cocos2d-js..*.js', "cocos2d-js-min.1111.js", filename)
print("file : ", num)

import os
import sys


def replace(filename, pattern, replace_str):
    with open(filename, 'r+', encoding="utf-8") as rf:
        lines = rf.readlines()
        with open('new', 'w+', encoding='utf-8') as wf:  #打开一个空文件
            for line in lines:
                content = re.sub(pattern, replace_str, line)
                wf.write(content)
    os.remove(filename)
    os.replace('new', filename)
    # with open(filename, 'rb+') as rf, open('new', 'wb+') as wf:
    #     lines = rf.readlines()
    #     for line in lines:
    #         content = re.sub(pattern, replace_str, line.decode("utf-8"))
    #         wf.write(content.encode("utf-8"))
    # os.remove(filename)
    # os.replace('new', filename)


curDir = os.path.dirname(sys.argv[0])
replace(os.path.join(curDir, "res", "config.xml"), r"main\..*\.js",
        "main.1111.js")
