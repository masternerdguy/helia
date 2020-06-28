import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';
import { WsService } from 'src/app/game/ws.service';
import { ServerJoinBody } from 'src/app/game/wsModels/join';
import { clientStart } from 'src/app/game/clientEngine';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  @ViewChild('username') username: ElementRef;
  @ViewChild('password') password: ElementRef;

  loginSuccess = false;

  screenWidth = window.screen.width;
  screenHeight = window.screen.height;

  constructor(
    private accountService: AccountService,
    private wsService: WsService
  ) {}

  ngOnInit(): void {}

  async login() {
    // try to sign in
    const s = await this.accountService.login({
      username: this.username.nativeElement.value,
      password: this.password.nativeElement.value,
    });

    if (s.success) {
      // rewrite to canvas
      this.loginSuccess = true;

      setTimeout(() => {
        // get back canvas
        const backCanvas = document.getElementsByClassName('backCanvas')[0] as any;

        // enter fullscreen
        const gameCanvas = document.getElementsByClassName('gameCanvas')[0] as any;
        console.log(gameCanvas);

        if (gameCanvas.webkitRequestFullScreen) {
          gameCanvas.webkitRequestFullScreen();
        } else {
          gameCanvas.mozRequestFullScreen();
        }

        // transfer control to game engine
        clientStart(this.wsService, gameCanvas, backCanvas, s.sid);
      });
    } else {
      alert('Login failed: ' + s.message);
    }
  }
}
