import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { RegisterModel } from './models/register';
import { environment } from 'src/environments/environment';
import { LoginModel, LoginResultModel } from './models/login';
import { ForgotModel, ForgotResultModel } from './models/forgot';

@Injectable({
  providedIn: 'root',
})
export class AccountService {
  constructor(private http: HttpClient) {}

  async register(a: RegisterModel): Promise<any> {
    return this.http
      .post<any>(environment.apiUrl + 'register', JSON.stringify(a))
      .toPromise();
  }

  async login(a: LoginModel): Promise<LoginResultModel> {
    return this.http
      .post<any>(environment.apiUrl + 'login', JSON.stringify(a))
      .toPromise();
  }

  async forgot(a: ForgotModel): Promise<ForgotResultModel> {
    return this.http
      .post<any>(environment.apiUrl + 'forgot', JSON.stringify(a))
      .toPromise();
  }
}
