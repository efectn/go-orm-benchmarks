# Results

### orm-benchmark (with no flags)

   200 times - Insert
  raw_stmt:     0.37s      1872632 ns/op     692 B/op     14 allocs/op
 beego_orm:     0.38s      1923870 ns/op    2405 B/op     56 allocs/op
       raw:     0.39s      1938406 ns/op    1130 B/op     16 allocs/op
      gorm:     0.39s      1955777 ns/op    6960 B/op     87 allocs/op
       bun:     0.40s      1994017 ns/op    1290 B/op     13 allocs/op
        pg:     0.40s      2014027 ns/op    7175 B/op     13 allocs/op
      xorm:     0.56s      2786996 ns/op    3511 B/op     96 allocs/op

   200 times - MultiInsert 100 row
       raw:     0.98s      4877568 ns/op  136605 B/op    923 allocs/op
  raw_stmt:     1.11s      5548251 ns/op  132549 B/op    923 allocs/op
        pg:     1.14s      5678735 ns/op    8696 B/op    114 allocs/op
 beego_orm:     1.16s      5796237 ns/op  179705 B/op   2746 allocs/op
       bun:     1.17s      5845266 ns/op    4488 B/op    213 allocs/op
      gorm:     1.48s      7406033 ns/op  293428 B/op   3727 allocs/op
      xorm:     1.90s      9508433 ns/op  285784 B/op   7422 allocs/op

   200 times - Update
  raw_stmt:     0.13s       652894 ns/op     809 B/op     14 allocs/op
       raw:     0.13s       667332 ns/op     792 B/op     13 allocs/op
      gorm:     0.38s      1917723 ns/op    6625 B/op     81 allocs/op
       bun:     0.39s      1927999 ns/op     604 B/op      4 allocs/op
        pg:     0.40s      2013945 ns/op     896 B/op     11 allocs/op
 beego_orm:     0.46s      2288381 ns/op    1803 B/op     47 allocs/op
      xorm:     0.52s      2591389 ns/op    2993 B/op    119 allocs/op

   200 times - Read
  raw_stmt:     0.13s       665066 ns/op    2267 B/op     50 allocs/op
       raw:     0.13s       670237 ns/op    2240 B/op     49 allocs/op
       bun:     0.14s       692506 ns/op    1330 B/op     18 allocs/op
        pg:     0.15s       726730 ns/op    1001 B/op     22 allocs/op
      gorm:     0.15s       727872 ns/op    5235 B/op    108 allocs/op
 beego_orm:     0.15s       770235 ns/op    2112 B/op     75 allocs/op
      xorm:     0.28s      1414234 ns/op    8319 B/op    237 allocs/op

   200 times - MultiRead limit 100
       raw:     0.18s       894095 ns/op   38358 B/op   1037 allocs/op
  raw_stmt:     0.18s       910734 ns/op   38390 B/op   1038 allocs/op
        pg:     0.20s       995147 ns/op   22406 B/op    631 allocs/op
       bun:     0.20s      1021603 ns/op   28735 B/op   1116 allocs/op
 beego_orm:     0.21s      1074391 ns/op   55257 B/op   3077 allocs/op
      gorm:     0.26s      1284915 ns/op   71716 B/op   3877 allocs/op
      xorm:     doesn't work

