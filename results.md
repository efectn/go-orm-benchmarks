# Results

- orm-benchmark (with no flags)
```
Reports:

ReadSlice
reform:		4342	274064 ns/op	4013 B/op	100 allocs/op
pgx_pool:	3062	433496 ns/op	43006 B/op	504 allocs/op
pgx:		2973	446280 ns/op	42947 B/op	504 allocs/op
raw:		2539	536161 ns/op	38339 B/op	1038 allocs/op
sqlc:		2106	573144 ns/op	62660 B/op	1150 allocs/op
pg:		2220	575978 ns/op	22259 B/op	629 allocs/op
gorp:		1893	643301 ns/op	57376 B/op	1494 allocs/op
sqlx:		1854	659654 ns/op	37480 B/op	1225 allocs/op
upper:		2200	678954 ns/op	4791 B/op	90 allocs/op
dbr:		1554	700492 ns/op	30816 B/op	1254 allocs/op
pop:		1821	706954 ns/op	68633 B/op	1306 allocs/op
beego:		1592	730664 ns/op	55199 B/op	3078 allocs/op
ent:		1543	740525 ns/op	77189 B/op	2035 allocs/op
bun:		1350	748726 ns/op	34037 B/op	1124 allocs/op
sqlboiler:	1425	771965 ns/op	66404 B/op	2259 allocs/op
gorm_prep:	1620	800750 ns/op	43153 B/op	2081 allocs/op
gorm:		1405	903997 ns/op	44308 B/op	2191 allocs/op
rel:		1110	1130177 ns/op	100640 B/op	2253 allocs/op
xorm:		932	1167266 ns/op	119374 B/op	4401 allocs/op
godb:		891	1209499 ns/op	75223 B/op	3084 allocs/op
zorm:		963	1211490 ns/op	161616 B/op	2949 allocs/op

Insert
pgx_pool:	2121	562769 ns/op	286 B/op	10 allocs/op
pgx:		2085	586815 ns/op	271 B/op	10 allocs/op
raw:		2215	586871 ns/op	705 B/op	13 allocs/op
sqlboiler:	2001	590325 ns/op	1557 B/op	35 allocs/op
beego:		2017	596681 ns/op	2384 B/op	57 allocs/op
reform:		1928	607740 ns/op	1774 B/op	51 allocs/op
sqlc:		1735	626644 ns/op	2771 B/op	61 allocs/op
gorp:		1851	631678 ns/op	1766 B/op	41 allocs/op
ent:		1989	635708 ns/op	4124 B/op	97 allocs/op
dbr:		2029	654464 ns/op	2688 B/op	65 allocs/op
pg:		1683	657252 ns/op	1461 B/op	10 allocs/op
gorm_prep:	2113	658672 ns/op	5159 B/op	65 allocs/op
bun:		1953	667183 ns/op	4995 B/op	13 allocs/op
gorm:		1779	756994 ns/op	7159 B/op	105 allocs/op
sqlx:		1455	813227 ns/op	856 B/op	19 allocs/op
rel:		1474	845961 ns/op	2606 B/op	45 allocs/op
godb:		1444	892213 ns/op	4504 B/op	115 allocs/op
xorm:		1447	917040 ns/op	3318 B/op	89 allocs/op
upper:		1328	933222 ns/op	5883 B/op	125 allocs/op
zorm:		1233	1054082 ns/op	3782 B/op	77 allocs/op
pop:		964	1231511 ns/op	9551 B/op	238 allocs/op

InsertMulti
pgx_pool:	712	1696255 ns/op	112953 B/op	43 allocs/op
pgx:		738	1775035 ns/op	112925 B/op	43 allocs/op
raw:		698	1831287 ns/op	183948 B/op	931 allocs/op
beego:		639	1911649 ns/op	177667 B/op	2745 allocs/op
gorm_prep:	615	1989480 ns/op	250993 B/op	1890 allocs/op
reform:		596	2108369 ns/op	458766 B/op	2745 allocs/op
pg:		535	2302945 ns/op	5413 B/op	112 allocs/op
bun:		534	2364169 ns/op	42470 B/op	219 allocs/op
ent:		541	2386419 ns/op	386586 B/op	4599 allocs/op
gorm:		468	2646983 ns/op	291355 B/op	5231 allocs/op
sqlx:		392	2767175 ns/op	169939 B/op	1551 allocs/op
xorm:		403	3139338 ns/op	248095 B/op	5414 allocs/op
godb:		387	3251076 ns/op	254018 B/op	5894 allocs/op
upper:		357	3320106 ns/op	322877 B/op	4204 allocs/op
zorm:		348	3568795 ns/op	199927 B/op	2780 allocs/op
rel:		414	3710117 ns/op	306898 B/op	3265 allocs/op
sqlboiler:	bulk-insert is not supported
sqlc:		bulk-insert is not supported
dbr:		bulk-insert is not supported
pop:		bulk-insert is not supported
gorp:		bulk-insert is not supported

Update
sqlc:		5400	209645 ns/op	878 B/op	14 allocs/op
raw:		4134	251857 ns/op	749 B/op	13 allocs/op
sqlx:		2714	497815 ns/op	872 B/op	20 allocs/op
gorp:		1966	581419 ns/op	1205 B/op	32 allocs/op
beego:		2172	588442 ns/op	1752 B/op	47 allocs/op
sqlboiler:	2179	588778 ns/op	902 B/op	17 allocs/op
reform:		2121	593651 ns/op	1775 B/op	51 allocs/op
pgx:		2214	600386 ns/op	273 B/op	10 allocs/op
pgx_pool:	1932	605994 ns/op	285 B/op	10 allocs/op
ent:		1792	611630 ns/op	4680 B/op	97 allocs/op
pop:		2130	627177 ns/op	6044 B/op	186 allocs/op
gorm_prep:	1971	639163 ns/op	5008 B/op	56 allocs/op
dbr:		1864	640274 ns/op	2651 B/op	57 allocs/op
bun:		2041	658366 ns/op	4729 B/op	5 allocs/op
pg:		1887	694214 ns/op	768 B/op	9 allocs/op
gorm:		1688	729914 ns/op	6752 B/op	99 allocs/op
rel:		1378	938263 ns/op	3048 B/op	45 allocs/op
xorm:		1114	938896 ns/op	3944 B/op	132 allocs/op
godb:		1246	944222 ns/op	5113 B/op	154 allocs/op
zorm:		1122	1110298 ns/op	3024 B/op	59 allocs/op
upper:		661	1985264 ns/op	16649 B/op	390 allocs/op

Read
pgx:		5856	212615 ns/op	893 B/op	8 allocs/op
pgx_pool:	5961	219021 ns/op	1078 B/op	9 allocs/op
sqlc:		5265	244057 ns/op	2078 B/op	51 allocs/op
beego:		4921	244872 ns/op	2096 B/op	76 allocs/op
raw:		4722	247777 ns/op	2061 B/op	50 allocs/op
pg:		4328	259957 ns/op	872 B/op	20 allocs/op
sqlboiler:	4612	268293 ns/op	945 B/op	14 allocs/op
dbr:		4938	276451 ns/op	2184 B/op	37 allocs/op
pop:		4832	279101 ns/op	3149 B/op	67 allocs/op
rel:		4670	283516 ns/op	2304 B/op	47 allocs/op
bun:		4461	285937 ns/op	5810 B/op	39 allocs/op
gorm_prep:	4383	287055 ns/op	4403 B/op	87 allocs/op
gorp:		4234	288366 ns/op	3878 B/op	194 allocs/op
reform:		4243	291442 ns/op	3198 B/op	86 allocs/op
ent:		3886	305047 ns/op	5605 B/op	144 allocs/op
gorm:		3843	337395 ns/op	4772 B/op	98 allocs/op
zorm:		3697	359331 ns/op	3016 B/op	64 allocs/op
sqlx:		2505	495102 ns/op	1976 B/op	43 allocs/op
xorm:		1828	549746 ns/op	4648 B/op	127 allocs/op
godb:		2426	551646 ns/op	4064 B/op	102 allocs/op
upper:		1846	609403 ns/op	5055 B/op	110 allocs/op
```
