# Results

- orm-benchmark (with no flags)

```
   
Reports:

   200 times - Insert
       pgx:     0.06s       288218 ns/op     244 B/op     10 allocs/op
  pgx_pool:     0.06s       292428 ns/op     265 B/op     10 allocs/op
     beego:     0.06s       308748 ns/op    2373 B/op     56 allocs/op
       raw:     0.07s       332530 ns/op    1001 B/op     15 allocs/op
 gorm_prep:     0.07s       335071 ns/op    5616 B/op     70 allocs/op
       ent:     0.07s       335864 ns/op    4116 B/op     97 allocs/op
       bun:     0.07s       348477 ns/op    5290 B/op     14 allocs/op
        pg:     0.07s       356719 ns/op    6273 B/op     10 allocs/op
      sqlc:     0.07s       364614 ns/op    3174 B/op     64 allocs/op
      gorp:     0.07s       365210 ns/op    2066 B/op     43 allocs/op
      gorm:     0.08s       380367 ns/op    6953 B/op     90 allocs/op
 sqlboiler:     0.08s       389067 ns/op    1870 B/op     36 allocs/op
    reform:     0.08s       390432 ns/op    2091 B/op     53 allocs/op
       dbr:     0.08s       406356 ns/op    2799 B/op     66 allocs/op
      sqlx:     0.09s       449923 ns/op     856 B/op     19 allocs/op
      zorm:     0.10s       491741 ns/op    3579 B/op     63 allocs/op
     upper:     0.11s       531634 ns/op    6208 B/op    130 allocs/op
       rel:     0.11s       542622 ns/op    2590 B/op     43 allocs/op
      godb:     0.11s       547584 ns/op    4708 B/op    116 allocs/op
      xorm:     0.11s       551200 ns/op    3443 B/op     90 allocs/op
       pop:     0.16s       792238 ns/op    9962 B/op    238 allocs/op

   200 times - MultiInsert 100 row
  pgx_pool:     0.22s      1078759 ns/op  112939 B/op     43 allocs/op
       pgx:     0.22s      1084597 ns/op  112983 B/op     43 allocs/op
    reform:     0.28s      1392798 ns/op  458909 B/op   2746 allocs/op
       raw:     0.29s      1458355 ns/op  184801 B/op    936 allocs/op
 gorm_prep:     0.29s      1467007 ns/op  227782 B/op   2279 allocs/op
     beego:     0.30s      1488197 ns/op  177749 B/op   2745 allocs/op
        pg:     0.34s      1707027 ns/op    8929 B/op    112 allocs/op
       bun:     0.35s      1745680 ns/op   42487 B/op    217 allocs/op
      sqlx:     0.39s      1955526 ns/op  170541 B/op   1551 allocs/op
       ent:     0.40s      2019994 ns/op  387612 B/op   4600 allocs/op
      gorm:     0.43s      2161288 ns/op  271113 B/op   3827 allocs/op
      godb:     0.48s      2408660 ns/op  255244 B/op   5892 allocs/op
      zorm:     0.49s      2441007 ns/op  199696 B/op   2767 allocs/op
     upper:     0.49s      2446815 ns/op  323194 B/op   4218 allocs/op
      xorm:     0.49s      2469398 ns/op  248125 B/op   5414 allocs/op
       rel:     0.52s      2604851 ns/op  306853 B/op   3261 allocs/op
      sqlc:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert

   200 times - Update
       raw:     0.03s       128427 ns/op     722 B/op     13 allocs/op
       ent:     0.03s       146657 ns/op    4618 B/op     98 allocs/op
  pgx_pool:     0.05s       272902 ns/op     302 B/op     10 allocs/op
      sqlx:     0.06s       288572 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.06s       298166 ns/op     924 B/op     17 allocs/op
     beego:     0.06s       315879 ns/op    1755 B/op     47 allocs/op
      sqlc:     0.06s       320962 ns/op     890 B/op     14 allocs/op
       pgx:     0.06s       323842 ns/op     282 B/op     10 allocs/op
      gorp:     0.07s       326973 ns/op    1223 B/op     32 allocs/op
 gorm_prep:     0.07s       339500 ns/op    5132 B/op     59 allocs/op
    reform:     0.07s       345597 ns/op    1747 B/op     51 allocs/op
       bun:     0.07s       361674 ns/op    4733 B/op      5 allocs/op
       dbr:     0.08s       380147 ns/op    2651 B/op     57 allocs/op
        pg:     0.08s       392705 ns/op     768 B/op      9 allocs/op
       pop:     0.08s       418039 ns/op    6019 B/op    186 allocs/op
      gorm:     0.08s       423767 ns/op    6625 B/op     81 allocs/op
      xorm:     0.10s       476808 ns/op    3944 B/op    132 allocs/op
      zorm:     0.10s       485742 ns/op    2872 B/op     51 allocs/op
      godb:     0.11s       533494 ns/op    5161 B/op    154 allocs/op
       rel:     0.11s       536177 ns/op    3027 B/op     44 allocs/op
     upper:     0.29s      1451682 ns/op   16728 B/op    391 allocs/op

   200 times - Read
  pgx_pool:     0.03s       127738 ns/op    1078 B/op      9 allocs/op
       pgx:     0.03s       130754 ns/op     887 B/op      8 allocs/op
      sqlc:     0.03s       136418 ns/op    2162 B/op     52 allocs/op
    reform:     0.03s       140318 ns/op    3232 B/op     86 allocs/op
     beego:     0.03s       143649 ns/op    2096 B/op     75 allocs/op
      gorp:     0.03s       145092 ns/op    3869 B/op    194 allocs/op
       pop:     0.03s       147520 ns/op    3572 B/op     67 allocs/op
        pg:     0.03s       149049 ns/op     873 B/op     20 allocs/op
 gorm_prep:     0.03s       152283 ns/op    4817 B/op     85 allocs/op
       ent:     0.03s       161192 ns/op    5426 B/op    150 allocs/op
       dbr:     0.03s       162338 ns/op    2184 B/op     37 allocs/op
       bun:     0.03s       162744 ns/op    5522 B/op     21 allocs/op
 sqlboiler:     0.03s       164566 ns/op    1116 B/op     14 allocs/op
       rel:     0.03s       166981 ns/op    2315 B/op     49 allocs/op
       raw:     0.03s       168982 ns/op    2098 B/op     50 allocs/op
      zorm:     0.04s       177516 ns/op    3024 B/op     63 allocs/op
      gorm:     0.04s       199754 ns/op    5153 B/op     96 allocs/op
      sqlx:     0.06s       284970 ns/op    1976 B/op     43 allocs/op
      xorm:     0.06s       300494 ns/op    4648 B/op    127 allocs/op
      godb:     0.06s       304218 ns/op    4112 B/op    102 allocs/op
     upper:     0.06s       322141 ns/op    5041 B/op    110 allocs/op

   200 times - MultiRead limit 100
    reform:     0.03s       136756 ns/op    2955 B/op     74 allocs/op
  pgx_pool:     0.05s       257077 ns/op   42993 B/op    504 allocs/op
       pgx:     0.05s       269928 ns/op   42919 B/op    504 allocs/op
       raw:     0.06s       314042 ns/op   38352 B/op   1038 allocs/op
     upper:     0.07s       360326 ns/op    4820 B/op     89 allocs/op
        pg:     0.07s       369103 ns/op   22258 B/op    629 allocs/op
      sqlc:     0.08s       401163 ns/op   73128 B/op   1251 allocs/op
      gorp:     0.08s       416928 ns/op   57385 B/op   1494 allocs/op
       pop:     0.09s       430371 ns/op   68773 B/op   1307 allocs/op
       ent:     0.09s       446087 ns/op   76126 B/op   2040 allocs/op
      sqlx:     0.09s       464533 ns/op   39064 B/op   1516 allocs/op
       bun:     0.10s       476843 ns/op   34034 B/op   1124 allocs/op
       dbr:     0.10s       487065 ns/op   32416 B/op   1545 allocs/op
 sqlboiler:     0.10s       521767 ns/op   66033 B/op   2260 allocs/op
     beego:     0.11s       536125 ns/op   55194 B/op   3077 allocs/op
      gorm:     0.12s       582023 ns/op   57181 B/op   2379 allocs/op
 gorm_prep:     0.12s       614838 ns/op   56083 B/op   2080 allocs/op
      zorm:     0.14s       719462 ns/op  163185 B/op   3237 allocs/op
      godb:     0.16s       788250 ns/op   75269 B/op   3084 allocs/op
      xorm:     0.16s       792235 ns/op  120208 B/op   4692 allocs/op
       rel:     0.16s       792658 ns/op  100616 B/op   2252 allocs/op
```

