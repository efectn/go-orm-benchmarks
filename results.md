# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
     upper:     0.59s      2951529 ns/op   15035 B/op    694 allocs/op
      gorp:     0.60s      2991030 ns/op    1832 B/op     39 allocs/op
      sqlc:     0.61s      3026546 ns/op    3294 B/op     63 allocs/op
       raw:     0.61s      3049297 ns/op    1039 B/op     16 allocs/op
  raw_stmt:     0.61s      3050785 ns/op     692 B/op     14 allocs/op
 beego_orm:     0.62s      3078870 ns/op    2415 B/op     56 allocs/op
       dbr:     0.62s      3085238 ns/op    2688 B/op     65 allocs/op
       ent:     0.62s      3108133 ns/op    4178 B/op     98 allocs/op
      gorm:     0.62s      3114189 ns/op    7280 B/op     91 allocs/op
    reform:     0.63s      3153798 ns/op    2984 B/op     59 allocs/op
 gorm_prep:     0.63s      3153898 ns/op    5886 B/op     71 allocs/op
        pg:     0.64s      3189698 ns/op    1031 B/op     10 allocs/op
       bun:     0.64s      3218783 ns/op    5327 B/op     14 allocs/op
      xorm:     0.66s      3296699 ns/op    3261 B/op     79 allocs/op
       rel:     0.67s      3350857 ns/op    2430 B/op     41 allocs/op
      godb:     0.68s      3394254 ns/op    4601 B/op    114 allocs/op
       pop:     0.70s      3494769 ns/op   15571 B/op    621 allocs/op

   200 times - MultiInsert 100 row
  raw_stmt:     0.83s      4131600 ns/op  188312 B/op    938 allocs/op
       raw:     0.84s      4207062 ns/op  192357 B/op    938 allocs/op
      gorp:     0.89s      4458305 ns/op  412389 B/op   4901 allocs/op
 gorm_prep:     0.92s      4586938 ns/op  257144 B/op   2486 allocs/op
    reform:     0.92s      4595145 ns/op  467516 B/op   2748 allocs/op
      gorm:     0.98s      4876716 ns/op  294717 B/op   4032 allocs/op
        pg:     0.98s      4904496 ns/op    8932 B/op    112 allocs/op
       ent:     0.99s      4964750 ns/op  412392 B/op   4901 allocs/op
       bun:     1.02s      5111698 ns/op   42289 B/op    217 allocs/op
       rel:     1.09s      5439037 ns/op  289894 B/op   3361 allocs/op
      xorm:     1.11s      5525189 ns/op  266586 B/op   5822 allocs/op
      godb:     1.11s      5546008 ns/op  259968 B/op   5893 allocs/op
 beego_orm:     1.18s      5915509 ns/op  179653 B/op   2746 allocs/op
      sqlc:     TBD
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
     upper:     TBD

   200 times - Update
       raw:     0.02s        90501 ns/op     724 B/op     13 allocs/op
  raw_stmt:     0.02s        97955 ns/op     738 B/op     14 allocs/op
       ent:     0.03s       136482 ns/op    4529 B/op     99 allocs/op
 gorm_prep:     0.48s      2407602 ns/op    5135 B/op     59 allocs/op
       pop:     0.55s      2760608 ns/op   12890 B/op    663 allocs/op
      sqlc:     0.60s      2987011 ns/op     893 B/op     14 allocs/op
      gorp:     0.61s      3053393 ns/op    1182 B/op     32 allocs/op
       dbr:     0.61s      3054068 ns/op    2635 B/op     56 allocs/op
    reform:     0.61s      3057119 ns/op    1783 B/op     51 allocs/op
 beego_orm:     0.63s      3126960 ns/op    1813 B/op     47 allocs/op
       bun:     0.63s      3163122 ns/op    4738 B/op      5 allocs/op
      xorm:     0.64s      3197053 ns/op    2892 B/op    105 allocs/op
        pg:     0.64s      3201897 ns/op     768 B/op      9 allocs/op
      gorm:     0.65s      3225643 ns/op    6625 B/op     81 allocs/op
       rel:     0.66s      3315603 ns/op    2507 B/op     40 allocs/op
      godb:     0.67s      3330137 ns/op    5160 B/op    154 allocs/op
     upper:     0.86s      4319402 ns/op   35243 B/op   1531 allocs/op

   200 times - Read
       raw:     0.02s       110852 ns/op    2190 B/op     49 allocs/op
      sqlc:     0.02s       115009 ns/op    2172 B/op     51 allocs/op
 beego_orm:     0.02s       115870 ns/op    2109 B/op     75 allocs/op
 gorm_prep:     0.02s       117617 ns/op    4961 B/op     98 allocs/op
  raw_stmt:     0.02s       117975 ns/op    2214 B/op     50 allocs/op
    reform:     0.03s       131378 ns/op    3453 B/op     85 allocs/op
      gorp:     0.03s       136895 ns/op    3428 B/op    137 allocs/op
        pg:     0.03s       143334 ns/op     873 B/op     20 allocs/op
       bun:     0.03s       152508 ns/op    5488 B/op     21 allocs/op
       rel:     0.03s       156666 ns/op    1824 B/op     47 allocs/op
       dbr:     0.03s       161827 ns/op    2184 B/op     37 allocs/op
       ent:     0.03s       161952 ns/op    5387 B/op    148 allocs/op
       pop:     0.04s       185806 ns/op    8622 B/op    442 allocs/op
      gorm:     0.04s       209840 ns/op    5298 B/op    109 allocs/op
      godb:     0.06s       289000 ns/op    4160 B/op    102 allocs/op
     upper:     0.07s       345652 ns/op    8694 B/op    324 allocs/op
      xorm:     0.08s       382003 ns/op    8777 B/op    216 allocs/op

   200 times - MultiRead limit 100
       raw:     0.06s       321835 ns/op   38360 B/op   1037 allocs/op
  raw_stmt:     0.07s       327875 ns/op   38390 B/op   1038 allocs/op
     upper:     0.07s       333117 ns/op    8388 B/op    303 allocs/op
    reform:     0.07s       358471 ns/op   56155 B/op   1270 allocs/op
      sqlc:     0.08s       381650 ns/op   73137 B/op   1250 allocs/op
        pg:     0.09s       448649 ns/op   22261 B/op    629 allocs/op
       ent:     0.10s       499320 ns/op   76044 B/op   2038 allocs/op
 beego_orm:     0.10s       524339 ns/op   55237 B/op   3077 allocs/op
       bun:     0.11s       529267 ns/op   32869 B/op   1118 allocs/op
       dbr:     0.11s       571297 ns/op   54872 B/op   1545 allocs/op
       pop:     0.12s       581282 ns/op   80728 B/op   1850 allocs/op
       rel:     0.16s       791356 ns/op   95352 B/op   2250 allocs/op
 gorm_prep:     0.16s       819667 ns/op   71325 B/op   3577 allocs/op
      godb:     0.17s       842183 ns/op   97709 B/op   3084 allocs/op
      gorm:     0.18s       916834 ns/op   71715 B/op   3877 allocs/op
      gorp:     0.18s       924254 ns/op  165449 B/op   3575 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
 beego_orm:     2.92s      2923521 ns/op    2409 B/op     56 allocs/op
    reform:     2.93s      2933613 ns/op    2666 B/op     58 allocs/op
      gorp:     2.94s      2942187 ns/op    1793 B/op     39 allocs/op
  raw_stmt:     2.96s      2959898 ns/op     713 B/op     14 allocs/op
       dbr:     2.98s      2975919 ns/op    2688 B/op     65 allocs/op
        pg:     2.99s      2992602 ns/op    1890 B/op     10 allocs/op
      gorm:     3.06s      3058018 ns/op    7011 B/op     91 allocs/op
       ent:     3.12s      3122250 ns/op    4220 B/op     99 allocs/op
      sqlc:     3.12s      3124531 ns/op    2957 B/op     62 allocs/op
      xorm:     3.13s      3125940 ns/op    2918 B/op     78 allocs/op
       raw:     3.15s      3149560 ns/op     769 B/op     13 allocs/op
       rel:     3.23s      3227596 ns/op    2424 B/op     42 allocs/op
     upper:     3.24s      3241685 ns/op   14507 B/op    684 allocs/op
       bun:     3.26s      3264239 ns/op    5059 B/op     14 allocs/op
 gorm_prep:     3.31s      3310234 ns/op    5400 B/op     69 allocs/op
      godb:     3.33s      3332358 ns/op    4590 B/op    115 allocs/op
       pop:     3.66s      3656848 ns/op   15193 B/op    620 allocs/op

  1000 times - MultiInsert 100 row
       raw:     4.05s      4045036 ns/op  191430 B/op    932 allocs/op
  raw_stmt:     4.30s      4301916 ns/op  187355 B/op    932 allocs/op
 beego_orm:     4.37s      4369579 ns/op  179801 B/op   2746 allocs/op
    reform:     4.49s      4494809 ns/op  467509 B/op   2748 allocs/op
 gorm_prep:     4.59s      4588322 ns/op  257083 B/op   2487 allocs/op
       ent:     4.74s      4741728 ns/op  412172 B/op   4899 allocs/op
      gorp:     4.78s      4780930 ns/op  412171 B/op   4899 allocs/op
       bun:     5.08s      5078058 ns/op   42369 B/op    218 allocs/op
        pg:     5.17s      5169943 ns/op    5487 B/op    112 allocs/op
      xorm:     5.24s      5240998 ns/op  266233 B/op   5821 allocs/op
      gorm:     5.28s      5280062 ns/op  294795 B/op   4033 allocs/op
       rel:     5.35s      5350123 ns/op  289908 B/op   3363 allocs/op
      godb:     5.58s      5575723 ns/op  260063 B/op   5895 allocs/op
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     TBD
     upper:     TBD

  1000 times - Update
       raw:     0.09s        92192 ns/op     753 B/op     13 allocs/op
  raw_stmt:     0.10s       103218 ns/op     769 B/op     14 allocs/op
       ent:     0.13s       129414 ns/op    4552 B/op     99 allocs/op
      sqlc:     2.89s      2885098 ns/op     881 B/op     14 allocs/op
 gorm_prep:     3.00s      3002992 ns/op    5113 B/op     59 allocs/op
       bun:     3.01s      3008360 ns/op    4730 B/op      5 allocs/op
      gorp:     3.09s      3088486 ns/op    1201 B/op     32 allocs/op
 beego_orm:     3.12s      3122460 ns/op    1801 B/op     47 allocs/op
    reform:     3.13s      3128104 ns/op    1801 B/op     51 allocs/op
      godb:     3.16s      3163809 ns/op    5162 B/op    154 allocs/op
       dbr:     3.16s      3164642 ns/op    2635 B/op     56 allocs/op
        pg:     3.18s      3182654 ns/op     768 B/op      9 allocs/op
      gorm:     3.23s      3228557 ns/op    6600 B/op     81 allocs/op
      xorm:     3.28s      3275626 ns/op    2888 B/op    105 allocs/op
       rel:     3.31s      3309862 ns/op    2504 B/op     40 allocs/op
       pop:     3.32s      3316528 ns/op   12916 B/op    663 allocs/op
     upper:     4.13s      4125830 ns/op   34914 B/op   1523 allocs/op

  1000 times - Read
       raw:     0.11s       105447 ns/op    2080 B/op     49 allocs/op
      sqlc:     0.11s       110138 ns/op    2199 B/op     51 allocs/op
  raw_stmt:     0.11s       111316 ns/op    2111 B/op     50 allocs/op
 beego_orm:     0.12s       120692 ns/op    2106 B/op     75 allocs/op
    reform:     0.12s       124219 ns/op    3456 B/op     85 allocs/op
      gorp:     0.13s       132393 ns/op    3439 B/op    137 allocs/op
 gorm_prep:     0.14s       139266 ns/op    4968 B/op     98 allocs/op
        pg:     0.14s       143262 ns/op     872 B/op     20 allocs/op
       ent:     0.15s       153645 ns/op    5407 B/op    148 allocs/op
       dbr:     0.16s       156372 ns/op    2184 B/op     37 allocs/op
       bun:     0.16s       157074 ns/op    5494 B/op     21 allocs/op
       rel:     0.16s       164066 ns/op    1824 B/op     47 allocs/op
       pop:     0.18s       181007 ns/op    8797 B/op    442 allocs/op
      gorm:     0.22s       216423 ns/op    5302 B/op    109 allocs/op
      godb:     0.29s       293456 ns/op    4161 B/op    102 allocs/op
     upper:     0.35s       349098 ns/op    8708 B/op    324 allocs/op
      xorm:     0.38s       378379 ns/op    8779 B/op    216 allocs/op

  1000 times - MultiRead limit 100
  raw_stmt:     0.32s       324404 ns/op   38396 B/op   1038 allocs/op
       raw:     0.34s       336438 ns/op   38364 B/op   1037 allocs/op
     upper:     0.34s       341458 ns/op    8392 B/op    303 allocs/op
    reform:     0.35s       351605 ns/op   56176 B/op   1270 allocs/op
      sqlc:     0.38s       376459 ns/op   73181 B/op   1250 allocs/op
        pg:     0.42s       416147 ns/op   24361 B/op    629 allocs/op
       ent:     0.44s       439749 ns/op   76063 B/op   2038 allocs/op
 beego_orm:     0.54s       535862 ns/op   55240 B/op   3077 allocs/op
       bun:     0.55s       546707 ns/op   32877 B/op   1118 allocs/op
       dbr:     0.55s       553895 ns/op   54872 B/op   1545 allocs/op
       pop:     0.59s       587103 ns/op   80019 B/op   1850 allocs/op
       rel:     0.79s       785404 ns/op   95352 B/op   2250 allocs/op
 gorm_prep:     0.84s       840215 ns/op   71245 B/op   3577 allocs/op
      godb:     0.87s       867492 ns/op   97735 B/op   3084 allocs/op
      gorp:     0.95s       951268 ns/op  165485 B/op   3575 allocs/op
      gorm:     0.96s       956051 ns/op   71642 B/op   3877 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
  raw_stmt:     5.96s      2977625 ns/op     717 B/op     14 allocs/op
    reform:     6.07s      3036360 ns/op    2624 B/op     58 allocs/op
       ent:     6.08s      3039014 ns/op    4229 B/op     99 allocs/op
       dbr:     6.18s      3089663 ns/op    2688 B/op     65 allocs/op
 gorm_prep:     6.22s      3109788 ns/op    5340 B/op     69 allocs/op
       bun:     6.26s      3129230 ns/op    5029 B/op     14 allocs/op
       raw:     6.28s      3141893 ns/op     737 B/op     13 allocs/op
 beego_orm:     6.32s      3158050 ns/op    2409 B/op     56 allocs/op
        pg:     6.38s      3188255 ns/op    1342 B/op     10 allocs/op
      gorm:     6.51s      3256442 ns/op    6997 B/op     91 allocs/op
      xorm:     6.52s      3259341 ns/op    2875 B/op     78 allocs/op
       rel:     6.58s      3291789 ns/op    2424 B/op     42 allocs/op
      sqlc:     6.60s      3300437 ns/op    2925 B/op     62 allocs/op
     upper:     6.65s      3325587 ns/op   14428 B/op    683 allocs/op
      godb:     6.71s      3353933 ns/op    4588 B/op    115 allocs/op
      gorp:     6.88s      3442144 ns/op    1792 B/op     39 allocs/op
       pop:     7.02s      3510216 ns/op   15137 B/op    620 allocs/op

  2000 times - MultiInsert 100 row
       raw:     8.31s      4156329 ns/op  191316 B/op    931 allocs/op
  raw_stmt:     8.40s      4199813 ns/op  187238 B/op    931 allocs/op
 beego_orm:     8.53s      4266801 ns/op  179696 B/op   2746 allocs/op
    reform:     9.09s      4544495 ns/op  467341 B/op   2747 allocs/op
 gorm_prep:     9.54s      4768636 ns/op  257057 B/op   2488 allocs/op
       ent:     9.63s      4817213 ns/op  412153 B/op   4899 allocs/op
       bun:     9.77s      4882696 ns/op   42402 B/op    218 allocs/op
        pg:     9.78s      4887817 ns/op    3876 B/op    112 allocs/op
      gorp:    10.24s      5119797 ns/op  412155 B/op   4899 allocs/op
      gorm:    10.36s      5180017 ns/op  294840 B/op   4034 allocs/op
      xorm:    10.65s      5323952 ns/op  266637 B/op   5822 allocs/op
      godb:    11.28s      5639790 ns/op  260038 B/op   5895 allocs/op
       rel:    11.37s      5687280 ns/op  289910 B/op   3363 allocs/op
       pop:     doesn't support bulk-insert
     upper:     TBD
      sqlc:     TBD
       dbr:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.21s       106550 ns/op     749 B/op     13 allocs/op
  raw_stmt:     0.21s       107179 ns/op     764 B/op     14 allocs/op
       ent:     0.25s       125996 ns/op    4556 B/op     99 allocs/op
      gorp:     5.97s      2986264 ns/op    1205 B/op     32 allocs/op
      sqlc:     5.99s      2996461 ns/op     877 B/op     14 allocs/op
    reform:     6.04s      3018746 ns/op    1805 B/op     51 allocs/op
 gorm_prep:     6.06s      3031192 ns/op    5121 B/op     59 allocs/op
        pg:     6.16s      3079701 ns/op     768 B/op      9 allocs/op
       bun:     6.22s      3112442 ns/op    4729 B/op      5 allocs/op
      gorm:     6.25s      3124880 ns/op    6604 B/op     81 allocs/op
 beego_orm:     6.28s      3141183 ns/op    1801 B/op     47 allocs/op
       pop:     6.31s      3153889 ns/op   12917 B/op    663 allocs/op
       dbr:     6.46s      3228423 ns/op    2635 B/op     56 allocs/op
      godb:     6.55s      3276181 ns/op    5161 B/op    154 allocs/op
       rel:     6.61s      3302634 ns/op    2504 B/op     40 allocs/op
      xorm:     6.71s      3355087 ns/op    2889 B/op    105 allocs/op
     upper:     8.07s      4035328 ns/op   34883 B/op   1523 allocs/op

  2000 times - Read
       raw:     0.23s       113192 ns/op    2087 B/op     49 allocs/op
      sqlc:     0.23s       115004 ns/op    2186 B/op     51 allocs/op
  raw_stmt:     0.23s       116494 ns/op    2114 B/op     50 allocs/op
 beego_orm:     0.24s       119078 ns/op    2105 B/op     75 allocs/op
    reform:     0.25s       125448 ns/op    3467 B/op     85 allocs/op
      gorp:     0.27s       134003 ns/op    3450 B/op    137 allocs/op
 gorm_prep:     0.28s       139527 ns/op    4971 B/op     98 allocs/op
        pg:     0.29s       143626 ns/op     872 B/op     20 allocs/op
       ent:     0.30s       150629 ns/op    5419 B/op    148 allocs/op
       bun:     0.31s       155047 ns/op    5492 B/op     21 allocs/op
       dbr:     0.32s       160539 ns/op    2184 B/op     37 allocs/op
       rel:     0.33s       163666 ns/op    1824 B/op     47 allocs/op
       pop:     0.37s       183384 ns/op    8749 B/op    442 allocs/op
      gorm:     0.43s       213169 ns/op    5306 B/op    109 allocs/op
      godb:     0.59s       292929 ns/op    4162 B/op    102 allocs/op
     upper:     0.69s       346601 ns/op    8711 B/op    324 allocs/op
      xorm:     0.74s       368540 ns/op    8781 B/op    216 allocs/op

  2000 times - MultiRead limit 100
       raw:     0.64s       317749 ns/op   38356 B/op   1037 allocs/op
  raw_stmt:     0.64s       319960 ns/op   38387 B/op   1038 allocs/op
     upper:     0.66s       330322 ns/op    8402 B/op    303 allocs/op
    reform:     0.72s       360891 ns/op   56149 B/op   1270 allocs/op
      sqlc:     0.74s       370047 ns/op   73172 B/op   1250 allocs/op
        pg:     0.81s       405874 ns/op   24359 B/op    629 allocs/op
       ent:     0.87s       435283 ns/op   76058 B/op   2038 allocs/op
 beego_orm:     1.03s       515418 ns/op   55236 B/op   3077 allocs/op
       bun:     1.06s       529792 ns/op   32884 B/op   1118 allocs/op
       dbr:     1.10s       548954 ns/op   54872 B/op   1545 allocs/op
       pop:     1.17s       583887 ns/op   79927 B/op   1849 allocs/op
       rel:     1.56s       782396 ns/op   95352 B/op   2250 allocs/op
 gorm_prep:     1.61s       804474 ns/op   71236 B/op   3577 allocs/op
      godb:     1.71s       852612 ns/op   97724 B/op   3084 allocs/op
      gorm:     1.88s       940138 ns/op   71629 B/op   3877 allocs/op
      gorp:     1.94s       968887 ns/op  165483 B/op   3575 allocs/op
      xorm:     doesn't work
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
  raw_stmt:    12.01s      3002522 ns/op     719 B/op     14 allocs/op
 beego_orm:    12.09s      3022997 ns/op    2408 B/op     56 allocs/op
    reform:    12.23s      3056810 ns/op    2602 B/op     58 allocs/op
       ent:    12.24s      3059942 ns/op    4230 B/op     99 allocs/op
      gorp:    12.25s      3063563 ns/op    1788 B/op     39 allocs/op
        pg:    12.35s      3087888 ns/op    1068 B/op     10 allocs/op
 gorm_prep:    12.35s      3088644 ns/op    5300 B/op     69 allocs/op
      sqlc:    12.38s      3094607 ns/op    2903 B/op     62 allocs/op
       dbr:    12.41s      3101695 ns/op    2688 B/op     65 allocs/op
       bun:    12.46s      3115317 ns/op    5010 B/op     14 allocs/op
      gorm:    12.51s      3126918 ns/op    6980 B/op     91 allocs/op
      godb:    12.95s      3237834 ns/op    4586 B/op    115 allocs/op
      xorm:    13.01s      3253057 ns/op    2854 B/op     78 allocs/op
       rel:    13.32s      3330119 ns/op    2424 B/op     42 allocs/op
     upper:    13.52s      3380797 ns/op   14414 B/op    682 allocs/op
       raw:    13.60s      3399530 ns/op     720 B/op     13 allocs/op
       pop:    13.98s      3494398 ns/op   15136 B/op    620 allocs/op

  4000 times - MultiInsert 100 row
  raw_stmt:    16.63s      4157382 ns/op  187179 B/op    931 allocs/op
 beego_orm:    17.52s      4379926 ns/op  179569 B/op   2746 allocs/op
       raw:    17.67s      4417005 ns/op  191258 B/op    931 allocs/op
    reform:    17.67s      4417369 ns/op  467370 B/op   2747 allocs/op
 gorm_prep:    18.75s      4686265 ns/op  257058 B/op   2488 allocs/op
      gorp:    19.30s      4825682 ns/op  412140 B/op   4899 allocs/op
       bun:    19.83s      4957073 ns/op   42414 B/op    219 allocs/op
       ent:    20.16s      5039215 ns/op  412142 B/op   4899 allocs/op
        pg:    20.21s      5051404 ns/op    3839 B/op    112 allocs/op
      gorm:    21.09s      5273030 ns/op  294868 B/op   4034 allocs/op
      xorm:    21.63s      5408257 ns/op  266568 B/op   5822 allocs/op
       rel:    21.76s      5440309 ns/op  289913 B/op   3363 allocs/op
      godb:    22.84s      5710840 ns/op  260025 B/op   5895 allocs/op
     upper:     TBD
       dbr:     doesn't support bulk-insert
      sqlc:     TBD
       pop:     doesn't support bulk-insert

  4000 times - Update
  raw_stmt:     0.42s       105004 ns/op     766 B/op     14 allocs/op
       raw:     0.42s       105708 ns/op     750 B/op     13 allocs/op
       ent:     0.51s       128295 ns/op    4556 B/op     99 allocs/op
    reform:    12.09s      3021676 ns/op    1805 B/op     51 allocs/op
 beego_orm:    12.12s      3030522 ns/op    1800 B/op     47 allocs/op
 gorm_prep:    12.27s      3068128 ns/op    5117 B/op     59 allocs/op
      sqlc:    12.29s      3072214 ns/op     878 B/op     14 allocs/op
       dbr:    12.36s      3089163 ns/op    2635 B/op     56 allocs/op
      gorp:    12.36s      3089352 ns/op    1204 B/op     32 allocs/op
        pg:    12.43s      3106954 ns/op     768 B/op      9 allocs/op
       pop:    12.58s      3144366 ns/op   12917 B/op    663 allocs/op
       bun:    12.74s      3184031 ns/op    4730 B/op      5 allocs/op
      gorm:    12.74s      3184901 ns/op    6604 B/op     81 allocs/op
      xorm:    13.04s      3260125 ns/op    2889 B/op    105 allocs/op
       rel:    13.16s      3290694 ns/op    2504 B/op     40 allocs/op
      godb:    13.17s      3291765 ns/op    5161 B/op    154 allocs/op
     upper:    16.35s      4086557 ns/op   34839 B/op   1522 allocs/op

  4000 times - Read
      sqlc:     0.44s       110184 ns/op    2190 B/op     51 allocs/op
       raw:     0.45s       111642 ns/op    2080 B/op     49 allocs/op
  raw_stmt:     0.46s       114036 ns/op    2112 B/op     50 allocs/op
 beego_orm:     0.48s       120056 ns/op    2104 B/op     75 allocs/op
    reform:     0.50s       123865 ns/op    3467 B/op     85 allocs/op
      gorp:     0.54s       133938 ns/op    3450 B/op    137 allocs/op
 gorm_prep:     0.56s       140534 ns/op    4972 B/op     98 allocs/op
        pg:     0.57s       143308 ns/op     872 B/op     20 allocs/op
       bun:     0.60s       150881 ns/op    5493 B/op     21 allocs/op
       ent:     0.61s       151551 ns/op    5418 B/op    148 allocs/op
       dbr:     0.63s       156274 ns/op    2184 B/op     37 allocs/op
       rel:     0.64s       161060 ns/op    1824 B/op     47 allocs/op
       pop:     0.73s       182434 ns/op    8711 B/op    442 allocs/op
      gorm:     0.83s       208299 ns/op    5303 B/op    109 allocs/op
      godb:     1.16s       289771 ns/op    4161 B/op    102 allocs/op
     upper:     1.39s       347123 ns/op    8712 B/op    324 allocs/op
      xorm:     1.46s       364538 ns/op    8780 B/op    216 allocs/op

  4000 times - MultiRead limit 100
       raw:     1.26s       316072 ns/op   38355 B/op   1037 allocs/op
  raw_stmt:     1.28s       319219 ns/op   38387 B/op   1038 allocs/op
     upper:     1.29s       322579 ns/op    8401 B/op    303 allocs/op
    reform:     1.41s       353159 ns/op   56157 B/op   1270 allocs/op
      sqlc:     1.48s       370139 ns/op   73176 B/op   1250 allocs/op
        pg:     1.60s       401090 ns/op   23308 B/op    629 allocs/op
       ent:     1.78s       443760 ns/op   76063 B/op   2038 allocs/op
 beego_orm:     2.09s       522492 ns/op   55228 B/op   3077 allocs/op
       bun:     2.12s       529040 ns/op   32885 B/op   1118 allocs/op
       dbr:     2.16s       540184 ns/op   54872 B/op   1545 allocs/op
       pop:     2.39s       596267 ns/op   80110 B/op   1850 allocs/op
       rel:     3.15s       787805 ns/op   95352 B/op   2250 allocs/op
 gorm_prep:     3.25s       812247 ns/op   71236 B/op   3577 allocs/op
      godb:     3.46s       865254 ns/op   97722 B/op   3084 allocs/op
      gorp:     3.70s       924067 ns/op  165467 B/op   3575 allocs/op
      gorm:     3.81s       951788 ns/op   71627 B/op   3877 allocs/op
      xorm:     doesn't work
```