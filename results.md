# Results

- orm-benchmark (with no flags)
```
Reports:

Update
sqlc:		8894	125174 ns/op	877 B/op	14 allocs/op
raw:		8701	127989 ns/op	750 B/op	13 allocs/op
pgx:		4094	289343 ns/op	270 B/op	10 allocs/op
pgx_pool:	4148	302471 ns/op	285 B/op	10 allocs/op
sqlx:		4053	303460 ns/op	872 B/op	20 allocs/op
gorp:		3681	305951 ns/op	1204 B/op	32 allocs/op
sqlboiler:	4064	318474 ns/op	902 B/op	17 allocs/op
beego:		3627	324575 ns/op	1752 B/op	47 allocs/op
gorm_prep:	3720	331909 ns/op	5008 B/op	56 allocs/op
ent:		3856	332479 ns/op	4677 B/op	97 allocs/op
reform:		3852	333816 ns/op	1773 B/op	51 allocs/op
pop:		3369	347575 ns/op	6047 B/op	186 allocs/op
pg:		3432	352951 ns/op	768 B/op	9 allocs/op
dbr:		3535	354966 ns/op	2651 B/op	57 allocs/op
bun:		3676	374690 ns/op	4729 B/op	5 allocs/op
gorm:		2924	389754 ns/op	6752 B/op	99 allocs/op
gen:		2696	447763 ns/op	13472 B/op	166 allocs/op
xorm:		2524	488879 ns/op	3944 B/op	132 allocs/op
rel:		2358	500243 ns/op	3048 B/op	45 allocs/op
zorm:		2416	507488 ns/op	3024 B/op	59 allocs/op
godb:		2373	514429 ns/op	5112 B/op	154 allocs/op
upper:		1027	1181716 ns/op	16664 B/op	390 allocs/op

Read
pgx:		8895	124135 ns/op	893 B/op	8 allocs/op
pgx_pool:	9255	125336 ns/op	1078 B/op	9 allocs/op
sqlc:		8650	142122 ns/op	2077 B/op	51 allocs/op
raw:		8420	142959 ns/op	2061 B/op	50 allocs/op
beego:		8223	148386 ns/op	2096 B/op	76 allocs/op
reform:		7891	152916 ns/op	3198 B/op	86 allocs/op
pop:		7957	161074 ns/op	3152 B/op	67 allocs/op
gorp:		8982	163616 ns/op	3878 B/op	194 allocs/op
ent:		7458	167530 ns/op	5605 B/op	144 allocs/op
gorm_prep:	6968	168634 ns/op	4403 B/op	87 allocs/op
sqlboiler:	8386	171179 ns/op	944 B/op	14 allocs/op
pg:		6914	174638 ns/op	872 B/op	20 allocs/op
bun:		6754	181052 ns/op	5810 B/op	39 allocs/op
dbr:		6892	181173 ns/op	2184 B/op	37 allocs/op
rel:		6735	184230 ns/op	2304 B/op	47 allocs/op
zorm:		6294	194217 ns/op	3016 B/op	64 allocs/op
gorm:		6402	225945 ns/op	4771 B/op	98 allocs/op
gen:		5041	250016 ns/op	10390 B/op	154 allocs/op
sqlx:		4057	303659 ns/op	1976 B/op	43 allocs/op
godb:		3759	315525 ns/op	4064 B/op	102 allocs/op
upper:		3736	328131 ns/op	5055 B/op	110 allocs/op
xorm:		3777	341177 ns/op	4648 B/op	127 allocs/op

ReadSlice
reform:		8152	153205 ns/op	4012 B/op	100 allocs/op
pgx_pool:	4412	282772 ns/op	43006 B/op	504 allocs/op
pgx:		3877	287635 ns/op	42949 B/op	504 allocs/op
upper:		3387	336150 ns/op	4790 B/op	90 allocs/op
raw:		3109	356571 ns/op	38341 B/op	1038 allocs/op
sqlc:		3106	381645 ns/op	62661 B/op	1150 allocs/op
pg:		2985	388400 ns/op	22635 B/op	629 allocs/op
gorp:		2560	441702 ns/op	57373 B/op	1494 allocs/op
ent:		2718	450938 ns/op	77190 B/op	2035 allocs/op
sqlx:		2622	451258 ns/op	37480 B/op	1225 allocs/op
pop:		2695	481166 ns/op	68380 B/op	1306 allocs/op
dbr:		2538	482435 ns/op	30816 B/op	1254 allocs/op
bun:		2390	492520 ns/op	34036 B/op	1124 allocs/op
beego:		2276	528558 ns/op	55198 B/op	3078 allocs/op
sqlboiler:	2205	555616 ns/op	66384 B/op	2259 allocs/op
gorm_prep:	1993	562963 ns/op	43149 B/op	2081 allocs/op
gorm:		1921	630590 ns/op	44301 B/op	2191 allocs/op
gen:		1664	676046 ns/op	50072 B/op	2248 allocs/op
zorm:		1498	746941 ns/op	161616 B/op	2949 allocs/op
rel:		1446	754700 ns/op	100640 B/op	2253 allocs/op
godb:		1570	780970 ns/op	75218 B/op	3084 allocs/op
xorm:		1281	842005 ns/op	119372 B/op	4401 allocs/op

Insert
pgx:		4239	291151 ns/op	272 B/op	10 allocs/op
raw:		3697	296976 ns/op	703 B/op	13 allocs/op
pgx_pool:	3486	297954 ns/op	288 B/op	10 allocs/op
gorp:		3492	311448 ns/op	1766 B/op	41 allocs/op
beego:		3663	311723 ns/op	2384 B/op	57 allocs/op
sqlc:		3786	316017 ns/op	2771 B/op	62 allocs/op
ent:		3958	316899 ns/op	4125 B/op	97 allocs/op
sqlboiler:	4110	318855 ns/op	1559 B/op	34 allocs/op
reform:		4126	322149 ns/op	1774 B/op	51 allocs/op
dbr:		3823	338669 ns/op	2688 B/op	65 allocs/op
gorm_prep:	3548	340538 ns/op	5159 B/op	65 allocs/op
pg:		3354	340722 ns/op	1128 B/op	10 allocs/op
bun:		3400	378758 ns/op	4994 B/op	13 allocs/op
gorm:		3124	381045 ns/op	7159 B/op	105 allocs/op
gen:		3080	412231 ns/op	10120 B/op	135 allocs/op
sqlx:		2646	439877 ns/op	856 B/op	19 allocs/op
xorm:		2539	470676 ns/op	3319 B/op	89 allocs/op
rel:		2384	497044 ns/op	2606 B/op	45 allocs/op
zorm:		2451	505682 ns/op	3783 B/op	77 allocs/op
godb:		2467	510017 ns/op	4504 B/op	115 allocs/op
upper:		2329	538232 ns/op	5879 B/op	125 allocs/op
pop:		1618	692142 ns/op	9555 B/op	238 allocs/op

InsertMulti
pgx:		1170	1099711 ns/op	112933 B/op	43 allocs/op
pgx_pool:	1134	1140986 ns/op	112937 B/op	43 allocs/op
raw:		1052	1267492 ns/op	183852 B/op	930 allocs/op
beego:		987	1341908 ns/op	177653 B/op	2745 allocs/op
gorm_prep:	804	1495817 ns/op	250990 B/op	1890 allocs/op
reform:		825	1496949 ns/op	458747 B/op	2745 allocs/op
pg:		768	1675093 ns/op	4776 B/op	112 allocs/op
ent:		723	1744214 ns/op	386569 B/op	4598 allocs/op
bun:		703	1806408 ns/op	42501 B/op	219 allocs/op
sqlx:		596	2015964 ns/op	169981 B/op	1551 allocs/op
gen:		604	2057610 ns/op	304667 B/op	5359 allocs/op
gorm:		607	2066170 ns/op	291347 B/op	5231 allocs/op
zorm:		526	2247303 ns/op	199928 B/op	2780 allocs/op
xorm:		530	2351565 ns/op	248053 B/op	5414 allocs/op
rel:		534	2433482 ns/op	306897 B/op	3265 allocs/op
upper:		480	2563274 ns/op	322888 B/op	4204 allocs/op
godb:		492	2571139 ns/op	254006 B/op	5894 allocs/op
pop:		bulk-insert is not supported
sqlboiler:	bulk-insert is not supported
sqlc:		bulk-insert is not supported
dbr:		bulk-insert is not supported
gorp:		bulk-insert is not supported
```
