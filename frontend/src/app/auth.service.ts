import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ApiUrlBuilder} from './api-url-builder';
import {ApiRequestBuilder} from './api-request-builder';
import {LoginResponse} from './interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly loginEndpoint: string = '/session/login';

  constructor(
    private readonly http: HttpClient
  ) {}

  public async login(login: string, password: string): Promise<LoginResponse> {
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.loginEndpoint),
      ApiRequestBuilder.getLoginRequest(login, password)
    ).toPromise() as LoginResponse;
  }
}
