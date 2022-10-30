import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';

@Component({
  selector: 'app-reset-password',
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.css'],
})
export class ResetPasswordComponent implements OnInit {
  @ViewChild('emailaddress') emailaddress: ElementRef;

  constructor(
    private accountService: AccountService
  ) {}

  ngOnInit(): void {}

  async submit() {
    // try to request reset
    const s = await this.accountService.forgot({
      emailaddress: this.emailaddress.nativeElement.value
    });

    // show result
    alert(s.message);

    // return home on success
    if (s.success) {
      window.location.href = "/";
    }
  }
}
