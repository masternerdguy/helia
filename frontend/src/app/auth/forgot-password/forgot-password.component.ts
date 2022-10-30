import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { AccountService } from '../account.service';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css'],
})
export class ForgotPasswordComponent implements OnInit {
  @ViewChild('emailaddress') emailaddress: ElementRef;

  constructor(private accountService: AccountService) {}

  ngOnInit(): void {}

  async submit() {
    // try to request reset
    const s = await this.accountService.forgot({
      emailaddress: this.emailaddress.nativeElement.value,
    });

    // show result
    alert(s.message);

    // return home on success
    if (s.success) {
      window.location.href = '/';
    }
  }
}
