import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';
import { WsService } from 'src/app/game/ws.service';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css'],
})
export class ForgotPasswordComponent implements OnInit {
  @ViewChild('emailaddress') emailaddress: ElementRef;

  constructor(
    private accountService: AccountService,
    private wsService: WsService
  ) {}

  ngOnInit(): void {}

  async submit() {
    // todo
  }
}
