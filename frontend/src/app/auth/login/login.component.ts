import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';
import { WsService } from 'src/app/game/ws.service';
import { ServerJoinBody } from 'src/app/game/wsModels/join';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  @ViewChild('username') username: ElementRef;
  @ViewChild('password') password: ElementRef;

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
      // test connect
      this.wsService.connect(s.sid, (d, ws) => {
        console.log({
          data: d,
          body: JSON.parse(d.body) as ServerJoinBody
        });
      });
    } else {
      alert('Login failed: ' + s.message);
    }
  }
}
