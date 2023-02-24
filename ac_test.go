package godsa

import (
	"testing"
	"unicode"
)

func TestAddKw(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("ac")
	trie.AddKeyword("de")

	emits := trie.ParseText("acde")
	trie.Print()
	for _, e := range emits {
		t.Log(e)
	}
}

func TestKeywordAndTextAreTheSame(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("abc")
	emits := trie.ParseText("abc")
	for _, e := range emits {
		t.Log(e)
	}
}

func TestVariousKeywordsOneMatch(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("abc")
	trie.AddKeyword("ert")
	trie.AddKeyword("be")
	emits := trie.ParseText("abeaert")
	for _, e := range emits {
		t.Log(e)
	}
}

func TestUshers(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("hers")
	trie.AddKeyword("his")
	trie.AddKeyword("she")
	trie.AddKeyword("he")
	emits := trie.ParseText("ushers")
	trie.Print()
	for _, e := range emits {
		t.Log(e)
	}
	if len(emits) != 3 {
		t.Error()
	}
}

func TestMisLeading(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("hers")
	emits := trie.ParseText("h he her hers")
	for _, e := range emits {
		t.Log(e)
	}
	if len(emits) != 1 {
		t.Error()
	}
}

func TestRecipes(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("veal")
	trie.AddKeyword("cauliflower")
	trie.AddKeyword("broccoli")
	trie.AddKeyword("tomatoes")
	emits := trie.ParseText("2 cauliflowers, 3 tomatoes, 4 slices of veal, 100g broccoli")
	for _, e := range emits {
		t.Log(e)
	}
	if len(emits) != 4 {
		t.Error()
	}
}

func TestLongAndShortOverlappingMatch(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("he")
	trie.AddKeyword("hehehehe")
	emits := trie.ParseText("hehehehehe")
	for _, e := range emits {
		t.Log(e)
	}
	if len(emits) != 7 {
		t.Error()
	}
}

func TestCNText(t *testing.T) {
	trie := NewTrie()
	trie.AddKeyword("消费")
	trie.AddKeyword("消费券")
	trie.AddKeyword("活动")
	emits := trie.ParseText(`本次消费券横跨餐饮、服务、娱乐、零售多个领域，经开区通过联合龙湖天街、大族广场、北京华联亦庄购物中心、文化创意生活广场四大商圈各类商家，重点围绕民生需求，开展形式多样、特色鲜明的系列促消费活动，激发消费潜力，进一步丰富消费供给，跑出消费“加速度”，点旺城市烟火气。此外，汽车消费活动补贴为1000万元，将持续到3月底，对辖区车企和消费者来说是一剂“强心针”，鼓励汽车销售企业更新迭代消费，优化汽车销售结构，促进汽车消费。`)
	trie.Print()
	for _, e := range emits {
		t.Log(e)
	}
}

func TestCi(t *testing.T) {
	t.Logf("%c", unicode.ToLower(rune('A')))
	t.Log(len([]rune("消费")))

	type I interface {
		s() int
	}

	var i I
	t.Log(i.s())
}
