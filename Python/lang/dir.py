import os


# 遍历文件夹
def walkFile(file):
    for root, dirs, files in os.walk(file):

        # root 表示当前正在访问的文件夹路径
        # dirs 表示该文件夹下的子目录名list
        # files 表示该文件夹下的文件list

        # 遍历文件
        for f in files:
            print(os.path.join(root, f))

        # 遍历所有的文件夹
        for d in dirs:
            print(os.path.join(root, d))


def list_dir(file_dir):
    '''
        通过 listdir 得到的是仅当前路径下的文件名，不包括子目录中的文件，如果需要得到所有文件需要递归
    '''
    dir_list = os.listdir(file_dir)
    for cur_file in dir_list:
        # 获取文件的绝对路径
        path = os.path.join(file_dir, cur_file)
        if os.path.isfile(path):  # 判断是否是文件还是目录需要用绝对路径
            print("{0} : is file!".format(cur_file))
        if os.path.isdir(path):
            print("{0} : is dir!".format(cur_file))
            # list_dir(path) # 递归子目录


def main():
    list_dir("D:\\temp")
    # value_channels = [1,1,1,1]
    # print("共打包%d个渠道包"%(len(value_channels)))


if __name__ == '__main__':
    main()
