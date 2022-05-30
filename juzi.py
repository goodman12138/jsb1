#coding=utf-8
class Book(object):
    def __init__(self, name, author, count, place):
        # 书名
        self.name = name
        # 作者
        self.author = author
        # 数量
        self.count = count
        # 位置
        self.place = place

    def __str__(self):
        if self.count == 0:
            count = 'OUT'
        else:
            count ='IN'
        return '%s' %(self.name)


class identity(object):
    # 定义存储用户名和密码的字典
    StudentDict = {}

    # 定义静态方法选择身份
    @staticmethod
    def Choice():
        print ('*' * 50)
        print ('\t\t\t\t欢迎使用图书管理系统')
        answer = input('  请选择您的身份:[1] 管理员 [2] 学生帐号 [3] 注册帐号\n')
        if answer == '1':
            # 如果选择是1，调用管理员的验证函数
            identity.ManagerIdentity()
        elif answer == '3':
            # 如果选择是3，选择注册一个新的学生帐号，并调用验证函数登陆
            name = input('注册新用户———请输入您的姓名：')
            cd = True
            while cd:
                cd = False
                # code = maskpass.askpass(prompt="code:", mask="*")
                code = input("请输入密码：")
                if (len(code) < 6):
                    print('密码小于六位，请重新注册...')
                    cd = True
                    continue
                elif (code.isalpha()):
                    print('密码为纯字母，请重新注册...')
                    cd = True
                    continue
                # password1=int(code)
                elif (code.isdigit()):
                    print('密码为纯数字，请重新注册...')
                    cd = True
                    continue
            identity.StudentDict[name] = code
            for item in identity.StudentDict.items():
                for i in range(len(item)):
                    str1 = item[i]
                    print(str1, end=' ')
                    with open(r'./通讯录.txt', 'a') as f:
                        f.write(str1)
                        f.write('\r\t')

            print ('***您已成功注册图书管理系统学生用户，请登陆：***')
            identity.StudentIdentity()
        elif answer =='2':
            # 如果选择是2，调用学生用户的验证函数
            identity.StudentIdentity()
        else:
            print ('***您的选择错误，不能登陆！***')

    # 定义类方法验证管理员信息，管理员的密码是123456，身份唯一
    @classmethod
    def ManagerIdentity(cls):
        # 管理员的验证函数
        passwd1 = input('请输入您的密码：')
        print ('')
        print('*' * 50)
        if passwd1 == '123456':
            print ('\033[5;31;2m\t\t您的身份是管理员，欢迎进入图书管理系统\033[0m')
            Menu.ManagerMenu()
        else :
            print ('您的密码错误！请重新输入：')
            cls.ManagerIdentity()

    # 登录
    @classmethod
    def StudentIdentity(cls):
        username =input('请输入您的帐号：')
        passwd2 = input('请输入您的密码：')
        print (' ')
        print ('*' * 50)
        for key in cls.StudentDict:
            if (key == username and cls.StudentDict[key] == passwd2):
                print('\033[5;31;2m\t\t您的身份是学生，欢迎您进入图书管理系统\) 033[0m')
                Menu.StudentMenu()
                return
        print ('您的用户名或身份不正确！请重新输入：')
        cls.StudentIdentity()


