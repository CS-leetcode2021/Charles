# RocksDB-I/O

---

1、page cache

    操作系统为了提高文件的IO性能，会增加一层page cache，用于缓存文件数据，大部分读写操作只需要访问page cache即可。
    不需要真正的发起IO，一般size为4k。

    读操作直接进行执行即可；
    写操作比较麻烦：
        如果命中page cache，直接修改并标记dirty；
        否则，刚好写入落在page size boundaries，直接分配page cache，标记dirty；
        否则，读入再修改，标记dirty。

    当内存紧缺时，会回收page cache。clean page可以直接回收，dirty page要write back后变为clean再回收。

    
2、write back

    write back 即将 dirty page 写回到 disk。write() 系统调用对于一般的文件而言(非 O_DIRECT|O_SYNC|
    O_DSYNC)，只是修改 page cache，所以只能保证进程挂掉时数据是完好的。 fdatasync()会将文件对应的dirty
    page写回disk，从而保证机器挂掉时数据是完好的，fsync() 除了dirty page外，还会把文件metadata(inode)
    写回 disk。

    除了主动 sync 外，操作系统也会帮助 write back，Linux 有如下几个配置：
    
    dirty_background_ratio：当 dirty pages 占 total available memory(free pages + reclaimable
    pages) 的比例超过该值时，会由后台 flusher 线程开始写回，默认为 10。 total available memory 不是 
    total system memory，因为还有 unreclaimable pages。dirty_background_bytes 类似，以字节为单位，
    设置其中一个，另一个会置零。
    
    dirty_writeback_centisecs：flusher 线程定期写回 dirty pages 的周期，单位为 1/100s，默认为 500。
    dirty_expire_centisecs：当 dirty pages 驻留时间超过该值时，会在下一次 flusher 线程工作时写回，单
    位为 1/100s，默认为 3000。
    dirty_ratio：和 dirty_background_ratio 意义相同，但是当超过该值时，发起写操作的进程要等待 flusher
    线程写回部分 dirty pages 保证 dirty_ratio 满足要求，相当于变为了同步写。 所以 dirty_ratio 也就是
    dirty pages 的上限。 dirty_bytes 和 dirty_background_bytes 类似。

3、block layer
    
    block layer会对I/O请求进行合并、排序等调度来提高性能，如电梯算法，然后通过队列下发任务给各个storage 
    device，每个storage device各有一个队列来保存I/O任务，互不影响。
    

4、mmap()

    mmap() 的行为比较复杂，而且底层实现也在变化，这里主要讨论 file-backed、MAP_SHARED这种。mmap()让
    程序以使用内存的形式读取和修改文件，省去了 read()/write() 的调用，原理是建立了页表到文件的映射，且
    没有立即读取文件内容，而是访问时触发 page fault 将文件加载到用户空间。

    mmap() 不能映射超过文件大小的区域，所以当文件大小发生变化时，就需要重新映射，如果是追加写，要先用 
    ftruncate()/fallocate()/lseek() 等增大文件大小，然后再映射写。

    可能在最初的实现中，mmap() 和 page cache 是不相关的，当触发 page fault 时直接从磁盘加载到用户
    空间，也就可以减少内核到用户空间的拷贝。 对文件映射的修改只有调用 msync() 或 munmap() 或者内存不
    够用发生 swap 时才会写回文件，所以使用 mmap() 来写文件时，如果文件大小超过内存可能因 swap 发生很
    多随机 I/O， 性能就会很差，而且带宽用的也不高，这时候需要搭配 msync() 和 madvise() 使用：

    msync()：写回指定部分的文件映射，MS_ASYNC 是异步写回，即相当于提交了写任务，MS_SYNC 是同步写回，
    即写到磁盘才返回。
    madvise()：告诉操作系统对映射操作的建议，MADV_DONTNEED 可以释放资源，降低内存使用。
    使用 mmap() 来读文件，也需要 madvise()，否则性能也会差：
        MADV_NORMAL：默认的行为，会发生少量 read-ahead 和 read-behind。
        MADV_SEQUENTIAL：read-ahead 更多的 pages，而且访问过的 page 会被很快释放。
        MADV_RANDOM：随机访问模式，不会进行 read-ahead 和 read-behind。
    mmap() 也有可见性问题，比如使用 mmap() 修改了共享文件映射，那么使用read()能否读到；使用write()
    修改了文件，mmap() 能否读到:
        不同进程相同映射的修改是立即可见的，可用于 IPC。
        mmap() 修改后要调用 msync(MS_ASYNC||MS_SYNC)，read() 才能读到修改。
        write() 修改后要调用 msync(MS_INVALIDATE)，mmap() 才能读到修改
    
    但是在 boltdb 中，是用共享文件 mmap() 作为 page pool，修改是调用 write()，也没调用 msync
    (MS_INVALIDATE)，这是怎么保证可见性的呢？原因是大多数操作系统(包括 Linux)都采用了 unified 
    virtual memory，或者叫 unified buffer cache，也就是 mmap() 和 page cache 会尽可能共享
    物理内存页，当触发 page fault 时，会先从 page cache 中查找对应的文件 page，read-ahead 等
    也是保存在 page cache(见 mmap, page cache 与 cgroup)。 所以 mmap() 和 read()/write() 
    是保持一致的(direct-io 有影响吗？)，也就不再需要同步，msync() 的用途只用于写回磁盘。这也影响了 
    mmap() 的写回，之前只有主动 msync()、munmap() 或者 swap 时写回， 现在因为已经是dirty page
    了，会被操作系统异步写回，fsync()/fdatasync() 也有效。

5、Writable File

    RocksDB 只会用顺序写，支持 direct I/O 和 mmap()，设置对应的 Options 即可。

    mmap() 追加写的时候，要扩大文件大小再重新映射，RocksDB 用的是 fallocate()，在 close 的时候要
    记得 ftruncate() 调整大小为真实大小。
    RocksDB 是读设备文件获取 direct I/O 对齐要求的，见 io_posix.cc:GetLogicalBufferSize()。
    WritableFileWriter 用 AlignedBuffer 缓冲写数据，AlignedBuffer 的起始地址和大小都是对齐的，
    以便使用 direct I/O。
    在一些场景下会先 fallocate() 预先分配空间然后再写，好像可以提高性能并且减少空间浪费，有几个配置
    也个这个功能有关，如 fallocate_with_keep_size 分配空间时不会改变文件大小，可用于提高顺序写性能。

    会用 sync_file_range() 定期刷盘，因为如果只靠操作系统异步刷盘的话，在 I/O 很重的情况下可能导致
    阻塞，原因见上面的 write back。但是最新的 1MB 不会刷，为了避免后面再追加写时重复写入。

6、Readable File

    RocksDB 中读文件有顺序读和随机读，顺序读不会用 mmap()。需要注意的是 readahead，操作系统默认会做
    readahead，能够提高顺序读的性能，但是对随机读毫无作用，可能还会降低性能，因为可能会把其他有用的page
    从page cache中挤出去，可以用 posix_fadvise(POSIX_FADV_RANDOM) 禁用 readahead，RocksDB 会
    对随机访问的文件自己做 readahead，见 ReadaheadRandomAccessFile 和 FilePrefetchBuffer。 
    posix_fadvise(POSIX_FADV_DONTNEED) 用于释放 page cache，也可以避免有用的 page 被挤出去。

