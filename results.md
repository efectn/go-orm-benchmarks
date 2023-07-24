# Results

- orm-benchmark (with no flags)
```
Reports:

InsertMulti
pgx_pool:       438     2971808 ns/op   113168 B/op     44 allocs/op
raw:            344     3032535 ns/op   184243 B/op     933 allocs/op
beego:          414     3109230 ns/op   177810 B/op     2745 allocs/op
gorm_prep:      393     3115722 ns/op   251113 B/op     1890 allocs/op
reform:         394     3149396 ns/op   459052 B/op     2746 allocs/op
ent:            349     3380779 ns/op   386646 B/op     4600 allocs/op
pg:             349     3508386 ns/op   6539 B/op       112 allocs/op
gen:            338     3624859 ns/op   304809 B/op     5359 allocs/op
sqlx:           343     3631249 ns/op   171548 B/op     1551 allocs/op
bun:            370     3664276 ns/op   42671 B/op      219 allocs/op
rel:            312     3732568 ns/op   306901 B/op     3265 allocs/op
upper:          282     3883268 ns/op   325569 B/op     4207 allocs/op
zorm:           308     3951400 ns/op   199929 B/op     2780 allocs/op
gorm:           238     4384769 ns/op   291501 B/op     5231 allocs/op
xorm:           224     4488046 ns/op   248461 B/op     5415 allocs/op
godb:           214     4706725 ns/op   254143 B/op     5894 allocs/op
pgx:            358     6817979 ns/op   113130 B/op     43 allocs/op
pop:            bulk-insert is not supported
sqlboiler:      bulk-insert is not supported
sqlc:           bulk-insert is not supported
gorp:           bulk-insert is not supported
dbr:            bulk-insert is not supported

Update
sqlc:           13990   84128 ns/op     877 B/op        14 allocs/op
raw:            14208   84819 ns/op     749 B/op        13 allocs/op
sqlx:           5720    178812 ns/op    872 B/op        20 allocs/op
reform:         745     1637956 ns/op   1776 B/op       51 allocs/op
sqlboiler:      732     1682801 ns/op   905 B/op        17 allocs/op
dbr:            758     1688043 ns/op   2651 B/op       57 allocs/op
beego:          698     1695577 ns/op   1752 B/op       47 allocs/op
pg:             672     1717302 ns/op   768 B/op        9 allocs/op
gorp:           672     1757154 ns/op   1212 B/op       32 allocs/op
gen:            668     1759213 ns/op   13472 B/op      166 allocs/op
gorm_prep:      732     1771883 ns/op   5008 B/op       56 allocs/op
pop:            685     1775140 ns/op   6045 B/op       186 allocs/op
pgx_pool:       728     1779697 ns/op   290 B/op        10 allocs/op
zorm:           690     1790709 ns/op   3024 B/op       59 allocs/op
ent:            735     1822547 ns/op   4681 B/op       97 allocs/op
bun:            720     1861019 ns/op   4731 B/op       5 allocs/op
rel:            656     1869423 ns/op   3048 B/op       45 allocs/op
pgx:            588     2102554 ns/op   267 B/op        10 allocs/op
gorm:           565     2163181 ns/op   6752 B/op       99 allocs/op
upper:          500     2296819 ns/op   16821 B/op      396 allocs/op
xorm:           432     2320718 ns/op   3944 B/op       132 allocs/op
godb:           544     2332299 ns/op   5112 B/op       154 allocs/op

Read
pgx:            14715   82206 ns/op     893 B/op        8 allocs/op
pgx_pool:       14556   83208 ns/op     1079 B/op       9 allocs/op
raw:            13690   86875 ns/op     2061 B/op       50 allocs/op
sqlc:           13342   87095 ns/op     2077 B/op       51 allocs/op
reform:         13088   91578 ns/op     3200 B/op       86 allocs/op
pop:            12604   93116 ns/op     3246 B/op       67 allocs/op
ent:            11398   99883 ns/op     5605 B/op       144 allocs/op
gorp:           9962    100944 ns/op    3877 B/op       194 allocs/op
beego:          13562   104061 ns/op    2098 B/op       76 allocs/op
pg:             10000   106387 ns/op    977 B/op        20 allocs/op
sqlboiler:      9590    111210 ns/op    957 B/op        14 allocs/op
bun:            10000   111689 ns/op    5824 B/op       39 allocs/op
dbr:            9127    112336 ns/op    2184 B/op       37 allocs/op
rel:            9196    116003 ns/op    2304 B/op       47 allocs/op
zorm:           9786    122160 ns/op    3016 B/op       64 allocs/op
gorm_prep:      8834    126717 ns/op    4411 B/op       87 allocs/op
gorm:           7953    150596 ns/op    4781 B/op       98 allocs/op
gen:            6963    160247 ns/op    10413 B/op      154 allocs/op
sqlx:           5980    182850 ns/op    1976 B/op       43 allocs/op
godb:           5841    190702 ns/op    4067 B/op       102 allocs/op
xorm:           5500    193652 ns/op    4651 B/op       127 allocs/op
upper:          5295    227699 ns/op    5106 B/op       112 allocs/op

ReadSlice
reform:         13243   93004 ns/op     4512 B/op       112 allocs/op
pgx:            7438    148692 ns/op    42949 B/op      504 allocs/op
pgx_pool:       6883    149681 ns/op    43007 B/op      504 allocs/op
raw:            6092    191391 ns/op    38341 B/op      1038 allocs/op
sqlc:           5407    217986 ns/op    62661 B/op      1150 allocs/op
upper:          4860    235242 ns/op    4843 B/op       92 allocs/op
pg:             5050    240542 ns/op    23520 B/op      629 allocs/op
sqlx:           4620    261532 ns/op    37480 B/op      1225 allocs/op
gorp:           4444    264584 ns/op    57404 B/op      1494 allocs/op
pop:            3909    271675 ns/op    69408 B/op      1307 allocs/op
ent:            3584    285559 ns/op    77190 B/op      2035 allocs/op
dbr:            3541    292902 ns/op    30816 B/op      1254 allocs/op
bun:            3669    324084 ns/op    34081 B/op      1124 allocs/op
sqlboiler:      3082    340183 ns/op    67173 B/op      2260 allocs/op
gorm_prep:      3258    340854 ns/op    43235 B/op      2081 allocs/op
beego:          3136    378430 ns/op    55225 B/op      3078 allocs/op
gorm:           2839    400370 ns/op    44398 B/op      2191 allocs/op
rel:            2446    470933 ns/op    100641 B/op     2253 allocs/op
zorm:           2532    470986 ns/op    161618 B/op     2949 allocs/op
godb:           2271    479831 ns/op    75260 B/op      3084 allocs/op
gen:            2096    503572 ns/op    50047 B/op      2247 allocs/op
xorm:           2091    554886 ns/op    119434 B/op     4401 allocs/op

Insert
reform:         723     1667341 ns/op   1781 B/op       51 allocs/op
ent:            603     1669617 ns/op   4133 B/op       97 allocs/op
sqlboiler:      711     1690820 ns/op   1569 B/op       34 allocs/op
raw:            753     1691616 ns/op   704 B/op        13 allocs/op
bun:            691     1720228 ns/op   5005 B/op       13 allocs/op
gorm_prep:      648     1730891 ns/op   5161 B/op       65 allocs/op
dbr:            717     1756766 ns/op   2688 B/op       65 allocs/op
gorp:           625     1766369 ns/op   1772 B/op       41 allocs/op
beego:          720     1767656 ns/op   2385 B/op       57 allocs/op
zorm:           672     1790357 ns/op   3780 B/op       77 allocs/op
pgx_pool:       691     1795782 ns/op   290 B/op        10 allocs/op
sqlx:           614     1807225 ns/op   856 B/op        19 allocs/op
pg:             652     1829659 ns/op   2520 B/op       10 allocs/op
rel:            687     1862237 ns/op   2604 B/op       45 allocs/op
gen:            643     1884134 ns/op   10123 B/op      135 allocs/op
upper:          655     1957702 ns/op   5939 B/op       127 allocs/op
pop:            622     2034033 ns/op   9553 B/op       238 allocs/op
xorm:           547     2289864 ns/op   3323 B/op       90 allocs/op
gorm:           570     2337354 ns/op   7161 B/op       105 allocs/op
godb:           564     2356383 ns/op   4507 B/op       115 allocs/op
sqlc:           552     2509708 ns/op   2760 B/op       61 allocs/op
pgx:            626     2637059 ns/op   266 B/op        10 allocs/op
```
