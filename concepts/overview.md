# bicon - Bible Context

Bicon is an idea to find more relevant verses through technology.

My favorite example is Mary and Martha (Luk 10, 38-42). The story is about Martha and Mary.
Mary listen to Jesus and Martha is hustling to welcome everyone also with food. The story
is about Jesus reaction.

Why are the desciples and Jesus in the house of Mary and Martha? Because there was a strategy
how the disciples enter a village or a city. Jesus teached them to enter a new place and be
welcomed in a house (Luk 10, 5-12). That was their base in the village.

Now the lost context, when I study the bible is verse 38:

"As Jesus and his disciples were on their way, he came to a village where a woman named
Martha opened her home to him."

This relationship is going deep, because there are a lot of stories in the bible about
Jesus and Martha, Maria and Lazarus.

This strategy and this house is an example given in the same chapter Luke 10. I only found
this context, because the holy spirit revealed it to me. And then I studied also the words
in greek. And on the chosen words (or strongs) you can also reveal this context.

So the Bicon Project (bible context) is reaching out for more context in the bible to
help understanding the scripture.

## Go Version

The go-version is the WebAssembly Backend (and maybe more) for this project. The target is
to have the Berean Study Bible in the go-files and libraries around this.

## Berean Study Bible

[Official Downloads](https://berean.bible/downloads.htm)

[Download](https://www.mysword.info/download-mysword/bibles) of bibles as SQLite Databases

The Holy Bible, Berean Study Bible - courtesy of Bible Hub (www.biblehub.com) BSB
1.9
5.4
28 March 2022

All editions of the Berean Bible, Courtesy of Bible Hub (www.biblehub.com, www.berean.bible):

- www.bereanbible.com - Berean Study Bible (BSB)
- www.literalbible.com - Berean Literal Bible (BLB)
- greekbible.org - Berean Greek Bible (BGB) {see below under Greek (Ancient) section}
- www.interlinearbible.com - Berean Interlinear Bible (BIB) {see below under Interlinear section}

The Holy Bible, Berean Interlinear Bible (OT/Hebrew and NT/Greek), 5 Layers - courtesy of Bible Hub (www.biblehub.com)
BIB+
8.6
42.8
13 October 2022

### Berean License

The Holy Bible, Berean Standard Bible, BSB
2016, 2020, 2023 by Bible Hub

Published by Bible Hub
Pittsburgh, PA 15045 USA
www.biblehub.com

Library of Congress Control Number: 2015961020

The Holy Bible, Berean Standard Bible, BSB is produced in cooperation with Bible Hub,
Discovery Bible, OpenBible.com, and the Berean Bible Translation Committee.
This text of God's Word has been dedicated to the public domain. The BSB text may be quoted
in any form (written, visual, audio, or electronic) without prior permission of the publisher.
You are free to copy, print, and distribute any portion of this text, or the full text itself. Free
resources are available at any of the Berean Bible sites

## MyBible Version

```
$ sqlite3 preparation/berean-bsb.bbl.mybible
sqlite> select * from Bible WHERE Book=1 AND Chapter=1 AND Verse=1;
1|1|1|<TS>The Creation<br><em><RX 43.1.1-5><RX 58.11.1-3></em><Ts> In the beginning God created the heavens and the earth.<CM>

$ sqlite3 preparation/berean-interlinear.bbl.mybible
sqlire> select * from Bible WHERE Book=1 AND Chapter=1 AND Verse=1;
1|1|1|<TS>The Creation<br>(<RX 43.1.1-5> <RX 58.11.1-3>)<Ts><rtl/> <Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q> <Q><H>בָּרָ֣א<WH1254><WTV-Qal-Perf-3ms><h><X>bā·rā<x><E> created <e><q> <Q><H>אֱלֹהִ֑ים<WH430><WTN-mp><h><X>’ĕ·lō·hîm<x><E> God <e><q> <Q><H>אֵ֥ת<WH853><WTDirObjM><h><X>’êṯ<x><E> - <e><q> <Q><H>הַשָּׁמַ֖יִם<WH8064><WTArt│N-mp><h><X>haš·šā·ma·yim<x><E> the heavens <e><q> <Q><H>וְאֵ֥ת<WH853><WTConj-w│DirObjM><h><X>wə·’êṯ<x><E> and <e><q> <Q><H>הָאָֽרֶץ׃<WH776><WTArt│N-fs><h><X>hā·’ā·reṣ<x><E> the earth . <e><q>
```

# Thank you

Thank you for all your work to provide this version!

Thank you

Bible Hub
Pittsburgh, PA 15045 USA
www.biblehub.com

and also thank you

Riversoft
www.mysword.info

