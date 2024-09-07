export class SoundtrackDataRepository {
    allTracks(): string[] {
        // thank you zakè (zakedrone.com) :)
        const o = [
            'zakè & Benoît Pioulard - eve - 01 eve.mp3',
            'zakè & Benoît Pioulard - eve - 02 frost.mp3',
            'zakè & Benoît Pioulard - eve - 03 pine.mp3',
            'zakè & City of Dawn - An Eternal Moment Hidden Away - 01 True North.mp3',
            'zakè & City of Dawn - An Eternal Moment Hidden Away - 02 Atrium.mp3',
            'zakè & City of Dawn - An Eternal Moment Hidden Away - 03 An Eternal Moment Hidden Away.mp3',
            'zakè & City of Dawn - An Eternal Moment Hidden Away - 04 Fragments of the Mosaic II.mp3',
            'zakè & T.R. Jordan - Stay With Me - 01 Shoreline.mp3',
            'zakè & T.R. Jordan - Stay With Me - 02 Infinite Sand.mp3',
            'zakè & T.R. Jordan - Stay With Me - 03 Longing Memories.mp3',
            'zakè & T.R. Jordan - Stay With Me - 04 Stay With Me -feat. marine eyes-.mp3',
            'zakè & Tyresta - Drift - 01 Monuments at Sea.mp3',
            'zakè & Tyresta - Drift - 02 Through the Mist.mp3',
            'zakè & Tyresta - Drift - 03 Ashes In the Wind.mp3',
            'zakè & Tyresta - Drift - 04 Drift.mp3',
            'zakè & Tyresta - Drift - 05 Reason to Believe.mp3',
            'zakè - Veta - 01 Veta.mp3',
            'zakè - Veta - 02 Bewrayeth, No. 2.mp3',
            'zakè - Veta - 03 Glory.mp3',
            'zakè - Veta - 04 Memorial.mp3'
        ];

        return o;
    }

    randomTrack(): string {
        // get track list
        const o = this.allTracks();

        // return a random track
        return o[o.length * Math.random() | 0]
    }
}