### orm-benchmark -multi=5

  1000 times - Insert
  raw_stmt:     1.90s      1899747 ns/op     713 B/op     14 allocs/op
 beego_orm:     1.90s      1904115 ns/op    2408 B/op     56 allocs/op
       raw:     1.92s      1917639 ns/op     763 B/op     13 allocs/op
       bun:     1.97s      1972166 ns/op     973 B/op     12 allocs/op
        pg:     1.98s      1976075 ns/op    2172 B/op     12 allocs/op
      gorm:     2.01s      2011269 ns/op    6699 B/op     87 allocs/op
      xorm:     2.62s      2619231 ns/op    3126 B/op     94 allocs/op

  1000 times - MultiInsert 100 row
       raw:     5.13s      5126197 ns/op  135378 B/op    917 allocs/op
  raw_stmt:     5.28s      5275308 ns/op  131303 B/op    917 allocs/op
 beego_orm:     5.82s      5815242 ns/op  179814 B/op   2746 allocs/op
       bun:     5.91s      5906738 ns/op    4275 B/op    214 allocs/op
        pg:     5.97s      5969001 ns/op    5619 B/op    114 allocs/op
      gorm:     8.28s      8283438 ns/op  293822 B/op   3728 allocs/op
      xorm:     8.69s      8685525 ns/op  285741 B/op   7421 allocs/op

  1000 times - Update
  raw_stmt:     0.66s       658758 ns/op     795 B/op     14 allocs/op
       raw:     0.67s       672173 ns/op     779 B/op     13 allocs/op
       bun:     1.91s      1911455 ns/op     592 B/op      4 allocs/op
        pg:     1.95s      1953156 ns/op    2020 B/op     11 allocs/op
 beego_orm:     1.97s      1973964 ns/op    1802 B/op     47 allocs/op
      gorm:     1.98s      1975183 ns/op    6600 B/op     81 allocs/op
      xorm:     2.67s      2665264 ns/op    2992 B/op    119 allocs/op

  1000 times - Read
 beego_orm:     0.65s       653197 ns/op    2105 B/op     75 allocs/op
       raw:     0.66s       661987 ns/op    2080 B/op     49 allocs/op
  raw_stmt:     0.66s       662631 ns/op    2111 B/op     50 allocs/op
       bun:     0.70s       700160 ns/op    1305 B/op     18 allocs/op
        pg:     0.72s       724049 ns/op    1002 B/op     22 allocs/op
      gorm:     0.79s       785298 ns/op    5238 B/op    108 allocs/op
      xorm:     1.40s      1403249 ns/op    8318 B/op    237 allocs/op

  1000 times - MultiRead limit 100
       raw:     0.88s       880073 ns/op   38364 B/op   1037 allocs/op
  raw_stmt:     0.89s       887865 ns/op   38396 B/op   1038 allocs/op
        pg:     0.98s       984238 ns/op   26677 B/op    631 allocs/op
       bun:     1.04s      1035875 ns/op   28787 B/op   1116 allocs/op
 beego_orm:     1.05s      1048641 ns/op   55252 B/op   3077 allocs/op
      gorm:     1.34s      1337241 ns/op   71642 B/op   3877 allocs/op
      xorm:     doesn't work

### orm-benchmark -multi=20

  4000 times - Insert
       raw:     7.46s      1866171 ns/op     719 B/op     13 allocs/op
  raw_stmt:     7.49s      1871428 ns/op     718 B/op     14 allocs/op
        pg:     7.79s      1947092 ns/op    1253 B/op     12 allocs/op
      gorm:     7.87s      1967935 ns/op    6654 B/op     87 allocs/op
       bun:     8.26s      2065789 ns/op     920 B/op     13 allocs/op
 beego_orm:     8.33s      2083489 ns/op    2409 B/op     56 allocs/op
      xorm:    10.78s      2693959 ns/op    3049 B/op     94 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    20.62s      5154921 ns/op  131076 B/op    916 allocs/op
 beego_orm:    22.67s      5667491 ns/op  179771 B/op   2746 allocs/op
       raw:    23.54s      5884413 ns/op  135154 B/op    916 allocs/op
       bun:    23.84s      5959424 ns/op    4263 B/op    214 allocs/op
        pg:    24.82s      6204737 ns/op    4511 B/op    114 allocs/op
      gorm:    31.76s      7939882 ns/op  293955 B/op   3729 allocs/op
      xorm:    33.74s      8435829 ns/op  285748 B/op   7421 allocs/op

  4000 times - Update
  raw_stmt:     2.61s       651574 ns/op     773 B/op     14 allocs/op
       raw:     2.62s       655903 ns/op     757 B/op     13 allocs/op
       bun:     7.66s      1915297 ns/op     584 B/op      4 allocs/op
        pg:     7.85s      1962301 ns/op     896 B/op     11 allocs/op
      gorm:     8.03s      2008527 ns/op    6604 B/op     81 allocs/op
 beego_orm:     8.08s      2019757 ns/op    1801 B/op     47 allocs/op
      xorm:    10.28s      2571165 ns/op    2993 B/op    119 allocs/op

  4000 times - Read
       raw:     2.59s       646737 ns/op    2080 B/op     49 allocs/op
  raw_stmt:     2.62s       654991 ns/op    2112 B/op     50 allocs/op
       bun:     2.74s       685250 ns/op    1307 B/op     18 allocs/op
        pg:     2.80s       700221 ns/op    1281 B/op     22 allocs/op
 beego_orm:     2.82s       704028 ns/op    2105 B/op     75 allocs/op
      gorm:     2.92s       729225 ns/op    5240 B/op    108 allocs/op
      xorm:     5.66s      1415778 ns/op    8318 B/op    237 allocs/op

  4000 times - MultiRead limit 100
       raw:     3.54s       886003 ns/op   38356 B/op   1037 allocs/op
  raw_stmt:     3.63s       906270 ns/op   38388 B/op   1038 allocs/op
        pg:     3.85s       963602 ns/op   23764 B/op    631 allocs/op
       bun:     4.02s      1005727 ns/op   28820 B/op   1116 allocs/op
 beego_orm:     4.37s      1092503 ns/op   55248 B/op   3077 allocs/op
      gorm:     5.12s      1280190 ns/op   71628 B/op   3877 allocs/op
      xorm:     doesn't work

