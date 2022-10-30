import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AccountService } from '../account.service';

@Component({
  selector: 'app-reset-password',
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.css'],
})
export class ResetPasswordComponent implements OnInit {
  @ViewChild('password') password: ElementRef;
  @ViewChild('confirmpassword') confirmpassword: ElementRef;

  constructor(
    private accountService: AccountService,
    private router: ActivatedRoute
  ) {}

  ngOnInit(): void {}

  async submit() {
    // try to request reset
    const s = await this.accountService.reset({
      password: this.password.nativeElement.value,
      confirmPassword: this.confirmpassword.nativeElement.value,
      userId: this.router.snapshot.queryParamMap.get("u"),
      token: this.router.snapshot.queryParamMap.get("t"),
    });

    // show result
    alert(s.message);

    // redirect to login on success
    if (s.success) {
      window.location.href = "/#/auth/signin";
    }
  }
}
