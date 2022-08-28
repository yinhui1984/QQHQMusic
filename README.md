# QQHQMusic
 QQ音乐高品质下载 以`flac`为主



破解原理来自于这里 https://github.com/QiuChenly/QQFlacMusicDownloader



搜索时输入的关键字可以是 歌曲名称 歌手 专辑 排行榜名称等

搜索结果默认返回最大50首 (筛选掉了低品质音乐)

播放调用的是ffplay, 确保电脑中安装了

批量下载时默认10首同时下载



> 下载时国外IP比国内IP快, 神奇的TX



示例

```
输入关键词进行搜索:
=> Beautiful Girls
┏━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃  # ┃ TITLE                                    ┃ SINGER           ┃ FORMAT ┃ SIZE(M) ┃ ALBUM                                        ┃
┣━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃  0 ┃ Beautiful Girls (Remix)                  ┃ 小美人           ┃ flac   ┃ 18 MB   ┃ 《Umbrella》                                 ┃
┃  1 ┃ Beautiful Girls                          ┃ Sean Kingston    ┃ flac   ┃ 27 MB   ┃ 《Me Love》                                  ┃
┃  2 ┃ Beautiful Girls (降调0.8x)               ┃ 常在梦里的人     ┃ flac   ┃ 13 MB   ┃ 《Love and loved》                           ┃
┃  3 ┃ Beautiful Girls (Bonus Track)            ┃ 少女时代         ┃ flac   ┃ 30 MB   ┃ 《The 1st Asia Tour Into The New World ...》 ┃
┃  4 ┃ Beautiful Girls (0.9x)                   ┃ 乔一魚           ┃ flac   ┃ 30 MB   ┃ 《Beautiful Girls》                          ┃
┃  5 ┃ Beautiful Girl (0.8x)                    ┃ J-I-E            ┃ mp3    ┃ 11 MB   ┃ 《U Make Me》                                ┃
┃  6 ┃ Beautiful Girls                          ┃ 宇宙无敌小钢蛋   ┃ flac   ┃ 15 MB   ┃ 《我与神明画押堵你心动一刹》                 ┃
┃  7 ┃ Beautiful Girls (BGM版)                  ┃ 五音旋律         ┃ mp3    ┃ 11 MB   ┃ 《超帅卡点》                                 ┃
┃  8 ┃ Beautiful Girls                          ┃ DJ新文           ┃ flac   ┃ 13 MB   ┃ 《Bet On Me》                                ┃
┃  9 ┃ Beautiful Girl                           ┃ Jazz Guitar Club ┃ flac   ┃ 37 MB   ┃ 《Satin: Jazzy Rock Ballads Instrumenta...》 ┃
┃ 10 ┃ Beautiful Girl                           ┃ DJ铁柱           ┃ mp3    ┃ 11 MB   ┃ 《EA7压迫旋律》                              ┃
┃ 11 ┃ Beautiful Girl                           ┃ 郭富城           ┃ flac   ┃ 20 MB   ┃ 《对你爱不完》                               ┃
┃ 12 ┃ beautiful girls (Gustixa Remix)          ┃ fawlin           ┃ flac   ┃ 29 MB   ┃ 《beautiful girls (Gustixa Remix)》          ┃
┃ 13 ┃ Beautiful Girl (Explicit)                ┃ DJake            ┃ flac   ┃ 18 MB   ┃ 《Beautiful Girl (Explicit)》                ┃
┃ 14 ┃ Beautiful Girls                          ┃ xxxCr3           ┃ flac   ┃ 11 MB   ┃ 《Ayo》                                      ┃
┃ 15 ┃ Beautiful Girls                          ┃ Danny Avila      ┃ flac   ┃ 22 MB   ┃ 《Beautiful Girls》                          ┃
┃ 16 ┃ Beautiful Girls                          ┃ Van Halen        ┃ flac   ┃ 30 MB   ┃ 《Van Halen 2》                              ┃
┃ 17 ┃ Beautiful Girls (Remix)                  ┃ 刘小蕊           ┃ flac   ┃ 16 MB   ┃ 《Glosses》                                  ┃
┃ 18 ┃ Beautiful Girls (Re-Recorded)            ┃ Sean Kingston    ┃ flac   ┃ 26 MB   ┃ 《Sean Kingston Hits (2007-2010) (The R...》 ┃
┃ 19 ┃ BEAUTIFUL GIRL                           ┃ TEEN TOP         ┃ flac   ┃ 29 MB   ┃ 《ROMAN》                                    ┃
┃ 20 ┃ Beautiful Girls (Originally Performed... ┃ Backing Business ┃ flac   ┃ 27 MB   ┃ 《Pristine Karaoke, Vol. 79》                ┃
┃ 21 ┃ Beautiful Girls (Live 1994)              ┃ David Lee Roth   ┃ mp3    ┃ 17 MB   ┃ 《Double Stack Thunder (Live 1994)》         ┃
┃ 22 ┃ Beautiful Girls (0.8X)                   ┃ 冠军             ┃ mp3    ┃ 12 MB   ┃ 《音乐拯救世界》                             ┃
┃ 23 ┃ Beautiful Girl                           ┃ Inxs             ┃ flac   ┃ 24 MB   ┃ 《Welcome to Wherever You Are》              ┃
┃ 24 ┃ Beautiful Girls (Andie Roy X Suntimec... ┃ Andie Roy        ┃ mp3    ┃ 13 MB   ┃ 《Beautiful Girls (Andie Roy X Suntimec...》 ┃
┃ 25 ┃ beautiful girls                          ┃ fawlin           ┃ flac   ┃ 26 MB   ┃ 《beautiful girls》                          ┃
┃ 26 ┃ Beautiful Girl                           ┃ The Hit Crew     ┃ flac   ┃ 25 MB   ┃ 《R&B Hits of the 80s, 90s and 2000s, V...》 ┃
┃ 27 ┃ Beautiful Girls (8-Bit Sean Kingston ... ┃ 8-Bit Arcade     ┃ flac   ┃ 24 MB   ┃ 《By Request, Vol. 188》                     ┃
┗━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛


选择和执行命令
    比如播放第1首:p 1
    比如下载第1首:d 1
    比如下载全部:da
    退出循环:q

=> p 1
2022/08/28 20:28:40 准备播放 Beautiful Girls-Sean Kingston, 获取链接:   http://ws.stream.qqmusic.qq.com/F000002KECPZ1eg9O9.flac?guid=QMD50&vkey=E789E7D842EACEEA732EC23E83E935CA3A25DF6BCFF35D4A88DEA449D8EE4CAB89BBDDC4A7841C74AB56F0CCA93FB96A843432E2750A52B4&uin=350577342&fromtag=140

=> d 1
2022/08/28 20:29:41 Downloading: Beautiful Girls-Sean Kingston [27 MB]
Downloading... 27 MB complete
已下载:  /Users/zhouyinhui/Music/QQHDMusic/单曲/Beautiful Girls-Sean Kingston.flac

=> da
2022/08/28 20:29:53 Downloading: Beautiful Girls-Sean Kingston [27 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls (BGM版)-五音旋律 [11 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls-宇宙无敌小钢蛋 [15 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls (Bonus Track)-少女时代 [30 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls-DJ新文 [13 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls (Remix)-小美人 [18 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girl-Jazz Guitar Club [37 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girl (0.8x)-J-I-E [11 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls (降调0.8x)-常在梦里的人 [13 MB]
2022/08/28 20:29:53 Downloading: Beautiful Girls (0.9x)-乔一魚 [30 MB]
2022/08/28 20:29:57 Downloading: Beautiful Girl-DJ铁柱 [11 MB]
2022/08/28 20:29:57 Downloading: Beautiful Girl-郭富城 [20 MB]
2022/08/28 20:29:58 Downloading: beautiful girls (Gustixa Remix)-fawlin [29 MB]
2022/08/28 20:30:00 Downloading: Beautiful Girl (Explicit)-DJake [18 MB]
2022/08/28 20:30:02 Downloading: Beautiful Girls-xxxCr3 [11 MB]
2022/08/28 20:30:02 Downloading: Beautiful Girls-Danny Avila [22 MB]
2022/08/28 20:30:04 Downloading: Beautiful Girls (Remix)-刘小蕊 [16 MB]
2022/08/28 20:30:05 Downloading: Beautiful Girls (Re-Recorded)-Sean Kingston [26 MB]
2022/08/28 20:30:05 Downloading: BEAUTIFUL GIRL-TEEN TOP [29 MB]
2022/08/28 20:30:07 Downloading: Beautiful Girls-Van Halen [30 MB]
2022/08/28 20:30:08 Downloading: Beautiful Girls (Live 1994)-David Lee Roth [17 MB]
2022/08/28 20:30:09 Downloading: Beautiful Girls (Originally Performed by Sean Kingston) (Instrumental Version)-Backing Business [27 MB]
2022/08/28 20:30:11 Downloading: Beautiful Girls (0.8X)-冠军 [12 MB]
2022/08/28 20:30:15 Downloading: Beautiful Girl-Inxs [24 MB]
2022/08/28 20:30:16 Downloading: Beautiful Girls (Andie Roy X Suntimechild Remix)-Andie Roy [13 MB]
2022/08/28 20:30:17 Downloading: Beautiful Girl-The Hit Crew [25 MB]
2022/08/28 20:30:17 Downloading: beautiful girls-fawlin [26 MB]
2022/08/28 20:30:19 Downloading: Beautiful Girls (8-Bit Sean Kingston Emulation)-8-Bit Arcade [24 MB]
2022/08/28 20:30:35 DOWNLOAD ALL DONE

=>
```



