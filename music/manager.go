package music

import "errors"

type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{
		musics: make([]Music, 0),
	}
}

// 列表长度
func (m *MusicManager) Len() int {
	return len(m.musics)
}

// 删除
func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index > m.Len() {
		return nil
	}
	removeMusic := &m.musics[index]

	if index < m.Len()-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]Music, 0)
	} else {
		m.musics = m.musics[:index-1]
	}

	return removeMusic
}

// 获取
func (m *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index > m.Len() {
		return nil, errors.New("Index out of range")
	}

	return &m.musics[index], nil
}

// 添加
func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}

// 查找
func (m *MusicManager) Find(name string) *Music {
	if m.Len() == 0 {
		return nil
	}

	for _, m1 := range m.musics {
		if m1.Name == name {
			return &m1
		}
	}

	return nil
}