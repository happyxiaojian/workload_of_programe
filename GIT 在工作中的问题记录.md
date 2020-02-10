# GIT 在工作中的问题记录

## git修改远程仓库地址

方法有三种：

1. 修改命令
   git remote origin set-url [url]

2. 先删后加
   git remote rm origin
   git remote add origin [url]

git clone 克隆远程的代码到本地
git remote -v  查看所有远程仓库
git remote branch_name 查看指定分支branch_name远程仓库地址
git remote set-url origin  通过命令直接修改远程地址

#### 设置一个名字为 upstream 的上游地址，也就是我们项目主仓库的地址
> git remote add upstream git@github.com:fe/github-flow.git

#### 把修改放到缓存中

git stash 将本地修改的文件放到缓存中
git stash pop 把缓存中的文件放回本地（会覆盖本地的文件）

//2019年8月23日11:19:46
问题： 如何创建 submodule

git pull = git fetch + git merge
注：git fetch不会进行合并，执行后需要手动执行git merge合并，而git pull拉取远程分之后直接与本地分支进行合并。
更准确地说，git pull是使用给定的参数运行git fetch，并调用git merge将检索到的分支头合并到当前分支中。

#### git pull的用法：
git pull <远程主机名> <远程分支名>:<本地分支名>
举例：将远程主机origin的master分支拉取过来，与本地的branchtest分支合并。
git pull origin master:branchtest
如果将冒号和后面的branchtest去掉，则表示将远程origin仓库的master分支拉取下来与本地当前分支合并。
以上的git pull操作如果用git fetch来表示：
git fetch origin master:brantest
git merge brantest

git checkout 其实是用版本库里的版本替换工作区的版本，无论工作区是修改还是删除，都可以“一键还原”。


git checkout -- readme.txt
意思就是，把readme.txt文件在工作区的修改全部撤销，这里有两种情况：

一种是readme.txt自修改后还没有被放到暂存区，现在，撤销修改就回到和版本库一模一样的状态；

一种是readme.txt已经添加到暂存区后，又作了修改，现在，撤销修改就回到添加到暂存区后的状态。

总之，就是让这个文件回到最近一次git commit或git add时的状态

git checkout -- file命令中的--很重要，没有--，就变成了“切换到另一个分支”的命令

请问，如何查看某一个分支的提交的log?

git log -p
git reset HEAD <file>
可以把暂存区的修改撤销掉（unstage），重新放回工作区


场景1：当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时，用命令git checkout -- file。

场景2：当你不但改乱了工作区某个文件的内容，还添加到了暂存区时，想丢弃修改，分两步，第一步用命令git reset HEAD <file>，就回到了场景1，第二步按场景1操作。

场景3：已经提交了不合适的修改到版本库时，想要撤销本次提交，参考版本回退一节，不过前提是没有推送到远程库。

#### 代码push不上去的问题

问题(Non-fast-forward)出现的原因是: git仓库中已有一部分代码, 它不允许你直接把你的代码覆盖上去。于是你有2个选择方式:
1. 强推，即利用强覆盖方式用你本地的代码替代git仓库内的内容:
git push -f
2. 先把git的东西fetch到你本地然后merge后再push
    - git fetch
    - git merge
这2句命令等价于
git pull
有时候remote有多个版本分支库，git pull需要指定开发版本的分支：
git pull origin master

#### git 修改本地和远程分支名称

> git branch -a #查看所有分支
>
> git branch -r #查看远程分支
>
> git branch -vv #查看本地分支所关联的远程分支
>
> git branch -m old_branch new_branch # 修改本地分支名称 Rename branch locally
>
> git push origin :old_branch # Delete the old branch
>
> git push --set-upstream origin new_branch 
>
> git push -u origin new_branch    # Push the new branch, set local branch to track the new remote

#### git 修改远程仓库地址

方法1. 
	1.1 git remote set-url origin [url]

方法2：
	2.1 git remote rm origin
	2.2 git remote add origin [url]

#### 建立本地分支与远程分支的映射关系

git branch -u origin/addFile

或者使用命令：

git branch --set-upstream-to origin/addFile
	

#### Git合并时遇到冲突或错误后取消合并

当合并分支时遇到错误或者冲突，分支旁边会多出“|MERGING”这个东西


所以需要先取消这次合并，使用“git merge --abort”命令

``` nginx
git stash
```


我们有时会遇到这样的情况，正在dev分支开发新功能，做到一半时有人过来反馈一个bug，让马上解决，但是新功能做到了一半你又不想提交，这时就可以使用git stash命令先把当前进度保存起来，然后切换到另一个分支去修改bug，修改完提交后，再切回dev分支，使用git stash pop来恢复之前的进度继续开发新功能。下面来看一下git stash命令的常见用法

