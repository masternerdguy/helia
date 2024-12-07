import { Component, ElementRef, ViewChild, OnInit } from '@angular/core';
import { AccountService } from '../account.service';

@Component({
    selector: 'app-signup',
    templateUrl: './signup.component.html',
    styleUrls: ['./signup.component.css'],
    standalone: false
})
export class SignupComponent implements OnInit {
  @ViewChild('emailaddress') emailaddress: ElementRef;
  @ViewChild('charactername') charactername: ElementRef;
  @ViewChild('password') password: ElementRef;
  @ViewChild('confirmpassword') confirmpassword: ElementRef;

  startid: string;

  constructor(private accountService: AccountService) {}

  ngOnInit(): void {}

  setStart(e: any) {
    this.startid = e.target.value;
  }

  async register() {
    let ok = true;

    // try to create account
    await this.accountService
      .register({
        emailaddress: this.emailaddress.nativeElement.value,
        charactername: this.charactername.nativeElement.value,
        startid: this.startid,
        password: this.password.nativeElement.value,
        confirmpassword: this.confirmpassword.nativeElement.value,
      })
      .catch((r) => {
        ok = false;

        // show error message
        alert(r.error);
      })
      .then(() => {
        if (ok) {
          // redirect to login page
          window.location.href = '/#/auth/signin';
        }
      });
  }
}
