import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ApiUrlBuilder} from '../helpers/api-url-builder';
import {ApiRequestBuilder} from '../helpers/api-request-builder';
import {SessionResponse} from '../interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly sessionEndpoint: string = '/session/';
  private readonly loginEndpoint: string = '/session/login';

  constructor(
    private readonly http: HttpClient
  ) {}

  public async login(login: string, password: string): Promise<SessionResponse> {
    return await this.http.post(
      ApiUrlBuilder.getApiUrlRequest(this.loginEndpoint),
      ApiRequestBuilder.getLoginRequest(login, password)
    ).toPromise() as SessionResponse;
  }

  public async getSession(): Promise<SessionResponse> {
    return await this.http.get(
      ApiUrlBuilder.getApiUrlRequest(this.sessionEndpoint)
    ).toPromise() as SessionResponse;
  }
}
