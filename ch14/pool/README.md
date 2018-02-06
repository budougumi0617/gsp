# 14.2.7 ワーカープール
CPUコア数分の`goroutine`をワーカーとして生成して処理する。

# Result
```bash
$ go run main.go
year=2 total=40458322 interest=458322 id=3
year=5 total=41118307 interest=1118307 id=3
year=6 total=41338299 interest=1338299 id=3
year=7 total=41558295 interest=1558295 id=3
year=1 total=40238328 interest=238328 id=1
year=9 total=41998284 interest=1998284 id=1
year=10 total=42218280 interest=2218280 id=1
year=11 total=42438261 interest=2438261 id=1
year=8 total=41778286 interest=1778286 id=3
year=13 total=42878258 interest=2878258 id=3
year=14 total=43098252 interest=3098252 id=3
year=15 total=43318246 interest=3318246 id=3
year=12 total=42658264 interest=2658264 id=1
year=3 total=40678317 interest=678317 id=0
year=18 total=43978228 interest=3978228 id=0
year=19 total=44198221 interest=4198221 id=0
year=16 total=43538236 interest=3538236 id=3
year=20 total=44418226 interest=4418226 id=0
year=17 total=43758231 interest=3758231 id=1
year=22 total=44858188 interest=4858188 id=0
year=23 total=45078196 interest=5078196 id=1
year=4 total=40898311 interest=898311 id=2
year=25 total=45518199 interest=5518199 id=1
year=26 total=45738179 interest=5738179 id=2
year=21 total=44638218 interest=4638218 id=3
year=28 total=46178169 interest=6178169 id=2
year=29 total=46398162 interest=6398162 id=3
year=24 total=45298192 interest=5298192 id=0
year=27 total=45958172 interest=5958172 id=1
year=32 total=47058142 interest=7058142 id=0
year=33 total=47278129 interest=7278129 id=1
year=34 total=47498129 interest=7498129 id=0
year=30 total=46618159 interest=6618159 id=2
year=35 total=47718127 interest=7718127 id=1
year=31 total=46838148 interest=6838148 id=3
```
