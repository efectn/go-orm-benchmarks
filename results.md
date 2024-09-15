# Results

- orm-benchmark (with no flags)

```
Reports:

goos: windows
goarch: amd64
cpu: AMD Ryzen 9 7900 12-Core Processor
psql (17rc1)

Read
raw:		56202	21695 ns/op	2093 B/op	50 allocs/op
pgx:		55284	22359 ns/op	776 B/op	18 allocs/op
beego:		53220	22512 ns/op	2115 B/op	75 allocs/op
sqlc:		52897	23177 ns/op	904 B/op	19 allocs/op
reform:		49404	23996 ns/op	3235 B/op	86 allocs/op
pgx_pool:	47500	24235 ns/op	963 B/op	19 allocs/op
pop:		42277	28525 ns/op	3304 B/op	66 allocs/op
ent:		41748	28880 ns/op	5685 B/op	145 allocs/op
gorm_prep:	37339	30453 ns/op	4584 B/op	89 allocs/op
sq:		    40525	31968 ns/op	11172 B/op	135 allocs/op
gorp:		40292	32044 ns/op	3909 B/op	194 allocs/op
sqlboiler:	31477	37785 ns/op	977 B/op	14 allocs/op
bun:		30330	39906 ns/op	6005 B/op	48 allocs/op
dbr:		28796	40637 ns/op	2184 B/op	36 allocs/op
rel:		29340	41130 ns/op	2336 B/op	47 allocs/op
jet:		29914	45625 ns/op	13082 B/op	273 allocs/op
zorm:		26124	46996 ns/op	3341 B/op	74 allocs/op
pg:		    21997	51967 ns/op	873 B/op	20 allocs/op
gorm:		20793	57503 ns/op	5000 B/op	102 allocs/op
sqlx:		18501	66069 ns/op	2008 B/op	43 allocs/op
godb:		16830	69508 ns/op	4101 B/op	102 allocs/op
xorm:		16609	70896 ns/op	5166 B/op	131 allocs/op
gen:		15051	74677 ns/op	10550 B/op	154 allocs/op
upper:		16030	75628 ns/op	5086 B/op	110 allocs/op

ReadSlice
reform:		44046	28640 ns/op	    4547 B/op	112 allocs/op
pgx:		19041	62141 ns/op	    30320 B/op	513 allocs/op
pgx_pool:	17004	68399 ns/op	    30380 B/op	513 allocs/op
sqlc:		17185	76693 ns/op	    54625 B/op	620 allocs/op
raw:		15013	79850 ns/op	    38373 B/op	1038 allocs/op
upper:		13708	86207 ns/op	    4822 B/op	90 allocs/op
ent:		10000	118526 ns/op	77397 B/op	2036 allocs/op
pop:		9259	121879 ns/op	78007 B/op	1306 allocs/op
sqlx:		9300	124204 ns/op	37512 B/op	1225 allocs/op
gorp:		9891	126972 ns/op	57593 B/op	1494 allocs/op
bun:		8560	139211 ns/op	34269 B/op	1124 allocs/op
pg:		    9058	141270 ns/op	23863 B/op	629 allocs/op
dbr:		8112	144384 ns/op	30944 B/op	1253 allocs/op
beego:		8048	144995 ns/op	55397 B/op	3077 allocs/op
sq:		    7861	147749 ns/op	152663 B/op	1829 allocs/op
gorm_prep:	7838	153621 ns/op	43685 B/op	2082 allocs/op
sqlboiler:	7518	156034 ns/op	67627 B/op	2260 allocs/op
gorm:		6540	183580 ns/op	44886 B/op	2196 allocs/op
gen:		5250	212713 ns/op	50117 B/op	2243 allocs/op
godb:		5587	216216 ns/op	75455 B/op	3084 allocs/op
xorm:		5161	233580 ns/op	121333 B/op	4407 allocs/op
zorm:		4990	236774 ns/op	169687 B/op	2959 allocs/op
rel:		5162	236917 ns/op	149041 B/op	2553 allocs/op
jet:		4590	255358 ns/op	192881 B/op	3642 allocs/op

Insert
raw:		16995	69016 ns/op	    703 B/op	13 allocs/op
sqlc:		16189	73017 ns/op	    280 B/op	8 allocs/op
pgx_pool:	16191	74545 ns/op	    298 B/op	8 allocs/op
pgx:		15680	75946 ns/op	    280 B/op	8 allocs/op
gorm_prep:	15361	77281 ns/op	    5198 B/op	66 allocs/op
sqlboiler:	15483	77582 ns/op	    1591 B/op	35 allocs/op
gorp:		15498	78233 ns/op	    1797 B/op	42 allocs/op
reform:		15205	78823 ns/op	    1777 B/op	51 allocs/op
beego:		15195	80149 ns/op	    2403 B/op	57 allocs/op
ent:		14725	81861 ns/op	    4150 B/op	98 allocs/op
jet:		13592	88131 ns/op	    3583 B/op	105 allocs/op
sq:		    14022	91575 ns/op	    9786 B/op	101 allocs/op
dbr:		11965	96681 ns/op	    2688 B/op	65 allocs/op
bun:		11092	106712 ns/op	5043 B/op	14 allocs/op
gorm:		10000	107692 ns/op	7215 B/op	105 allocs/op
sqlx:		10000	111007 ns/op	856 B/op	19 allocs/op
pg:		    10000	113662 ns/op	795 B/op	10 allocs/op
gen:		10000	117440 ns/op	10050 B/op	132 allocs/op
xorm:		9663	121127 ns/op	3123 B/op	88 allocs/op
godb:		9294	126215 ns/op	4541 B/op	116 allocs/op
zorm:		9032	128129 ns/op	5150 B/op	104 allocs/op
rel:		8763	128676 ns/op	2639 B/op	45 allocs/op
upper:		8616	134489 ns/op	5912 B/op	126 allocs/op
pop:		6121	187828 ns/op	9456 B/op	235 allocs/op

InsertMulti
sqlc:		4275	293523 ns/op	67617 B/op	642 allocs/op
raw:		3435	356378 ns/op	187041 B/op	929 allocs/op
sq:		    2940	401974 ns/op	237553 B/op	1706 allocs/op
beego:		2547	403591 ns/op	177805 B/op	2744 allocs/op
pgx:		4642	407619 ns/op	48373 B/op	40 allocs/op
pgx_pool:	2886	416123 ns/op	48468 B/op	39 allocs/op
gorm_prep:	2622	458123 ns/op	254607 B/op	1890 allocs/op
reform:		2494	469998 ns/op	462567 B/op	2746 allocs/op
ent:		2256	529270 ns/op	396508 B/op	4596 allocs/op
sqlx:		1735	638379 ns/op	170480 B/op	1550 allocs/op
bun:		1814	658984 ns/op	42610 B/op	219 allocs/op
zorm:		1614	682796 ns/op	212289 B/op	2808 allocs/op
jet:		1647	689081 ns/op	338335 B/op	6493 allocs/op
xorm:		1657	700879 ns/op	259138 B/op	5518 allocs/op
pg:		    1569	725557 ns/op	4037 B/op	112 allocs/op
gorm:		1618	729755 ns/op	276386 B/op	5230 allocs/op
gen:		1552	767914 ns/op	289585 B/op	5354 allocs/op
rel:		1432	769388 ns/op	312562 B/op	3265 allocs/op
upper:		1456	789388 ns/op	328157 B/op	4204 allocs/op
godb:		1374	833119 ns/op	260963 B/op	5894 allocs/op
sqlboiler:	bulk-insert is not supported
pop:		bulk-insert is not supported
gorp:		bulk-insert is not supported
dbr:		bulk-insert is not supported

Update
raw:		59931	19930 ns/op	    749 B/op	13 allocs/op
sqlc:		47966	22559 ns/op	    288 B/op	8 allocs/op
sqlx:		20326	60751 ns/op	    872 B/op	20 allocs/op
sqlboiler:	15908	75355 ns/op	    902 B/op	17 allocs/op
gorp:		15728	76933 ns/op	    1206 B/op	32 allocs/op
gorm_prep:	15747	77430 ns/op	    5072 B/op	56 allocs/op
pgx_pool:	15159	78342 ns/op	    305 B/op	8 allocs/op
pgx:		15430	79285 ns/op	    288 B/op	8 allocs/op
beego:		15148	79409 ns/op	    1738 B/op	46 allocs/op
reform:		15296	79819 ns/op	    1776 B/op	51 allocs/op
ent:		13944	83493 ns/op	    4725 B/op	98 allocs/op
jet:		13460	85227 ns/op	    4565 B/op	126 allocs/op
sq:		    14203	86057 ns/op	    7422 B/op	90 allocs/op
pop:		13476	88313 ns/op	    5862 B/op	184 allocs/op
dbr:		10000	107907 ns/op	2651 B/op	57 allocs/op
bun:		10000	113244 ns/op	4767 B/op	5 allocs/op
pg:		    9854	122477 ns/op	768 B/op	9 allocs/op
gorm:		10000	123730 ns/op	6832 B/op	99 allocs/op
gen:		9397	135241 ns/op	13408 B/op	161 allocs/op
xorm:		8818	136194 ns/op	4308 B/op	145 allocs/op
rel:		8492	141163 ns/op	3048 B/op	45 allocs/op
zorm:		8816	142342 ns/op	4453 B/op	84 allocs/op
godb:		8348	147051 ns/op	5150 B/op	154 allocs/op
upper:		4118	287930 ns/op	16751 B/op	390 allocs/op
```
