import { Component, ElementRef, ViewChild, OnInit } from '@angular/core';
import { AccountService } from '../account.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
})
export class SignupComponent implements OnInit {
  @ViewChild('charactername') charactername: ElementRef;
  @ViewChild('password') password: ElementRef;

  constructor(private accountService: AccountService) {}

  ngOnInit(): void {}

  async register() {
    // try to create account
    await this.accountService.register({
      charactername: this.charactername.nativeElement.value,
      password: this.password.nativeElement.value,
    });

    // redirect to login page
    window.location.href = '/auth/signin';
  }
}
