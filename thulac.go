package gothulac

/*
#cgo CXXFLAGS: -I./include -fPIC
#include "thulac_lib.h"
*/
import "C"
import (
	"errors"
	"regexp"
	"strings"
)

type WordInfo struct {
	Word      string
	WordClass string
	Start     int
	End       int
}

var splitRegex *regexp.Regexp

func init() {
	splitRegex = regexp.MustCompile("^(.*?)_([a-z]+)$")
}

func LoadModel(modelPath string, userDictPath string, preAllocSize int, t2s int, justSeg int) {
	C.init(C.CString(modelPath), C.CString(userDictPath), C.int(preAllocSize), C.int(t2s), C.int(justSeg))
}

func CutSentence(sentence string) ([]WordInfo, error) {
	state := int(C.seg(C.CString(sentence)))
	if state == 0 {
		return nil, errors.New("segmentation error")
	}
	resultPtr := C.getResult()
	defer C.freeResult()
	goResult := C.GoString(resultPtr)
	resultList := strings.Split(goResult, " ")
	cutWordList := []WordInfo{}
	wordStartPos := 0
	for _, v := range resultList {
		m := splitRegex.FindAllStringSubmatch(v, -1)
		wordChars := strings.Split(m[0][1], "")
		wordInfo := WordInfo{
			Word:      m[0][1],
			WordClass: m[0][2],
			Start:     wordStartPos,
			End:       wordStartPos + len(wordChars),
		}
		cutWordList = append(cutWordList, wordInfo)
		wordStartPos += len(wordChars)
	}
	return cutWordList, nil
}
