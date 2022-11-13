package music

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("NewMusicManager 测试失败")
	}

	if mm.len() != 0 {
		t.Error("NewMusicManager 测试失败, 不是空的")
	}

	m0 := &Music{
		"1", "Hero", "xxx", "http://xxx.xxx.xx", "MP3",
	}
	mm.Add(m0)

	if mm.len() != 1 {
		t.Error("NewMusicManager Add 测试失败")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("NewMusicManager Find 测试失败")
	}
	if m.Id != m0.Id || m.Name != m0.Name ||
		m.Artist != m0.Artist || m.Source != m0.Source ||
		m.Type != m0.Type {
			t.Error("NewMusicManager Find 测试失败  不匹配")
	}

	m, err := mm.Get(0)
	if err != nil {
		t.Error("NewMusicManager Get ", err)
	}
	if m == nil {
		t.Error("NewMusicManager Get 测试失败")
	}
	
	m = mm.Remove(0)
	if mm.len() != 0 {
		t.Error("NewMusicManager Remove 测试失败")
	}
}