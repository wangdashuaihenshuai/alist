package op

import (
	"regexp"
	"strconv"
	"strings"
)

type MovieNameInfo struct {
	Name        string
	EnglistName string
	Meta        []string
	Type        string
	Raw         string
}

var char = []string{}

var urlReg = regexp.MustCompile(`[a-zA-Z]+\.[a-zA-Z]+\.[com|cn|net]\]`)

var replaceRegs = []*regexp.Regexp{
	regexp.MustCompile(`\d+届-`),
	regexp.MustCompile(`\d+x\d+`),
	regexp.MustCompile(`\d+x\d+`),
	regexp.MustCompile(`no\.\d+`),
}

var replaceWords = []string{
	"【十万度v信 shiwandus】",
	"【十万v信 shiwandus】",
	"【",
	"】",
	"-",
	"]",
	"[",
	"(",
	")",
}

var metas = []string{
	"4k",
	"aac",
	"10bit",
	"中字",
	"国语",
	"mnhd-frds",
	"3audio",
	"1080p",
	"x265",
	"x264",
	"2audio",
	"hd中英双字",
	"bd1080p",
	"x264",
	"chd_eng",
	"双语",
	"720p",
	"chi_eng",
	"bdrip",
	"双语",
	"字幕",
	"hr-hdtv",
	"dts",
	"dual-audio",
	"hr-hdtv",
	"1024x576",
	"x264",
	"dvdrip",
	"2audios-cmct",
	"双音轨",
	"ac3",
	"完整版",
	"加长版",
	"无水印",
	"bluray",
	"x264",
	"国英音轨",
	"flac-cmct",
	"flac",
	"dvdrip",
	"unrated",
	"bluray",
	"ac3",
	"hr-hdtv",
	"4audios",
	"cmct",
	"dc",
	"repack",
	"人人影视",
}

func includeMeta(word string) (string, bool) {
	for _, m := range metas {
		if strings.Contains(word, m) {
			return m, true
		}
	}

	return "", false
}

func splitChars(r rune) bool {
	return r == '.' || r == '(' || r == ')'
}

func splitName(name string) []string {
	ret := []string{}
	for _, v := range strings.FieldsFunc(name, splitChars) {
		if v != "" {
			ret = append(ret, v)
		}
	}

	if len(ret) <= 0 {
		return []string{name}
	}

	return ret
}

func replaceName(word string) string {
	for _, w := range replaceWords {
		word = strings.ReplaceAll(word, w, " ")
	}

	for _, r := range replaceRegs {
		word = r.ReplaceAllString(word, " ")
	}

	return word
}

var videoFileExtensions = []string{
	".3g2",
	".3gp",
	".3gp2",
	".3gpp",
	".asf",
	".asx",
	".avi",
	".flv",
	".m2ts",
	".mkv",
	".mov",
	".mp4",
	".mpg",
	".mpeg",
	".rm",
	".swf",
	".vob",
	".wmv",
	".m4v",
	".m4p",
	".m4b",
	".m4r",
	".mts",
	".ts",
	".tp",
	".trp",
	".webm",
	".f4v",
	".ogv",
	".ogg",
}

func isVideoType(t string) bool {
	for _, vt := range videoFileExtensions {
		if vt == "."+t {
			return true
		}

	}

	return false
}

func IsNumberVideoName(name string) bool {
	words := strings.Split(name, ".")
	if len(words) <= 1 {
		return false
	}

	fileType := words[len(words)-1]
	if !isVideoType(fileType) {
		return false
	}

	fileName := strings.Join(words[:len(words)-1], ".")
	_, err := strconv.ParseFloat(fileName, 64)
	return err == nil
}

func FilterVideoName(name string) string {
	words := splitName(name)
	if len(words) <= 1 {
		return name
	}

	fileType := words[len(words)-1]
	if !isVideoType(fileType) {
		return name
	}

	name = strings.ToLower(name)
	name = urlReg.ReplaceAllString(name, "")
	words = splitName(name)

	if len(words) == 2 {
		return replaceName(words[0])
	}

	formatWords := []string{}
	for _, w := range words[0 : len(words)-1] {
		_, ok := includeMeta(w)
		if !ok {
			formatWords = append(formatWords, w)
		}
	}

	n, err := strconv.ParseFloat(words[0], 64)
	if err == nil && n < 1500 {
		formatWords = formatWords[1:]
	}
	formatWords = append(formatWords, fileType)

	return replaceName(strings.Join(formatWords, "."))
}
