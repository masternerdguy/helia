import { SoundtrackDataRepository } from '../data/soundtrackData';
import { Howl } from 'howler';

let playing: boolean = false;
let volume: number = 0.75;
let sound: Howl = undefined;

export class Jukebox {
  begin() {
    // begin playback
    playTrack(nextTrack());

    // mark as started
    playing = true;
  }
}

function nextTrack(): string {
  return new SoundtrackDataRepository().randomTrack();
}

function playTrack(t: string) {
  // load next track and hook following
  sound = new Howl({
    src: ['/assets/music/' + t],
    html5: true,
    loop: false,
    volume: volume,
    onend: function () {
      // to prevent overflows
      setTimeout(() => {
        // play next random track
        playTrack(nextTrack());
      }, 100);
    },
  });

  // start playing
  sound.play();
}