### orm-benchmark -multi=50

 10000 times - Insert
       raw:    18.82s      1882463 ns/op     713 B/op     13 allocs/op
  raw_stmt:    18.87s      1886864 ns/op     720 B/op     14 allocs/op
       bun:    19.60s      1959859 ns/op     907 B/op     13 allocs/op
      gorm:    19.67s      1967267 ns/op    6655 B/op     88 allocs/op
        pg:    19.74s      1974130 ns/op    1055 B/op     12 allocs/op
 beego_orm:    19.85s      1984567 ns/op    2409 B/op     56 allocs/op
      xorm:    25.89s      2589262 ns/op    3035 B/op     94 allocs/op

 10000 times - MultiInsert 100 row
  raw_stmt:    48.24s      4824085 ns/op  131034 B/op    916 allocs/op
       raw:    48.92s      4891580 ns/op  135114 B/op    916 allocs/op
 beego_orm:    56.10s      5610480 ns/op  179821 B/op   2746 allocs/op
       bun:    59.21s      5921031 ns/op    4245 B/op    214 allocs/op
        pg:    60.79s      6078986 ns/op    4396 B/op    114 allocs/op
      gorm:    72.20s      7219802 ns/op  293984 B/op   3729 allocs/op
      xorm:    81.02s      8102211 ns/op  285707 B/op   7421 allocs/op

 10000 times - Update
  raw_stmt:     6.44s       643858 ns/op     768 B/op     14 allocs/op
       raw:     6.56s       656429 ns/op     753 B/op     13 allocs/op
 beego_orm:    19.42s      1942096 ns/op    1801 B/op     47 allocs/op
       bun:    19.54s      1953862 ns/op     585 B/op      4 allocs/op
      gorm:    19.72s      1971955 ns/op    6604 B/op     81 allocs/op
        pg:    19.76s      1975545 ns/op    1121 B/op     11 allocs/op
      xorm:    26.50s      2649894 ns/op    2993 B/op    119 allocs/op

 10000 times - Read
 beego_orm:     6.56s       655910 ns/op    2105 B/op     75 allocs/op
  raw_stmt:     6.57s       657027 ns/op    2110 B/op     50 allocs/op
       raw:     6.61s       661139 ns/op    2079 B/op     49 allocs/op
       bun:     6.87s       687084 ns/op    1307 B/op     18 allocs/op
        pg:     7.09s       708644 ns/op    1225 B/op     22 allocs/op
      gorm:     7.29s       728913 ns/op    5242 B/op    108 allocs/op
      xorm:    14.00s      1400469 ns/op    8318 B/op    237 allocs/op

 10000 times - MultiRead limit 100
       raw:     8.80s       880235 ns/op   38357 B/op   1037 allocs/op
  raw_stmt:     8.80s       880457 ns/op   38389 B/op   1038 allocs/op
        pg:     9.71s       970842 ns/op   24755 B/op    631 allocs/op
       bun:    10.13s      1012692 ns/op   28802 B/op   1116 allocs/op
 beego_orm:    10.52s      1051815 ns/op   55250 B/op   3077 allocs/op
      gorm:    13.09s      1309194 ns/op   71629 B/op   3877 allocs/op
      xorm:     doesn't work
