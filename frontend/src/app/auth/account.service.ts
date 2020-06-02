import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { RegisterModel } from './models/register';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AccountService {
  constructor(private http: HttpClient) { }

  async register(a: RegisterModel): Promise<boolean> {
    return await this.http.post<boolean>(environment.apiUrl + 'register', JSON.stringify(a)).toPromise();
  }
}
