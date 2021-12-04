import { Component, ElementRef, ViewChild, OnInit } from '@angular/core';
import { AccountService } from '../account.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
})
export class SignupComponent implements OnInit {
  @ViewChild('emailaddress') emailaddress: ElementRef;
  @ViewChild('charactername') charactername: ElementRef;
  @ViewChild('startid') startid: ElementRef;
  @ViewChild('password') password: ElementRef;

  constructor(private accountService: AccountService) {}

  ngOnInit(): void {}

  async register() {
    // try to create account
    await this.accountService.register({
      emailaddress: this.emailaddress.nativeElement.value,
      charactername: this.charactername.nativeElement.value,
      startid: this.startid.nativeElement.value,
      password: this.password.nativeElement.value,
    });

    // redirect to login page
    window.location.href = '/auth/signin';
  }
}
