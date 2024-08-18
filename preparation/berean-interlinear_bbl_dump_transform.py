import codecs
start_string = "INSERT INTO Bible VALUES("
end_string = ");"

input_example = "INSERT INTO Bible VALUES(1,1,1,'<TS>The Creation<br>(<RX 43.1.1-5> <RX 58.11.1-3>)<Ts><rtl/> <Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q> <Q><H>בָּרָ֣א<WH1254><WTV-Qal-Perf-3ms><h><X>bā·rā<x><E> created <e><q> <Q><H>אֱלֹהִ֑ים<WH430><WTN-mp><h><X>’ĕ·lō·hîm<x><E> God <e><q> <Q><H>אֵ֥ת<WH853><WTDirObjM><h><X>’êṯ<x><E> - <e><q> <Q><H>הַשָּׁמַ֖יִם<WH8064><WTArt│N-mp><h><X>haš·šā·ma·yim<x><E> the heavens <e><q> <Q><H>וְאֵ֥ת<WH853><WTConj-w│DirObjM><h><X>wə·’êṯ<x><E> and <e><q> <Q><H>הָאָֽרֶץ׃<WH776><WTArt│N-fs><h><X>hā·’ā·reṣ<x><E> the earth . <e><q>');"
output_example = '"1_1_1": {1, 1, 1, genesis11},'

with open("bereaninterlinear.go", "w") as output_file:
    output_file.write("package base\n\n")
    output_file.write("func GetBereanInterlinear() map[string]VerseType {\n")
    output_file.write("    bereanInterlinear := map[string]VerseType{\n")
    with open("berean-interlinear.bbl.dump.sql") as input_file:
        for line in input_file.readlines():
            if line.find(start_string) < 0:
                continue
            line = line[25:-2]
            data = line.split(",", 3)
            data[3] = data[3][1:-2]
            data[3] = codecs.encode(data[3], 'unicode-escape').decode().replace("\\x", "\\u00")
            format_string = "        \"{book}_{chapter}_{verse}\": {o}{book},{chapter},{verse},\"{text}\"{c},\n"
            output_line = format_string.format(book=data[0], chapter=data[1], verse=data[2], text=data[3], o='{', c='}')
            output_file.write(output_line)
    output_file.write("    }\n\n    return bereanInterlinear;\n}\n\n")

