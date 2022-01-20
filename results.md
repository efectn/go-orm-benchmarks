# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
      sqlc:     0.46s      2295978 ns/op    3292 B/op     63 allocs/op
 sqlboiler:     0.46s      2318850 ns/op    1987 B/op     36 allocs/op
        pg:     0.54s      2701578 ns/op    6278 B/op     10 allocs/op
     beego:     0.60s      3008184 ns/op    2378 B/op     56 allocs/op
    reform:     0.62s      3088464 ns/op    2196 B/op     54 allocs/op
       ent:     0.62s      3091174 ns/op    4177 B/op     98 allocs/op
       dbr:     0.62s      3099481 ns/op    3086 B/op     66 allocs/op
       bun:     0.63s      3152326 ns/op    5329 B/op     14 allocs/op
      godb:     0.64s      3195663 ns/op    4962 B/op    116 allocs/op
       rel:     0.64s      3219935 ns/op    2430 B/op     41 allocs/op
       raw:     0.66s      3284551 ns/op    1035 B/op     16 allocs/op
      gorm:     0.69s      3437694 ns/op    7094 B/op     91 allocs/op
      gorp:     0.70s      3507176 ns/op    2186 B/op     42 allocs/op
     upper:     0.70s      3520865 ns/op   14774 B/op    694 allocs/op
       pop:     0.71s      3529825 ns/op   10578 B/op    249 allocs/op
      xorm:     0.89s      4447943 ns/op    3051 B/op     78 allocs/op

   200 times - MultiInsert 100 row
     beego:     0.77s      3834880 ns/op  179910 B/op   2747 allocs/op
       raw:     0.82s      4100560 ns/op  192355 B/op    938 allocs/op
    reform:     0.90s      4483683 ns/op  466616 B/op   2749 allocs/op
       ent:     0.94s      4679457 ns/op  412429 B/op   4901 allocs/op
       bun:     0.98s      4901280 ns/op   42791 B/op    218 allocs/op
        pg:     0.99s      4940827 ns/op    8932 B/op    112 allocs/op
      godb:     1.01s      5046343 ns/op  259974 B/op   5893 allocs/op
       rel:     1.07s      5351714 ns/op  289893 B/op   3361 allocs/op
      gorm:     1.28s      6394514 ns/op  294718 B/op   4032 allocs/op
     upper:     1.36s      6802582 ns/op  547944 B/op  19623 allocs/op
      xorm:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

   200 times - Update
       raw:     0.02s       105174 ns/op     722 B/op     13 allocs/op
       ent:     0.02s       123426 ns/op    4530 B/op     99 allocs/op
    reform:     0.60s      2997694 ns/op    1748 B/op     51 allocs/op
      gorp:     0.60s      3013140 ns/op    1223 B/op     32 allocs/op
 sqlboiler:     0.61s      3052671 ns/op     924 B/op     17 allocs/op
      gorm:     0.61s      3065857 ns/op    6625 B/op     81 allocs/op
     beego:     0.62s      3091212 ns/op    1761 B/op     47 allocs/op
       dbr:     0.62s      3105721 ns/op    2651 B/op     57 allocs/op
        pg:     0.63s      3126072 ns/op     768 B/op      9 allocs/op
       bun:     0.63s      3140933 ns/op    4733 B/op      5 allocs/op
       pop:     0.63s      3148729 ns/op    6563 B/op    198 allocs/op
       rel:     0.64s      3215577 ns/op    2387 B/op     40 allocs/op
      xorm:     0.65s      3234338 ns/op    2875 B/op    104 allocs/op
      godb:     0.65s      3269060 ns/op    5161 B/op    154 allocs/op
     upper:     0.81s      4026101 ns/op   35310 B/op   1530 allocs/op
      sqlc:     0.90s      4502990 ns/op     891 B/op     14 allocs/op

   200 times - Read
    reform:     0.02s       100179 ns/op    3237 B/op     85 allocs/op
       raw:     0.02s       101687 ns/op    2186 B/op     49 allocs/op
      sqlc:     0.02s       110313 ns/op    2168 B/op     51 allocs/op
     beego:     0.02s       117642 ns/op    2093 B/op     75 allocs/op
       pop:     0.02s       122990 ns/op    3735 B/op     70 allocs/op
        pg:     0.03s       129428 ns/op     873 B/op     20 allocs/op
      gorp:     0.03s       130331 ns/op    3875 B/op    193 allocs/op
       bun:     0.03s       144827 ns/op    5488 B/op     21 allocs/op
 sqlboiler:     0.03s       147770 ns/op    1118 B/op     14 allocs/op
       rel:     0.03s       149384 ns/op    1824 B/op     47 allocs/op
       ent:     0.03s       150209 ns/op    5428 B/op    148 allocs/op
       dbr:     0.03s       157330 ns/op    2184 B/op     37 allocs/op
      gorm:     0.04s       197803 ns/op    5298 B/op    109 allocs/op
      godb:     0.06s       280118 ns/op    4112 B/op    102 allocs/op
     upper:     0.06s       299603 ns/op    8675 B/op    324 allocs/op
      xorm:     0.07s       354736 ns/op    8752 B/op    215 allocs/op

   200 times - MultiRead limit 100
    reform:     0.02s       116330 ns/op    2961 B/op     73 allocs/op
      xorm:     0.04s       194957 ns/op    3296 B/op     77 allocs/op
       raw:     0.06s       313028 ns/op   38359 B/op   1037 allocs/op
     upper:     0.06s       314719 ns/op    8430 B/op    303 allocs/op
      sqlc:     0.07s       373799 ns/op   73134 B/op   1250 allocs/op
        pg:     0.08s       393843 ns/op   22263 B/op    629 allocs/op
 sqlboiler:     0.08s       407842 ns/op   58598 B/op   1259 allocs/op
      gorp:     0.08s       420988 ns/op   57390 B/op   1493 allocs/op
       ent:     0.09s       436852 ns/op   76042 B/op   2038 allocs/op
     beego:     0.10s       512406 ns/op   55208 B/op   3077 allocs/op
       dbr:     0.10s       514976 ns/op   32416 B/op   1545 allocs/op
       bun:     0.11s       531101 ns/op   32914 B/op   1118 allocs/op
       pop:     0.11s       543156 ns/op   77100 B/op   1510 allocs/op
       rel:     0.15s       767541 ns/op   95352 B/op   2250 allocs/op
      godb:     0.16s       810343 ns/op   75269 B/op   3084 allocs/op
      gorm:     0.19s       962779 ns/op   71716 B/op   3877 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
    reform:     2.80s      2804790 ns/op    1860 B/op     51 allocs/op
     beego:     2.82s      2824471 ns/op    2375 B/op     56 allocs/op
       dbr:     2.92s      2915986 ns/op    2768 B/op     65 allocs/op
       raw:     2.97s      2968879 ns/op     770 B/op     13 allocs/op
      xorm:     2.98s      2981423 ns/op    2709 B/op     77 allocs/op
 sqlboiler:     3.00s      3002842 ns/op    1650 B/op     34 allocs/op
      sqlc:     3.01s      3013386 ns/op    2955 B/op     62 allocs/op
       ent:     3.03s      3029559 ns/op    4220 B/op     99 allocs/op
      gorp:     3.05s      3046146 ns/op    1856 B/op     41 allocs/op
       rel:     3.20s      3197608 ns/op    2424 B/op     42 allocs/op
      godb:     3.21s      3213417 ns/op    4642 B/op    115 allocs/op
      gorm:     3.32s      3316146 ns/op    6982 B/op     91 allocs/op
     upper:     3.34s      3342917 ns/op   14391 B/op    684 allocs/op
       bun:     3.45s      3453909 ns/op    5067 B/op     14 allocs/op
        pg:     3.53s      3531590 ns/op    1891 B/op     10 allocs/op
       pop:     3.70s      3696167 ns/op   10153 B/op    248 allocs/op

  1000 times - MultiInsert 100 row
     beego:     4.12s      4124178 ns/op  179652 B/op   2746 allocs/op
       raw:     4.20s      4195976 ns/op  191432 B/op    932 allocs/op
    reform:     4.44s      4444236 ns/op  466402 B/op   2748 allocs/op
       ent:     4.70s      4704817 ns/op  412186 B/op   4899 allocs/op
       bun:     4.79s      4790215 ns/op   42542 B/op    218 allocs/op
        pg:     4.79s      4793534 ns/op    3316 B/op    112 allocs/op
      gorm:     5.19s      5188752 ns/op  294794 B/op   4033 allocs/op
       rel:     5.53s      5526241 ns/op  289910 B/op   3363 allocs/op
      godb:     5.63s      5628786 ns/op  260043 B/op   5895 allocs/op
     upper:     6.63s      6628465 ns/op  546290 B/op  19507 allocs/op
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      xorm:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.10s        97438 ns/op     753 B/op     13 allocs/op
       ent:     0.13s       127821 ns/op    4561 B/op     99 allocs/op
      gorp:     2.80s      2797201 ns/op    1209 B/op     32 allocs/op
 sqlboiler:     2.88s      2880961 ns/op     907 B/op     17 allocs/op
      sqlc:     2.93s      2925622 ns/op     880 B/op     14 allocs/op
    reform:     2.99s      2985758 ns/op    1778 B/op     51 allocs/op
     beego:     3.01s      3013789 ns/op    1752 B/op     47 allocs/op
        pg:     3.04s      3035302 ns/op     768 B/op      9 allocs/op
       rel:     3.04s      3038639 ns/op    2425 B/op     40 allocs/op
       bun:     3.05s      3048797 ns/op    4729 B/op      5 allocs/op
      gorm:     3.14s      3135903 ns/op    6600 B/op     81 allocs/op
       dbr:     3.17s      3171700 ns/op    2651 B/op     57 allocs/op
      godb:     3.20s      3197678 ns/op    5161 B/op    154 allocs/op
      xorm:     3.35s      3351478 ns/op    2872 B/op    104 allocs/op
       pop:     3.47s      3474504 ns/op    6589 B/op    198 allocs/op
     upper:     4.09s      4091842 ns/op   34782 B/op   1523 allocs/op

  1000 times - Read
       raw:     0.10s       104672 ns/op    2080 B/op     49 allocs/op
     beego:     0.11s       114724 ns/op    2089 B/op     75 allocs/op
      sqlc:     0.12s       115013 ns/op    2198 B/op     51 allocs/op
    reform:     0.12s       124166 ns/op    3200 B/op     85 allocs/op
      gorp:     0.13s       133908 ns/op    3903 B/op    193 allocs/op
       pop:     0.14s       135084 ns/op    3622 B/op     70 allocs/op
        pg:     0.14s       137553 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.15s       154717 ns/op     971 B/op     14 allocs/op
       bun:     0.16s       155336 ns/op    5498 B/op     21 allocs/op
       dbr:     0.16s       160550 ns/op    2184 B/op     37 allocs/op
       rel:     0.16s       161290 ns/op    1824 B/op     47 allocs/op
       ent:     0.16s       161417 ns/op    5423 B/op    148 allocs/op
      gorm:     0.23s       225077 ns/op    5302 B/op    109 allocs/op
      godb:     0.31s       309833 ns/op    4112 B/op    102 allocs/op
     upper:     0.38s       375795 ns/op    8688 B/op    324 allocs/op
      xorm:     0.40s       403748 ns/op    8754 B/op    215 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.12s       115431 ns/op    2943 B/op     73 allocs/op
      xorm:     0.19s       193967 ns/op    3278 B/op     77 allocs/op
       raw:     0.31s       309767 ns/op   38364 B/op   1037 allocs/op
     upper:     0.35s       353330 ns/op    8402 B/op    303 allocs/op
        pg:     0.39s       394906 ns/op   26535 B/op    629 allocs/op
      sqlc:     0.44s       437637 ns/op   73180 B/op   1250 allocs/op
 sqlboiler:     0.47s       465491 ns/op   58867 B/op   1259 allocs/op
       ent:     0.51s       509056 ns/op   76071 B/op   2038 allocs/op
      gorp:     0.51s       510037 ns/op   57406 B/op   1493 allocs/op
       dbr:     0.61s       609577 ns/op   32416 B/op   1545 allocs/op
     beego:     0.61s       612597 ns/op   55198 B/op   3077 allocs/op
       bun:     0.62s       619632 ns/op   32899 B/op   1118 allocs/op
       pop:     0.64s       639741 ns/op   75892 B/op   1509 allocs/op
       rel:     0.91s       914624 ns/op   95352 B/op   2250 allocs/op
      godb:     0.94s       940930 ns/op   75282 B/op   3084 allocs/op
      gorm:     1.08s      1080015 ns/op   71642 B/op   3877 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
    reform:     5.81s      2906075 ns/op    1818 B/op     51 allocs/op
      gorp:     5.86s      2929743 ns/op    1824 B/op     41 allocs/op
       ent:     5.88s      2939070 ns/op    4230 B/op     99 allocs/op
     beego:     5.93s      2966956 ns/op    2376 B/op     56 allocs/op
 sqlboiler:     5.95s      2974914 ns/op    1617 B/op     34 allocs/op
       dbr:     5.96s      2977652 ns/op    2727 B/op     65 allocs/op
      sqlc:     6.00s      2999005 ns/op    2925 B/op     62 allocs/op
      gorm:     6.09s      3044196 ns/op    6979 B/op     91 allocs/op
       bun:     6.15s      3074719 ns/op    5029 B/op     14 allocs/op
       rel:     6.17s      3084052 ns/op    2424 B/op     42 allocs/op
        pg:     6.24s      3117709 ns/op     818 B/op     10 allocs/op
      xorm:     6.26s      3129622 ns/op    2667 B/op     77 allocs/op
       raw:     6.53s      3262962 ns/op     737 B/op     13 allocs/op
     upper:     6.55s      3274439 ns/op   14355 B/op    683 allocs/op
       pop:     6.75s      3373425 ns/op   10117 B/op    248 allocs/op
      godb:     7.09s      3544493 ns/op    4594 B/op    115 allocs/op

  2000 times - MultiInsert 100 row
       raw:     8.26s      4131288 ns/op  191315 B/op    931 allocs/op
     beego:     8.68s      4339003 ns/op  179677 B/op   2746 allocs/op
    reform:     9.39s      4692541 ns/op  466390 B/op   2748 allocs/op
       ent:     9.50s      4747857 ns/op  412153 B/op   4899 allocs/op
        pg:     9.63s      4814764 ns/op    3315 B/op    112 allocs/op
       bun:     9.91s      4956780 ns/op   42478 B/op    219 allocs/op
      gorm:    10.58s      5291788 ns/op  294841 B/op   4034 allocs/op
      godb:    11.23s      5612906 ns/op  260060 B/op   5895 allocs/op
       rel:    11.32s      5660863 ns/op  289913 B/op   3363 allocs/op
     upper:    13.46s      6730799 ns/op  546173 B/op  19493 allocs/op
      xorm:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.21s       102874 ns/op     748 B/op     13 allocs/op
       ent:     0.26s       131052 ns/op    4557 B/op     99 allocs/op
 sqlboiler:     5.77s      2882681 ns/op     901 B/op     17 allocs/op
      gorp:     5.77s      2886459 ns/op    1205 B/op     32 allocs/op
     beego:     5.83s      2914467 ns/op    1753 B/op     47 allocs/op
    reform:     5.86s      2928875 ns/op    1773 B/op     51 allocs/op
      sqlc:     5.93s      2963733 ns/op     876 B/op     14 allocs/op
       pop:     5.95s      2972650 ns/op    6583 B/op    198 allocs/op
        pg:     5.97s      2983479 ns/op     768 B/op      9 allocs/op
       dbr:     6.08s      3038700 ns/op    2651 B/op     57 allocs/op
      gorm:     6.19s      3096611 ns/op    6604 B/op     81 allocs/op
      xorm:     6.23s      3116197 ns/op    2873 B/op    104 allocs/op
      godb:     6.44s      3218352 ns/op    5161 B/op    154 allocs/op
       bun:     6.81s      3403877 ns/op    4730 B/op      5 allocs/op
       rel:     7.07s      3537107 ns/op    2384 B/op     40 allocs/op
     upper:     7.92s      3959778 ns/op   34727 B/op   1523 allocs/op

  2000 times - Read
       raw:     0.22s       112173 ns/op    2087 B/op     49 allocs/op
      sqlc:     0.23s       113801 ns/op    2186 B/op     51 allocs/op
     beego:     0.23s       116859 ns/op    2089 B/op     75 allocs/op
    reform:     0.25s       125673 ns/op    3212 B/op     85 allocs/op
      gorp:     0.27s       134251 ns/op    3890 B/op    193 allocs/op
       pop:     0.28s       138120 ns/op    3581 B/op     70 allocs/op
        pg:     0.29s       142572 ns/op     872 B/op     20 allocs/op
       bun:     0.31s       154626 ns/op    5496 B/op     21 allocs/op
 sqlboiler:     0.31s       154713 ns/op     956 B/op     14 allocs/op
       ent:     0.32s       158499 ns/op    5419 B/op    148 allocs/op
       dbr:     0.32s       158941 ns/op    2184 B/op     37 allocs/op
       rel:     0.32s       162019 ns/op    1824 B/op     47 allocs/op
      gorm:     0.45s       227405 ns/op    5306 B/op    109 allocs/op
      godb:     0.65s       325791 ns/op    4114 B/op    102 allocs/op
     upper:     0.75s       373201 ns/op    8695 B/op    324 allocs/op
      xorm:     0.81s       403193 ns/op    8756 B/op    215 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.23s       116747 ns/op    2938 B/op     73 allocs/op
      xorm:     0.41s       203681 ns/op    3275 B/op     77 allocs/op
     upper:     0.71s       354482 ns/op    8395 B/op    303 allocs/op
       raw:     0.72s       358756 ns/op   38356 B/op   1037 allocs/op
      sqlc:     0.84s       422308 ns/op   73172 B/op   1250 allocs/op
 sqlboiler:     0.94s       468907 ns/op   58651 B/op   1259 allocs/op
        pg:     0.94s       471156 ns/op   22786 B/op    629 allocs/op
      gorp:     1.00s       501425 ns/op   57399 B/op   1493 allocs/op
       ent:     1.04s       520025 ns/op   76063 B/op   2038 allocs/op
       dbr:     1.21s       605021 ns/op   32416 B/op   1545 allocs/op
     beego:     1.23s       616328 ns/op   55199 B/op   3077 allocs/op
       bun:     1.24s       622374 ns/op   32899 B/op   1118 allocs/op
       pop:     1.30s       649474 ns/op   75821 B/op   1509 allocs/op
       rel:     1.79s       894424 ns/op   95352 B/op   2250 allocs/op
      godb:     1.88s       942200 ns/op   75287 B/op   3084 allocs/op
      gorm:     2.16s      1080108 ns/op   71630 B/op   3877 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
       raw:    11.60s      2899684 ns/op     720 B/op     13 allocs/op
    reform:    11.77s      2941454 ns/op    1796 B/op     51 allocs/op
      gorp:    11.79s      2947211 ns/op    1802 B/op     41 allocs/op
       ent:    11.80s      2950143 ns/op    4230 B/op     99 allocs/op
       dbr:    12.06s      3014671 ns/op    2707 B/op     65 allocs/op
      sqlc:    12.10s      3024474 ns/op    2903 B/op     62 allocs/op
 sqlboiler:    12.11s      3028471 ns/op    1594 B/op     34 allocs/op
       bun:    12.17s      3041652 ns/op    5012 B/op     14 allocs/op
      xorm:    12.37s      3092313 ns/op    2648 B/op     77 allocs/op
        pg:    12.43s      3107421 ns/op     806 B/op     10 allocs/op
      gorm:    12.48s      3120269 ns/op    6982 B/op     91 allocs/op
     beego:    12.54s      3134933 ns/op    2376 B/op     56 allocs/op
      godb:    12.56s      3139276 ns/op    4573 B/op    115 allocs/op
     upper:    13.24s      3309224 ns/op   14330 B/op    682 allocs/op
       pop:    13.45s      3363355 ns/op   10112 B/op    248 allocs/op
       rel:    14.23s      3558389 ns/op    2424 B/op     42 allocs/op

  4000 times - MultiInsert 100 row
       raw:    16.60s      4148822 ns/op  191259 B/op    931 allocs/op
     beego:    16.94s      4235940 ns/op  179617 B/op   2746 allocs/op
    reform:    17.97s      4493121 ns/op  466352 B/op   2747 allocs/op
       ent:    19.29s      4821323 ns/op  412143 B/op   4899 allocs/op
       bun:    19.44s      4859012 ns/op   42409 B/op    219 allocs/op
        pg:    19.97s      4991967 ns/op    3595 B/op    112 allocs/op
      gorm:    21.48s      5370545 ns/op  294868 B/op   4034 allocs/op
       rel:    22.32s      5580700 ns/op  289914 B/op   3363 allocs/op
      godb:    22.55s      5637113 ns/op  259997 B/op   5895 allocs/op
     upper:    26.79s      6697677 ns/op  545938 B/op  19485 allocs/op
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      xorm:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.42s       105054 ns/op     750 B/op     13 allocs/op
       ent:     0.52s       128906 ns/op    4558 B/op     99 allocs/op
      sqlc:    11.60s      2899772 ns/op     878 B/op     14 allocs/op
 sqlboiler:    11.61s      2902704 ns/op     901 B/op     17 allocs/op
    reform:    11.71s      2926893 ns/op    1775 B/op     51 allocs/op
     beego:    11.71s      2927845 ns/op    1752 B/op     47 allocs/op
      gorp:    11.80s      2950816 ns/op    1204 B/op     32 allocs/op
       dbr:    11.84s      2959964 ns/op    2651 B/op     57 allocs/op
       pop:    12.01s      3002106 ns/op    6583 B/op    198 allocs/op
      gorm:    12.32s      3079383 ns/op    6604 B/op     81 allocs/op
      xorm:    12.45s      3113117 ns/op    2872 B/op    104 allocs/op
        pg:    12.54s      3134497 ns/op     768 B/op      9 allocs/op
       bun:    12.55s      3137002 ns/op    4730 B/op      5 allocs/op
       rel:    12.61s      3151848 ns/op    2384 B/op     40 allocs/op
      godb:    14.74s      3685113 ns/op    5161 B/op    154 allocs/op
     upper:    16.13s      4032826 ns/op   34681 B/op   1522 allocs/op

  4000 times - Read
       raw:     0.46s       114290 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.46s       115783 ns/op    2190 B/op     51 allocs/op
     beego:     0.48s       120233 ns/op    2088 B/op     75 allocs/op
    reform:     0.50s       125172 ns/op    3213 B/op     85 allocs/op
       pop:     0.53s       133705 ns/op    3553 B/op     70 allocs/op
      gorp:     0.54s       134596 ns/op    3896 B/op    193 allocs/op
        pg:     0.56s       138831 ns/op     872 B/op     20 allocs/op
       bun:     0.62s       155269 ns/op    5493 B/op     21 allocs/op
       dbr:     0.63s       158484 ns/op    2184 B/op     37 allocs/op
       ent:     0.64s       159790 ns/op    5418 B/op    148 allocs/op
 sqlboiler:     0.65s       161999 ns/op     965 B/op     14 allocs/op
       rel:     0.65s       162498 ns/op    1824 B/op     47 allocs/op
      gorm:     0.94s       233844 ns/op    5304 B/op    109 allocs/op
      godb:     1.11s       278508 ns/op    4113 B/op    102 allocs/op
     upper:     1.51s       377915 ns/op    8697 B/op    324 allocs/op
      xorm:     1.63s       407517 ns/op    8755 B/op    215 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.47s       117124 ns/op    2944 B/op     73 allocs/op
      xorm:     0.78s       195170 ns/op    3274 B/op     77 allocs/op
     upper:     1.42s       355861 ns/op    8390 B/op    303 allocs/op
       raw:     1.45s       362607 ns/op   38355 B/op   1037 allocs/op
        pg:     1.57s       392794 ns/op   23834 B/op    629 allocs/op
      sqlc:     1.72s       429824 ns/op   73176 B/op   1250 allocs/op
 sqlboiler:     1.84s       460480 ns/op   58731 B/op   1259 allocs/op
      gorp:     2.02s       504338 ns/op   57403 B/op   1493 allocs/op
       ent:     2.04s       510775 ns/op   76059 B/op   2038 allocs/op
       dbr:     2.42s       604583 ns/op   32416 B/op   1545 allocs/op
     beego:     2.48s       619675 ns/op   55191 B/op   3077 allocs/op
       bun:     2.49s       622581 ns/op   32891 B/op   1118 allocs/op
       pop:     2.56s       640857 ns/op   75816 B/op   1509 allocs/op
      godb:     3.26s       815760 ns/op   75272 B/op   3084 allocs/op
       rel:     3.52s       880888 ns/op   95352 B/op   2250 allocs/op
      gorm:     4.39s      1096908 ns/op   71628 B/op   3877 allocs/op
```