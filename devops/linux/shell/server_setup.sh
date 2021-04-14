#!/bin/bash

#********************************************************************
#Author: liangdu1992@gmail.com
#website： www.z-gour.com
#Date： 2021-04-14
#FileName： server_setup.sh
#Description： Annotated script
#********************************************************************
set -e
set -u
set -o pipefail


# Determine OS platform
UNAME=$(uname | tr "[:upper:]" "[:lower:]")


if [ X"$UNAME" = X"linux" ];then
    source /etc/os-release
    SERVER_OS=$ID
else
  echo  -e "\033platform is not linux,script not support now...:[0m"

fi

	# 更改源设置
source_setup(){
	echo -e "\033[3xm setup linux server source configs....\033[0m"
}
# 更新设置
update_setup(){
		echo -e "\033[3xm setup linux server security configs....\033[0m"
    if [ X$SERVER_OS = X"ubuntu" ];then
      echo "begin to update ubuntu"
    elif [ X$SERVER_OS = X"Centos" ];then
      echo "begin to update centos"
    else
      echo "os $SERVER_OS not support now..."
    fi
}


	# 服务器安全设置，包括端口修改，root账户禁止登录，关闭selinux
	#
security_setup(){
	echo -e "\033[3xm setup linux server security configs....\033[0m"

}


	#
	# 服务器时间和日期设置，包括时区，ntp服务等
timezone_setup(){
	echo -e "\033[3xm setup linux server time-zone configs....\033[0m"

}

	## 用户管理 set在
useradmin_setup(){
	echo -e "\033[3xm setup linux server user-administrator configs....\033[0m"

}

	# 文件描述符和文件打开数操作
fd_setup(){
	echo -e "\033[3xm setup linux server filedescripter configs....\033[0m"

}

	#服务器网络调优,包括TCP 复用等
network_setup(){
	echo -e "\033[3xm setup linux server network configs....\033[0m"

}


	# 虚拟化设置，docker容器等服务安装
virtual_setup(){
	echo -e "\033[3xm setup linux server security configs....\033[0m"
}

	# 内核优化
kernel_setup(){
	echo -e "\033[3xm setup linux server security configs....\033[0m"

}

prof_setup(){

	echo -e "\033[3xm setup linux server security configs....\033[0m"

}



main()
{

  echo -e "=================================================================="
	echo -e "\033[1;33mbegin to set_up linux server $HOSTNAME  which operating system is $SERVER_OS.... [0m"
  update_setup

 
	echo -e "=================================================================="
	echo -e "\033[32mCongratulations! server $HOSTNAME init succeeded!\033[0m"
	echo -e "=================================================================="
  
}

main