- orm-benchmark -multi=5

```
  
Reports:

  1000 times - Insert
  pgx_pool:     0.30s       301308 ns/op     281 B/op     10 allocs/op
       pgx:     0.31s       310009 ns/op     265 B/op     10 allocs/op
       raw:     0.32s       317333 ns/op     764 B/op     13 allocs/op
 sqlboiler:     0.33s       325095 ns/op    1621 B/op     35 allocs/op
     beego:     0.34s       337910 ns/op    2375 B/op     56 allocs/op
 gorm_prep:     0.34s       343628 ns/op    5312 B/op     68 allocs/op
      gorp:     0.35s       345889 ns/op    1827 B/op     42 allocs/op
    reform:     0.35s       351201 ns/op    1836 B/op     51 allocs/op
       ent:     0.36s       357610 ns/op    4118 B/op     98 allocs/op
      sqlc:     0.36s       358690 ns/op    2928 B/op     63 allocs/op
        pg:     0.36s       361165 ns/op    1889 B/op     10 allocs/op
       dbr:     0.37s       366148 ns/op    2711 B/op     65 allocs/op
       bun:     0.39s       392081 ns/op    5052 B/op     14 allocs/op
      gorm:     0.40s       401754 ns/op    6983 B/op     90 allocs/op
      sqlx:     0.48s       475955 ns/op     856 B/op     19 allocs/op
      zorm:     0.49s       494254 ns/op    3569 B/op     64 allocs/op
      xorm:     0.51s       505737 ns/op    3344 B/op     89 allocs/op
     upper:     0.51s       512385 ns/op    5944 B/op    126 allocs/op
       rel:     0.52s       520836 ns/op    2584 B/op     44 allocs/op
      godb:     0.53s       530191 ns/op    4575 B/op    116 allocs/op
       pop:     0.70s       702280 ns/op    9588 B/op    237 allocs/op

  1000 times - MultiInsert 100 row
  pgx_pool:     1.12s      1117424 ns/op  112951 B/op     43 allocs/op
       pgx:     1.13s      1129315 ns/op  112976 B/op     43 allocs/op
       raw:     1.28s      1276575 ns/op  183878 B/op    930 allocs/op
     beego:     1.34s      1341589 ns/op  177651 B/op   2745 allocs/op
 gorm_prep:     1.53s      1533648 ns/op  227705 B/op   2281 allocs/op
    reform:     1.56s      1559633 ns/op  458786 B/op   2745 allocs/op
        pg:     1.74s      1737976 ns/op    4437 B/op    112 allocs/op
       ent:     1.82s      1817903 ns/op  387396 B/op   4598 allocs/op
       bun:     1.83s      1827828 ns/op   42416 B/op    218 allocs/op
      sqlx:     2.02s      2016918 ns/op  170116 B/op   1551 allocs/op
      gorm:     2.08s      2080665 ns/op  271500 B/op   3829 allocs/op
      zorm:     2.30s      2295446 ns/op  199696 B/op   2767 allocs/op
      xorm:     2.34s      2339171 ns/op  248054 B/op   5414 allocs/op
       rel:     2.39s      2392962 ns/op  306869 B/op   3263 allocs/op
      godb:     2.44s      2439242 ns/op  255242 B/op   5894 allocs/op
     upper:     2.48s      2482581 ns/op  322949 B/op   4207 allocs/op
 sqlboiler:     doesn't support bulk-insert
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  1000 times - Update
       raw:     0.12s       121338 ns/op     753 B/op     13 allocs/op
       ent:     0.16s       160569 ns/op    4609 B/op     98 allocs/op
      sqlx:     0.28s       278525 ns/op     872 B/op     20 allocs/op
       pgx:     0.31s       310886 ns/op     272 B/op     10 allocs/op
      sqlc:     0.32s       319122 ns/op     881 B/op     14 allocs/op
 gorm_prep:     0.32s       323182 ns/op    5113 B/op     59 allocs/op
  pgx_pool:     0.33s       330588 ns/op     287 B/op     10 allocs/op
 sqlboiler:     0.34s       339045 ns/op     907 B/op     17 allocs/op
    reform:     0.34s       340033 ns/op    1776 B/op     51 allocs/op
     beego:     0.34s       341075 ns/op    1752 B/op     47 allocs/op
       dbr:     0.34s       344240 ns/op    2651 B/op     57 allocs/op
       pop:     0.36s       361596 ns/op    6051 B/op    186 allocs/op
        pg:     0.37s       367950 ns/op     768 B/op      9 allocs/op
      gorp:     0.37s       372293 ns/op    1209 B/op     32 allocs/op
       bun:     0.39s       394314 ns/op    4728 B/op      5 allocs/op
      gorm:     0.41s       408548 ns/op    6600 B/op     81 allocs/op
      xorm:     0.51s       506738 ns/op    3944 B/op    132 allocs/op
       rel:     0.52s       517481 ns/op    3024 B/op     44 allocs/op
      zorm:     0.52s       519464 ns/op    2872 B/op     51 allocs/op
      godb:     0.59s       588193 ns/op    5160 B/op    154 allocs/op
     upper:     1.15s      1153561 ns/op   16636 B/op    390 allocs/op

  1000 times - Read
       pgx:     0.12s       116838 ns/op     889 B/op      8 allocs/op
  pgx_pool:     0.12s       122644 ns/op    1076 B/op      9 allocs/op
      sqlc:     0.13s       132190 ns/op    2176 B/op     52 allocs/op
       raw:     0.14s       136207 ns/op    2057 B/op     50 allocs/op
     beego:     0.14s       136289 ns/op    2089 B/op     75 allocs/op
    reform:     0.14s       141815 ns/op    3193 B/op     86 allocs/op
       pop:     0.15s       145114 ns/op    3254 B/op     67 allocs/op
      gorp:     0.15s       145262 ns/op    3881 B/op    194 allocs/op
        pg:     0.15s       148843 ns/op     872 B/op     20 allocs/op
 sqlboiler:     0.15s       153443 ns/op     971 B/op     14 allocs/op
 gorm_prep:     0.16s       157881 ns/op    4820 B/op     85 allocs/op
       rel:     0.16s       159918 ns/op    2312 B/op     49 allocs/op
       ent:     0.16s       164895 ns/op    5448 B/op    150 allocs/op
       dbr:     0.17s       167189 ns/op    2184 B/op     37 allocs/op
       bun:     0.17s       168254 ns/op    5520 B/op     21 allocs/op
      zorm:     0.17s       173538 ns/op    3024 B/op     63 allocs/op
      gorm:     0.20s       198160 ns/op    5157 B/op     96 allocs/op
      sqlx:     0.27s       271027 ns/op    1976 B/op     43 allocs/op
     upper:     0.30s       303902 ns/op    5055 B/op    110 allocs/op
      godb:     0.31s       309591 ns/op    4112 B/op    102 allocs/op
      xorm:     0.32s       320204 ns/op    4648 B/op    127 allocs/op

  1000 times - MultiRead limit 100
    reform:     0.14s       137480 ns/op    2920 B/op     74 allocs/op
       pgx:     0.26s       255189 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.28s       277344 ns/op   43009 B/op    504 allocs/op
     upper:     0.30s       304443 ns/op    4783 B/op     89 allocs/op
       raw:     0.35s       353420 ns/op   38342 B/op   1038 allocs/op
      sqlc:     0.38s       376593 ns/op   73158 B/op   1251 allocs/op
        pg:     0.39s       387992 ns/op   22258 B/op    629 allocs/op
      gorp:     0.42s       416951 ns/op   57376 B/op   1494 allocs/op
       pop:     0.43s       425264 ns/op   68674 B/op   1306 allocs/op
      sqlx:     0.44s       437359 ns/op   39064 B/op   1516 allocs/op
       ent:     0.45s       445709 ns/op   76097 B/op   2040 allocs/op
       dbr:     0.49s       486526 ns/op   32416 B/op   1545 allocs/op
       bun:     0.49s       488834 ns/op   34025 B/op   1124 allocs/op
     beego:     0.50s       498300 ns/op   55186 B/op   3077 allocs/op
 sqlboiler:     0.55s       550929 ns/op   66543 B/op   2260 allocs/op
 gorm_prep:     0.57s       568956 ns/op   56065 B/op   2079 allocs/op
      gorm:     0.61s       610644 ns/op   57188 B/op   2379 allocs/op
      zorm:     0.73s       728000 ns/op  163185 B/op   3237 allocs/op
       rel:     0.74s       740989 ns/op  100616 B/op   2252 allocs/op
      xorm:     0.81s       813190 ns/op  120181 B/op   4692 allocs/op
      godb:     0.82s       820185 ns/op   75267 B/op   3084 allocs/op
```

