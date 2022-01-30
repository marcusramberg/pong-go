package game

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type Sfx struct {
	audioContext *audio.Context
	volume128    int
	sounds       map[string]*audio.Player
}

func NewSfx() (*Sfx, error) {

	soundFiles := map[string]string{"beep": "assets/ping_pong_8bit_beeep.wav", "peep": "assets/ping_pong_8bit_peeeeeep.wav", "plop": "assets/ping_pong_8bit_plop.wav"}
	sounds  := make(map[string]*audio.Player)
	audioContext := audio.NewContext(44100)
	for k, v := range soundFiles {
		f, err := ebitenutil.OpenFile(v)
		if err != nil {
			return nil, err
		}
		d, err := wav.Decode(audioContext, f)
		if err != nil {
			return nil, err
		}
		audioPlayer, err := audio.NewPlayer(audioContext, d)
		if err != nil {
			return nil, err
		}
		sounds[k] = audioPlayer
	}
	return &Sfx{audioContext: audioContext, volume128: 128, sounds: sounds}, nil

}

func (sfx *Sfx) Play(sound string) error {
	player:= sfx.sounds[sound]
	if player == nil { return errors.New("Sound not found") }
	err := player.Rewind()
	if err != nil {
 	  return err
	}
	player.Play()
	return nil
}
