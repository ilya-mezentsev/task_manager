import { Component, OnInit } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {RoleNavigationService} from '../services/role-navigation.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public login: string = '';
  public password: string = '';

  public constructor(
    private readonly authService: AuthService,
    private readonly roleNavigation: RoleNavigationService
  ) {}

  public tryLogin(): void {
    this.authService.login(this.login, this.password)
      .then(session => this.roleNavigation.navigateUser(session))
      .catch(err => console.log(err));
  }

  ngOnInit() {
    this.authService.getSession()
      .then(session => this.roleNavigation.navigateUser(session))
      .catch(err => console.log(err));
  }
}
