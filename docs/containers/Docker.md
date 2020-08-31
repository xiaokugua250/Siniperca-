# docker 资料
      
      ### docker 资源隔离和限制
      namespace 是用来做资源隔离的，在 Linux 内核上有七种 namespace，docker 中用到了前六种。第七种 cgroup namespace 在 docker 本身并没有用到，但是在 runC 实现中实现了 cgroup namespace。  
      1.mount  
      2.uts  
      3.pids  
      4.network  
      5.user  
      6.ipc  
      7.cgroup  
      
          第一个是 mout namespace。mout namespace 就是保证容器看到的文件系统的视图，是容器镜像提供的一个文件系统，也就是说它看不见宿主机上的其它文件，除了通过 -v 参数 bound 的那种模式，是可以把宿主机上面的一些目录和文件，让它在容器里面可见的；
      
       
      
          第二个是 uts namespace，这个 namespace 主要是隔离了 hostname 和 domain；
      
       
      
          第三个是 pid namespace，这个 namespace 是保证了容器的 init 进程是以 1 号进程来启动的；
      
       
      
          第四个是网络 namespace，除了容器用 host 网络这种模式之外，其他所有的网络模式都有一个自己的 network namespace 的文件；
      
       
      
          第五个是 user namespace，这个 namespace 是控制用户 UID 和 GID 在容器内部和宿主机上的一个映射，不过这个 namespace 用的比较少；
      
       
      
          第六个是 IPC namespace，这个 namespace 是控制了进程兼通信的一些东西，比方说信号量；
      
       
      
          第七个是 cgroup namespace，上图右边有两张示意图，分别是表示开启和关闭 cgroup namespace。用 cgroup namespace 带来的一个好处是容器中看到的 cgroup 视图是以根的形式来呈现的，这样的话就和宿主机上面进程看到的 cgroup namespace 的一个视图方式是相同的；另外一个好处是让容器内部使用 cgroup 会变得更安全。
      
       
      
      这里我们简单用 unshare 示例一下 namespace 创立的过程。容器中 namespace 的创建其实都是用 unshare 这个系统调用来创建的。
      
      
      2. cgroup
      
      
      两种 cgroup 驱动
      
      cgroup 主要是做资源限制的，docker 容器有两种 cgroup 驱动：一种是 systemd 的，另外一种是 cgroupfs 的。
      
      
       
      
          cgroupfs 比较好理解。比如说要限制内存是多少、要用 CPU share 为多少？其实直接把 pid 写入对应的一个 cgroup 文件，然后把对应需要限制的资源也写入相应的 memory cgroup 文件和 CPU 的 cgroup 文件就可以了；
      
      
          另外一个是 systemd 的一个 cgroup 驱动。这个驱动是因为 systemd 本身可以提供一个 cgroup 管理方式。所以如果用 systemd 做 cgroup 驱动的话，所有的写 cgroup 操作都必须通过 systemd 的接口来完成，不能手动更改 cgroup 的文件。
          
      容器中常用的 cgroup
      
      接下来看一下容器中常用的 cgroup。Linux 内核本身是提供了很多种 cgroup，但是 docker 容器用到的大概只有下面六种：
      
      
      
          第一个是 CPU，CPU 一般会去设置 cpu share 和 cupset，控制 CPU 的使用率；
      
          第二个是 memory，是控制进程内存的使用量；
      
          第三个 device ，device 控制了你可以在容器中看到的 device 设备；
      
          第四个 freezer。它和第三个 cgroup（device）都是为了安全的。当你停止容器的时候，freezer 会把当前的进程全部都写入 cgroup，然后把所有的进程都冻结掉，这样做的目的是：防止你在停止的时候，有进程会去做 fork。这样的话就相当于防止进程逃逸到宿主机上面去，是为安全考虑；
      
          第五个是 blkio，blkio 主要是限制容器用到的磁盘的一些 IOPS 还有 bps 的速率限制。因为 cgroup 不唯一的话，blkio 只能限制同步 io，docker io 是没办法限制的；
      
          第六个是 pid cgroup，pid cgroup 限制的是容器里面可以用到的最大进程数量。
      也有一部分是 docker 容器没有用到的 cgroup。
      
      
      容器中常用的和不常用的，这个区别是对 docker 来说的，因为对于 runC 来说，除了最下面的 rdma，所有的 cgroup 其实都是在 runC 里面支持的，但是 docker 并没有开启这部分支持，所以说 docker 容器是不支持下图这些 cgroup 的。
      不常用的cgroup有
      
      1.net_cls
      2.net_prio
      3.hugetlb
