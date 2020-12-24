import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { RegisterModel } from './models/register';
import { environment } from 'src/environments/environment';
import { LoginModel, LoginResultModel } from './models/login';

@Injectable({
  providedIn: 'root',
})
export class AccountService {
  constructor(private http: HttpClient) {}

  async register(a: RegisterModel): Promise<any> {
    return await this.http
      .post<any>(environment.apiUrl + 'register', JSON.stringify(a))
      .toPromise();
  }

  async login(a: LoginModel): Promise<LoginResultModel> {
    return await this.http
      .post<LoginResultModel>(environment.apiUrl + 'login', JSON.stringify(a))
      .toPromise();
  }
}