1. git stash

	工作进度，会把暂存区和工作区的改动保存起来。执行完这个命令后，在运行git status命令，就会发现当前是一个干净的工作区，没有任何改动。使用git stash save 'message...'可以添加一些注释

2. git stash list
	显示保存进度的列表。也就意味着，git stash命令可以多次执行。

3. git stash pop [–index] [stash_id]

4. git stash pop 恢复最新的进度到工作区。git默认会把工作区和暂存区的改动都恢复到工作区。

5. git stash pop --index 恢复最新的进度到工作区和暂存区。（尝试将原来暂存区的改动还恢复到暂存区）

6. git stash pop stash@{1}恢复指定的进度到工作区。stash_id是通过git stash list命令得到的
	通过git stash pop命令恢复进度后，会删除当前进度。
	
7. git stash apply [–index] [stash_id]
	除了不删除恢复的进度之外，其余和git stash pop 命令一样。

8. git stash drop [stash_id]
	删除一个存储的进度。如果不指定stash_id，则默认删除最新的存储进度。

9. git stash clear
	删除所有存储的进度。使用git log 查看提交的信息,记住commit id.

#### git 合并某个提交commit 到指定的分支上

> git checkout 要修改的分支
>
> git cherry-pick 某个commit id   // 把某个commit id的提交合并到当前分支

今天将feature分支的代码merge到develop分支后我后悔了，
因为feature分支的功能还没有全部开发完成，我在feature分支上
commit是可以的，但是这之后我又把它merge到了develop分支这
就不合适了。

> 第一步：git checkout到你要恢复的那个分支上
> git checkout develop
>
> 第二步：git reflog 查出要回退到merge前的版本号
> git reflog
>
> 第三步：git reset --hard [版本号]就回退到merge前的代码状态了
> git reset --hard f82cfd2

#### git 工作流 

1. fork 先在远程将主仓库中的代码fork到自己的托管空间下

2. git clone git@github.com:xxxx/github-flow.git 把自己托管空间下的代码克隆到本地

3. git remote -v 查看当前远端仓库的地址 输出如下：

   origin  git@github.com:xxx/github-flow.git (fetch)
   origin  git@github.com:xxx/github-flow.git (push)

4. git remote add upstream git@github.com:fe/github-flow.git
   		|
   		| // 设置一个名字为 upstream 的上游地址，也就是我们项目主仓库的地址
   		|

5. git remote -v 
   		| //查看当前远端仓库的地址
   		| origin  git@github.com:xxx/github-flow.git (fetch)
   		| origin  git@github.com:xxx/github-flow.git (push)
   		| upstream    git@github.com:fe/github-flow.git (fetch)
   		| upstream    git@github.com:fe/github-flow.git (push)
   		|	

6. git fetch upstream 
   		|
    		| //将上游主仓库中master分支的代码同步到本地
   		|

7. git checkout master 切换到本地的master分支
   		|
   		|

8.  git merge upstream/master 把上游仓库的master分支合并到本地的master分支
   		|
   		|

9. git push origin master 本地仓库master分支推送到远端（origin）仓库的master分支
   		|
   		|

10. git checkout -b feat/feedback develop 基于develop 分支，新建一个分支出来
    		|
    		| 新建feature 功能分支
    		|

现在的问题是，开发新的活动没有切换新的分支，导致代码混乱 -- git 只是的加强

