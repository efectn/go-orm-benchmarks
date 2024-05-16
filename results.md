# Results

- orm-benchmark (with no flags)
```
Reports:

Read
pgx:		9114	145028 ns/op	893 B/op	8 allocs/op
pgx_pool:	8848	149457 ns/op	1078 B/op	9 allocs/op
raw:		8554	153093 ns/op	2076 B/op	50 allocs/op
sqlc:		8664	153785 ns/op	2093 B/op	51 allocs/op
beego:		8460	155170 ns/op	2112 B/op	76 allocs/op
reform:		8043	161340 ns/op	3214 B/op	86 allocs/op
gorp:		7502	166495 ns/op	3892 B/op	194 allocs/op
pop:		7990	169894 ns/op	3192 B/op	67 allocs/op
sq:		6944	173734 ns/op	11143 B/op	135 allocs/op
ent:		6648	174768 ns/op	5621 B/op	144 allocs/op
pg:		6639	176713 ns/op	872 B/op	20 allocs/op
gorm_prep:	7336	177299 ns/op	4420 B/op	87 allocs/op
dbr:		7473	177754 ns/op	2168 B/op	36 allocs/op
bun:		7220	178798 ns/op	5830 B/op	39 allocs/op
sqlboiler:	7272	179096 ns/op	947 B/op	14 allocs/op
rel:		7070	180911 ns/op	2320 B/op	47 allocs/op
zorm:		6560	188995 ns/op	3032 B/op	64 allocs/op
jet:		6108	197136 ns/op	12939 B/op	273 allocs/op
gorm:		5414	220944 ns/op	4790 B/op	98 allocs/op
gen:		5218	239745 ns/op	10333 B/op	153 allocs/op
sqlx:		4071	325161 ns/op	1992 B/op	43 allocs/op
godb:		3846	337488 ns/op	4081 B/op	102 allocs/op
xorm:		3936	341079 ns/op	4665 B/op	127 allocs/op
upper:		3844	347903 ns/op	5070 B/op	110 allocs/op

ReadSlice
reform:		7532	164142 ns/op	4029 B/op	100 allocs/op
pgx:		4627	247717 ns/op	42949 B/op	504 allocs/op
pgx_pool:	4688	251315 ns/op	43005 B/op	504 allocs/op
raw:		3846	304145 ns/op	38356 B/op	1038 allocs/op
sqlc:		3612	331861 ns/op	62677 B/op	1150 allocs/op
pg:		3439	341483 ns/op	22262 B/op	629 allocs/op
upper:		3376	356030 ns/op	4805 B/op	90 allocs/op
sqlx:		3243	374309 ns/op	37496 B/op	1225 allocs/op
gorp:		3096	377523 ns/op	57402 B/op	1494 allocs/op
ent:		2956	382742 ns/op	77206 B/op	2035 allocs/op
pop:		3046	384796 ns/op	69040 B/op	1306 allocs/op
sq:		2844	416537 ns/op	144937 B/op	1829 allocs/op
dbr:		2806	421140 ns/op	30800 B/op	1253 allocs/op
bun:		2682	427746 ns/op	34062 B/op	1124 allocs/op
beego:		2660	442137 ns/op	55217 B/op	3078 allocs/op
sqlboiler:	2692	446555 ns/op	66781 B/op	2260 allocs/op
gorm_prep:	2464	468689 ns/op	43190 B/op	2081 allocs/op
gorm:		2210	532142 ns/op	44348 B/op	2191 allocs/op
gen:		1996	557703 ns/op	49933 B/op	2246 allocs/op
zorm:		2062	579764 ns/op	161632 B/op	2949 allocs/op
jet:		1920	623028 ns/op	199857 B/op	3642 allocs/op
rel:		1860	634956 ns/op	100656 B/op	2253 allocs/op
godb:		1824	647892 ns/op	75247 B/op	3084 allocs/op
xorm:		1840	654576 ns/op	119408 B/op	4401 allocs/op

Insert
pgx_pool:	3970	332017 ns/op	286 B/op	10 allocs/op
raw:		3614	332762 ns/op	703 B/op	13 allocs/op
sqlc:		3560	345224 ns/op	2786 B/op	61 allocs/op
jet:		3157	345751 ns/op	3575 B/op	105 allocs/op
reform:		3405	346586 ns/op	1773 B/op	51 allocs/op
sq:		3512	353173 ns/op	9759 B/op	100 allocs/op
dbr:		3166	354059 ns/op	2688 B/op	65 allocs/op
pgx:		4138	354485 ns/op	271 B/op	10 allocs/op
ent:		3502	355447 ns/op	4141 B/op	97 allocs/op
sqlboiler:	3483	357483 ns/op	1574 B/op	34 allocs/op
gorm_prep:	3645	363544 ns/op	5176 B/op	65 allocs/op
gorp:		3366	363730 ns/op	1782 B/op	41 allocs/op
beego:		3398	370699 ns/op	2400 B/op	57 allocs/op
pg:		3367	376303 ns/op	1128 B/op	10 allocs/op
bun:		3067	393554 ns/op	5012 B/op	13 allocs/op
gorm:		2925	415275 ns/op	7176 B/op	105 allocs/op
gen:		2896	424180 ns/op	10057 B/op	134 allocs/op
sqlx:		2326	512990 ns/op	856 B/op	19 allocs/op
upper:		2326	516648 ns/op	5895 B/op	125 allocs/op
xorm:		2450	524665 ns/op	3311 B/op	88 allocs/op
rel:		2503	529844 ns/op	2623 B/op	45 allocs/op
godb:		2300	544408 ns/op	4521 B/op	115 allocs/op
zorm:		2377	549614 ns/op	3798 B/op	77 allocs/op
pop:		1549	758753 ns/op	9576 B/op	237 allocs/op

InsertMulti
pgx:		1273	921021 ns/op	112900 B/op	42 allocs/op
raw:		1236	938650 ns/op	183817 B/op	930 allocs/op
pgx_pool:	1329	951908 ns/op	112947 B/op	42 allocs/op
sq:		1195	1037596 ns/op	227249 B/op	1706 allocs/op
beego:		1180	1046170 ns/op	177709 B/op	2745 allocs/op
gorm_prep:	1074	1094747 ns/op	251036 B/op	1890 allocs/op
reform:		1066	1159516 ns/op	458844 B/op	2746 allocs/op
ent:		930	1316729 ns/op	386568 B/op	4598 allocs/op
pg:		912	1340212 ns/op	3315 B/op	112 allocs/op
jet:		889	1374731 ns/op	327550 B/op	6493 allocs/op
bun:		822	1496440 ns/op	42496 B/op	219 allocs/op
gorm:		814	1506417 ns/op	291408 B/op	5231 allocs/op
sqlx:		829	1510806 ns/op	170680 B/op	1551 allocs/op
gen:		808	1562062 ns/op	304641 B/op	5358 allocs/op
zorm:		747	1616409 ns/op	199927 B/op	2780 allocs/op
xorm:		651	1707319 ns/op	248245 B/op	5414 allocs/op
rel:		718	1742936 ns/op	306914 B/op	3265 allocs/op
godb:		694	1789560 ns/op	254057 B/op	5894 allocs/op
upper:		660	1832200 ns/op	322889 B/op	4204 allocs/op
pop:		bulk-insert is not supported
gorp:		bulk-insert is not supported
sqlboiler:	bulk-insert is not supported
dbr:		bulk-insert is not supported
sqlc:		bulk-insert is not supported

Update
sqlc:		8803	147845 ns/op	877 B/op	14 allocs/op
raw:		8768	149098 ns/op	749 B/op	13 allocs/op
sqlx:		3828	316141 ns/op	872 B/op	20 allocs/op
pgx:		3477	329730 ns/op	268 B/op	10 allocs/op
pgx_pool:	3702	340443 ns/op	284 B/op	10 allocs/op
gorp:		3525	342475 ns/op	1206 B/op	32 allocs/op
ent:		3686	349895 ns/op	4676 B/op	97 allocs/op
beego:		3530	349954 ns/op	1752 B/op	47 allocs/op
sqlboiler:	3604	350815 ns/op	901 B/op	17 allocs/op
reform:		3424	356697 ns/op	1775 B/op	51 allocs/op
sq:		3240	357928 ns/op	7413 B/op	90 allocs/op
gorm_prep:	3625	365797 ns/op	5008 B/op	56 allocs/op
jet:		3736	366678 ns/op	4558 B/op	126 allocs/op
dbr:		3802	371884 ns/op	2651 B/op	57 allocs/op
bun:		3037	381250 ns/op	4730 B/op	5 allocs/op
pg:		3388	399181 ns/op	768 B/op	9 allocs/op
pop:		3278	401652 ns/op	6049 B/op	186 allocs/op
gorm:		3172	412440 ns/op	6752 B/op	99 allocs/op
gen:		2617	430305 ns/op	13392 B/op	165 allocs/op
rel:		2520	522942 ns/op	3048 B/op	45 allocs/op
xorm:		2299	533499 ns/op	3944 B/op	132 allocs/op
zorm:		2325	540689 ns/op	3024 B/op	59 allocs/op
godb:		2182	559443 ns/op	5129 B/op	154 allocs/op
upper:		975	1267440 ns/op	16703 B/op	390 allocs/op
```
