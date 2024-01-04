# Results

- orm-benchmark (with no flags)
```
Reports:

Insert
raw:		4398	275215 ns/op	703 B/op	13 allocs/op
ent:		3212	323791 ns/op	4126 B/op	97 allocs/op
sqlboiler:	3459	346336 ns/op	1559 B/op	34 allocs/op
pgx_pool:	4028	351897 ns/op	288 B/op	10 allocs/op
jet:		3624	366378 ns/op	3575 B/op	105 allocs/op
beego:		2892	367038 ns/op	2384 B/op	57 allocs/op
gorm_prep:	3391	375917 ns/op	5160 B/op	65 allocs/op
reform:		3411	389070 ns/op	1773 B/op	51 allocs/op
sqlc:		3255	400183 ns/op	2770 B/op	61 allocs/op
pg:		2857	407129 ns/op	794 B/op	10 allocs/op
dbr:		2677	419241 ns/op	2688 B/op	65 allocs/op
pgx:		3603	428586 ns/op	271 B/op	10 allocs/op
sqlx:		2984	429482 ns/op	856 B/op	19 allocs/op
gen:		3081	431377 ns/op	10121 B/op	135 allocs/op
bun:		2796	460865 ns/op	4996 B/op	13 allocs/op
gorp:		2504	469179 ns/op	1767 B/op	41 allocs/op
gorm:		2421	485324 ns/op	7160 B/op	105 allocs/op
upper:		2334	596531 ns/op	5879 B/op	125 allocs/op
rel:		2115	610205 ns/op	2606 B/op	45 allocs/op
xorm:		2176	624164 ns/op	3320 B/op	89 allocs/op
zorm:		2371	645388 ns/op	3782 B/op	77 allocs/op
godb:		1989	654434 ns/op	4505 B/op	115 allocs/op
pop:		1521	810242 ns/op	9555 B/op	238 allocs/op

InsertMulti
raw:		1400	919564 ns/op	183799 B/op	930 allocs/op
pgx_pool:	1150	1015330 ns/op	112931 B/op	42 allocs/op
beego:		1188	1068948 ns/op	177662 B/op	2745 allocs/op
pgx:		1153	1074291 ns/op	112904 B/op	42 allocs/op
reform:		1032	1156813 ns/op	458849 B/op	2746 allocs/op
gorm_prep:	1011	1195540 ns/op	251011 B/op	1890 allocs/op
ent:		960	1381233 ns/op	386555 B/op	4598 allocs/op
pg:		844	1430083 ns/op	5889 B/op	112 allocs/op
jet:		849	1519508 ns/op	327542 B/op	6493 allocs/op
bun:		810	1528402 ns/op	42461 B/op	219 allocs/op
gorm:		782	1552918 ns/op	291371 B/op	5231 allocs/op
sqlx:		868	1558473 ns/op	170153 B/op	1551 allocs/op
gen:		788	1569529 ns/op	304694 B/op	5359 allocs/op
zorm:		698	1744919 ns/op	199928 B/op	2780 allocs/op
rel:		591	1789279 ns/op	306899 B/op	3265 allocs/op
upper:		668	1856845 ns/op	322888 B/op	4204 allocs/op
xorm:		700	1861732 ns/op	248156 B/op	5414 allocs/op
godb:		662	1906071 ns/op	254030 B/op	5894 allocs/op
dbr:		bulk-insert is not supported
pop:		bulk-insert is not supported
sqlc:		bulk-insert is not supported
gorp:		bulk-insert is not supported
sqlboiler:	bulk-insert is not supported

Update
sqlc:		8998	147141 ns/op	878 B/op	14 allocs/op
raw:		8665	149048 ns/op	750 B/op	13 allocs/op
sqlx:		4146	315593 ns/op	872 B/op	20 allocs/op
sqlboiler:	4503	349513 ns/op	903 B/op	17 allocs/op
jet:		3295	356447 ns/op	4557 B/op	126 allocs/op
reform:		3772	368711 ns/op	1775 B/op	51 allocs/op
gorp:		2893	375445 ns/op	1207 B/op	32 allocs/op
beego:		2728	376515 ns/op	1752 B/op	47 allocs/op
pgx_pool:	3468	376919 ns/op	286 B/op	10 allocs/op
pgx:		2853	399255 ns/op	271 B/op	10 allocs/op
pop:		3217	407075 ns/op	6049 B/op	186 allocs/op
pg:		2730	414197 ns/op	768 B/op	9 allocs/op
gorm_prep:	2755	423510 ns/op	5008 B/op	56 allocs/op
dbr:		3630	436941 ns/op	2651 B/op	57 allocs/op
bun:		3088	469881 ns/op	4729 B/op	5 allocs/op
ent:		2209	473125 ns/op	4677 B/op	97 allocs/op
gen:		2583	503572 ns/op	13472 B/op	166 allocs/op
rel:		2097	543030 ns/op	3048 B/op	45 allocs/op
zorm:		2443	568159 ns/op	3024 B/op	59 allocs/op
gorm:		1984	592409 ns/op	6752 B/op	99 allocs/op
xorm:		2169	629606 ns/op	3945 B/op	132 allocs/op
godb:		2091	687962 ns/op	5113 B/op	154 allocs/op
upper:		956	1223537 ns/op	16658 B/op	390 allocs/op

Read
pgx:		9414	142781 ns/op	893 B/op	8 allocs/op
pgx_pool:	8665	146917 ns/op	1078 B/op	9 allocs/op
sqlc:		8772	153631 ns/op	2077 B/op	51 allocs/op
raw:		8550	154318 ns/op	2061 B/op	50 allocs/op
beego:		8347	156060 ns/op	2096 B/op	76 allocs/op
reform:		7948	162923 ns/op	3198 B/op	86 allocs/op
pop:		7494	166692 ns/op	3167 B/op	67 allocs/op
gorp:		7684	167069 ns/op	3877 B/op	194 allocs/op
ent:		7442	172432 ns/op	5605 B/op	144 allocs/op
dbr:		7372	176732 ns/op	2184 B/op	37 allocs/op
bun:		7486	176850 ns/op	5813 B/op	39 allocs/op
gorm_prep:	6738	177739 ns/op	4405 B/op	87 allocs/op
pg:		7148	177916 ns/op	872 B/op	20 allocs/op
sqlboiler:	7479	180938 ns/op	940 B/op	14 allocs/op
rel:		7384	181446 ns/op	2304 B/op	47 allocs/op
zorm:		6576	188636 ns/op	3016 B/op	64 allocs/op
jet:		6666	196269 ns/op	12923 B/op	273 allocs/op
gorm:		5691	217876 ns/op	4773 B/op	98 allocs/op
gen:		5048	245333 ns/op	10396 B/op	154 allocs/op
godb:		4124	330458 ns/op	4064 B/op	102 allocs/op
sqlx:		3832	331170 ns/op	1976 B/op	43 allocs/op
upper:		3834	345827 ns/op	5055 B/op	110 allocs/op
xorm:		3900	346415 ns/op	4649 B/op	127 allocs/op

ReadSlice
reform:		7484	161505 ns/op	4013 B/op	100 allocs/op
pgx:		4562	244213 ns/op	42947 B/op	504 allocs/op
pgx_pool:	4669	253915 ns/op	43005 B/op	504 allocs/op
raw:		3950	300570 ns/op	38340 B/op	1038 allocs/op
sqlc:		3457	333091 ns/op	62661 B/op	1150 allocs/op
pg:		3572	342831 ns/op	23772 B/op	629 allocs/op
upper:		3588	351432 ns/op	4791 B/op	90 allocs/op
sqlx:		3139	372574 ns/op	37480 B/op	1225 allocs/op
gorp:		3150	377954 ns/op	57383 B/op	1494 allocs/op
ent:		3025	378687 ns/op	77189 B/op	2035 allocs/op
pop:		3146	379131 ns/op	68700 B/op	1306 allocs/op
dbr:		2834	421422 ns/op	30816 B/op	1254 allocs/op
bun:		2764	426921 ns/op	34042 B/op	1124 allocs/op
beego:		2641	430898 ns/op	55199 B/op	3078 allocs/op
sqlboiler:	2696	438854 ns/op	66324 B/op	2259 allocs/op
gorm_prep:	2337	472229 ns/op	43168 B/op	2081 allocs/op
gorm:		2265	526073 ns/op	44320 B/op	2191 allocs/op
gen:		1888	569909 ns/op	50021 B/op	2247 allocs/op
zorm:		2090	575987 ns/op	161617 B/op	2949 allocs/op
jet:		1975	606258 ns/op	191846 B/op	3642 allocs/op
godb:		1886	647417 ns/op	75225 B/op	3084 allocs/op
rel:		1862	652038 ns/op	100640 B/op	2253 allocs/op
xorm:		1752	654864 ns/op	119387 B/op	4401 allocs/op
```
