import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';
import { WsService } from 'src/app/game/ws.service';
import { clientStart } from 'src/app/game/clientEngine';
import * as $ from 'jquery';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css'],
    standalone: false
})
export class LoginComponent implements OnInit {
  @ViewChild('emailaddress') emailaddress: ElementRef;
  @ViewChild('password') password: ElementRef;

  loginSuccess = false;

  constructor(
    private accountService: AccountService,
    private wsService: WsService,
  ) {}

  ngOnInit(): void {}

  async login() {
    // try to sign in
    const s = await this.accountService.login({
      emailaddress: this.emailaddress.nativeElement.value,
      password: this.password.nativeElement.value,
    });

    if (s.success) {
      // rewrite to canvas
      this.loginSuccess = true;

      setTimeout(() => {
        // remove extra stuff from page to avoid interference with main canvas
        const layoutContainer = document.getElementsByClassName(
          'helia-game-top',
        )[0].parentNode.parentNode as any;

        layoutContainer.removeAttribute('class');

        $('#site-header').remove();
        $('#site-container').removeAttr('id');
        $('#site-footer').remove();
        $('router-outlet').remove();
        $('br').remove();

        // prevent overflow causing scroll in body
        const documentBody = document.getElementsByTagName('body')[0] as any;
        documentBody.style.overflow = 'hidden';
        documentBody.style.height = '100%';

        setTimeout(() => {
          // get back canvas
          const backCanvas = document.getElementsByClassName(
            'backCanvas',
          )[0] as any;

          // get main canvas
          const gameCanvas = document.getElementsByClassName(
            'gameCanvas',
          )[0] as any;

          setTimeout(() => {
            // transfer control to game engine
            clientStart(this.wsService, gameCanvas, backCanvas, s.sid);
          }, 100);
        }, 100);
      }, 100);
    } else {
      alert('Login failed: ' + s.message);
    }
  }
}
