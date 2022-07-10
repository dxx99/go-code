# mysql实战45讲

## 1. 基础篇
### 1.1 一条sql查询如何执行
[链接](https://excalidraw.com/#json=YAoxWVyLifjAHstRdNLLI,3shnK3oLNTjFO0r8oPaiGA)
![](./images/mysql/mysql架构图.jpg)

#### 1.1.1 连接器
- tcp连接之后，在进行用户账号密码权限认证，通过之后在权限表中查出所有权限
- 之后这个连接里面的权限判断逻辑，都依赖此时读到的权限。(修改权限，原来的连接权限不变)
- 空闲连接啥时候被回收 `show variables like '%wait_timeout%'` 单位秒，默认8小时

**怎么解决长连接导致内存占用过大问题**  
- 定期断开连接
- 版本大于5.7可以通过 mysql_reset_connection() 初始化连接资源

#### 1.1.2 查询缓存
- 鸡肋产物，redis完美解决

#### 1.1.3 分析器
- 词法解析，语法解析

#### 1.1.4 优化器
- 多个索引时候，决定用哪一个
- 多表关联时，决定表的连接顺序
- 等等

#### 1.1.5 执行器
- 先要判断你对这个表有没有执行权限(表权限判断)
- 选择表对应的引擎提供的接口操作


### 1.2 一条sql更新是如何执行的
- 与查询基本差不多，需要失效所有的表相关的缓存
- 找到数据进行更新
- 设计重要的日志记录（重做日志，归档日志）

#### 1.2.1 重要的日志模块： redo log
- 是innodb引擎层的日志
- WAL(write ahead logging)技术
  - 磁盘顺序写比随机写快
  - 组提交机制可以大幅降低磁盘的IOPS消耗(fsync)
- 刷脏页策略如何控制
  - 定时刷
  - buffer pool不足
  - mysql 正常关闭
  - redo log满的时候

#### 1.2.2 重要的日志模块：binlog
- 为归档日志，在server层做记录，没有crash-safe能力
- 数据更新执行器流程  
![](./images/mysql/数据更新执行器流程.png)

**思考？**
- 为什么有了redo log还需要binlog, 谈一下你对这两种日志的理解，以及它们的区别？
  - redo log 是innodb所有，binlog是所有的引擎所有，在sever层
  - redo log 是物理日志，记录某个数据页修改了什么，binlog是逻辑日志，记录语句的原始逻辑，比如"给ID=2的字段c加1"
  - redo log是循环写，空间固定会用完，binlog是追加写，不会覆盖以前的日志

#### 1.2.3 两阶段提交
- redo log分为两个阶段prepare和commit阶段，拆成两份的目的是为了保证两份日志之间逻辑一致
- 数据要做定期的备份

**思考**
- 怎么让数据库恢复到半个月内任意一秒状态？
  - 先找到最近的全量备份，将这个备份恢复到临时库
  - 从备份的开始时间算，将备份的binlog依次取出，重放到指定的时刻
  - 然后再把临时的数据库恢复到线上
- 日志为啥需要"二阶段提交"？(1 prepare阶段 2 写binlog 3 commit)
  - redo log与binlog提交的先后顺序都会导致日志数据的不一致
  - 当在2之前崩溃时, 重启恢复：后发现没有commit，回滚。备份恢复：没有binlog 。
  - 当在3之前崩溃 ,重启恢复：虽没有commit，但满足prepare和binlog完整，所以重启后会自动commit。备份：有binlog.
- 什么场景需要用到binlog和redo log来恢复数据?
  - 误操作恢复数据
  - 扩容恢复从库都会有问题 


### 1.3 事务隔离：为什么你改了我还看不到？
- 事务特性(ACID)
  - 原子性
  - 一致性
  - 隔离性
  - 持久性
- 多个事务同时执行可能出现的问题
  - 脏读(dirty read) (读到其他事务未提交的数据)
  - 不可重复读(non-repeatable read) (前后读取的记录内部不一致)
  - 幻读(phantom read) (前后读取的记录行数不一致)
- 隔离级别
  - 读未提交(read uncommitted)  (没有视图，直接返回最小行数据)
  - 读已提交(read committed) 视图(事务启动时创建)
  - 可重复读(repeatable read) 视图(在sql语句执行时创建) 事务在整个执行的过程中前后看到的数据是一致的
  - 串行化(serializable)   加锁
- 如何查看事务的隔离级别
  - `show variables like 'transaction-isolation';`
- RR级别的隔离如何解决幻读
  - `for update` 添加排他锁，当前读，也就是写锁
  - 使用可串行化(serializable)

#### 1.3.1 事务隔离级别的实现
- MVCC(多版本并发控制)
- undo log(回滚日志)

#### 1.3.2 事务的启动方式
- begin, start transaction
- rollback
- commit
- set autocommit=0 关闭自动提交事务，减少显示设置开启事务
- 修改事务的隔离级别
  - `set global transaction isolation level read committed;`
  - 或者通过修改配置

#### 1.3.3 事务到底是隔离还是不隔离？
- `begin/start transaction` 并不是事务的起点，一致性快照的起点
  - 一致性快照是在执行第一个快照读语句创建
  - 一致性快照是在执行`start transaction with consistent snapshot`
- 在mysql有"两个"视图的概念
  - 一个是view, 通过create view....
  - 一个是MVCC的一致性读视图，consistent read view，用在RC、RR隔离中实现
    -没有物理结构，用来定义事务执行期能看到什么数据

#### 1.3.4 快照在MVCC里是怎样工作的
- 快照是基于整库实现的，主要是一个数组（当前正在"活跃"的所有事物）的拷贝
- Innodb的每个事务都有一个唯一的ID,叫transaction id，是事务开始的时候申请的，按顺序严格递增
- 每行数据有多个版本，每次事务更新数据的时候都会生成一个新的版本，并且把transaction id 赋值给这个版本的数据，旧数据保留，新版本可以直接拿到这个版本数据
- 也就是数据表中的一行记录，其实可能有多个版本记录，每个版本记录有自己的row trx_id

**回滚日志（undo log）**
- 多版本并不是物理上真实存在的，而是通过undo log计算出来的
  - 比如需要V2版本的数据（需要通过v4版本依次执行U3、U2算出来）

**可重复读的隔离级别定义**
- 一个事务启动时，能够看到所有已提交的事务结果，但是之后，在执行期间，其实事务的更新对它不可见
- 事务执行之后，其他事务的更新对它虽然不可见，但是数据版本还是可见的，因为数据库实际上存储的是最新版本的数据。但是对于该事务来说，需要根据版本号以及Undo Logs计算出他需要的版本对应的数据

**事务的可重复读能力是怎么实现的？**
- 核心是一致性读（consistent read）
- 事务更新数据的时候只能用*当前读*
- 如果当前事物的行锁被其他事务占用的话，就需要进入锁等待

**读提交与可重复读的逻辑类似，主要区别**
- 可重复读隔离级别下，只需要在事务开始时候创建一致性视图，之后事务里的其他查询都共用这个一致性视图
- 在读提交的隔离级别下，每一条语句执行都会重新算出一个新的视图

**思考**
- 回滚日志什么时候删除?
  - 系统中没有比回滚日志更早的日志就可以删除原因，因为read-view在commit之后就会被删掉
  - 没有事务再需要用到这个回滚日志，会被purge进程删掉
- 怎么看innodb中的长事务？对系统的危害有哪些？
  - information_schema.INNODB_TRX
  - 长事务的危害，导致undo log一直不回被删掉，导致大量的记录会被保留，占用额外的存储空间
  - 占用锁资源，拖垮数据库
- 如何避免系统的长事务？
  - 应用端思考
    - 开是否开启auto_commit, 开启general_log日志，去确定业务逻辑是否有问题
    - 确定是否有不必要的只读事务，把不必要的select从事务中去掉
    - 根据业务本身预估设置max_execution_time的最长时间，必须单条语句执行过长
  - 从数据库端思考
    - 监控information_schema.INNODB_TRX的表的数据，设置报警，超时kill动作
    - 利用pt-kill工具监控长事务
    - 保证undo log表空间足够
- 为啥表结构没有可重复读？
  - 因为表结构没有对应的行数据，也没有row trx_id,因此只能遵循当前的读逻辑

### 1.4 深入浅出索引
- 索引的目的是为了提高数据的查询效率

#### 1.4.1 索引常见的模型
- **hash表**，适合等值查询，不适合区间查询
- **有序数组**，等值查询与范围查询都很优秀，添加数据要移动数组，且空间有限
- **二叉搜索树**，
  - 为了维持O(log(N))的查询效率，需要维持这个数的平衡
  - 二叉搜索的效率高，但不用的原因主要是因为索引除了存在内存中外，还有可能在磁盘上，树太高，磁盘寻址太慢
  - 为了减少数据块的访问次数，因此我们就应该使用N叉树，这里的N取决于数据块的大小
  - N叉树的读写性能优点，以及适配磁盘的访问模式，非常适合数据库存储引擎
- **跳表(skiplist)**，增加向前指针的链表(多级索引)，随机化数据结构，可以进行二分查找的有序链表
- **LSM树**，使用顺序写代替随机写来提高性能，与此同时微弱降低读性能

#### 1.4.2 Innodb索引模型
- B+树索引模型
  - 每一个索引在Innodb里面对应一颗B+树
  - 索引类型分为主键索引(聚族索引)，非主键索引(二级索引)
- 主键索引与普通索引的查询有啥不同？
  - 非主键查询可能需要进行回表，索引覆盖之后就不用回表

#### 1.4.3 索引维护
- 维护索引数据的有序性
  - 在插入的时候需要进行维护，推荐使用自增主键，好处是每次插入数据追加，不必需要数据移动
  - 自增主键 `not null primary key auto_increment`
- 业务字段做主键问题
  - 逻辑的字段不容易保证数据的自增，这样写数据成本相对高一点
  - 如果字段大小控制不好，导致二级索引需要占用更多的额外空间
- **页分裂**
  - 页数据满，需要申请额外的页，移动部分数据，这个过程称为页分裂，会影响数据页的空间利用率
- **页合并**
  - 相连的两个页由于数据的删除，导致空间利用率很低，会进行页的合并

#### 1.4.4 经典范围查询如果做回表
- `select * from t1 where k between 3 and 5`
  - 先找k索引树上的3，通过Id进行回表
  - 再顺序访问k上的4，通过ID进行回表
  - 再顺序访问k上的5，通过ID进行回表
  - 最后访问k上的6，不满足条件退出循环。
- 如何避免回表查询
  - 覆盖索引，减少树的搜索，一种常见的优化手段
- 最左前缀原则
  - 联合索引是先根据第一个字段排序，如果第一个字段相同，再根据第二个字段排序
  - 创建联合索引要考虑索引字段的排序，尽量保证索引的复用能力
  - 如果(a,b)字段分开查询的频率都比较高，就要考虑字段的空间，来建立联合索引
- 索引下推
  - 仅能利用使用最左前缀原则，利用二级索引的值进行判断，减少回表查询
  
**思考？**
- 什么场景适合用业务字段来表示主键
  - 只有一个索引（这个就没有二级索引占用额外的空间）
  - 该索引必须是唯一索引（这样就不会重复）
- Innodb为什么要用B+树来进行索引
  - 更好配合的磁盘的读写特性，减少单次查询磁盘的访问次数
- 没有主键索引建普通索引，Innodb是如何进行查询
  - 如果删除，新建主键索引，会同时去修改普通索引对应的主键索引，性能消耗比较大。
  - 删除重建普通索引貌似影响不大，不过要注意在业务低谷期操作，避免影响业务。


### 1.5 全局锁和表锁
- 锁的目的
  - 处理多用户访问共享资源的并发问题
- 锁分类
  - 全局锁、表锁、行锁

#### 1.5.1 全局锁
- 对数据库整个实例加锁
  - `flush tables with read lock` (让数据库只读)
  - `unlock tables` (解除数据库的全局锁)
  - 应用场景，坐全库逻辑备份
- 官方的mysqldump对数据进行备份，如何保证数据的一致性
  - 导数据之前，会启用一个事务，来确保拿到数据的一致性视图
  - 为啥还需要FTWRL, 因为并不是所有的引擎都支持
- 为啥不使用 `set global readonly=true` 来设置只读
  - readonly有可能用到其他业务，比如主从
  - 客户端异常断开，FTWRL会释放锁，但是readonly不会，所以风险更高

#### 1.5.2 表级锁
- 分两种
  - 表锁
  - 元数据锁(meta data lock)MDL
- 表锁语法
  - lock tables ... read/write
  - write锁是排它锁，意味着其他线程不能读写
  - read锁是共享锁，意味着其他线程只读不写，本线程也只能读不能写
- MDL不显示使用
  - 一个访问就会自动加上，如果在查询的过程中，表结构数据发生变化，导致数据内容对不上，肯定不行
  - 在做增删改查的时候加上MDL读锁，在对表结构修改时加上写锁
  - MDL锁只有在事务提交的时候才释放，要小心防止锁住线上数据
- 如何给你查询频繁的表添加字段
  - 防止阻塞，添加等待时间，如果在规定的时间还没拿到锁就放弃，后面再重试
  - `alter table t1 nowait/wait n add column`

#### 1.5.3 行级锁
- 是有引擎自己实现了，并不是所有引擎都支持，MyiSAM就不支持行锁
- Innodb行锁的默认超时时间为 innodb_lock_wait_timeout设置为on

**二阶段锁**
- 行锁实在需要的时候加上的，但不是立马就释放，而是等到事务提交才释放
- 锁的添加与释放分为两个阶段，之间不允许交叉加锁和释放锁
- 如果你的事务可能要锁多行，就要把最有可能造成冲突的锁，影响并发度的锁放在后面

**Innodb死锁的产生，如何解决死锁问题**
- 由于两个事务之间资源循环依赖，涉及的线程都在等待别的线程资源释放，导致死锁的产生
- 两种策略解决死锁
  - 进入等待，一直到超时 innodb_lock_wait_timeout
  - 发起死锁检测，发现死锁，主动回滚死锁链条中的某一个事务，让其他事务能执行 innodb_deadlock_detect
- 死锁检测原理
  - 构建一个以事务为顶点、锁为边的有向图，判断有向图是否存在环，存在即有死锁
- 回滚原理
  - 选择插入更新或者删除的行数最少的事务回滚，于 INFORMATION_SCHEMA.INNODB_TRX 表中的 trx_weight 字段来判断。
  
**怎么解决热点数据更新导致的性能问题？**
- 关闭死锁检测（风险：大量超时）
- 做并发控制
- 优化业务逻辑


## 2. 实战篇
### 2.1 普通索引与唯一索引的选择？
#### 2.1.1 查询过程
- 以如下语句进行优化
`select * from t where k = 5;`

**普通索引**
- 先通过普通索引k查找到满足k=5的记录，然后再去判断下一条记录是否满足k=5, 然后回表查询记录
- 由于innodb的最小单位是page, 默认大小为16kb, 所以下一条记录的判断绝大多的时候不需再次进行磁盘Io, 除非分页

**唯一索引**
- 由于唯一索引的唯一性，查询到唯一满足条件的时候就会停止检索，然后回表查询到对应记录
- 因为唯一索引比普通索引查询性能好一丢丢。理论上cpu的操作性能损耗应忽略

#### 2.1.2 change buffer(可以持久化数据)
- 当需要数据更新的时候
  - 如果数据在内存中，直接更新
  - 如果不在内存中，在不影响一致性的前提，更新操作会缓存再change buffer中，这样就不需要从磁盘中读取数据
- 在下次访问数据页的时候(merge操作)
  - 将数据缓存到内存中，然后执行change buffer中与这个页的操作，更新内存中的数据
- 什么情况下会触发merge操作(将change buffer的数据应用到原始数据页)
  - 读操作
  - 后台线程定期merge
  - 数据库的正常关闭
- change buffer主要为了解决什么问题
  - 减少读磁盘(随机读)
- 什么情况下可以使用change buffer?
  - 由于唯一索引必须判断这个k=4这条记录否是存在索引中，所以必须先将数据读入到内存中。因为唯一索引不能使用change buffer
  - 普通索引可以使用change buffer

#### 2.1.3 思考题？
- 某次写使用了change buffer机制，之后主机异常，是否会造成change buffer数据丢失？
  - 不会丢失
  - 虽然只更新内存，但在事务提交的时候，change buffer记录在redo log
  - 系统恢复会通过redo log恢复change buffer的数据
- merge的执行流程如何？会把数据写回磁盘？
  - 从磁盘中读入数据页到内存(老版本的数据页)
  - 从change buffer里找到这个数据页的change buffer记录(可能多条)，依次执行，得到新的数据页
  - 写入redo log, 这个redo log包含了数据的变更和change buffer记录的变更。
  - merge就执行完成了，因此merge不会写入磁盘，写入磁盘是脏页的刷新逻辑


### 2.2 mysql为什么会选错索引
#### 2.2.1 索引选择的工作原理
- 优化器选择索引的逻辑
  - 扫描的行数(扫描行数越少->访问磁盘越少->消耗cpu越少)
    - 根据统计信息估算记录数，也就是"区分度"，不同的值越多，区分度越高
    - 查看索引的基数 `show index from t4;` "cardinality"
    - 通过采样统计来确定
      - 默认选择N页，统计不同值，得到平均値，然后再乘以索引的页面数，这样就可以得到索引的基数
      - 当变更的数超过 1/M 页的时候，会触发重新做一次索引统计
    - 两种存储索引统计的方式，通过innodb_stats_persistent来设置
      - on 表示统计信息回持久化存储，默认N=20,M=10
      - off 表示统计信息只存储在内存中，默认N=8,M=16
  - 优化器还要考虑索引回表查询的代价
    - `analyze table t4` 可以用来重新统计索引信息
  - 是否使用临时表(内存不够，使用基于磁盘的临时表)
    - union查询
    - order by 与 group by的子句不一样时
    - distinct查询并加上order by时
    - 子查询
  - 是否进行排序
    - `explain select * from t4 where a between 1 and 1000 and b between 50000 and 100000 order by b limit 1;` 使用索引b, 而不是索引a
    - 因为使用索引a, 优化器会认为使用文件排序需要额外的消耗，所以选错索引

#### 2.2.2 索引选错异常处理
- 使用force index 强行选择索引
  - `explain select * from t4 force index(a) where a between 1 and 1000 and b between 50000 and 100000 order by b limit 1;`
  - 不够优美，变更索引名需要调整等问题
- 修改sql, 使其命中我们期望的索引
  - `explain select * from t4 where a between 1 and 1000 and b between 50000 and 100000 order by b,a limit 1;`
  - 调整 order by 使其选择任何一个索引都需要排序
  - `explain select * from  (select * from t4 where (a between 1 and 1000)  and (b between 50000 and 100000) order by b limit 100)alias limit 1;`
  - 诱导编译器选择索引b代价很高
- 新建更适合的索引，来提供给优化器做选择，或删掉误用的索引
  - 删掉不合适的索引
  - 建立联合索引避免回表查询等

### 2.3 怎么给字段添加索引
- 使用索引的原则
  - 区分度越高越好，因为高区分度，意味着复用的键值越少
- 索引创建的使用场景
  - 直接创建完整索引，要注意空间的占用情况
  - 前缀索引，节省空间，但会增加查询扫描次数，并且不能使用索引覆盖
  - 倒叙索引，可以解决前缀索引区分度不够问题，不支撑范围查询
  - hash索引，查询性能稳定，需要额外的存储和计算消耗，不支持范围查询

### 2.4 为什么我们的mysql会"抖"一下
#### 2.4.1 脏页的概念
- 脏页/干净页的理解
  - 当内存数据页跟磁盘数据页内容不一至的时候，我们称这个内存页为"脏页"
  - 内存数据写入磁盘后，内存和磁盘上的数据页的内容就一致，我们称"干净页"
  - 脏页与干净页都在内存中
- innodb在什么情况下，会触发刷脏(flush)过程？
  - redo log写满了(停止更新操作，checkpoint往前推进，redo log留出空间)
  - 系统内存不足，需要淘汰脏页
  - 系统空闲，触发刷脏
  - 系统正常关闭，触发刷脏

#### 2.4.2 刷脏对系统性能的影响
- 明显影响性能的情况？
  - redo log写满，更新全部堵住，写性能为零，对敏感性业务不能接受
  - 一个查询淘汰的脏页个数太多，会导致查询的响应时间明显变长
- 缓冲池中内存的三种状态
  - 没有使用
  - 使用了的干净页
  - 使用了的脏页
- 读取申请内存不够的淘汰策略
  - 先淘汰最久不使用的数据页从内存中淘汰
  - 如果是干净页直接释放复用，如果是脏页，必须先将脏页刷磁盘变成干净页之后再复用
  
#### 2.4.3 innodb刷脏的控制策略
- `show variables like '%innodb_io_capacity%'`
  - innodb_io_capacity表示磁盘的能力，建议设置成磁盘的iops
- innodb刷盘速度参考因素
  - 刷脏比例 `innodb_max_dirty_pages_pct` 默认75%
  - redo log写盘熟读
- 查看innodb的脏比
  - `select VARIABLE_VALUE into @a from performance_schema.global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_dirty'; select VARIABLE_VALUE into @b from performance_schema.global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_total'; select @a/@b;`
- innodb连刷机制
  - `show variables like '%innodb_flush_neighbors%';`
  - 会让查询抖动更大，现在磁盘io增大，建议关闭这个
- WAL技术
  - 将数据库的随机写转成顺序写，大大提升数据库性能

### 2.5 为什么表数据删除一半，表文件不变
#### 2.5.1 表结构的定义与数据存储
- mysql的innodb表包含部分
  - 表结构定义(8.0可以把表结构定义放在系统数据表中)
  - 数据
- innodb_file_per_table用来控制表数据存在共享表空间，也可以单独的文件
  - OFF表示，表数据放在系统共享表空间，也就是跟数据字典放一起
  - ON表示，每个Innodb表数据存储在一个.idb为后缀的文件（建议）
    - 独立在删除 drop table命令之后，系统可以直接删除文件，空间得到释放

#### 2.5.2 数据删除的执行流程
- 如果删掉一条记录R4
  - 会标记这个位置为删除，但磁盘文件的大小并不缩小
  - 如果后面插入的记录会在R4这个位置，会复用这个位置
- 如果删掉整个数据页上的所有记录
  - 这样整个数据页就可以被复用
  - 数据页复用跟记录复用不同，整个页复用可以存储到任何位置
- 相邻两个数据页利用率，页合并操作
  - 系统会将两个页的数据合并到一个也上，另外一个页就被标记为可复用
- 如果使用delete删除数据
  - 所有的页被标记为可复用，但磁盘上文件大小不变
  - 也就是说delete不会回收磁盘空间
  - 实际上不止删除数据会造成空洞，插入数据也会（随机的插入，就可能造成索引数据的分裂）

#### 2.5.3 如何解决磁盘不回收的问题？
- 重建表
  - `alter table A engine=Innodb`
  - mysql会自动完成转存数据，交换表名，删除旧表的操作
- 重建表的时候，往临时表插入数据，旧表又有数据写入如何解决？
  - mysql5.6 引入 online DDL
  - 建立一个临时文件，扫描表 A 主键的所有数据页；
  - 用数据页中表 A 的记录生成 B+ 树，存储到临时文件中；
  - 生成临时文件的过程中，将所有对 A 的操作记录在一个日志文件（row log）中，对应的是图中 state2 的状态；
  - 临时文件生成后，将日志文件中的操作应用到临时文件，得到一个逻辑数据上与表 A 相同的数据文件，对应的就是图中 state3 的状态；
  - 用临时文件替换表 A 的数据文件。
  - 由于日志文件的记录和重放操作功能的存在，在重建表的时候，允许对表A做增删改操作
  - 注意：会消耗IO和cpu资源
  
#### 2.5.4 online与inplace的区别
- tmp_table是由server创建的,`alter table t engine=innodb,ALGORITHM=copy;`
- tmp_file 是由innodb创建的，整个ddl都在innodb内部完成，是一种原地操作
- tmp_file `alter table t engine=innodb,ALGORITHM=inplace;`
- 给innodb添加全文索引字段过程
  - `alter table t add fulltext(file_name);`
  - 这个过程是inplace,但会阻塞增删改操作，非online
  - ddl过程如果是online, 则一点是inplace
  - 反之来未必，添加全文索引(fulltext index) 和 空间索引(spatial index)就属于这种情况

#### 2.5.5 optimize table / analyze table / alter table / truncate区别
- alter table 是recreate table操作
- analyze table 不是重建建表，只是对索引信息做重新的统计，没有修改数据，这个过程中加MDL读锁
- optimize table 等于 recreate + analyze
- truncate table 等于 drop + create

#### 2.5.6 思考题？
- 进行收缩表，结果占用空间更大
  - 本来很紧凑
  - 重新收缩，15/16整理数据，1/16留给update
  - 未整理之前的页已经占用了90%以上，收缩后，文件反而变大


### 2.7 count(*)这么慢，如何优化？
#### 2.7.1 count(*)实现方式
- innodb与myisam实现的原理
  - innodb 需要把数据一行一行地从引擎里面读出来，然后累积记数
  - Myisam引擎把表的总行数存在磁盘上
    - 每次直接返回，效率高，但这是没有过滤条件的count(*)
  - innodb引擎为什么不存储多少行数据，方便查询
    - 由于多版本并发控制(MVCC)的原因，Innodb也不知道返回多少行
    - 在事务的多个版本读取的行数不同
- innodb扫描行的优化
  - 扫描主键索引与其他索引一样，因此会选择扫描最小那棵树来遍历
  - 在保证逻辑正确的前提下，尽量减少扫描的数据量，是数据库系统设计的通用法则之一。
- `show table status` 的行数是估算出来的
  - 误差40%-50%
  - 不能直接替换扫描的行数

#### 2.7.2 经常使用count(*)如何优化
- 思路： 找一个地方，记录表的行数
  - redis缓存起来
    - 丢失，
    - 并发数据不精准，并发系统无法控制不同线程执行时刻
  - 找一个额外表存储起来
    - 解决丢失问题
    - 事务隔离的可视性问题

#### 2.7.3 不同的count逻辑不同问题
- count(*)、count(主键)、count(1)表示返回满足条件的结果集总行数
- count(字段)，表示满足条件，参数"字段"不为null的总个数
- 执行效率
  - count(字段)<count(主键id)<count(1)<count(*)
  - count(主键id)需要解析id, 然后判断null, 再累加
  - count(1) 不需要解析，只需判断行不为null，然后累积
  - count(*) mysql进行优化，认为一定非空，直接按行累积

#### 2.7.5 思考题？从并发性能来看，先插入操作记录还是先更新计数表
- 先插入记录，再更新计数统计表
- 更新统计表涉及到行数，先插入再更新能最大程度减少事务之间的锁等待，提高并发度


### 2.8 日志与索引的相关问题
- mysql修改操作redo log 与 binlog的二阶段提交[1.2.3]

#### 2.8.1 追问二阶段提交
- mysql怎么知道binlog是完整的
  - statement格式的binlog, 最后有一个commit
  - row格式的binlog, 最后会有一个xid event
- redo log 与 binlog是怎么串联起来的
  - 通过一个共同的字段 xid
  - 崩溃恢复的时候，会按顺序扫描redo log
  - 如果碰到有prepare、commit的redo log直接提交
  - 如果只有prepare, 没有commit的redo log, 会拿着xid去binlog找对应事务
- 处于 prepare 阶段的 redo log 加上完整 binlog，重启就能恢复，MySQL 为什么要这么设计?
  - 这个策略会 解决主从和备份数据的一致性问题，因为这些操作依赖binlog
  - 需要保证binlog的完整性
- 为什么需要二阶段提交？干脆先 redo log 写完，再写 binlog。崩溃恢复的时候，必须得两个日志都完整才可以。是不是一样的逻辑？
  - 二阶段提交是经典的分布式系统的问题
  - 如果直接写redo log, binlog失败，redo log又不能回滚，会造成数据的不一致性
- 只保留binlog可以吗？
  - 不能支持崩溃恢复
- 只有redo log可以吗？
  - 不可以，redo log是一个循环写，会覆盖原来的记录
  - 没办法起到归档作用
  - mysql系统的高可用就是binlog的复用
- redo log一般要设置多大？
  - 太小，很容易写满，不能不强刷redo log, wal技术发挥不出来
- 正常运行中的实例，数据写入后的最终落盘，是从 redo log 更新过来的还是从 buffer pool 更新过来的呢？
  - redo log并没有记录数据页的完整数据，没有能力更新数据磁盘页
  - 脏页的刷新跟redo log没有关系
  - 崩溃恢复的时候，脏页会丢失，会先将数据读到内存，然后用redo log更新为脏页
- redo log buffer 是什么？是先修改内存，还是先写 redo log 文件？
  - redo log buffer是一个内存，先保存redo日志
  - 真正的日志是redo log file(文件名ib_logfile+数字)，是在commit语句的时候做的
  - 事务在执行的过程中不会主动刷盘，用来减少io的消耗

#### 2.8.2 思考题，跟新一条相同的数据，mysql如何处理？
- InnoDB 认真执行了“把这个值修改成 (1,2)"这个操作，该加锁的加锁，该更新的更新。
- 主要是事务的隔离性


### 2.9 order by的工作原理
- 全字段排序
- rowid排序

#### 2.9.1 全字段排序
- `select name,city,age from t5 where city="hangzhou" order by name limit 100`的执行流程
  - 初始化sort_buffer,确定放入name,city,age这三个字段
  - 从索引city找到对应的主键id
  - 到主键id索引回表查询数据，取name,city,age三个字段，存储到sort_buffer
  - 重复执行索引city, 主键id回表查询
  - 对sort_buffer进行快速排序【可能在内存】
  - 按排序的结果取100行返回给客户端
- sort_buffer
  - sort_buffer【是mysql为每个线程分配一块内存用于排序】
  - 通过 `show variables like '%sort_buffer_size%';`
  - 如果排序的数据量太大，内存放不下，就会利用磁盘临时文件辅助排序
- 怎么判断排序是否使用了临时文件
  - number_of_tmp_files 超过需求排序的数据量的大小
  - sort_buffer_size 越小，需要分成的份数越多，number_of_tmp_files 的值就越大。

#### 2.9.2 rowid排序
- 单行数据很大，返回的数据量很大，一直使用临时文件，
- mysql对单行长度太大的优化
  - 初始化 sort_buffer，确定放入两个字段，即 name 和 id；
  - 从索引 city 找到第一个满足 city='杭州’条件的主键 id，也就是图中的 ID_X；
  - 到主键 id 索引取出整行，取 name、id 这两个字段，存入 sort_buffer 中；
  - 从索引 city 取下一个记录的主键 id；
  - 重复步骤 3、4 直到不满足 city='杭州’条件为止，也就是图中的 ID_Y；
  - 对 sort_buffer 中的数据按照字段 name 进行排序；
  - 遍历排序结果，取前 1000 行，并按照 id 的值回到原表中取出 city、name 和 age 三个字段返回给客户端。

#### 2.9.3 全字段排序 Vs rowid排序
- 选择原理
  - 如果内存足够大，会优先全字段排序，尽量减少磁盘io
  - 如果内存太小，会影响排序效率，才会采用rowid排序算法
  - 对innodb表来说，rowid排序会要求回表多造成磁盘读，因此不会优先选择

#### 2.9.4 如果减少文件排序
- 添加索引(联合索引)，将数据变成有序
- `alter table t5 add index city_name_age(city, name, age);` 联合索引，减少回表查询
- 分析 `using index` 表示使用了覆盖索引，减少了回表查询

#### 2.9.5 思考题
- `select * from t where city in ('杭州',"苏州") order by name limit 100;` 会有排序过程吗
  - 会
  - 杭州的name是排好序，苏州的name也是排好序，但是加起来就不满足排序
- 如果维护一个数据库端排序的方案，不需要额外排序
  - 分别查询，归并排序


### 2.10 如何正确显示随机消息
#### 2.10.1 order by rand() 排序随机取几条数据的背后逻辑
- `select word from words order by rand() limit 3` Using temporary(使用临时表); Using filesort(需要排序)
  - 创建临时表，使用memory引擎，会有两个字段，一个double(存随机数), varchar(存word单词)，表中没有建索引
  - 按主键取word值，调用rand()函数，将这两个值存在临时表中
  - 然后将没有索引的临时表按照随机数排序
  - 初始化sort_buffer, 将临时表的数据扫描添加到sort_buffer中，然后根据随机数排序
  - 排完序之后取前三条结果返回位置信息，这样就可以直接从内存中取出word值
  - `show variables like '%slow_query_log%';` 查看慢日志看扫描的行数

#### 2.10.2 mysql如何定位一行数据？
- 没有主键如何回表
  - 删掉主键，mysql会自己生成一个长度为6字节的rowid来作为主键
  - 排序过程中rowid的来历，标识数据行的信息
    - 没有主键，rowid是系统生成，有主键，rowid就是主键id
    - memory引擎不是索引组织表，可以认为是一个数组，rowid就是数组的下标
- 总结order by rand()
  - 使用了内存临时表
  - 内存临时表排序的时候使用了rowid排序方法

#### 2.10.3 磁盘临时表
- 什么时候会使用
  - 数据超过 tmp_table_size的限制，就会使用磁盘临时表， 默认16M
  - `show variables like '%tmp_table_size%';`
  - 磁盘临时表使用的引擎是innodb, 用`internal_tmp_dist_storage_engine`控制
  - 使用磁盘临时表就是一个没有显示索引的innodb表排序过程
- 为什么临时文件排序使用优先队列排序？而不是归并排序算法
  - 使用优先队列就可以只得到前3个，不需要排所有的
- 优先队列排序算法
  - 对于这 10000 个准备排序的 (R,rowid)，先取前三行，构造成一个堆；
  - 取下一个行 (R’,rowid’)，跟当前堆里面最大的 R 比较，如果 R’小于 R，把这个 (R,rowid) 从堆中去掉，换成 (R’,rowid’)；
  - 重复第 2 步，直到第 10000 个 (R’,rowid’) 完成比较。
  - 总是保证堆顶最大值

#### 2.10.4 如果优化随机排序算法？
- 思路1
  - 取得这个表的主键 id 的最大值 M 和最小值 N;
  - 用随机函数生成一个最大值到最小值之间的数 X = (M-N)*rand() + N;
  - 取不小于 X 的第一个 ID 的行。
  - 问题：空洞问题，会取不到值，概率不均衡
- 思路2
  - 取得整个表的行数，并记为 C
  - 取得 Y = floor(C * rand())。 floor 函数在这里的作用，就是取整数部分。
  - 再用 limit Y,1 取得一行。
  - 注意：mysql处理limit的做法就是按顺序一个一个读出来，丢掉前Y个，然后把一个记录作为返回结果，因此需要扫描y+c+1行，代价相对高一些
- 取三个值的随机算法
  - 取得整个表的行数，记为 C；
  - 根据相同的随机方法得到 Y1、Y2、Y3；
  - 再执行三个 limit Y, 1 语句得到三行数据。
- 思路3
  - `select * from words where id >= (select floor(rand()*(select max(id) from words))) order by id limit 10;`
  - 单词全部取出，使用redis缓存排序来处理业务，尽量不要让数据库做业务逻辑
  - 可以重建索引解决空洞问题

### 2.11 为什么mysql逻辑相同，性能差异这么大？
#### 2.11.1 条件字段函数操作
- `select count(*) from trade_log where month(t_modified)=7` t_modified上有索引
  - 不会走索引
  - `select * from words where id-1=999;` 索引失效

#### 2.11.2 隐式类型转换
- `select "10">9` 如果等于1,说明字符串转成了int, 否则int转成了字符串
- `explain select * from tradelog where tradeid=3423422` tradeid有索引且类型为varchar(32)
  - 会发生隐式类型转换，这样索引就会失效就会扫描全表
  - 对优化器执行的代码如下 `select * from tradelog where CAST(tradid AS signed int) = 110717;`
  - 也就是对索引字段操作，优化器会放弃走树搜索功能

#### 2.11.3 隐式字符编码转换
  - 关联表，一个是utf8, 一个是utf8mb4
  - `select * from trade_detail where tradeid=$L2.tradeid.value;`
    - 实际触发隐式转换 `select * from trade_detail  where CONVERT(traideid USING utf8mb4)=$L2.tradeid.value; `
    - 连接过程中要求在被驱动表的索引字段上加函数操作，直接导致全表查询
  - 如果解决字符串转换问题
    - 将表的字符集修改成相同

#### 2.11.4 另外慢查询场景
- 100万数据的表，有一个b字段，int(10), 其中等于1234567890的数据有10万条
  - `select * from t where b ="1234567890abcd";` 这个查询的执行流程，会走索引吗？
  - 在传给引擎执行的时候，做了字符截断。因为引擎里面这个行只定义了长度是 10，所以只截了前 10 个字节，就是’1234567890’进去做匹配；
  - 这样满足条件的数据有 10 万行；
  - 因为是 select *， 所以要做 10 万次回表；
  - 但是每次回表以后查出整行，到 server 层一判断，b 的值都不是’1234567890abcd’;
  - 返回结果是空。

#### 2.11.4 总结
- 对索引字段做函数操作，可能破坏索引值的有序性，因此优化器就决定放弃走树搜索功能
- 业务代码升级，执行explain是一个很好的习惯


### 2.12 为什么查一行语句，也执行这么慢？
#### 2.12.1 等到MDL锁
- 查询长时间不返回
  - `show processlist` 查看waiting for table metadata lock
  - 一个线程正在表t上请求或持有MDL写锁，把select语句堵住了

#### 2.12.2 等flush
- **flush表的动作**
  - `flush tables words with read lock;` 和 `flush tables with read lock;`
  - flush表动作->关闭已打开的表对象，同时将查询缓存中的结果清空（会等到所有正在运行的sql请求）
- 产生阻塞的流程
  - `select sleep(1) from words;` 默认要执行10万秒
  - `flush tables words` 需要关闭查询对象，就需要等待上一个结束，因此会阻塞
  - `select * from words where id =1` 也会被阻塞，被flush命令阻塞
  - `show processlist` 会出现 waiting for table flush

#### 2.12.3 等行锁
- 场景复现
  - session A `begin; update words set word="adad" where id =1;`
  - session B `select * from words where id=1 lock in share mode;`
  - session A启动事务，占用写锁，但未提交
  - session B查询需要获取读锁
- 如果解决
  - 找到对应的查询线程, 断开这个链接
  - 链接被断开的时候，会自动回滚这个连接里面正在执行的线程，也就释放了id=1的行锁
  - `show processlist; kill query processlist_id;`

#### 2.12.4 慢查询
- 没有索引，只是一条条扫描
  - `select * from t where c=50000 limit 1;`
  - 可以看慢查询扫描的行数
  - 坏查询不一定是慢查询，线上一般超过1ms都不是好查询
- 回滚日志太多，导致查询时间差的问题
  - session B 执行10万次 `update words set c=c+1 where id=1;`
  - session A `select * from words where id =1;` 会执行10次的回滚日志，所以比较耗时 
  - session A `select * from words where id =1 lock in share mode;`
  - `lock in share mode` 表示当前读，会直接读到结果不需要执行回滚日志

#### 2.12.5 总结
- 概念
  - 表锁
  - 行锁
  - 一致性读 `lock in share mode`
- 思考题？
  - `select * from t where c= 5 for update;` 会等待行锁释放之后，返回查询结果
  - `select * from t where c=5 for update nowait` 不等待，直接提示锁冲突，不返回结果
  - `select * from t where c=5 for update wait 5` 等待5秒，如果行锁仍未释放，则提示行锁冲突，不返回结果
  - `select * from t where c=5 for update skip locked` 直接返回结果，忽略行锁记录


### 2.13 幻读是什么？幻读有什么问题？
#### 2.13.1 幻读是什么？
- 一个事务在前后两次查询同一范围的时候，后一次查询看到前一次查询没有看到的行
  - 在rr级别下，普通查询时快照读，看不到别的事务插入的数据，因此，幻读在"当前读"才会出现
  - 修改不能称为幻读，幻读专指"新插入的行"

#### 2.13.2 幻读有什么问题？

#### 2.13.3 如何解决幻读？
- 间隙锁（Gap lock）只有在rr隔离级别下才生效
  - 产生幻读的原因是，行锁只能锁住行
  - 新插入记录这个动作，要更新的事记录之间的"间隙"
  - 锁的就是两个值之间的空隙
- `select * from t where d=5 for update`背后的逻辑
  - 不光锁住数据库已有的记录(n)个行锁
  - 还同时增加(n+1)个间隙锁，这样确保无法再插入新的记录
- 间隙锁和行锁合成 next-key lock
  - next-key lock 是一个前开后闭的区间
  - 间隙锁为开区间，next-key lock为前开后闭区间

#### 2.13.4 锁的冲突关系
- 读锁/写锁
  - 读锁与写锁冲突
  - 写锁与写锁冲突
  - 读锁与读锁不冲突
- 间隙锁(gap lock)
  - 间隙锁之间不存在冲突关系
  - 间隙锁与往这个间隙插入一个记录存在冲突关系


#### 2.13.5 思考
- rc模式，binlog_format=row组合为啥要这样用
  - row模式下保存的是每一行的前后记录，虽然占空间，但是不会因为保存命令而造成幻读
- 即使给所有的行加锁也解决不了幻读的影响
- 行锁确实比较直观，判断规则也相对简单，间隙锁的引入会影响系统的并发度，也增加了锁分析的复杂度


### 2.14 为什么我只改变一行，锁这么多？
#### 2.14.1 rr级别下，加锁的规则（两原则，两优化，一bug）
- 5.x-5.7.24, 8.0-8.0.13版本中的情况
  - 原则1: 加锁的基本单位是next-key lock, 是前开后闭区间
  - 原则2: 查找过程中访问到的对象才会加锁
  - 优化1: 索引上的等值查询，给唯一索引加锁时，next-key lock退化为行锁
  - 优化2: 索引上的等值查询，向右遍历时且最后一个值不满足等值条件的时候，next-key lock 退化为间隙锁
  - 一个bug: 唯一索引上的范围查询会访问到不满足条件的第一个值为止

#### 2.14.2 各种案例分析
- 等值查询间隙锁
![](images/幻读/1.png)

- 非唯一索引等值锁
![](images/幻读/2.png)

- 主键索引范围锁
  ![](images/幻读/3.png)

- 非唯一索引范围锁
  ![](images/幻读/4.png)

- 唯一索引范围锁bug
  ![](images/幻读/5.png)

- 非唯一索引上存在"等值"的例子
 ![](images/幻读/6.png)

- limit语句加锁
  ![](images/幻读/7.png)

- 一个死锁的例子
  ![](images/幻读/8.png)

### 2.15 mysql有哪些"饮鸩止渴"提高性能的方法？
#### 2.15.1 短连接风暴
- 短连接模型存在的风险
  - 一旦数据库处理的慢一些，连接数就会暴涨
  - 通过参数max_connections来查看最大的连接数
  - 查看现有的mysql连接 `select * from information_schema.processlist;`
  - `kill id;` 即可删掉连接
- 处理的方式
  - 第一种：先处理掉那些占着连接但是不工作的线程 `kill id`
  - 第二种：减少连接过程的消耗
    - 短时间大量的连接申请，跳过权限验证阶段, 重启加上参数 --skip-grant-tables
    - 开启--skip-grant-tables参数，mysql默认会开启skip-networking参数，只能被本地客户端连接

#### 2.15.2 慢查询的性能问题
- 索引没有设计好
  - 修改索引执行 alter table语句
  - 理解做法： 
    - 从库关闭 set sql_log_bin=off, 然后alter table添加索引
    - 切换主从设备
    - 在原来的主库执行set sql_log_bin=off, 修改alter table添加索引
- sql语句没有写好
  - `select * from t where id+1=1000;` 修改sql
- mysql选错索引
  - force index 强制指定索引
  
#### 2.15.3 如何提前发现索引没设计好和语句没有写好
- 测试环境开启慢sql, 设置 long_query_time=0, 确保每个语句记录都写到慢日志
- 测试表中模拟线上做回归测试
- 观察慢日志，留意rows_examined字段师傅与预期一致
- [慢日志分析工具](https://www.percona.com/doc/percona-toolkit/3.0/pt-query-digest.html)

#### 2.15.4 Qps突增问题
- 业务突然出现高峰，或者应用程序bug, 导致某个语句qps突然暴涨
  - 一种由全新业务的bug导致的，直接去掉这个功能
  - 如果新业务使用新的数据库，直接删掉账号，断开现有的连接
  - 如果这个功能跟现有的主体部署一起，只能通过处理语句来限制


### 2.16 mysql怎么保证数据不丢失
- 重要结论
  - 只要redo log和binlog保证持久化到磁盘，就能确保mysql异常重启后，数据可以恢复

#### 2.16.1 binlog的写入机制
- 写入逻辑
  - 事务执行中，先把日志写入到binlog cache, 
  - 事务提交的时候，将binlog cache写入到binlog磁盘文件中
  - 一个事务的binlog不能被拆分，无论事务多大，都要保证一次性写入
- binlog cache
  - 每个线程一片内存，`show variables like '%binlog_cache_size%'` 查看大小
  - 如果超过这个大小就要暂存到磁盘中
  - 事务提交将binlog cache写入文件中，会清空binlog cache的缓存

#### 2.16.2 binlog写盘分析
![](./images/binlog/binlog_write_disk.png)
- 分析
  - 图中的 write，指的就是指把日志写入到文件系统的 page cache，并没有把数据持久化到磁盘，所以速度比较快。
  - 图中的 fsync，才是将数据持久化到磁盘的操作。一般情况下，我们认为 fsync 才占磁盘的 IOPS。
- write 和 fsync的时机，由参数 `show variables like '%sync_binlog%';` 控制
  - 0，表示每次提交事务都只write， 不fsync
  - 1，表示每次提交事务都会 fsync
  - N(n>1) 表示累积N个事务提交之后，才fsync
- 性能优化
  - sync_binlog设置一个比较大的值，可以提升性能，解决一定的io问题
  - 一般在100-1000之间
  - 对应的风险：如果主机异常重启，会丢失最近N个事务的binlog日志

#### 2.16.3 redo log的写入机制


