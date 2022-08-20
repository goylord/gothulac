package gothulac

import "testing"

func TestCutSentence(t *testing.T) {
	modelPath := "/program/python/python-git/thulac/models"
	userDictPath := ""
	preAllocSize := 1024 * 1024 * 16
	t2s := 0
	justSeg := 0

	LoadModel(modelPath, userDictPath, preAllocSize, t2s, justSeg)
	testcase := make(map[string]int)
	testcase["这里是中华人民_共和国的领土"] = 8
	testcase["钓鱼岛是中国的"] = 4
	for k, v := range testcase {
		words, err := CutSentence(k)
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(words) != v {
			t.Errorf("【%s】期望的分词长度是%d， 实际是%d\n", k, v, len(words))
		}
	}

}
