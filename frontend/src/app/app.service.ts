import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class AppService {
  constructor(private http: HttpClient) {}

  async health(): Promise<string> {
    let response = await fetch(environment.apiUrl + "health");
    return response.text();
  }
}
