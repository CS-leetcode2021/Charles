# linux 

-----------------------------------------------------------------------
### 2020-11-20-issue

[安装linux版本的微信（已解决）](https://github.com/wszqkzqk/deepin-wine-ubuntu)
-	华为云wechat仓库已经取消，现在使用github上的拉取直接安装

-----------------------------------------------------------------------
### 2020-11-22-issue

[关于安装rust（已解决）](：https://www.rust-lang.org/ )

1、直接在终端下载
```
	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```
2、安装以后需要配置PATH路径
```
    	export /home/jackey/.cargo/bin
```
-----------------------------------------------------------------------
### 2020-11-24-issue

[make tikv过程中出现gflags依赖缺少（已解决）](https://blog.csdn.net/calvinpaean/article/details/99761765)

安装步骤：
```
	git clone https://github.com/gflags/gflags.git
	cd gflags
	mkdir build && cd build
	cmake -DCMAKE_INSTALL_PREFIX=/usr/local -DBUILD_SHARED_LIBS=ON -DGFLAGS_NAMESPACE=gflags ../ 
	make -j4
	sudo make install
```
-----------------------------------------------------------------------
### 2020-12-09-issue
-   安装百度云盘
-   [参考链接](https://www.cnblogs.com/rickzhai/p/12444153.html)
-----------------------------------------------------------------------
### 2020-12-11-issues
- 解决linux系统代理问题
- [文章参考链接](https://www.rumosky.com/archives/556.html)

- 代理详解-包括购买
- [网站地址](https://tlanyan.me/v2ray-tutorial/)