class Menu(object):
    # 定义菜单类，学生和管理员有不同的选择
    @staticmethod
    def StudentMenu():
        print ('')
        print ('[1]查看所有藏书\n')
        print ('[2]借阅图书\n')
        print ('[3]归还图书\n')
        print ('[0]退出系统\n')
        while True:
            studentchoice = input('请选择您要进行的操作：')
            if studentchoice == '1':
                # 学生查看所有图书部分
                choice.StudentSee()
            elif studentchoice =='2':
                # 学生借书部分
                choice.StudentBorrow()
            elif studentchoice == '3':
                choice.StudentGive()
            elif studentchoice =='0':
                # 学生退出菜单部分
                print ('')
                print ('\033[5;31;2m***欢迎再次使用图书管理系统***\033[0m')
                break
            else:
                print ('您的输入不正确，请输入[1]查看所有藏书[2]借书[3]还书[0]退出系统\n')


    @staticmethod
    def ManagerMenu():
        # 定义管理员选择菜单函数
        print ('')
        print ('[1]查看所有藏书\n')
        print ('[2]查询图书\n')
        print ('[3]增加图书\n')
        print("[4]进入用户管理页面\n")
        print ('[0]退出系统\n')
        while True:
            managerchoice = input('请选择您要进行的操作：')
            if managerchoice == '1':
                # 管理员查看图书部分
                choice.StudentSee()
            elif managerchoice == '2':
                # 管理员查询图书部分
                choice.ManagerFind()
            elif managerchoice == '3':
                # 管理员增加图书部分
                choice.ManagerAdd()
            elif managerchoice == '4':
                Menu.guanli()
            elif managerchoice == '0':
                # 退出菜单部分
                print('')
                print ('\033[5;31;2m***欢迎再次使用图书管理系统***\033[0m')
                break
            else:
                print ('您的输入不正确，请输入[1]查看所有图书[2]查询图书[3]增加图书[0]退出系统\n')

    @staticmethod
    def guanli():
        while True:
            print("1:增加")
            print("2:修改")
            print("3:删除")
            print("0:退出")
            managerchoice = input('请选择您要进行的操作：')
            if managerchoice == '1':
                # 增加用户
                namm=input("请输入添加的用户名：")
                passwd=input("请输入添加的密码：")
                identity.StudentDict[namm]=passwd
            elif managerchoice == '2':
                # 修改用户
                cd=True
                while cd:
                    cd=False
                    d=input("请输入您要更改的用户名：")
                    for ass  in identity.StudentDict :
                        if d==ass:
                            s=input("请重新更改密码：")
                            identity.StudentDict[d]='s'
                            print("更改成功！！！")
                            exit()
                        else:
                            cd=True
                            print("没有此用户")
            elif managerchoice == '3':
                # 删除用户
                cd=True
                while cd:
                    cd=False
                    s=input("请输入要删除的用户名：")
                    for d in identity.StudentDict :
                        if d==s :
                            identity.StudentDict.pop(d)
                            print('删除成功！！')
                            exit()
                        else:
                            cd=True
                            print("该用户不存在请重新输入！！")
            elif managerchoice == '0':
                # 退出菜单部分
                print('')
                print ('\033[5;31;2m***欢迎再次使用图书管理系统***\033[0m')
                break


class choice(object):

    # 增加图书
    @staticmethod
    def ManagerAdd():
        name = input('请输入书籍的名称：')
        author = input('请输入书的作者：')
        count = input('请输入书的数量：')
        place = input('请输入书所在的楼层：')

        BookList.append(Book(name, author, count, place))
        print ('%s 添加成功！' % name)
        print ('')
        print ('图书馆所有的书籍有：')
        # 调用StudentSee方法，显示图书馆中所有的书
        choice.StudentSee()

    # 查询图书
    @staticmethod
    def ManagerFind():
        name = input('请输入要查询的书名：')
        for book in BookList:
            if name == book.name :
                print ('《%s》 作者：%s  数量：%s  楼层：%s！' % (name,book.author,book.count,book.place))
                return book
        else:
            print( '《%s》没有找到！' % name)
            return None
    # 遍历图书馆中的所有书
    @staticmethod
    def StudentSee():
        print ('图书馆中所有的书有：')
        for book in BookList:
            print ('《%s》 作者：%s \t数量：%s \t楼层：%s '% (book,book.author,book.count,book.place))

    # 学生借书部分
    @staticmethod
    def StudentBorrow():
        name = input('请输入您想借阅的书名：')
        BookResult = choice.ManagerFind()
        if BookResult:
            if BookResult.count == 0 or BookResult.count < 0:
                print ('***该书已经被借出，请稍后再来！***')
            else:
                BookResult.count > 0
                print('您已成功借到《%s》...' % (name))
                BookResult.count  -= 1

        else:
            print ('抱歉，图书馆暂时没有《%s》这本书...' % name)

    # 学生还书部分
    @staticmethod
    def StudentGive():
        name = input('请输入您要归还的书名：')
        BookResult1 = choice.ManagerFind()
        BookResult1.count += 1
        print ('您的书已经归还成功,欢迎下次使用...')

# 类外主程序部分
# 定义一个存放书籍信息的列表
BookList = []

# 实例化初始书籍信息，存放在列表中
book1 = BookList.append(Book('C语言','盖伦',5,'1Floor'))
book2 = BookList.append(Book('python','金铲铲',2, '2Floor'))
book3 = BookList.append(Book('数据库','嘉文',1,'3Floor'))

# 调用identity类中的静态方法
identity.Choice()