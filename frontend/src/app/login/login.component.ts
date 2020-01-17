import { Component, OnInit } from '@angular/core';
import {AuthService} from '../auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public login: string = '';
  public password: string = '';

  public constructor(
    private readonly authService: AuthService
  ) {}

  public tryLogin(): void {
    this.authService.login(this.login, this.password)
      .then(() => {
        console.log('nice');
      })
      .catch(err => console.log(err));
  }

  ngOnInit() {
    this.authService.getSession()
      .then(r => console.log(r))
      .catch(err => console.log(err));
  }
}