#### git 命令大全

	git init                                                  # 初始化本地git仓库（创建新仓库）
	git config --global user.name "xxx"                       # 配置用户名
	git config --global user.email "xxx@xxx.com"              # 配置邮件
	git config --global color.ui true                         # git status等命令自动着色
	git config --global color.status auto
	git config --global color.diff auto
	git config --global color.branch auto
	git config --global color.interactive auto
	git config --global --unset http.proxy                    # remove  proxy configuration on git
	git clone git+ssh://git@192.168.53.168/VT.git             # clone远程仓库
	git status                                                # 查看当前版本状态（是否修改）
	git add xyz                                               # 添加xyz文件至index
	git add .                                                 # 增加当前子目录下所有更改过的文件至index
	git commit -m 'xxx'                                       # 提交
	git commit --amend -m 'xxx'                               # 合并上一次提交（用于反复修改）
	git commit -am 'xxx'                                      # 将add和commit合为一步
	git rm xxx                                                # 删除index中的文件
	git rm -r *                                               # 递归删除
	git log                                                   # 显示提交日志
	git log -1                                                # 显示1行日志 -n为n行
	git log -5
	git log --stat                                            # 显示提交日志及相关变动文件
	git log -p -m
	git show dfb02e6e4f2f7b573337763e5c0013802e392818         # 显示某个提交的详细内容
	git show dfb02                                            # 可只用commitid的前几位
	git show HEAD                                             # 显示HEAD提交日志
	git show HEAD^                                            # 显示HEAD的父（上一个版本）的提交日志 ^^为上两个版本 ^5为上5个版本
	git tag                                                   # 显示已存在的tag
	git tag -a v2.0 -m 'xxx'                                  # 增加v2.0的tag
	git show v2.0                                             # 显示v2.0的日志及详细内容
	git log v2.0                                              # 显示v2.0的日志
	git diff                                                  # 显示所有未添加至index的变更
	git diff --cached                                         # 显示所有已添加index但还未commit的变更
	git diff HEAD^                                            # 比较与上一个版本的差异
	git diff HEAD -- ./lib                                    # 比较与HEAD版本lib目录的差异
	git diff origin/master..master                            # 比较远程分支master上有本地分支master上没有的
	git diff origin/master..master --stat                     # 只显示差异的文件，不显示具体内容
	git remote add origin git+ssh://git@192.168.53.168/VT.git # 增加远程定义（用于push/pull/fetch）
	git branch                                                # 显示本地分支
	git branch --contains 50089                               # 显示包含提交50089的分支
	git branch -a                                             # 显示所有分支
	git branch -r                                             # 显示所有原创分支
	git branch --merged                                       # 显示所有已合并到当前分支的分支
	git branch --no-merged                                    # 显示所有未合并到当前分支的分支
	git branch -m master master_copy                          # 本地分支改名
	git checkout -b master_copy                               # 从当前分支创建新分支master_copy并检出
	git checkout -b master master_copy                        # 上面的完整版
	git checkout features/performance                         # 检出已存在的features/performance分支
	git checkout --track hotfixes/BJVEP933                    # 检出远程分支hotfixes/BJVEP933并创建本地跟踪分支
	git checkout v2.0                                         # 检出版本v2.0
	git checkout -b devel origin/develop                      # 从远程分支develop创建新本地分支devel并检出
	git checkout -- README                                    # 检出head版本的README文件（可用于修改错误回退）
	git merge origin/master                                   # 合并远程master分支至当前分支
	git cherry-pick ff44785404a8e                             # 合并提交ff44785404a8e的修改
	git push origin master                                    # 将当前分支push到远程master分支
	git push origin :hotfixes/BJVEP933                        # 删除远程仓库的hotfixes/BJVEP933分支
	git push --tags                                           # 把所有tag推送到远程仓库
	git fetch                                                 # 获取所有远程分支（不更新本地分支，另需merge）
	git fetch --prune                                         # 获取所有原创分支并清除服务器上已删掉的分支
	git pull origin master                                    # 获取远程分支master并merge到当前分支
	git mv README README2                                     # 重命名文件README为README2
	git reset --hard HEAD                                     # 将当前版本重置为HEAD（通常用于merge失败回退）
	git rebase
	git branch -d hotfixes/BJVEP933                           # 删除分支hotfixes/BJVEP933（本分支修改已合并到其他分支）
	git branch -D hotfixes/BJVEP933                           # 强制删除分支hotfixes/BJVEP933
	git ls-files                                              # 列出git index包含的文件
	git show-branch                                           # 图示当前分支历史
	git show-branch --all                                     # 图示所有分支历史
	git whatchanged                                           # 显示提交历史对应的文件修改
	git revert dfb02e6e4f2f7b573337763e5c0013802e392818       # 撤销提交dfb02e6e4f2f7b573337763e5c0013802e392818
	git ls-tree HEAD                                          # 内部命令：显示某个git对象
	git rev-parse v2.0                                        # 内部命令：显示某个ref对于的SHA1 HASH
	git reflog                                                # 显示所有提交，包括孤立节点
	git show HEAD@{5}
	git show master@{yesterday}                               # 显示master分支昨天的状态
	git log --pretty=format:'%h %s' --graph                   # 图示提交日志
	git show HEAD~3
	git show -s --pretty=raw 2be7fcb476
	git stash                                                 # 暂存当前修改，将所有至为HEAD状态
	git stash list                                            # 查看所有暂存
	git stash show -p stash@{0}                               # 参考第一次暂存
	git stash apply stash@{0}                                 # 应用第一次暂存
	git grep "delete from"                                    # 文件中搜索文本“delete from”
	git grep -e '#define' --and -e SORT_DIRENT
	git gc
	git fsck
		
		