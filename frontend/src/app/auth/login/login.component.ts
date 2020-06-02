import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  @ViewChild('username') username: ElementRef;
  @ViewChild('password') password: ElementRef;

  constructor(private accountService: AccountService) { }

  ngOnInit(): void {
  }

  async login() {
    // try to sign in
    const s = await this.accountService.login({
      username: this.username.nativeElement.value,
      password: this.password.nativeElement.value
    });

    console.log(s);
  }
}
