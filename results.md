# Results

- orm-benchmark (with no flags)
```
Reports:

Read
sqlc:		8448	148618 ns/op	904 B/op	19 allocs/op
pgx:		8600	148944 ns/op	776 B/op	18 allocs/op
pgx_pool:	7509	151412 ns/op	963 B/op	19 allocs/op
beego:		8461	153944 ns/op	2112 B/op	75 allocs/op
raw:		8208	155169 ns/op	2093 B/op	50 allocs/op
reform:		8284	164604 ns/op	3230 B/op	86 allocs/op
gorp:		7862	166784 ns/op	3909 B/op	194 allocs/op
pop:		8088	169084 ns/op	3206 B/op	66 allocs/op
pg:		7572	172384 ns/op	872 B/op	20 allocs/op
sq:		7442	175209 ns/op	11160 B/op	135 allocs/op
bun:		7082	175719 ns/op	5845 B/op	39 allocs/op
dbr:		7629	176276 ns/op	2184 B/op	36 allocs/op
sqlboiler:	7356	176344 ns/op	957 B/op	14 allocs/op
ent:		7432	177505 ns/op	5685 B/op	145 allocs/op
rel:		7408	179028 ns/op	2336 B/op	47 allocs/op
gorm_prep:	6666	179496 ns/op	4565 B/op	89 allocs/op
zorm:		6795	193242 ns/op	3337 B/op	74 allocs/op
jet:		6574	194693 ns/op	13067 B/op	273 allocs/op
gorm:		5752	223097 ns/op	4983 B/op	102 allocs/op
gen:		5293	241580 ns/op	10508 B/op	154 allocs/op
sqlx:		4130	319489 ns/op	2008 B/op	43 allocs/op
godb:		3882	336949 ns/op	4097 B/op	102 allocs/op
xorm:		3566	350352 ns/op	5161 B/op	131 allocs/op
upper:		3532	350440 ns/op	5087 B/op	110 allocs/op

ReadSlice
reform:		7446	159781 ns/op	4044 B/op	100 allocs/op
pgx:		4573	247183 ns/op	30320 B/op	513 allocs/op
pgx_pool:	4633	247645 ns/op	30380 B/op	513 allocs/op
sqlc:		4596	259784 ns/op	54624 B/op	620 allocs/op
raw:		3896	296835 ns/op	38374 B/op	1038 allocs/op
pg:		3427	339762 ns/op	23002 B/op	629 allocs/op
upper:		3492	349657 ns/op	4823 B/op	90 allocs/op
ent:		3240	364374 ns/op	77398 B/op	2036 allocs/op
sqlx:		3196	368843 ns/op	37512 B/op	1225 allocs/op
gorp:		3386	370691 ns/op	57545 B/op	1494 allocs/op
pop:		3070	381803 ns/op	77133 B/op	1306 allocs/op
sq:		2968	399622 ns/op	152519 B/op	1829 allocs/op
bun:		2877	414916 ns/op	34207 B/op	1124 allocs/op
dbr:		2799	415880 ns/op	30944 B/op	1253 allocs/op
beego:		2721	431007 ns/op	55351 B/op	3077 allocs/op
sqlboiler:	2749	431614 ns/op	67006 B/op	2260 allocs/op
gorm_prep:	2487	461743 ns/op	43532 B/op	2082 allocs/op
gorm:		2293	515118 ns/op	44752 B/op	2196 allocs/op
gen:		2046	540283 ns/op	50267 B/op	2247 allocs/op
xorm:		2100	566538 ns/op	121237 B/op	4407 allocs/op
zorm:		2068	572100 ns/op	169548 B/op	2959 allocs/op
jet:		2029	585397 ns/op	192714 B/op	3642 allocs/op
godb:		1839	637968 ns/op	75392 B/op	3084 allocs/op
rel:		1785	666717 ns/op	149041 B/op	2553 allocs/op

Insert
raw:		4507	292461 ns/op	704 B/op	13 allocs/op
pgx:		4158	294418 ns/op	280 B/op	8 allocs/op
sqlc:		4150	298137 ns/op	280 B/op	8 allocs/op
pgx_pool:	4075	298807 ns/op	298 B/op	8 allocs/op
beego:		4190	300981 ns/op	2400 B/op	56 allocs/op
sqlboiler:	3470	312346 ns/op	1591 B/op	34 allocs/op
gorm_prep:	4120	319339 ns/op	5192 B/op	65 allocs/op
jet:		4076	319977 ns/op	3584 B/op	105 allocs/op
reform:		3774	320355 ns/op	1775 B/op	51 allocs/op
sq:		4378	330198 ns/op	9777 B/op	100 allocs/op
gorp:		3480	331171 ns/op	1798 B/op	41 allocs/op
ent:		3637	332290 ns/op	4151 B/op	97 allocs/op
bun:		3572	336491 ns/op	5029 B/op	13 allocs/op
pg:		3733	343780 ns/op	1095 B/op	10 allocs/op
dbr:		3303	355087 ns/op	2688 B/op	65 allocs/op
gen:		2950	376749 ns/op	10041 B/op	131 allocs/op
gorm:		3283	409477 ns/op	7209 B/op	105 allocs/op
rel:		2821	445630 ns/op	2639 B/op	45 allocs/op
sqlx:		2977	445687 ns/op	856 B/op	19 allocs/op
upper:		2426	484638 ns/op	5913 B/op	125 allocs/op
zorm:		2728	489784 ns/op	5145 B/op	104 allocs/op
godb:		2509	498311 ns/op	4537 B/op	115 allocs/op
xorm:		2488	506730 ns/op	3120 B/op	87 allocs/op
pop:		1840	708639 ns/op	9421 B/op	234 allocs/op

InsertMulti
pgx:		1707	728816 ns/op	47532 B/op	38 allocs/op
pgx_pool:	1640	733620 ns/op	47998 B/op	38 allocs/op
sqlc:		1687	761902 ns/op	66433 B/op	639 allocs/op
raw:		1402	851151 ns/op	187127 B/op	930 allocs/op
sq:		1293	949722 ns/op	237251 B/op	1706 allocs/op
beego:		1314	969644 ns/op	177652 B/op	2744 allocs/op
gorm_prep:	1184	1047137 ns/op	254382 B/op	1890 allocs/op
reform:		1184	1123778 ns/op	462209 B/op	2746 allocs/op
ent:		1022	1161757 ns/op	396528 B/op	4597 allocs/op
jet:		927	1358204 ns/op	338333 B/op	6493 allocs/op
pg:		900	1382670 ns/op	4563 B/op	112 allocs/op
sqlx:		867	1396858 ns/op	170457 B/op	1550 allocs/op
bun:		865	1434560 ns/op	42580 B/op	219 allocs/op
gorm:		844	1444648 ns/op	276182 B/op	5230 allocs/op
gen:		826	1463510 ns/op	289370 B/op	5354 allocs/op
zorm:		750	1598643 ns/op	212131 B/op	2808 allocs/op
rel:		770	1625136 ns/op	312562 B/op	3265 allocs/op
upper:		756	1643732 ns/op	328153 B/op	4204 allocs/op
godb:		715	1753339 ns/op	260734 B/op	5894 allocs/op
xorm:		717	1757731 ns/op	258931 B/op	5518 allocs/op
gorp:		bulk-insert is not supported
sqlboiler:	bulk-insert is not supported
dbr:		bulk-insert is not supported
pop:		bulk-insert is not supported

Update
raw:		8934	147057 ns/op	749 B/op	13 allocs/op
sqlc:		8311	148037 ns/op	288 B/op	8 allocs/op
pgx:		4188	301977 ns/op	288 B/op	8 allocs/op
pgx_pool:	3891	304645 ns/op	305 B/op	8 allocs/op
beego:		4108	306440 ns/op	1736 B/op	46 allocs/op
sqlx:		3846	314410 ns/op	872 B/op	20 allocs/op
sq:		3429	314900 ns/op	7417 B/op	90 allocs/op
reform:		3711	316066 ns/op	1773 B/op	51 allocs/op
sqlboiler:	4436	321791 ns/op	901 B/op	17 allocs/op
gorm_prep:	4154	324145 ns/op	5072 B/op	56 allocs/op
jet:		4131	326424 ns/op	4565 B/op	126 allocs/op
gorp:		3295	327641 ns/op	1205 B/op	32 allocs/op
pop:		3597	335513 ns/op	5857 B/op	184 allocs/op
bun:		3705	336445 ns/op	4762 B/op	5 allocs/op
ent:		3654	339315 ns/op	4725 B/op	98 allocs/op
pg:		3789	353320 ns/op	768 B/op	9 allocs/op
dbr:		3787	356261 ns/op	2651 B/op	57 allocs/op
gorm:		2986	389390 ns/op	6832 B/op	99 allocs/op
gen:		3373	403848 ns/op	13408 B/op	161 allocs/op
rel:		2775	471517 ns/op	3048 B/op	45 allocs/op
xorm:		2366	517855 ns/op	4305 B/op	145 allocs/op
zorm:		2434	524936 ns/op	4449 B/op	84 allocs/op
godb:		2534	532693 ns/op	5145 B/op	154 allocs/op
upper:		1030	1156016 ns/op	16743 B/op	390 allocs/op
```