- orm-benchmark -multi=10

```
  
Reports:

  2000 times - Insert
     beego:     0.62s       308585 ns/op    2375 B/op     56 allocs/op
       raw:     0.63s       313321 ns/op     734 B/op     13 allocs/op
       pgx:     0.63s       315178 ns/op     269 B/op     10 allocs/op
  pgx_pool:     0.64s       319990 ns/op     285 B/op     10 allocs/op
 gorm_prep:     0.67s       334833 ns/op    5299 B/op     68 allocs/op
       ent:     0.68s       340463 ns/op    4119 B/op     98 allocs/op
      gorp:     0.69s       344323 ns/op    1798 B/op     42 allocs/op
 sqlboiler:     0.70s       349510 ns/op    1592 B/op     35 allocs/op
    reform:     0.70s       352449 ns/op    1807 B/op     51 allocs/op
      sqlc:     0.72s       360523 ns/op    2900 B/op     63 allocs/op
        pg:     0.72s       362347 ns/op     817 B/op     10 allocs/op
       dbr:     0.74s       368570 ns/op    2699 B/op     65 allocs/op
       bun:     0.77s       382930 ns/op    5023 B/op     14 allocs/op
      gorm:     0.84s       419043 ns/op    6960 B/op     90 allocs/op
      sqlx:     0.93s       463751 ns/op     856 B/op     19 allocs/op
       rel:     0.97s       483551 ns/op    2584 B/op     44 allocs/op
      xorm:     1.01s       505163 ns/op    3332 B/op     89 allocs/op
      zorm:     1.01s       506301 ns/op    3568 B/op     64 allocs/op
     upper:     1.06s       531028 ns/op    5914 B/op    126 allocs/op
      godb:     1.07s       536448 ns/op    4564 B/op    116 allocs/op
       pop:     1.41s       704283 ns/op    9568 B/op    237 allocs/op

  2000 times - MultiInsert 100 row
  pgx_pool:     2.20s      1101453 ns/op  112948 B/op     43 allocs/op
       pgx:     2.23s      1116291 ns/op  112926 B/op     43 allocs/op
       raw:     2.61s      1303008 ns/op  183763 B/op    929 allocs/op
     beego:     2.61s      1306603 ns/op  177599 B/op   2745 allocs/op
    reform:     2.94s      1469534 ns/op  458812 B/op   2746 allocs/op
 gorm_prep:     2.94s      1471604 ns/op  227688 B/op   2281 allocs/op
        pg:     3.48s      1741177 ns/op    4363 B/op    112 allocs/op
       ent:     3.49s      1744539 ns/op  387371 B/op   4598 allocs/op
       bun:     3.68s      1838299 ns/op   42431 B/op    218 allocs/op
      sqlx:     3.96s      1980629 ns/op  170711 B/op   1551 allocs/op
      gorm:     4.15s      2075483 ns/op  271587 B/op   3829 allocs/op
      xorm:     4.63s      2313284 ns/op  248055 B/op   5414 allocs/op
      zorm:     4.65s      2324915 ns/op  199696 B/op   2767 allocs/op
       rel:     4.73s      2363592 ns/op  306871 B/op   3263 allocs/op
      godb:     4.95s      2473256 ns/op  255250 B/op   5894 allocs/op
     upper:     5.00s      2499999 ns/op  322915 B/op   4205 allocs/op
       dbr:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert

  2000 times - Update
       raw:     0.24s       122312 ns/op     748 B/op     13 allocs/op
       ent:     0.29s       145675 ns/op    4608 B/op     98 allocs/op
      sqlx:     0.57s       286412 ns/op     872 B/op     20 allocs/op
     beego:     0.60s       300401 ns/op    1753 B/op     47 allocs/op
       pgx:     0.62s       310680 ns/op     268 B/op     10 allocs/op
 sqlboiler:     0.63s       316386 ns/op     901 B/op     17 allocs/op
      gorp:     0.65s       325199 ns/op    1205 B/op     32 allocs/op
  pgx_pool:     0.65s       325557 ns/op     285 B/op     10 allocs/op
      sqlc:     0.66s       330942 ns/op     876 B/op     14 allocs/op
    reform:     0.68s       340866 ns/op    1772 B/op     51 allocs/op
       pop:     0.72s       362486 ns/op    6047 B/op    186 allocs/op
        pg:     0.74s       368247 ns/op     768 B/op      9 allocs/op
 gorm_prep:     0.74s       368339 ns/op    5121 B/op     59 allocs/op
       dbr:     0.74s       369367 ns/op    2651 B/op     57 allocs/op
       bun:     0.77s       383124 ns/op    4729 B/op      5 allocs/op
      gorm:     0.85s       423775 ns/op    6604 B/op     81 allocs/op
      xorm:     0.98s       491757 ns/op    3944 B/op    132 allocs/op
      zorm:     0.98s       491875 ns/op    2872 B/op     51 allocs/op
       rel:     1.06s       531360 ns/op    3024 B/op     44 allocs/op
      godb:     1.07s       532867 ns/op    5161 B/op    154 allocs/op
     upper:     2.33s      1165450 ns/op   16624 B/op    390 allocs/op

  2000 times - Read
       pgx:     0.24s       119492 ns/op     892 B/op      8 allocs/op
  pgx_pool:     0.24s       121239 ns/op    1080 B/op      9 allocs/op
      sqlc:     0.28s       138570 ns/op    2172 B/op     52 allocs/op
     beego:     0.28s       140086 ns/op    2088 B/op     75 allocs/op
       raw:     0.28s       140513 ns/op    2065 B/op     50 allocs/op
    reform:     0.29s       144967 ns/op    3197 B/op     86 allocs/op
      gorp:     0.30s       148108 ns/op    3876 B/op    194 allocs/op
       pop:     0.30s       152326 ns/op    3222 B/op     67 allocs/op
        pg:     0.32s       159726 ns/op     872 B/op     20 allocs/op
       bun:     0.32s       161474 ns/op    5522 B/op     21 allocs/op
 gorm_prep:     0.33s       163124 ns/op    4816 B/op     85 allocs/op
 sqlboiler:     0.33s       163279 ns/op     974 B/op     14 allocs/op
       ent:     0.33s       163423 ns/op    5452 B/op    150 allocs/op
       rel:     0.34s       168996 ns/op    2312 B/op     49 allocs/op
       dbr:     0.35s       173010 ns/op    2184 B/op     37 allocs/op
      zorm:     0.37s       182981 ns/op    3024 B/op     63 allocs/op
      gorm:     0.41s       204365 ns/op    5152 B/op     96 allocs/op
      sqlx:     0.57s       286319 ns/op    1976 B/op     43 allocs/op
      xorm:     0.62s       310348 ns/op    4648 B/op    127 allocs/op
     upper:     0.63s       312764 ns/op    5054 B/op    110 allocs/op
      godb:     0.63s       315995 ns/op    4112 B/op    102 allocs/op

  2000 times - MultiRead limit 100
    reform:     0.27s       134961 ns/op    2924 B/op     74 allocs/op
       pgx:     0.53s       263769 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     0.54s       271652 ns/op   43009 B/op    504 allocs/op
     upper:     0.63s       314136 ns/op    4776 B/op     89 allocs/op
       raw:     0.66s       330877 ns/op   38342 B/op   1038 allocs/op
        pg:     0.70s       351680 ns/op   22821 B/op    629 allocs/op
      sqlc:     0.79s       393485 ns/op   73158 B/op   1251 allocs/op
      gorp:     0.82s       407874 ns/op   57373 B/op   1494 allocs/op
      sqlx:     0.83s       415995 ns/op   39064 B/op   1516 allocs/op
       ent:     0.87s       436347 ns/op   76093 B/op   2040 allocs/op
       pop:     0.88s       439123 ns/op   68656 B/op   1306 allocs/op
       dbr:     0.95s       472796 ns/op   32416 B/op   1545 allocs/op
     beego:     0.97s       486268 ns/op   55188 B/op   3077 allocs/op
       bun:     0.97s       486692 ns/op   34030 B/op   1124 allocs/op
 sqlboiler:     1.05s       523815 ns/op   66232 B/op   2259 allocs/op
 gorm_prep:     1.14s       572207 ns/op   56060 B/op   2079 allocs/op
      gorm:     1.25s       625416 ns/op   57174 B/op   2379 allocs/op
      zorm:     1.43s       717197 ns/op  163185 B/op   3237 allocs/op
       rel:     1.48s       740479 ns/op  100616 B/op   2252 allocs/op
      godb:     1.53s       763958 ns/op   75268 B/op   3084 allocs/op
      xorm:     1.62s       809025 ns/op  120177 B/op   4692 allocs/op
```

