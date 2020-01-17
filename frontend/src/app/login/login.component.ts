import { Component, OnInit } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {RoleNavigationService} from '../services/role-navigation.service';
import {SessionResponse, ApiErrorResponse} from '../interfaces/api';
import {StorageService} from '../services/storage.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public login: string = '';
  public password: string = '';
  private readonly statusError: string = 'error';

  public constructor(
    private readonly storage: StorageService,
    private readonly authService: AuthService,
    private readonly roleNavigation: RoleNavigationService
  ) {}

  public tryLogin(): void {
    this.authService.login(this.login, this.password)
      .then(session => this.saveSessionAndNavigateUser(session))
      .catch(err => console.log(err));
  }

  private saveSessionAndNavigateUser(session: SessionResponse | ApiErrorResponse): void {
    if (session.status === this.statusError) {
      console.log((session as ApiErrorResponse).error_detail);
      return;
    }

    session = session as SessionResponse;
    this.storage.saveSession(session.data);
    this.roleNavigation.navigateUser(session);
  }

  ngOnInit() {
    this.authService.getSession()
      .then(session => this.saveSessionAndNavigateUser(session))
      .catch(err => console.log(err));
  }
}
