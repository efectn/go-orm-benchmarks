# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
       raw:     0.55s      2752652 ns/op    1039 B/op     16 allocs/op
       ent:     0.67s      3352429 ns/op    4172 B/op     99 allocs/op
      gorp:     0.67s      3359378 ns/op    2097 B/op     43 allocs/op
    reform:     0.68s      3414137 ns/op    2115 B/op     54 allocs/op
       dbr:     0.70s      3478146 ns/op    2797 B/op     66 allocs/op
        pg:     0.72s      3599470 ns/op    6280 B/op     11 allocs/op
       bun:     0.73s      3646819 ns/op    5322 B/op     14 allocs/op
      gorm:     0.75s      3727902 ns/op    7176 B/op     90 allocs/op
 sqlboiler:     0.76s      3777524 ns/op    1906 B/op     37 allocs/op
       pop:     0.76s      3786137 ns/op   10506 B/op    250 allocs/op
      sqlc:     0.77s      3847178 ns/op    3209 B/op     64 allocs/op
     beego:     0.79s      3926057 ns/op    2378 B/op     56 allocs/op
      godb:     0.79s      3949843 ns/op    4702 B/op    116 allocs/op
      xorm:     0.80s      4007111 ns/op    3446 B/op     89 allocs/op
     upper:     0.84s      4215853 ns/op   14224 B/op    685 allocs/op
       rel:     0.86s      4283517 ns/op    2511 B/op     41 allocs/op

   200 times - MultiInsert 100 row
       raw:     1.24s      6212363 ns/op  192354 B/op    938 allocs/op
     beego:     1.28s      6378317 ns/op  179858 B/op   2747 allocs/op
        pg:     1.28s      6423552 ns/op    8932 B/op    112 allocs/op
    reform:     1.38s      6909158 ns/op  466615 B/op   2749 allocs/op
       bun:     1.49s      7442687 ns/op   42588 B/op    217 allocs/op
       ent:     1.54s      7693575 ns/op  412426 B/op   4902 allocs/op
      gorm:     1.60s      7990129 ns/op  272176 B/op   3728 allocs/op
      godb:     1.72s      8586919 ns/op  260048 B/op   5894 allocs/op
       rel:     1.81s      9042387 ns/op  303975 B/op   3261 allocs/op
      xorm:     1.99s      9966053 ns/op  255114 B/op   5417 allocs/op
     upper:     2.14s     10711570 ns/op  547019 B/op  19610 allocs/op
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

   200 times - Update
       raw:     0.06s       304950 ns/op     722 B/op     13 allocs/op
       ent:     0.07s       343798 ns/op    4530 B/op     99 allocs/op
    reform:     0.65s      3271410 ns/op    1751 B/op     51 allocs/op
       dbr:     0.66s      3289595 ns/op    2651 B/op     57 allocs/op
      gorp:     0.68s      3382558 ns/op    1221 B/op     32 allocs/op
 sqlboiler:     0.69s      3461369 ns/op     926 B/op     17 allocs/op
       pop:     0.72s      3593452 ns/op    6562 B/op    198 allocs/op
      gorm:     0.75s      3730265 ns/op    6624 B/op     81 allocs/op
        pg:     0.76s      3814982 ns/op     768 B/op      9 allocs/op
      xorm:     0.77s      3872616 ns/op    3653 B/op    126 allocs/op
       bun:     0.78s      3922279 ns/op    4733 B/op      5 allocs/op
     beego:     0.79s      3951965 ns/op    1761 B/op     47 allocs/op
      sqlc:     0.80s      3993444 ns/op     890 B/op     14 allocs/op
      godb:     0.85s      4240579 ns/op    5161 B/op    154 allocs/op
       rel:     0.91s      4541783 ns/op    2531 B/op     41 allocs/op
     upper:     1.07s      5360679 ns/op   33824 B/op   1510 allocs/op

   200 times - Read
       pop:     0.02s       118089 ns/op    3729 B/op     71 allocs/op
      sqlc:     0.03s       126747 ns/op    2162 B/op     52 allocs/op
    reform:     0.03s       135440 ns/op    3232 B/op     86 allocs/op
       rel:     0.03s       137262 ns/op    1803 B/op     45 allocs/op
      gorp:     0.03s       144737 ns/op    3870 B/op    194 allocs/op
       raw:     0.03s       153841 ns/op    2099 B/op     50 allocs/op
     beego:     0.04s       177885 ns/op    2096 B/op     75 allocs/op
        pg:     0.04s       188658 ns/op     873 B/op     20 allocs/op
       bun:     0.04s       215349 ns/op    5493 B/op     21 allocs/op
       dbr:     0.06s       303842 ns/op    2184 B/op     37 allocs/op
      xorm:     0.06s       321212 ns/op    4618 B/op    125 allocs/op
      gorm:     0.07s       325065 ns/op    5154 B/op     94 allocs/op
       ent:     0.08s       375319 ns/op    5424 B/op    149 allocs/op
 sqlboiler:     0.08s       388808 ns/op    1119 B/op     14 allocs/op
      godb:     0.09s       441338 ns/op    4112 B/op    102 allocs/op
     upper:     0.18s       911060 ns/op    8127 B/op    315 allocs/op

   200 times - MultiRead limit 100
    reform:     0.02s       107973 ns/op    2955 B/op     74 allocs/op
     upper:     0.08s       381366 ns/op    7896 B/op    294 allocs/op
        pg:     0.11s       568338 ns/op   27513 B/op    629 allocs/op
      sqlc:     0.11s       569410 ns/op   73128 B/op   1251 allocs/op
       raw:     0.15s       741537 ns/op   38354 B/op   1038 allocs/op
 sqlboiler:     0.16s       780131 ns/op   58614 B/op   1260 allocs/op
       dbr:     0.16s       804551 ns/op   32416 B/op   1545 allocs/op
      gorm:     0.18s       888645 ns/op   57220 B/op   2278 allocs/op
      gorp:     0.18s       900421 ns/op   57398 B/op   1494 allocs/op
       bun:     0.23s      1132938 ns/op   34029 B/op   1124 allocs/op
       pop:     0.23s      1141635 ns/op   77345 B/op   1511 allocs/op
       ent:     0.25s      1262233 ns/op   76036 B/op   2039 allocs/op
       rel:     0.32s      1602469 ns/op   95321 B/op   2248 allocs/op
     beego:     0.34s      1679280 ns/op   55207 B/op   3077 allocs/op
      godb:     0.37s      1865592 ns/op   75289 B/op   3084 allocs/op
      xorm:     0.42s      2091036 ns/op  120022 B/op   4687 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
       ent:     3.28s      3281510 ns/op    4214 B/op    100 allocs/op
      sqlc:     3.29s      3292224 ns/op    2935 B/op     63 allocs/op
        pg:     3.32s      3317896 ns/op    1890 B/op     10 allocs/op
      gorp:     3.37s      3370382 ns/op    1834 B/op     42 allocs/op
    reform:     3.44s      3441771 ns/op    1845 B/op     51 allocs/op
       raw:     3.50s      3504481 ns/op     770 B/op     13 allocs/op
 sqlboiler:     3.53s      3533927 ns/op    1628 B/op     35 allocs/op
       dbr:     3.57s      3571490 ns/op    2709 B/op     65 allocs/op
     beego:     3.68s      3679622 ns/op    2376 B/op     56 allocs/op
      gorm:     3.74s      3737412 ns/op    6976 B/op     90 allocs/op
       bun:     3.75s      3745253 ns/op    5068 B/op     14 allocs/op
       rel:     3.79s      3794838 ns/op    2504 B/op     42 allocs/op
      godb:     3.92s      3919214 ns/op    4578 B/op    115 allocs/op
      xorm:     4.25s      4248362 ns/op    3344 B/op     89 allocs/op
     upper:     4.43s      4427989 ns/op   13843 B/op    675 allocs/op
       pop:     4.66s      4657351 ns/op   10152 B/op    249 allocs/op

  1000 times - MultiInsert 100 row
       raw:     6.81s      6811252 ns/op  191429 B/op    932 allocs/op
        pg:     6.90s      6896056 ns/op    5487 B/op    112 allocs/op
     beego:     6.98s      6977997 ns/op  179746 B/op   2746 allocs/op
    reform:     7.05s      7046445 ns/op  466480 B/op   2748 allocs/op
       bun:     7.05s      7052686 ns/op   42472 B/op    218 allocs/op
       ent:     7.49s      7486468 ns/op  412167 B/op   4900 allocs/op
      gorm:     7.92s      7924576 ns/op  272238 B/op   3729 allocs/op
       rel:     8.37s      8372648 ns/op  303990 B/op   3263 allocs/op
      godb:     9.01s      9007316 ns/op  260060 B/op   5895 allocs/op
      xorm:    10.29s     10289078 ns/op  255179 B/op   5417 allocs/op
     upper:    12.31s     12311868 ns/op  545459 B/op  19495 allocs/op
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  1000 times - Update
       ent:     0.17s       165678 ns/op    4561 B/op     99 allocs/op
       raw:     0.17s       169520 ns/op     753 B/op     13 allocs/op
      gorp:     3.20s      3199590 ns/op    1209 B/op     32 allocs/op
       dbr:     3.26s      3259352 ns/op    2651 B/op     57 allocs/op
    reform:     3.28s      3283593 ns/op    1777 B/op     51 allocs/op
       pop:     3.46s      3455398 ns/op    6591 B/op    198 allocs/op
       bun:     3.50s      3500472 ns/op    4732 B/op      5 allocs/op
      gorm:     3.50s      3501583 ns/op    6600 B/op     81 allocs/op
        pg:     3.60s      3597592 ns/op     768 B/op      9 allocs/op
     beego:     3.61s      3608103 ns/op    1752 B/op     47 allocs/op
 sqlboiler:     3.70s      3698171 ns/op     907 B/op     17 allocs/op
      xorm:     3.86s      3855258 ns/op    3649 B/op    126 allocs/op
      sqlc:     3.86s      3858753 ns/op     881 B/op     14 allocs/op
       rel:     3.97s      3968831 ns/op    2528 B/op     41 allocs/op
      godb:     4.14s      4143434 ns/op    5161 B/op    154 allocs/op
     upper:     6.77s      6766179 ns/op   33570 B/op   1503 allocs/op

  1000 times - Read
     beego:     0.14s       136168 ns/op    2089 B/op     75 allocs/op
      gorp:     0.14s       137646 ns/op    3881 B/op    194 allocs/op
       bun:     0.14s       138936 ns/op    5493 B/op     21 allocs/op
      sqlc:     0.14s       143865 ns/op    2176 B/op     52 allocs/op
       rel:     0.14s       144795 ns/op    1800 B/op     45 allocs/op
        pg:     0.15s       147126 ns/op     872 B/op     20 allocs/op
       raw:     0.15s       152479 ns/op    2058 B/op     50 allocs/op
 sqlboiler:     0.16s       158449 ns/op     971 B/op     14 allocs/op
       ent:     0.17s       173621 ns/op    5401 B/op    149 allocs/op
       pop:     0.17s       174173 ns/op    3599 B/op     71 allocs/op
       dbr:     0.21s       208544 ns/op    2184 B/op     37 allocs/op
      godb:     0.29s       292068 ns/op    4112 B/op    102 allocs/op
    reform:     0.32s       323246 ns/op    3195 B/op     86 allocs/op
     upper:     0.42s       415778 ns/op    8136 B/op    315 allocs/op
      gorm:     0.44s       444751 ns/op    5160 B/op     94 allocs/op
      xorm:     0.52s       515186 ns/op    4617 B/op    125 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.19s       194686 ns/op    2921 B/op     74 allocs/op
        pg:     0.46s       462610 ns/op   22265 B/op    629 allocs/op
     upper:     0.47s       470344 ns/op    7850 B/op    294 allocs/op
       raw:     0.47s       470561 ns/op   38342 B/op   1038 allocs/op
 sqlboiler:     0.49s       490828 ns/op   58903 B/op   1260 allocs/op
      sqlc:     0.53s       532440 ns/op   73159 B/op   1251 allocs/op
      gorp:     0.71s       706455 ns/op   57386 B/op   1494 allocs/op
       pop:     0.73s       725720 ns/op   76058 B/op   1511 allocs/op
       ent:     0.74s       735015 ns/op   76049 B/op   2039 allocs/op
       bun:     0.85s       850035 ns/op   34019 B/op   1124 allocs/op
       dbr:     0.97s       974918 ns/op   32416 B/op   1545 allocs/op
     beego:     1.03s      1034051 ns/op   55200 B/op   3077 allocs/op
      gorm:     1.29s      1286199 ns/op   57205 B/op   2278 allocs/op
       rel:     1.30s      1300818 ns/op   95321 B/op   2248 allocs/op
      xorm:     1.80s      1798396 ns/op  119998 B/op   4687 allocs/op
      godb:     1.85s      1845555 ns/op   75284 B/op   3084 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
       raw:     6.51s      3254416 ns/op     737 B/op     13 allocs/op
 sqlboiler:     6.62s      3307520 ns/op    1595 B/op     35 allocs/op
      sqlc:     6.63s      3314577 ns/op    2903 B/op     63 allocs/op
       ent:     6.67s      3335150 ns/op    4215 B/op    100 allocs/op
    reform:     6.75s      3377125 ns/op    1811 B/op     51 allocs/op
      gorp:     6.79s      3396968 ns/op    1801 B/op     42 allocs/op
     beego:     6.88s      3438687 ns/op    2377 B/op     56 allocs/op
       bun:     6.90s      3450958 ns/op    5011 B/op     14 allocs/op
      gorm:     7.10s      3549871 ns/op    6953 B/op     90 allocs/op
       dbr:     7.66s      3831927 ns/op    2698 B/op     65 allocs/op
        pg:     7.71s      3856146 ns/op    1342 B/op     10 allocs/op
      xorm:     7.93s      3965390 ns/op    3336 B/op     89 allocs/op
     upper:     8.36s      4178747 ns/op   13800 B/op    674 allocs/op
       rel:     8.67s      4333981 ns/op    2504 B/op     42 allocs/op
      godb:     8.71s      4356965 ns/op    4566 B/op    115 allocs/op
       pop:     9.20s      4599542 ns/op   10103 B/op    249 allocs/op

  2000 times - MultiInsert 100 row
       raw:    12.77s      6382989 ns/op  191315 B/op    931 allocs/op
     beego:    12.79s      6397123 ns/op  179773 B/op   2746 allocs/op
    reform:    13.79s      6895057 ns/op  466470 B/op   2748 allocs/op
        pg:    13.96s      6982496 ns/op    3876 B/op    112 allocs/op
       bun:    14.44s      7219623 ns/op   42441 B/op    219 allocs/op
       ent:    14.49s      7244104 ns/op  412141 B/op   4900 allocs/op
      gorm:    17.34s      8669810 ns/op  272284 B/op   3729 allocs/op
       rel:    18.53s      9263969 ns/op  303993 B/op   3263 allocs/op
      godb:    18.68s      9339545 ns/op  260074 B/op   5895 allocs/op
      xorm:    20.77s     10385097 ns/op  255214 B/op   5417 allocs/op
     upper:    22.26s     11130821 ns/op  545256 B/op  19480 allocs/op
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  2000 times - Update
       ent:     0.27s       134339 ns/op    4557 B/op     99 allocs/op
       raw:     0.32s       161238 ns/op     748 B/op     13 allocs/op
 sqlboiler:     6.68s      3340365 ns/op     902 B/op     17 allocs/op
      sqlc:     6.76s      3379268 ns/op     876 B/op     14 allocs/op
    reform:     6.79s      3393221 ns/op    1773 B/op     51 allocs/op
     beego:     6.82s      3410439 ns/op    1753 B/op     47 allocs/op
       pop:     7.05s      3526231 ns/op    6584 B/op    198 allocs/op
       dbr:     7.16s      3579520 ns/op    2651 B/op     57 allocs/op
      gorm:     7.24s      3620211 ns/op    6604 B/op     81 allocs/op
        pg:     7.26s      3629699 ns/op     768 B/op      9 allocs/op
      gorp:     7.65s      3823163 ns/op    1205 B/op     32 allocs/op
       rel:     7.79s      3894594 ns/op    2528 B/op     41 allocs/op
       bun:     7.91s      3955603 ns/op    4731 B/op      5 allocs/op
      xorm:     8.13s      4064715 ns/op    3649 B/op    126 allocs/op
      godb:     8.85s      4423999 ns/op    5162 B/op    154 allocs/op
     upper:    12.25s      6123706 ns/op   33523 B/op   1503 allocs/op

  2000 times - Read
      gorp:     0.22s       112474 ns/op    3876 B/op    194 allocs/op
       raw:     0.24s       121224 ns/op    2065 B/op     50 allocs/op
        pg:     0.26s       130743 ns/op     872 B/op     20 allocs/op
       pop:     0.29s       147187 ns/op    3586 B/op     71 allocs/op
       bun:     0.30s       149788 ns/op    5492 B/op     21 allocs/op
       rel:     0.30s       150577 ns/op    1800 B/op     45 allocs/op
    reform:     0.31s       152991 ns/op    3198 B/op     86 allocs/op
 sqlboiler:     0.32s       158520 ns/op     956 B/op     14 allocs/op
       ent:     0.35s       174430 ns/op    5405 B/op    149 allocs/op
      sqlc:     0.37s       185700 ns/op    2172 B/op     52 allocs/op
       dbr:     0.41s       204742 ns/op    2184 B/op     37 allocs/op
     beego:     0.51s       257279 ns/op    2089 B/op     75 allocs/op
      gorm:     0.52s       259500 ns/op    5154 B/op     94 allocs/op
      godb:     0.64s       320668 ns/op    4113 B/op    102 allocs/op
     upper:     0.74s       370381 ns/op    8134 B/op    315 allocs/op
      xorm:     0.75s       376563 ns/op    4617 B/op    125 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.56s       278826 ns/op    2926 B/op     74 allocs/op
     upper:     0.66s       329868 ns/op    7841 B/op    294 allocs/op
       raw:     0.90s       449667 ns/op   38342 B/op   1038 allocs/op
      sqlc:     1.00s       498300 ns/op   73158 B/op   1251 allocs/op
        pg:     1.03s       513254 ns/op   22264 B/op    629 allocs/op
 sqlboiler:     1.08s       541683 ns/op   58804 B/op   1260 allocs/op
       ent:     1.21s       602875 ns/op   76049 B/op   2039 allocs/op
       bun:     1.25s       624352 ns/op   34020 B/op   1124 allocs/op
      gorp:     1.41s       706243 ns/op   57389 B/op   1494 allocs/op
       pop:     1.42s       708719 ns/op   75963 B/op   1511 allocs/op
       dbr:     1.46s       731694 ns/op   32416 B/op   1545 allocs/op
     beego:     1.63s       815788 ns/op   55208 B/op   3077 allocs/op
      gorm:     2.51s      1253647 ns/op   57190 B/op   2278 allocs/op
       rel:     3.52s      1762226 ns/op   95321 B/op   2248 allocs/op
      xorm:     3.79s      1894936 ns/op  119993 B/op   4687 allocs/op
      godb:     4.19s      2092958 ns/op   75289 B/op   3084 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
       raw:    13.02s      3255212 ns/op     720 B/op     13 allocs/op
      gorp:    13.16s      3288827 ns/op    1783 B/op     42 allocs/op
      sqlc:    13.18s      3294672 ns/op    2885 B/op     63 allocs/op
 sqlboiler:    13.19s      3298647 ns/op    1576 B/op     35 allocs/op
       dbr:    13.47s      3367938 ns/op    2693 B/op     65 allocs/op
        pg:    13.58s      3393871 ns/op     806 B/op     10 allocs/op
       bun:    14.03s      3506665 ns/op    5014 B/op     14 allocs/op
      gorm:    14.56s      3639400 ns/op    6933 B/op     90 allocs/op
       ent:    14.77s      3693396 ns/op    4216 B/op    100 allocs/op
     beego:    14.91s      3726279 ns/op    2376 B/op     56 allocs/op
    reform:    15.27s      3816701 ns/op    1792 B/op     51 allocs/op
       rel:    15.30s      3824917 ns/op    2504 B/op     42 allocs/op
      xorm:    15.32s      3829436 ns/op    3327 B/op     89 allocs/op
      godb:    15.85s      3962599 ns/op    4559 B/op    115 allocs/op
     upper:    16.78s      4194212 ns/op   13775 B/op    673 allocs/op
       pop:    18.49s      4622682 ns/op   10084 B/op    249 allocs/op

  4000 times - MultiInsert 100 row
       raw:    25.01s      6253405 ns/op  191258 B/op    931 allocs/op
     beego:    25.12s      6279130 ns/op  179717 B/op   2746 allocs/op
        pg:    26.51s      6626510 ns/op    4120 B/op    112 allocs/op
    reform:    27.32s      6831205 ns/op  466461 B/op   2748 allocs/op
       bun:    28.48s      7121107 ns/op   42484 B/op    219 allocs/op
       ent:    29.63s      7406418 ns/op  412129 B/op   4900 allocs/op
      gorm:    32.42s      8105585 ns/op  272300 B/op   3729 allocs/op
       rel:    33.48s      8370838 ns/op  303994 B/op   3263 allocs/op
      godb:    35.75s      8938434 ns/op  260078 B/op   5895 allocs/op
      xorm:    40.19s     10046385 ns/op  255168 B/op   5417 allocs/op
     upper:    46.22s     11555653 ns/op  545167 B/op  19473 allocs/op
      gorp:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.42s       106196 ns/op     750 B/op     13 allocs/op
       ent:     0.56s       140521 ns/op    4558 B/op     99 allocs/op
    reform:    13.42s      3355531 ns/op    1775 B/op     51 allocs/op
      gorp:    13.45s      3362383 ns/op    1204 B/op     32 allocs/op
      sqlc:    13.47s      3368367 ns/op     876 B/op     14 allocs/op
     beego:    13.60s      3399569 ns/op    1753 B/op     47 allocs/op
        pg:    14.26s      3564773 ns/op     768 B/op      9 allocs/op
       bun:    14.33s      3581336 ns/op    4732 B/op      5 allocs/op
      gorm:    14.49s      3621870 ns/op    6604 B/op     81 allocs/op
 sqlboiler:    14.61s      3652173 ns/op     901 B/op     17 allocs/op
       dbr:    14.95s      3737299 ns/op    2651 B/op     57 allocs/op
      xorm:    15.11s      3777118 ns/op    3650 B/op    126 allocs/op
       pop:    15.57s      3892037 ns/op    6586 B/op    198 allocs/op
      godb:    16.16s      4040401 ns/op    5162 B/op    154 allocs/op
       rel:    16.84s      4209090 ns/op    2528 B/op     41 allocs/op
     upper:    25.06s      6263904 ns/op   33503 B/op   1502 allocs/op

  4000 times - Read
      sqlc:     0.49s       122556 ns/op    2172 B/op     52 allocs/op
       raw:     0.49s       123397 ns/op    2062 B/op     50 allocs/op
    reform:     0.51s       128650 ns/op    3199 B/op     86 allocs/op
       pop:     0.56s       139620 ns/op    3555 B/op     71 allocs/op
       bun:     0.57s       142709 ns/op    5495 B/op     21 allocs/op
       rel:     0.59s       146392 ns/op    1800 B/op     45 allocs/op
 sqlboiler:     0.63s       157321 ns/op     965 B/op     14 allocs/op
     beego:     0.63s       158172 ns/op    2088 B/op     75 allocs/op
        pg:     0.64s       160565 ns/op     872 B/op     20 allocs/op
       ent:     0.67s       168133 ns/op    5404 B/op    149 allocs/op
       dbr:     0.70s       174635 ns/op    2184 B/op     37 allocs/op
      gorp:     0.72s       180514 ns/op    3878 B/op    194 allocs/op
      gorm:     0.91s       227725 ns/op    5152 B/op     94 allocs/op
      godb:     1.22s       306069 ns/op    4113 B/op    102 allocs/op
      xorm:     1.31s       326767 ns/op    4618 B/op    125 allocs/op
     upper:     2.18s       545960 ns/op    8136 B/op    315 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.72s       180644 ns/op    2927 B/op     74 allocs/op
       raw:     1.48s       370397 ns/op   38342 B/op   1038 allocs/op
     upper:     1.53s       382464 ns/op    7841 B/op    294 allocs/op
        pg:     1.87s       466911 ns/op   23856 B/op    629 allocs/op
      sqlc:     1.88s       468887 ns/op   73158 B/op   1251 allocs/op
 sqlboiler:     2.06s       516048 ns/op   58594 B/op   1260 allocs/op
       ent:     2.35s       587629 ns/op   76045 B/op   2039 allocs/op
      gorp:     2.37s       593522 ns/op   57389 B/op   1494 allocs/op
       pop:     2.74s       684419 ns/op   75921 B/op   1511 allocs/op
       bun:     2.80s       699047 ns/op   34025 B/op   1124 allocs/op
       dbr:     2.85s       711905 ns/op   32416 B/op   1545 allocs/op
     beego:     2.87s       717417 ns/op   55201 B/op   3077 allocs/op
      gorm:     3.66s       913801 ns/op   57181 B/op   2278 allocs/op
       rel:     6.59s      1647903 ns/op   95321 B/op   2248 allocs/op
      godb:     7.41s      1851727 ns/op   75289 B/op   3084 allocs/op
      xorm:     7.54s      1884039 ns/op  120000 B/op   4687 allocs/op
```