- orm-benchmark -multi=20

```
  
Reports:

  4000 times - Insert
       pgx:     1.19s       297190 ns/op     270 B/op     10 allocs/op
       raw:     1.22s       305228 ns/op     719 B/op     13 allocs/op
  pgx_pool:     1.24s       310789 ns/op     286 B/op     10 allocs/op
    reform:     1.30s       324223 ns/op    1790 B/op     51 allocs/op
 sqlboiler:     1.32s       328874 ns/op    1574 B/op     35 allocs/op
     beego:     1.35s       338476 ns/op    2376 B/op     56 allocs/op
      sqlc:     1.36s       340464 ns/op    2884 B/op     63 allocs/op
      gorp:     1.37s       342145 ns/op    1782 B/op     42 allocs/op
       ent:     1.37s       342487 ns/op    4120 B/op     98 allocs/op
       dbr:     1.42s       354542 ns/op    2693 B/op     65 allocs/op
        pg:     1.42s       355602 ns/op     805 B/op     10 allocs/op
 gorm_prep:     1.42s       356229 ns/op    5275 B/op     68 allocs/op
      gorm:     1.62s       406242 ns/op    6947 B/op     90 allocs/op
       bun:     1.63s       407442 ns/op    5008 B/op     14 allocs/op
      sqlx:     1.88s       470188 ns/op     856 B/op     19 allocs/op
       rel:     1.99s       497495 ns/op    2584 B/op     44 allocs/op
      godb:     2.06s       514866 ns/op    4558 B/op    116 allocs/op
      zorm:     2.06s       515033 ns/op    3568 B/op     64 allocs/op
      xorm:     2.08s       519718 ns/op    3326 B/op     89 allocs/op
     upper:     2.09s       522665 ns/op    5896 B/op    126 allocs/op
       pop:     2.85s       712810 ns/op    9537 B/op    237 allocs/op

  4000 times - MultiInsert 100 row
       pgx:     4.03s      1008675 ns/op  112917 B/op     43 allocs/op
  pgx_pool:     4.21s      1053306 ns/op  112944 B/op     43 allocs/op
       raw:     4.59s      1148386 ns/op  183706 B/op    929 allocs/op
     beego:     5.04s      1259028 ns/op  177629 B/op   2745 allocs/op
 gorm_prep:     5.70s      1424409 ns/op  227686 B/op   2281 allocs/op
    reform:     5.80s      1450976 ns/op  458783 B/op   2745 allocs/op
        pg:     6.54s      1634556 ns/op    3576 B/op    112 allocs/op
       ent:     6.81s      1702440 ns/op  387356 B/op   4598 allocs/op
       bun:     6.89s      1721412 ns/op   42403 B/op    219 allocs/op
      sqlx:     7.88s      1969890 ns/op  170586 B/op   1551 allocs/op
      gorm:     7.98s      1994370 ns/op  271625 B/op   3829 allocs/op
      zorm:     8.80s      2199182 ns/op  199696 B/op   2767 allocs/op
      xorm:     9.31s      2326913 ns/op  248025 B/op   5414 allocs/op
       rel:     9.45s      2361905 ns/op  306872 B/op   3263 allocs/op
      godb:     9.46s      2365321 ns/op  255232 B/op   5894 allocs/op
     upper:     9.54s      2386234 ns/op  322903 B/op   4205 allocs/op
       dbr:     doesn't support bulk-insert
 sqlboiler:     doesn't support bulk-insert
      gorp:     doesn't support bulk-insert
      sqlc:     doesn't support bulk-insert
       pop:     doesn't support bulk-insert

  4000 times - Update
       raw:     0.48s       119276 ns/op     750 B/op     13 allocs/op
       ent:     0.59s       146612 ns/op    4606 B/op     98 allocs/op
      sqlx:     1.13s       281512 ns/op     872 B/op     20 allocs/op
      sqlc:     1.22s       305081 ns/op     878 B/op     14 allocs/op
  pgx_pool:     1.22s       306242 ns/op     286 B/op     10 allocs/op
 sqlboiler:     1.25s       311615 ns/op     901 B/op     17 allocs/op
    reform:     1.26s       315557 ns/op    1775 B/op     51 allocs/op
 gorm_prep:     1.27s       316955 ns/op    5116 B/op     59 allocs/op
       pgx:     1.30s       324184 ns/op     270 B/op     10 allocs/op
      gorp:     1.32s       329623 ns/op    1204 B/op     32 allocs/op
     beego:     1.37s       343347 ns/op    1752 B/op     47 allocs/op
       pop:     1.39s       348702 ns/op    6046 B/op    186 allocs/op
       dbr:     1.52s       379205 ns/op    2651 B/op     57 allocs/op
        pg:     1.52s       380612 ns/op     768 B/op      9 allocs/op
       bun:     1.56s       389276 ns/op    4728 B/op      5 allocs/op
      gorm:     1.68s       419388 ns/op    6604 B/op     81 allocs/op
       rel:     2.03s       508129 ns/op    3024 B/op     44 allocs/op
      godb:     2.11s       526271 ns/op    5160 B/op    154 allocs/op
      xorm:     2.11s       527579 ns/op    3944 B/op    132 allocs/op
      zorm:     2.16s       538985 ns/op    2872 B/op     51 allocs/op
     upper:     4.66s      1164592 ns/op   16621 B/op    390 allocs/op

  4000 times - Read
       pgx:     0.50s       124188 ns/op     894 B/op      8 allocs/op
  pgx_pool:     0.50s       126151 ns/op    1079 B/op      9 allocs/op
       raw:     0.51s       128045 ns/op    2062 B/op     50 allocs/op
      sqlc:     0.54s       135029 ns/op    2172 B/op     52 allocs/op
     beego:     0.56s       138760 ns/op    2088 B/op     75 allocs/op
    reform:     0.56s       140150 ns/op    3199 B/op     86 allocs/op
       pop:     0.58s       146196 ns/op    3181 B/op     67 allocs/op
      gorp:     0.59s       148310 ns/op    3878 B/op    194 allocs/op
        pg:     0.62s       155051 ns/op     872 B/op     20 allocs/op
 gorm_prep:     0.64s       159484 ns/op    4816 B/op     85 allocs/op
       ent:     0.65s       161684 ns/op    5452 B/op    150 allocs/op
 sqlboiler:     0.65s       162555 ns/op     956 B/op     14 allocs/op
       bun:     0.66s       164072 ns/op    5521 B/op     21 allocs/op
       rel:     0.66s       166236 ns/op    2312 B/op     49 allocs/op
       dbr:     0.67s       168521 ns/op    2184 B/op     37 allocs/op
      zorm:     0.73s       181706 ns/op    3024 B/op     63 allocs/op
      gorm:     0.79s       196946 ns/op    5148 B/op     96 allocs/op
      sqlx:     1.15s       286977 ns/op    1976 B/op     43 allocs/op
      godb:     1.19s       296962 ns/op    4112 B/op    102 allocs/op
     upper:     1.23s       308675 ns/op    5055 B/op    110 allocs/op
      xorm:     1.24s       310436 ns/op    4648 B/op    127 allocs/op

  4000 times - MultiRead limit 100
    reform:     0.52s       129211 ns/op    2926 B/op     74 allocs/op
       pgx:     1.05s       261658 ns/op   42949 B/op    504 allocs/op
  pgx_pool:     1.06s       265681 ns/op   43007 B/op    504 allocs/op
     upper:     1.27s       316530 ns/op    4777 B/op     89 allocs/op
       raw:     1.28s       318893 ns/op   38341 B/op   1038 allocs/op
        pg:     1.48s       370925 ns/op   22540 B/op    629 allocs/op
      sqlc:     1.51s       378040 ns/op   73158 B/op   1251 allocs/op
      gorp:     1.66s       414557 ns/op   57370 B/op   1494 allocs/op
       pop:     1.67s       417057 ns/op   68638 B/op   1306 allocs/op
       ent:     1.73s       432800 ns/op   76094 B/op   2040 allocs/op
      sqlx:     1.73s       433062 ns/op   39064 B/op   1516 allocs/op
       dbr:     1.83s       458296 ns/op   32416 B/op   1545 allocs/op
       bun:     1.89s       472394 ns/op   34026 B/op   1124 allocs/op
     beego:     1.99s       497784 ns/op   55188 B/op   3077 allocs/op
 sqlboiler:     2.01s       501611 ns/op   66295 B/op   2259 allocs/op
 gorm_prep:     2.26s       565648 ns/op   56058 B/op   2079 allocs/op
      gorm:     2.47s       616833 ns/op   57161 B/op   2379 allocs/op
      zorm:     2.84s       708800 ns/op  163185 B/op   3237 allocs/op
       rel:     2.98s       743931 ns/op  100616 B/op   2252 allocs/op
      godb:     2.99s       748189 ns/op   75265 B/op   3084 allocs/op
      xorm:     3.12s       780087 ns/op  120175 B/op   4692 allocs/op
```
